package callback

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/Raj63/golang-microservices/services/web-portal/pkg/infrastructure/rest/middlewares/auth0/authenticator"
)

// Handler for our callback.
func Handler(auth *authenticator.Authenticator) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		if ctx.Query("state") != session.Get("state") {
			ctx.String(http.StatusBadRequest, "Invalid state parameter.")
			return
		}

		// Exchange an authorization code for a token.
		token, err := auth.Exchange(ctx.Request.Context(), ctx.Query("code"))
		if err != nil {
			ctx.String(http.StatusUnauthorized, "Failed to exchange an authorization code for a token.")
			return
		}

		idToken, err := auth.VerifyIDToken(ctx.Request.Context(), token)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "Failed to verify ID Token.")
			return
		}

		var profile map[string]interface{}
		if err := idToken.Claims(&profile); err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}
		fmt.Printf("profile: %+v\n", profile)

		// Retrieve user roles
		apiAccessToken, err := auth.GetAuth0AccessToken(ctx)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "Failed to get Api access token.")
			return
		}
		userRoles, err := auth.GetUserRoles(ctx, apiAccessToken, profile["sub"].(string))
		if err != nil {
			ctx.String(http.StatusInternalServerError, "Failed to retrieve user roles.")
			return
		}
		fmt.Printf("userRoles: %+v\n", userRoles)
		fmt.Println("accessToken set:", token.AccessToken)
		fmt.Println("id_token", token.Extra("id_token"))
		//session.Set("access_token", token.AccessToken)
		session.Set("id_token", token.Extra("id_token"))
		//session.Set("profile", profile)
		//session.Set("user_roles", userRoles)
		fmt.Println("session set")
		if err := session.Save(); err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}

		fmt.Println("session saved")
		// Redirect to logged in page.
		ctx.Redirect(http.StatusTemporaryRedirect, "/v1/protected")
	}
}
