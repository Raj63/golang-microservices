// Package routes contains all routes of the application
package routes

import (
	"github.com/Raj63/go-sdk/logger"
	sdksql "github.com/Raj63/go-sdk/sql"

	// swaggerFiles for documentation
	_ "github.com/Raj63/golang-microservices/services/web-portal/docs"
	"github.com/Raj63/golang-microservices/services/web-portal/pkg/infrastructure/rest/adapter"
	"github.com/Raj63/golang-microservices/services/web-portal/pkg/infrastructure/rest/middlewares/jwt"
	"github.com/Raj63/golang-microservices/services/web-portal/pkg/integration/grpc/invoices"
	"github.com/gin-gonic/gin"
)

// DI is the dependency injection entity for ApplicationV1Router
type DI struct {
	Router              *gin.Engine
	DB                  *sdksql.DB
	Logger              *logger.Logger
	JWTMiddleware       jwt.Middleware
	InvoicesGRPCService invoices.Service
}

//	@title			Golang Integration Sample
//	@version		2.0
//	@description	Documentation's Golang Integration Sample
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
	routerV1 := di.Router.Group("/v1")
	{
		InvoiceRoutes(routerV1, di.JWTMiddleware, adapter.InvoiceAdapter(&adapter.InvoiceAdapterDI{
			Logger:         di.Logger,
			InvoiceService: di.InvoicesGRPCService,
		}))
	}
}
