build:
	docker build -t invoices .

run:
	docker-compose down
	docker-compose up

reload:
	docker-compose down
	docker image rm invoices
	docker image rm invoices_app
	docker build -t invoices .
	docker-compose up

test:
	go test -count=1 -failfast -v -race ./... -coverprofile=coverage.out && go tool cover -html=coverage.out -o coverage.html && go tool cover -func coverage.out

generate-mocks:
	mockgen -source=pkg/service/investors.go -destination=pkg/service/mocks/investors.go -package mocks
	mockgen -source=pkg/service/invoices.go -destination=pkg/service/mocks/invoices.go -package mocks
	mockgen -source=pkg/service/issuers.go -destination=pkg/service/mocks/issuers.go -package mocks
	
generate:
	protoc -I../../api -Iapi --go_opt=paths=source_relative --go_out=./api/ --go-grpc_opt=paths=source_relative --go-grpc_out=./api/ --grpc-gateway_opt=logtostderr=true,paths=source_relative --grpc-gateway_out=./api/ --openapiv2_out=:docs api/*.proto

db-update:
	docker run web-portal update-db