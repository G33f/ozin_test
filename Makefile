DC=docker-compose

DEPLOY_PATH=./deploy

MOCKS_DESTINATION=mocks

.PHONY: all
all: compose-up-postgresql

.PHONY: compose-up-in-memory
compose-up-in-memory:
	 STORAGE_TYPE=InMemory $(DC) -f $(DEPLOY_PATH)/docker-compose.yaml up

.PHONY: compose-up-postgresql
compose-up-postgresql:
	STORAGE_TYPE=PostgreSQL $(DC) -f $(DEPLOY_PATH)/docker-compose.yaml up

.PHONY: compose-start
compose-start:
	$(DC) -f $(DEPLOY_PATH)/docker-compose.yaml start

.PHONY: compose-stop
compose-stop:
	$(DC) -f $(DEPLOY_PATH)/docker-compose.yaml stop

.PHONY: compose-down
compose-down:
	$(DC) -f $(DEPLOY_PATH)/docker-compose.yaml down
	docker rmi deploy-short_url

.PHONY: proto-gen
proto-gen:
	protoc --go_out=. --go-grpc_out=require_unimplemented_servers=false:. proto/shortener.proto

.PHONY: mocks
mocks:
	rm -rf $(MOCKS_DESTINATION)
	mockgen -source=internal/repo/interface.go -destination=mocks/mock_repo/mock_repo.go