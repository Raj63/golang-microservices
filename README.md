# golang-microservices

A sample production grade REST API application written in Golang with Invoice managment business functionalities.

- Microservices
- Docker Containers
- REST Api + GRPC for internal Communication
- Distributed logging
- Distributed tracing
- Monitoring (grafana + Prometheus)
- Performance analysis(pprof + benchmark testing + autocannon)
- Postgres database with Auto Migration and Seeding
- Debugging capacity with inspector
- Load .env files for configuration settings
- 80% Unit testing coverage
- Mocks (mockery)
- Autoi generated Swagger OpenApi & Markdown Docs
- ThunderClient requests for Testing APIs
- Makefile for easy commands
- Docker-Compose for local running
- Enforcing Coding standards with `https://github.com/tekwizely/pre-commit-golang`

## Generate Swagger documentation for the Server

- Run command `make generate`

## Test Application Server

- Run command `make test`
  
## Build Application Server

- Run command `make build`

## Run Application Server

- Run command `make run`

## Reload Application Server

- Run command `make reload`

### When we run above run or reload command, the db container and app container will be up and running then we can access APIs provided by the App

- The terminal shell looks like this
![alt](screenshots/run-cli.png)

- The swagger documentation UI will be available at `http://localhost:8080/swagger/index.html`
![alt](screenshots/swagger-ui.png)

- The PProf will be avilable at `http://localhost:8080/debug/pprof`
![alt](screenshots/pprof.png)

- The Inspector UI will be available at `hhttp://localhost:8080/_inspector`
![alt](screenshots/inspector.png)

- The metrics api will be available at `hhttp://localhost:8080/metrics`
![alt](screenshots/metrics.png)

- The Newrelic UI will look like this
![alt](screenshots/newrelic.png)

- The Thunder-Client dashboard will look like this
![alt](screenshots/thunderclient.png)

- The docker state will look like this
![alt](screenshots/docker-ps.png)

- The Coding standard will be enforeced while commiting the code to the repository
![alt](screenshots/precommit-1.png)
![alt](screenshots/precommit-2.png)
