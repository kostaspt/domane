.PHONY: all
all:
	docker-compose build --parallel
	docker-compose -f ./docker-compose.yml up -d --force-recreate --remove-orphans

.PHONY: client
client:
	docker-compose -f ./docker-compose.yml -f ./docker-compose-local.yml up -d --build --force-recreate --remove-orphans client

.PHONY: api
api:
	docker-compose -f ./docker-compose.yml -f ./docker-compose-local.yml up -d --build --force-recreate --remove-orphans api