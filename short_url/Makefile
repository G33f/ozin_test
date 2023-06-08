.PHONY: all
all: shortener

.PHONY: proto-gen
proto-gen:
	protoc --go_out=. --go-grpc_out=require_unimplemented_servers=false:. proto/shortener.proto

.PHONY: shortener
shortener: proto-gen
	go run cmd/api/main.go

.PHONY: shortener-postgreSQL
shortener-postgreSQL: proto-gen
	go run cmd/api/main.go -storage_type=PostgreSQL

.PHONY: shortener-in-memory
shortener-in-memory: proto-gen
	go run cmd/api/main.go -storage_type=InMemory
