package jwt

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Middleware represents the Auth0 JWT handler middleware
type Middleware interface {
	AuthMiddleware(next http.Handler, roles ...Role) http.Handler
	GinAuthMiddleware(roles ...Role) gin.HandlerFunc
}
