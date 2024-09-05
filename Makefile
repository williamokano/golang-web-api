.DEFAULT_GOAL := help
.PHONY : help

help:
	@echo "Web API development. Available commands:"
	@echo
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		sort | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: new-migration
new-migration:  ## Create a new migration NAME=migration-name
	docker run -v $(PWD)/migrations:/migrations migrate/migrate create -ext sql -dir /migrations $(NAME)

.PHONY: docs
docs:  ## Generate the swagger updated documentation
	go generate ./...
