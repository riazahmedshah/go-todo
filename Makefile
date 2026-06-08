include .env
export $(shell sed 's/=.*//' .env)

.DEFAULT_GOAL := build

.PHONY: fmt vet build

fmt: 
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build -o main