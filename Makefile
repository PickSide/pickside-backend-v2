include .env
export

run: build
	@./bin/pickside-service

build:
	@go build -o ./bin/pickside-service cmd/api/main.go

up:
	migrate -path db/migrations/ -database "mysql://$(DSN)" -verbose up
down:
	@migrate -path db/migrations/ -database "mysql://$(DSN)" -verbose down

migration:
	@migrate create -ext sql -dir cmd/db/migrations $(filter-out $@,$(MAKECMDGOALS))

drop:
	@go run cmd/drop/main.go

seed:
	@go run db/seed/main.go
