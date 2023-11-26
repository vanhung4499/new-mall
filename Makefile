# amd64 arm64
ARCH = amd64

# linux darwin windows
OS = darwin

DIR := $(shell pwd)
OUTPUT = main

CONTAINER_NAME = mall-api
IMAGE_NAME = mall-api:latest


.PHONY: run
run: build
	./$(OUTPUT)

.PHONY: build
build:
	echo "build project to ./$(OUTPUT)"
	go build -o ./$(OUTPUT) ./cmd

.PHONY: compose-up
compose-up:
	docker-compose up -d
	echo "docker-compose start success"

.PHONY: compose-down
compose-down:
	docker-compose down
	echo "docker-compose stop success"

default: run