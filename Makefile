.PHONY: serve-local
serve-local:
	docker-compose build --parallel
	docker-compose up --force-recreate --remove-orphans