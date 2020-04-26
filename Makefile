.PHONY: help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'


.PHONY: test-integration
test-integration: build-docker ## Run integration tests
	docker-compose -f docker-compose.yaml -f docker-compose.test.yaml run web

.PHONY: build-docker
build-docker: ## Build a docker
	docker-compose -f docker-compose.yaml -f docker-compose.test.yaml build
