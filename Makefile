.PHONY: usage test build

OK_COLOR=\033[32;01m
NO_COLOR=\033[0m

GO := go
GO_LINTER := golint
GOFLAGS ?=
ROOT_DIR := $(realpath .)

DOCKER_COMPOSE := docker-compose

LOCAL_VARIABLES ?= $(shell while read -r line; do printf "$$line" | sed 's/ /\\ /g' | awk '{print}'; done < $(ROOT_DIR)/.env)

PKGS = $(shell $(GO) list ./...)

## usage: show available actions
usage: Makefile
	@echo "to use make call:"
	@echo "make <action>"
	@echo ""
	@echo "list of available actions:"
	@sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/ /'

## build: build all
build: test
	@echo "$(OK_COLOR)==> Building binary (linux/amd64/lambda-api)...$(NO_COLOR)"
	@echo GOOS=linux GOARCH=amd64 $(GO) build -v -o bin/linux_amd64/lambda-api ./cmd/lambda-api
	@GOOS=linux GOARCH=amd64 $(GO) build -v $(BUILDFLAGS) -o bin/linux_amd64/lambda-api ./cmd/lambda-api

## up: start services
up: build
	@echo "$(OK_COLOR)==> Starting services...$(NO_COLOR)"
	$(DOCKER_COMPOSE) up

## down: stop services
down:
	@echo "$(OK_COLOR)==> Stopping services...$(NO_COLOR)"
	$(DOCKER_COMPOSE) down

## test: run unit tests
test:
	@echo "$(OK_COLOR)==> Running tests with envs[$(LOCAL_VARIABLES)]:$(NO_COLOR)"
	@$(LOCAL_VARIABLES) $(GO) test $(GOFLAGS) $(PKGS)
