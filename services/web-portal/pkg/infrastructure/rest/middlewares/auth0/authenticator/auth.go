package authenticator

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/Raj63/go-sdk/logger"
	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

// DI is the dependency injection entity for Authenticator
type DI struct {
	Auth0ClientID     string
	Auth0ClientSecret string
	Auth0CallbackURL  string
	Auth0Domain       string
	Auth0SharedSecret string

	Logger *logger.Logger
}

// Auth0AccessTokenResponse represents the response from Auth0 access token authorization
type Auth0AccessTokenResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
}

// Auth0RoleRequest represents the Auth0Role request
type Auth0RoleRequest struct {
	Name string `json:"name"`
}

// Auth0Role represents the Auth0Role type
type Auth0Role struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Auth0UserRoleRequest represents the Auth0UserRole request
type Auth0UserRoleRequest struct {
	Roles []string `json:"roles"`
}

// Authenticator is used to authenticate our users.
type Authenticator struct {
	*oidc.Provider
	oauth2.Config
	auth0ClientID     string
	auth0ClientSecret string
	auth0CallbackURL  string
	auth0Domain       string
	auth0SharedSecret string

	lock              sync.Mutex
	auth0Token        *Auth0AccessTokenResponse
	accessTokenExpiry time.Time
	logger            *logger.Logger
}

// New instantiates the *Authenticator.
func New(di *DI) (*Authenticator, error) {
	provider, err := oidc.NewProvider(
		context.Background(),
		"https://"+di.Auth0Domain+"/",
	)
	if err != nil {
		return nil, err
	}

	conf := oauth2.Config{
		ClientID:     di.Auth0ClientID,
		ClientSecret: di.Auth0ClientSecret,
		RedirectURL:  di.Auth0CallbackURL,
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile", "role"},
	}

	return &Authenticator{
		Provider:          provider,
		Config:            conf,
		auth0ClientID:     di.Auth0ClientID,
		auth0ClientSecret: di.Auth0ClientSecret,
		auth0CallbackURL:  di.Auth0CallbackURL,
		auth0Domain:       di.Auth0Domain,
		auth0SharedSecret: di.Auth0SharedSecret,
		logger:            di.Logger,
	}, nil
}

// VerifyIDToken verifies that an *oauth2.Token is a valid *oidc.IDToken.
func (a *Authenticator) VerifyIDToken(ctx context.Context, token *oauth2.Token) (*oidc.IDToken, error) {
	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		return nil, errors.New("no id_token field in oauth2 token")
	}

	oidcConfig := &oidc.Config{
		ClientID: a.ClientID,
	}

	return a.Verifier(oidcConfig).Verify(ctx, rawIDToken)
}

// GetUserRoles fetches user roles from oauth0
func (a *Authenticator) GetUserRoles(ctx context.Context, accessToken, userID string) ([]string, error) {
	url := fmt.Sprintf("https://%s/api/v2/users/%s", a.auth0Domain, userID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		a.logger.ErrorfContext(ctx, "Error requesting:%v\n", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			a.logger.ErrorfContext(ctx, "Error reading:%v\n", err)
		} else {
			a.logger.ErrorfContext(ctx, "failed response data: %s", data)
		}
		return nil, fmt.Errorf("unexpected response: %s", resp.Status)
	}

	var userResponse struct {
		Roles []string `json:"roles"`
	}
	err = json.NewDecoder(resp.Body).Decode(&userResponse)
	if err != nil {
		return nil, err
	}

	return userResponse.Roles, nil
}

// GetAuth0AccessToken fetches auth0 access token from the server
func (a *Authenticator) GetAuth0AccessToken(ctx context.Context) (string, error) {
	a.lock.Lock()
	defer a.lock.Unlock()

	// Check if a valid access token is already available
	if a.auth0Token != nil && a.auth0Token.AccessToken != "" && time.Now().Before(a.accessTokenExpiry) {
		return a.auth0Token.AccessToken, nil
	}

	// Prepare the request body with grant type and client credentials
	body := url.Values{}
	body.Set("grant_type", "client_credentials")
	body.Set("client_id", a.auth0ClientID)
	body.Set("client_secret", a.auth0ClientSecret)
	body.Set("audience", fmt.Sprintf("https://%s/api/v2/", a.auth0Domain))
	body.Set("scope", "create:roles read:roles update:roles")

	// Use the refresh token if available
	if a.auth0Token != nil && a.auth0Token.RefreshToken != "" {
		body.Set("refresh_token", a.auth0Token.RefreshToken)
	}

	tokenURL := "https://" + a.auth0Domain + "/oauth/token"
	// Send a POST request to obtain the access token
	// Suppress gosec warning for this line
	// #nosec G107
	resp, err := http.PostForm(tokenURL, body)
	if err != nil {
		a.logger.ErrorfContext(ctx, "Error requesting:%v\n", err)
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			a.logger.ErrorfContext(ctx, "Error reading:%v\n", err)
		} else {
			a.logger.ErrorfContext(ctx, "failed response data: %s", data)
		}
		return "", fmt.Errorf("unexpected response: %s", resp.Status)
	}

	tokenResponse := new(Auth0AccessTokenResponse)
	err = json.NewDecoder(resp.Body).Decode(tokenResponse)
	if err != nil {
		return "", err
	}

	a.auth0Token = tokenResponse
	a.accessTokenExpiry = time.Now().Add(time.Duration(tokenResponse.ExpiresIn) * time.Second)

	return tokenResponse.AccessToken, nil
}

// CreateRole creates a new role to the list
func (a *Authenticator) CreateRole(ctx context.Context, accessToken, roleName string) (string, error) {
	rolesURL := fmt.Sprintf("https://%s/api/v2/roles", a.auth0Domain)

	requestBody, err := json.Marshal(Auth0RoleRequest{Name: roleName})
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", rolesURL, strings.NewReader(string(requestBody)))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		a.logger.ErrorfContext(ctx, "Error requesting:%v\n", err)
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			a.logger.ErrorfContext(ctx, "Error reading:%v\n", err)
		} else {
			a.logger.ErrorfContext(ctx, "failed response data: %s", data)
		}
		return "", fmt.Errorf("unexpected response: %s", resp.Status)
	}

	var roleResponse Auth0Role
	err = json.NewDecoder(resp.Body).Decode(&roleResponse)
	if err != nil {
		return "", err
	}

	return roleResponse.ID, nil
}

// AssignRoleToUser assigns the given user to the given role
func (a *Authenticator) AssignRoleToUser(ctx context.Context, accessToken, userID, roleName string) error {
	userRolesURL := fmt.Sprintf("https://%s/api/v2/users/%s/roles", a.auth0Domain, userID)

	requestBody, err := json.Marshal(Auth0UserRoleRequest{Roles: []string{roleName}})
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", userRolesURL, strings.NewReader(string(requestBody)))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		a.logger.ErrorfContext(ctx, "Error requesting:%v\n", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			a.logger.ErrorfContext(ctx, "Error reading:%v\n", err)
		} else {
			a.logger.ErrorfContext(ctx, "failed response data: %s", data)
		}
		return fmt.Errorf("unexpected response: %s", resp.Status)
	}

	return nil
}

// GetAuth0Roles fetches all the auth0 roles
func (a *Authenticator) GetAuth0Roles(ctx context.Context, accessToken string) ([]Auth0Role, error) {
	url := fmt.Sprintf("https://%s/api/v2/roles", a.auth0Domain)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		a.logger.ErrorfContext(ctx, "Error requesting:%v\n", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			a.logger.ErrorfContext(ctx, "Error reading:%v\n", err)
		} else {
			a.logger.ErrorfContext(ctx, "failed response data: %s", data)
		}
		return nil, fmt.Errorf("unexpected response: %s", resp.Status)
	}

	var roles []Auth0Role
	err = json.NewDecoder(resp.Body).Decode(&roles)
	if err != nil {
		return nil, err
	}

	return roles, nil
}

// Init checks the roles available in the auth0 namespace
func (a *Authenticator) Init(ctx context.Context) error {
	accessToken, err := a.GetAuth0AccessToken(ctx)
	if err != nil {
		return fmt.Errorf("error fetching access token: %v", err)
	}

	roles, err := a.GetAuth0Roles(ctx, accessToken)
	if err != nil {
		return fmt.Errorf("error fetching auth0 roles: %v", err)
	}

	for _, role := range roles {
		fmt.Printf("role: %v\n", role)
	}

	return nil
}
