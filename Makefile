include .env
export

run:
	@./bin/pickside-service

run_is: build_is
	@./bin/image-service
	
build: build_api build_is

build_api:
	@echo "Building API service..."
	@go build -o ./bin/pickside-service cmd/api/main.go

build_is:
	@echo "Building Upload service..."
	@go build -o ./bin/image-service cmd/image-service/main.go

up:
	migrate -path db/migrations/ -database "mysql://$(DSN)" -verbose up
down:
	migrate -path db/migrations/ -database "mysql://$(DSN)" -verbose down

migration:
	@migrate create -ext sql -dir cmd/db/migrations $(filter-out $@,$(MAKECMDGOALS))

drop:
	@go run cmd/drop/main.go

seed:
	@go run db/seed/main.go
