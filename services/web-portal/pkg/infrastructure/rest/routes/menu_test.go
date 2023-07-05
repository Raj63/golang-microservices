// Package routes contains all routes of the application
package routes_test

/*
import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	domainInvoice "github.com/Raj63/golang-microservices/services/web-portal/pkg/domain/invoice"
	"github.com/gin-gonic/gin"

	appErr "github.com/Raj63/golang-microservices/services/web-portal/pkg/domain/errors"
	mockRepository "github.com/Raj63/golang-microservices/services/web-portal/pkg/infrastructure/mocks/repository"
	"github.com/Raj63/golang-microservices/services/web-portal/pkg/infrastructure/repository"
	errorsController "github.com/Raj63/golang-microservices/services/web-portal/pkg/infrastructure/rest/controllers/errors"
	invoiceController "github.com/Raj63/golang-microservices/services/web-portal/pkg/infrastructure/rest/controllers/invoice"
	"github.com/Raj63/golang-microservices/services/web-portal/pkg/infrastructure/rest/routes"
	"github.com/brianvoe/gofakeit"
	"github.com/golang/mock/gomock"
)

func getTestRouter() (*gin.Engine, *gin.RouterGroup) {
	// initialize the router
	router := gin.Default()
	// the application errors will be processed here before returning to the caller
	router.Use(errorsController.Handler)

	return router, router.Group("/v1")
}

func TestInvoiceRoutes(t *testing.T) {

	invoiceName := gofakeit.BeerHop()
	invoiceDesc1 := gofakeit.BeerName()
	invoicePrice1 := gofakeit.Float64()
	invoiceName2 := gofakeit.BeerHop()
	invoiceDesc2 := gofakeit.BeerName()
	invoicePrice2 := gofakeit.Float64()
	invoiceName3 := gofakeit.BeerHop()
	invoiceDesc3 := gofakeit.BeerName()
	invoicePrice3 := gofakeit.Float64()
	type args struct {
		method       string
		endpoint     string
		body         interface{}
		mockrepoFn   func() repository.Invoices
		outputStatus int
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "Add new Invoice successfully",
			args: args{
				method:   "POST",
				endpoint: "/v1/invoices/",
				body: invoiceController.NewInvoiceRequest{
					Name:        invoiceName,
					Description: invoiceDesc1,
					Price:       float64(invoicePrice1),
				},
				outputStatus: http.StatusCreated,
				mockrepoFn: func() repository.Invoices {
					mRepository := mockRepository.NewMockInvoices(gomock.NewController(t))
					mRepository.EXPECT().Create(gomock.Any(), gomock.Any()).AnyTimes().Return(&domainInvoice.Invoice{
						ID:          gofakeit.Int64(),
						Name:        invoiceName,
						Description: invoiceDesc1,
						Price:       invoicePrice1,
						CreatedAt:   time.Now(),
						UpdatedAt:   time.Now(),
					}, nil)
					return mRepository
				},
			},
		},
		{
			name: "Add new Invoice failed due to missing invoice desc validation error",
			args: args{
				method:   "POST",
				endpoint: "/v1/invoices/",
				body: invoiceController.NewInvoiceRequest{
					Name:  invoiceName,
					Price: float64(invoicePrice1),
				},
				outputStatus: http.StatusBadRequest,
				mockrepoFn: func() repository.Invoices {
					mRepository := mockRepository.NewMockInvoices(gomock.NewController(t))
					return mRepository
				},
			},
		},
		{
			name: "Fetch Invoice by ID successfully",
			args: args{
				method:       "GET",
				endpoint:     "/v1/invoices/1",
				outputStatus: http.StatusOK,
				mockrepoFn: func() repository.Invoices {
					mRepository := mockRepository.NewMockInvoices(gomock.NewController(t))
					mRepository.EXPECT().GetByID(gomock.Any(), gomock.Any()).AnyTimes().Return(&domainInvoice.Invoice{
						ID:          gofakeit.Int64(),
						Name:        invoiceName,
						Description: invoiceDesc1,
						Price:       invoicePrice1,
						CreatedAt:   time.Now(),
						UpdatedAt:   time.Now(),
					}, nil)
					return mRepository
				},
			},
		},
		{
			name: "Failed to fetch Invoice by ID due to repository error",
			args: args{
				method:       "GET",
				endpoint:     "/v1/invoices/1",
				outputStatus: http.StatusInternalServerError,
				mockrepoFn: func() repository.Invoices {
					mRepository := mockRepository.NewMockInvoices(gomock.NewController(t))
					mRepository.EXPECT().GetByID(gomock.Any(), gomock.Any()).AnyTimes().Return(nil, appErr.NewAppErrorWithType(appErr.RepositoryError))
					return mRepository
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			if tt.args.body != nil {
				err := json.NewEncoder(&buf).Encode(tt.args.body)
				if err != nil {
					log.Fatal(err)
				}
			}

			req, err := http.NewRequest(tt.args.method, tt.args.endpoint, &buf)
			if err != nil {
				t.Errorf("Error creating a new request: %v", err)
			}
			rr := httptest.NewRecorder()
			router, routerV1 := getTestRouter()
			routes.InvoiceRoutes(routerV1, &invoiceController.Controller{InvoiceService: invoiceService.Service{InvoiceRepository: tt.args.mockrepoFn()}})
			router.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.args.outputStatus {
				t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.", tt.args.outputStatus, status)
			}
		})
	}
}
*/
