SHELL := /bin/bash

APP_NAME := portal-news
PKG := ./...
MAIN := ./cmd/api
PORT ?= 9898

.PHONY: bootstrap run test tidy build

bootstrap:
	go mod tidy

run:
	PORT=$(PORT) go run $(MAIN)

build:
	go build -o bin/$(APP_NAME) $(MAIN)

test:
	go test -v $(PKG)
