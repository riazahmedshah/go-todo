include .env
export $(shell sed 's/=.*//' .env)

.DEFAULT_GOAL := build

.PHONY: fmt vet build

fmt: 
	go fmt ./...

vet: fmt
	go vet ./...

migrate: vet
	tern migrate --migrations migrations --conn-string $(DATABASE_URL)

build: migrate
	go build -o main