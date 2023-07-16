// Package routes contains all routes of the application
package routes

import (
	"encoding/gob"

	"github.com/Raj63/go-sdk/logger"
	sdksql "github.com/Raj63/go-sdk/sql"
	"github.com/gin-contrib/sessions/cookie"

	"github.com/gin-contrib/sessions"

	//"github.com/gorilla/sessions"
	// swaggerFiles for documentation
	_ "github.com/Raj63/golang-microservices/services/web-portal/docs"
	"github.com/Raj63/golang-microservices/services/web-portal/pkg/infrastructure/config"
	"github.com/Raj63/golang-microservices/services/web-portal/pkg/infrastructure/rest/adapter"
	"github.com/Raj63/golang-microservices/services/web-portal/pkg/infrastructure/rest/middlewares/auth0/authenticator"
	"github.com/Raj63/golang-microservices/services/web-portal/pkg/infrastructure/rest/middlewares/auth0/token"
	"github.com/Raj63/golang-microservices/services/web-portal/pkg/integration/grpc/invoices"
	"github.com/gin-gonic/gin"
)

// DI is the dependency injection entity for ApplicationV1Router
type DI struct {
	Router              *gin.Engine
	DB                  *sdksql.DB
	Logger              *logger.Logger
	Config              config.AppConfig
	Authenticator       *authenticator.Authenticator
	InvoicesGRPCService invoices.Service
}

//	@title			Web-Portal
//	@version		2.0
//	@description	Documentation's Web-Portal
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	Rajesh Kumar Biswas
//	@contact.url	https://github.com/Raj63
//	@contact.email	biswas.rajesh63@gmail.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// ApplicationV1Router is a function that contains all routes of the application
//
//	@host		localhost:8080
//	@BasePath	/v1
func ApplicationV1Router(di DI) {
	// the application errors will be processed here before returning to the caller
	// di.Router.Use(errorsController.Handler)

	// To store custom types in our cookies,
	// we must first register them using gob.Register
	gob.Register(map[string]interface{}{})
	store := cookie.NewStore([]byte("secret"))

	routerV1 := di.Router.Group("/v1")
	{
		// Use sessions handler for initializing Sessions from context
		routerV1.Use(sessions.Sessions("auth-session", store))

		// Register Login routes
		LoginRoutes(routerV1, di.Config, di.Authenticator)

		// Use the token validation middleware for protected routes
		routerV1.Use(token.ValidationMiddleware(di.Authenticator))

		// Register the example protected routes
		ProtectedRoutes(routerV1)

		// Register logout routes
		LogoutRoutes(routerV1, di.Config)

		// Register invoice routes
		InvoiceRoutes(routerV1, adapter.InvoiceAdapter(&adapter.InvoiceAdapterDI{
			Logger:         di.Logger,
			InvoiceService: di.InvoicesGRPCService,
		}))
	}
}
