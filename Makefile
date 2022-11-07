

# msg := $(shell echo $${HOME})

.DEFAULT_GOAL := help

.PHONY: help
help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n\nTargets:\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-10s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

dev: ## run server in dev environment
	docker-compose -f docker-compose.dev.yaml --env-file dev.env build --force-rm
	docker-compose -f docker-compose.dev.yaml --env-file dev.env up && \
		docker-compose -f docker-compose.dev.yaml --env-file dev.env down
prod: ## run server in prod environment


.PHONY: compose_dev_test
compose_dev_test: compose_dev ## commands for compose
	go test -v ./...

.PHONY:compose_dev_test_and_run
compose_dev_test_and_run: compose_dev_test ## commands for compose
	air

.PHONY: compose_dev
compose_dev:
	migrate -source file:///mashi/src/delivery/infra/db/postgres/migrations/ -database postgres://mashi:123456789@db/mashi?sslmode=disable up

.PHONY: migrate
migrate:
	migrate create -ext sql -dir ./src/delivery/infra/db/postgres/migrations/ $(name)
