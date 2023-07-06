// Package invoice contains the invoice controller
package invoice

// NewInvoiceRequest is a struct that contains the new invoice request information
type NewInvoiceRequest struct {
	Description string `json:"description" example:"Something" binding:"required"`
	Number      string `json:"number" example:"RF-0017" binding:"required"`
	Status      string `json:"status" example:" DRAFT, OPEN, FILLED, LOCKED, APPROVED" binding:"required"`
	Amount      *Money `json:"amount" binding:"required"`
}

// Money is a struct that represents currency code and value
type Money struct {
	Amount int64  `json:"amount" example:"1200" binding:"required"`
	Code   string `json:"code" example:"INR" binding:"required"`
}
