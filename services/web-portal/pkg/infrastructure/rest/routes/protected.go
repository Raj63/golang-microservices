package routes

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// ProtectedRoutes contains example routes which are protected by Authetications.
func ProtectedRoutes(router *gin.RouterGroup) {
	router.GET("/protected", ProtectedRouteHandler())
}

// ProtectedRouteHandler ensures the token validation.
func ProtectedRouteHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Access token has already been validated by the middleware
		// You can access the user profile and roles from the session

		session := sessions.Default(ctx)
		profile := session.Get("profile")
		userRoles := session.Get("user_roles")

		// Example response for an authenticated user
		ctx.JSON(http.StatusOK, gin.H{"message": "Authenticated user", "profile": profile, "roles": userRoles})
	}
}
