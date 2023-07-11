// Package routes contains all routes of the application
package routes

import (
	invoiceController "github.com/Raj63/golang-microservices/services/web-portal/pkg/infrastructure/rest/controllers/invoice"
	"github.com/Raj63/golang-microservices/services/web-portal/pkg/infrastructure/rest/middlewares/jwt"
	"github.com/gin-gonic/gin"
)

// InvoiceRoutes is a function that contains all invoice routes
func InvoiceRoutes(router *gin.RouterGroup, jwtMiddleware jwt.Middleware, controller *invoiceController.Controller) {

	routerInvoice := router.Group("/invoices")
	{
		routerInvoice.POST("/", jwtMiddleware.GinAuthMiddleware(jwt.ADMIN), controller.CreateInvoice)
		routerInvoice.GET("/:id", jwtMiddleware.GinAuthMiddleware(jwt.ADMIN, jwt.USER), controller.GetInvoiceByID)
	}

}
