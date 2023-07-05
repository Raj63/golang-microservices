// Package routes contains all routes of the application
package routes

import (
	invoiceController "github.com/Raj63/golang-microservices/services/web-portal/pkg/infrastructure/rest/controllers/invoice"
	"github.com/gin-gonic/gin"
)

// InvoiceRoutes is a function that contains all invoice routes
func InvoiceRoutes(router *gin.RouterGroup, controller *invoiceController.Controller) {

	routerInvoice := router.Group("/invoices")
	{
		routerInvoice.POST("/", controller.CreateInvoice)
		routerInvoice.GET("/:id", controller.GetInvoiceByID)
	}

}
