.PHONY: help, dev, prod, test, migrate, migrate-up, migrate-down

# msg := $(shell echo $${HOME})

.DEFAULT_GOAL := help

help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n\nTargets:\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-10s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

dev: docker_down docker_build docker_up ## run server in dev environment
	
prod: ## run server in prod environment
	echo "prod"

migrate: ## create sql migration files
	migrate create -ext sql -dir ./src/delivery/infra/db/postgres/migrations/ $(name)
migrate-up:
	docker exec mashi_backend_1 migrate -source file:///mashi/src/delivery/infra/db/postgres/migrations/ -database postgres://mashi:123456789@db/mashi?sslmode=disable up
migrate-down:
	docker exec mashi_backend_1 migrate -source file:///mashi/src/delivery/infra/db/postgres/migrations/ -database postgres://mashi:123456789@db/mashi?sslmode=disable down 1
test:
	docker exec mashi_backend_1 go test -v ./...


# DOCKER-COMPOSE
compose_backend_run:
	air
in:
	docker exec -it mashi_backend_1 bash

# DOCKER
docker_build:
	docker-compose -f docker/docker-compose.dev.yaml --env-file dev.env --compatibility build
docker_up:
	docker-compose -f docker/docker-compose.dev.yaml --env-file dev.env --compatibility up \
	&& docker-compose -f docker/docker-compose.dev.yaml --env-file dev.env --compatibility down
docker_down:
	docker-compose -f docker/docker-compose.dev.yaml --env-file dev.env --compatibility down