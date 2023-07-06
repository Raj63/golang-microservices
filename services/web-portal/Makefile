.PHONY: dep build test generate-mocks generate-docs precommit db-update

dep:
	go mod tidy

build:
	docker build -t web-portal .

test:
	go test -count=1 -failfast -v -race ./... -coverprofile=coverage.out && go tool cover -html=coverage.out -o coverage.html && go tool cover -func coverage.out

generate-mocks:
	mockgen -source=pkg/domain/invoice/invoice.go -destination=pkg/domain/mocks/invoice/invoice.go -package mocks
	mockgen -source=pkg/integration/grpc/invoices/invoices.go -destination=pkg/integration/grpc/mocks/invoices/invoices.go -package mocks

generate-docs:
	swag init -g pkg/infrastructure/rest/routes/routes.go

db-update:
	docker run web-portal update-db

precommit: generate-docs generate-mocks dep test
