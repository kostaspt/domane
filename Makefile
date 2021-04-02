.PHONY: all
all:
	docker-compose build --parallel
	docker-compose -f ./docker-compose.yml up -d --force-recreate --remove-orphans

.PHONY: api
api:
	docker-compose -f ./docker-compose.yml -f ./docker-compose-local.yml up -d --force-recreate --remove-orphans api