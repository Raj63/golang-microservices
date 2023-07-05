// Package invoice contains the invoice controller
package invoice

import (
	"errors"

	"github.com/Raj63/go-sdk/logger"
	domainErrors "github.com/Raj63/golang-microservices/services/web-portal/pkg/domain/errors"
	domainInvoice "github.com/Raj63/golang-microservices/services/web-portal/pkg/domain/invoice"
	"github.com/Raj63/golang-microservices/services/web-portal/pkg/infrastructure/rest/controllers"
	"github.com/google/uuid"

	"net/http"

	"github.com/gin-gonic/gin"
)

// Controller is a struct that contains the invoice service
type Controller struct {
	Logger         *logger.Logger
	InvoiceService domainInvoice.Service
}

// CreateInvoice godoc
//
//	@Tags			invoices
//	@Summary		Create New Invoice
//	@Description	Create new invoice on the system
//	@Accept			json
//	@Produce		json
//	@Param			data	body		NewInvoiceRequest	true	"body data"
//	@Success		201		{object}	domainInvoice.Invoice
//	@Failure		400		{object}	MessageResponse
//	@Failure		500		{object}	MessageResponse
//	@Router			/invoices [post]
func (c *Controller) CreateInvoice(ctx *gin.Context) {
	var request NewInvoiceRequest

	if err := controllers.BindJSON(ctx, &request); err != nil {
		appError := domainErrors.NewAppError(err, domainErrors.ValidationError)
		_ = ctx.Error(appError)
		return
	}
	newInvoice := domainInvoice.Invoice{
		Description: request.Description,
	}

	var result *domainInvoice.Invoice
	var err error

	result, err = c.InvoiceService.Create(ctx.Request.Context(), &newInvoice)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, result)
}

// GetInvoiceByID godoc
//
//	@Tags			invoices
//	@Summary		Get invoice by ID
//	@Description	Get Invoice by ID on the system
//	@Param			invoice_id	path		string	true	"id of invoice"
//	@Success		200			{object}	domainInvoice.Invoice
//	@Failure		400			{object}	MessageResponse
//	@Failure		500			{object}	MessageResponse
//	@Router			/invoices/{invoice_id} [get]
func (c *Controller) GetInvoiceByID(ctx *gin.Context) {
	invoiceID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		appError := domainErrors.NewAppError(errors.New("invoice id is invalid"), domainErrors.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	domainInvoice, err := c.InvoiceService.Get(ctx.Request.Context(), invoiceID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, domainInvoice)
}
