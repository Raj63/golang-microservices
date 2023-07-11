package auth0

import (
	"context"
	"net/http"
	"net/url"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"github.com/Raj63/go-sdk/logger"
	contextkey "github.com/Raj63/golang-microservices/services/web-portal/pkg/infrastructure/rest/middlewares/context"
	ijwt "github.com/Raj63/golang-microservices/services/web-portal/pkg/infrastructure/rest/middlewares/jwt"
)

// DI is the dependency Injection entity
type DI struct {
	Auth0Domain    string
	Auth0Audience  string
	Auth0Namespace string
	// AllowedOrigin value. Ex: "http://localhost:3000"
	AllowedOrigin string
	Logger        *logger.Logger
	Enabled       bool
}

type auth0Jwt struct {
	auth0Domain    string
	auth0Audience  string
	auth0Namespace string
	allowedOrigin  string
	logger         *logger.Logger
	middleware     *jwtmiddleware.JWTMiddleware
}

// NewJWT returns implementatin of jwt.Middleware interface
func NewJWT(di *DI) (ijwt.Middleware, error) {
	if !di.Enabled {
		return &auth0Jwt{}, nil
	}

	issuerURL, err := url.Parse("https://" + di.Auth0Domain + "/")
	if err != nil {
		di.Logger.Errorf("Failed to parse the issuer url: %v", err)
		return nil, err
	}

	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

	jwtValidator, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		[]string{di.Auth0Audience},
		validator.WithCustomClaims(
			func() validator.CustomClaims {
				return &CustomClaims{}
			},
		),
		validator.WithAllowedClockSkew(time.Minute),
	)
	if err != nil {
		di.Logger.Errorf("Failed to set up the jwt validator")
		return nil, err
	}

	errorHandler := func(w http.ResponseWriter, r *http.Request, err error) {
		di.Logger.Errorf("Encountered error while validating JWT: %v", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)

		_, err = w.Write([]byte(`{"message":"Failed to validate JWT."}`))
		if err != nil {
			di.Logger.Errorf("Error writing: %v", err)
		}
	}

	middleware := jwtmiddleware.New(
		jwtValidator.ValidateToken,
		jwtmiddleware.WithErrorHandler(errorHandler),
	)

	return &auth0Jwt{
		auth0Domain:    di.Auth0Domain,
		auth0Audience:  di.Auth0Audience,
		allowedOrigin:  di.AllowedOrigin,
		auth0Namespace: di.Auth0Namespace,
		logger:         di.Logger,
		middleware:     middleware,
	}, nil
}

// AuthMiddleware implements the JWT Auth middleware interface
func (j *auth0Jwt) AuthMiddleware(next http.Handler, roles ...ijwt.Role) http.Handler {
	if j.middleware == nil {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Serve the next handler
			next.ServeHTTP(w, r)
		})
	}
	return j.middleware.CheckJWT(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// CORS Headers.
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Origin", j.allowedOrigin)
		w.Header().Set("Access-Control-Allow-Headers", "Authorization")
		w.Header().Set("Content-Type", "application/json")

		// Check Authorization by roles
		role, ok := j.isAuthorizedUser(r.Context(), roles...)
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			_, err := w.Write([]byte(`{"message":"Unauthorized user."}`))
			if err != nil {
				j.logger.Errorf("Error writing: %v", err)
			}
			return
		}

		// Pass the role to the next middleware or handler
		ctx := context.WithValue(r.Context(), contextkey.RoleContextKey, role)
		r = r.WithContext(ctx)

		// Serve the next handler
		next.ServeHTTP(w, r)
	}))
}

// GinAuthMiddleware implements the JWT Gin compatible Auth middleware interface
func (j *auth0Jwt) GinAuthMiddleware(roles ...ijwt.Role) gin.HandlerFunc {
	if j.middleware == nil {
		return func(c *gin.Context) {
			// Serve the next handler
			c.Next()
		}
	}
	return func(c *gin.Context) {
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// CORS Headers.
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Header("Access-Control-Allow-Origin", j.allowedOrigin)
			c.Header("Access-Control-Allow-Headers", "Authorization")
			c.Header("Content-Type", "application/json")

			// Check Authorization by roles
			role, ok := j.isAuthorizedUser(r.Context(), roles...)
			if ok {
				// Pass the role to the next middleware or handler
				c.Set(string(contextkey.RoleContextKey), role)

				// Serve the next handler
				c.Next()
				return
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized user."})
		})

		j.middleware.CheckJWT(handler).ServeHTTP(c.Writer, c.Request)
	}
}

func (j *auth0Jwt) isAuthorizedUser(ctx context.Context, roles ...ijwt.Role) (string, bool) {
	token := ctx.Value("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	myClaim := "https://" + j.auth0Namespace + "/role"
	var role string
	if claims[myClaim] != nil {
		role = claims[myClaim].(string)
	}

	// if no roles are passed then return true and value
	if len(roles) == 0 {
		return role, true
	}

	roleType := ijwt.ToRole(role)
	for _, r := range roles {
		if r == roleType {
			return role, true
		}
	}
	return role, false
}
