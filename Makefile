include .env
export


export PROJECT_ROOT=$(shell pwd)


env-up:
	@docker compose up -d agilewebapp-postgres

env-down:
	@docker compose down agilewebapp-postgres

env-cleanup:
	@read -p "Clean all environment volume files? [y/N]: " ans; \
	if [ "$$ans" = "y" ]; then \
		docker compose down agilewebapp-postgres port-forwarder && \
		rm -rf out/pgdata && \
		echo "Environment files cleaned up"; \
	else \
		echo "Cancelled"; \
	fi

env-port-forward:
	@docker compose up -d port-forwarder

env-port-close:
	@docker compose down port-forwarder

migrate-create:
	@if [ -z "$(seq)" ]; then \
		echo "missing seq parameter. Example: make migrate-create seq=init"; \
		exit 1; \
	fi; \
	docker compose run --rm agilewebapp-postgres-migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq "$(seq)"

migrate-up:
	@make migrate-action action=up

migrate-down:
	@make migrate-action action=down

migrate-action:
	@if [ -z "$(action)" ]; then \
		echo "missing action parameter. Example: make migrate-action action=up"; \
		exit 1; \
	fi; \
	docker compose run --rm agilewebapp-postgres-migrate \
		-path /migrations \
		-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@agilewebapp-postgres:5432/${POSTGRES_DB}?sslmode=disable \
		"$(action)"

agilewebapp-run:
	@export LOGGER_FOLDER=${PROJECT_ROOT}/out/logs && \
	export POSTGRES_HOST=localhost && \
	go run cmd/agilewebapp/main.go