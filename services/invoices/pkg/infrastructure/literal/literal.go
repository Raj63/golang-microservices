package literal

// const for mislleneous literals
const (
	EmptyString = ""
)

// const for validation errors
const (
	InvalidInvoiceID             = "Invalid input: Invoice ID"
	InvalidInvoiceNumber         = "Invalid input: Invoice Number"
	InvalidInvoiceDescription    = "Invalid input: Invoice Description"
	InvalidInvoiceAmount         = "Invalid input: Invoice Amount"
	InvalidInvoiceAmountValue    = "Invalid input: Invoice Amount value"
	InvalidInvoiceAmountCurrency = "Invalid input: Invoice Currency code"
	InvalidIssuerID              = "Invalid input: Issuer ID"
	InvalidPagingOption          = "Invalid input: Paging Option"
	InvalidInvestorID            = "Invalid input: Investor ID"
)

// const for service errors
const (
	ErrCreateInvoice    = "Invoices cannot be created"
	ErrGetInvoice       = "Invoices cannot be fetched"
	ErrApproveTrade     = "Trade cannot be approved"
	ErrGetIssuerBalance = "Issuer Balance cannot be fetched"
	ErrGetInvestorList  = "Investor List cannot be fetched"
	ErrPlacingBids      = "Placing Bid cannot be performed"
)
