DC=docker-compose

#SHORT_URL_PATH=./short_url

DEPLOY_PATH=./deploy

.PHONY: all
all: compose-up

.PHONY: compose-up
compose-up:
	$(DC) -f $(DEPLOY_PATH)/docker-compose.yaml up

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
