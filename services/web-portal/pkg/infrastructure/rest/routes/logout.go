package routes

import (
	"github.com/Raj63/golang-microservices/services/web-portal/pkg/infrastructure/config"
	"github.com/Raj63/golang-microservices/services/web-portal/pkg/infrastructure/rest/middlewares/auth0/logout"
	"github.com/gin-gonic/gin"
)

// LogoutRoutes is a function that contains logout routes
func LogoutRoutes(router *gin.RouterGroup, config config.AppConfig) {
	router.GET("/logout", logout.Handler(&logout.DI{
		Auth0ClientID: config.Auth0Config.Auth0ClientID,
		Auth0Domain:   config.Auth0Config.Auth0Domain,
	}))
}
