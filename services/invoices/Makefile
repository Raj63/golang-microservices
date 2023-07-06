.PHONY: dep build test generate-mocks generate-proto precommit db-update

dep:
	go mod tidy

build:
	docker build -t invoices .

test:
	go test -count=1 -failfast -v -race ./... -coverprofile=coverage.out && go tool cover -html=coverage.out -o coverage.html && go tool cover -func coverage.out

generate-mocks:
	mockgen -source=pkg/service/investor.go -destination=pkg/service/mocks/investor.go -package mocks
	mockgen -source=pkg/service/invoice.go -destination=pkg/service/mocks/invoice.go -package mocks
	mockgen -source=pkg/service/issuer.go -destination=pkg/service/mocks/issuer.go -package mocks

generate-proto:
	protoc -I../../api -Iapi --go_opt=paths=source_relative --go_out=./api/ --go-grpc_opt=paths=source_relative --go-grpc_out=./api/ --grpc-gateway_opt=logtostderr=true,paths=source_relative --grpc-gateway_out=./api/ --openapiv2_out=:docs api/*.proto

db-update:
	docker run web-portal update-db

precommit: generate-proto generate-mocks dep test
