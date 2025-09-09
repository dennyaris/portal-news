SHELL := /bin/bash

APP_NAME := news-portal-gorm
PKG := ./...
MAIN := ./cmd/api
PORT ?= 8080

.PHONY: bootstrap run test tidy build

bootstrap:
	go mod tidy

run:
	PORT=$(PORT) go run $(MAIN)

build:
	go build -o bin/$(APP_NAME) $(MAIN)

test:
	go test -v $(PKG)

tidy:
	go mod tidy
