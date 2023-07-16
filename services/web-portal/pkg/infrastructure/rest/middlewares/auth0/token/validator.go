package token

import (
	"net/http"

	"github.com/Raj63/golang-microservices/services/web-portal/pkg/infrastructure/rest/middlewares/auth0/authenticator"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

// ValidationMiddleware is a middleware that will validate the user request token
func ValidationMiddleware(auth *authenticator.Authenticator) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		idToken := session.Get("id_token")
		token := &oauth2.Token{}
		token = token.WithExtra(map[string]interface{}{"id_token": idToken})
		// Verify the ID token using the VerifyIDToken method
		tokenClaims, err := auth.VerifyIDToken(ctx.Request.Context(), token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}
		var profile map[string]interface{}
		if err := tokenClaims.Claims(&profile); err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing profile claim"})
			return
		}
		ctx.Set("profile", profile)

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
		ctx.Set("roles", userRoles)

		// Token is valid, continue processing the request
		ctx.Next()
	}
}
