#!/usr/bin/make

GOCMD=$(shell which go)
GOMOD=$(shell which go) mod
GOLINT=$(shell which golint)
GODOC=$(shell which doc)
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOLIST=$(GOCMD) list
GOVET=$(GOCMD) vet
GORUN=$(GOCMD) run

compile:
	go mod tidy
	go build -o ./urldisp ./cmd/url-dispatcher
	go build -o ./urlapi ./cmd/api

run-api:
	go run cmd/api/main.go
run-disp:
	go run cmd/url-dispatcher/main.go

help:
	@echo 'Usage: make <OPTIONS> ... <TARGETS>'
	@echo ''
	@echo 'Available targets are:'
	@echo ''
	@echo '    compile                 	Build executable file.'
	@echo '    run-api                	Start api without compile.'
	@echo '    run-disp               	Start dispatcher without compile.'
	@echo '    help                    	Show this help screen.'
	@echo '    unit         			Run unit tests.'
	@echo '    test-integration        	Run integration tests.'
	@echo '    local-up                	Run service by docker compose'
	@echo '    local-down              	Stop service by docker compose'
	@echo '    local-restart           	Restart service by docker compose'
	@echo '    gen-mocks               	generate mocks for db controller'
	@echo ''
	@echo 'Targets run by default are: fmt deps vet lint build test-unit.'
	@echo ''

unit:
	$(GOTEST) ./...

test-integration:
	$(GOTEST) ./test -v -count=1 -tags 'integration' -timeout 20m

local-up:
	./deployments/docker-compose up -d

local-down:
	./deployments/docker-compose down

local-restart: | local-down local-up

gen-mocks:
	mockgen -destination=internal/mocks/mock_db.go -package=mocks url-healthcheck/pkg/db DB
