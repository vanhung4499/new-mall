# amd64 arm64
ARCH = amd64

# linux darwin windows
OS = darwin

DIR := $(shell pwd)
OUTPUT = main

CONTAINER_NAME = new_mall_server
IMAGE_NAME = new_mall:3.0

GO = go
GO_BUILD = $(GO) build
GO_BUILD_FLAGS = -v
GO_BUILD_LDFLAGS = -X main.version=$(VERSION)

.PHONY: tools
tools:
	cd $(AGENT_SOURCE_PATH) && make deps
	cd $(AGENT_SOURCE_PATH) && \
	GOOS=$(OS) GOARCH=$(ARCH) $(GO_BUILD) $(GO_BUILD_FLAGS) -ldflags "$(GO_BUILD_LDFLAGS)" -o $(TOOLS_PATH)/$(BINARY)-$(VERSION)-$(OS)-$(ARCH) ./cmd

.PHONY: run
test:
	@make build
	@./$(OUTPUT)

.PHONY: build
build:
	@echo "build project to ./$(OUTPUT)"
	$(GO_BUILD) \
	-toolexec="$(AGENT_PATH) -config $(AGENT_CONFIG)" \
	-a -o ./$(OUTPUT) ./cmd

.PHONY: env-up
env-up:
	docker-compose up -d
	@echo "env start success"

.PHONY: env-down
env-down:
	docker-compose down
	@echo "env stop success"


.PHONY: docker-up
docker-up:
	docker build \
	-t $(IMAGE_NAME) \
	-f ./Dockerfile \
	./
	docker run \
	-it \
	--name $(CONTAINER_NAME) \
	--network host \
	-d $(IMAGE_NAME)
	@echo "container run success at localhost:5001"

.PHONY: docker-down
docker-down:
	docker stop $(CONTAINER_NAME)
	docker rm $(CONTAINER_NAME)
	docker rmi $(IMAGE_NAME)
	@echo "container stop && rm success"

default: run