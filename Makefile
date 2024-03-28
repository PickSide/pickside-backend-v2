include .env
export

run: build_api
	@./bin/pickside-service

run_fs: build_fs
	@./bin/file-service

run_nfs: build_nfs
	@./bin/notifications-service
	
run_qs: build_qs
	@./bin/queue-service


build_api:
	@echo "Building API service..."
	@go build -o ./bin/pickside-service cmd/api/main.go

build_fs:
	@echo "Building File service..."
	@go build -o ./bin/file-service cmd/file-service/main.go

build_nfs:
	@echo "Building Notifications service..."
	@go build -o ./bin/notifications-service cmd/notifications-service/main.go

build_qs:
	@echo "Building Queue service..."
	@go build -v -o ./bin/queue-service cmd/queue-service/main.go

build: build_api build_fs build_nfs build_qs

up:
	migrate -path db/migrations/ -database "mysql://$(DSN)" -verbose up
down:
	migrate -path db/migrations/ -database "mysql://$(DSN)" -verbose down

migration:
	@migrate create -ext sql -dir cmd/db/migrations $(filter-out $@,$(MAKECMDGOALS)) -verbose down

drop:
	@go run cmd/drop/main.go

seed:
	@go run db/seed/main.go
