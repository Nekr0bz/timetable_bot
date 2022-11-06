args = `arg="$(filter-out $@,$(MAKECMDGOALS))" && echo $${arg:-${1}}`


COMPOSE_CMD=docker-compose
COMPOSE_DEV = $(COMPOSE_CMD) --profile dev
COMPOSE_APP = $(COMPOSE_CMD) --profile app
COMPOSE_APP_CMD = $(COMPOSE_APP) run --rm
COMPOSE_TOOLS_CMD = $(COMPOSE_CMD) --profile tools run --rm

# COLORS
GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
RESET  := $(shell tput -Txterm sgr0)


TARGET_MAX_CHAR_NUM=20
## Show help
help:
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
		helpMessage = match(lastLine, /^## (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
			printf "  ${YELLOW}%-$(TARGET_MAX_CHAR_NUM)s${RESET} ${GREEN}%s${RESET}\n", helpCommand, helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)


## Build the project
build:
	$(COMPOSE_DEV) build

## Run the database
db:
	$(COMPOSE_CMD) up db

## Run the development environment
dev:
	$(COMPOSE_DEV) up

## Run the application environment
app:
	$(COMPOSE_APP) up

## Run the bot
bot:
	$(COMPOSE_APP_CMD) bot

## Run the scheduler
scheduler:
	$(COMPOSE_APP_CMD) scheduler

## Create a new migration
create_migration:
	$(COMPOSE_TOOLS_CMD) create-migration $(call args)

## Run migrations
migrate:
	$(COMPOSE_TOOLS_CMD) migrate

