

# msg := $(shell echo $${HOME})

.DEFAULT_GOAL := help

.PHONY: help
help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n\nTargets:\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-10s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

dev: ## run server in dev environment
	cp dev.env .env
	cp docker-compose.dev.yaml docker-compose.yaml
	docker-compose build --force-rm
	docker-compose up
prod: ## run server in prod environment
	cp docker-compose.prod.yaml docker-compose.yaml



