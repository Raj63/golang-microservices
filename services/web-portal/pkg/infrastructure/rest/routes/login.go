package routes

import (
	"github.com/Raj63/golang-microservices/services/web-portal/pkg/infrastructure/config"
	"github.com/Raj63/golang-microservices/services/web-portal/pkg/infrastructure/rest/middlewares/auth0/authenticator"
	"github.com/Raj63/golang-microservices/services/web-portal/pkg/infrastructure/rest/middlewares/auth0/callback"
	"github.com/Raj63/golang-microservices/services/web-portal/pkg/infrastructure/rest/middlewares/auth0/login"
	"github.com/gin-gonic/gin"
)

// LoginRoutes is a function that contains all login related routes
func LoginRoutes(router *gin.RouterGroup, config config.AppConfig, authprovider *authenticator.Authenticator) {
	router.GET("/login", login.Handler(authprovider))
	router.GET("/callback", callback.Handler(authprovider))
}
