GO_FLAGS   ?=
NAME       := kuevs
OUTPUT_BIN ?= bin/${NAME}
PACKAGE    := github.com/openqt/$(NAME)
GIT_REV    ?= $(shell git rev-parse --short HEAD)
SOURCE_DATE_EPOCH ?= $(shell date +%s)
DATE       ?= $(shell date -u -d @${SOURCE_DATE_EPOCH} +"%Y-%m-%dT%H:%M:%SZ")
VERSION    ?= v0.3
IMG_NAME   := openqt/$(NAME)
IMAGE      := ${IMG_NAME}:${VERSION}

default: build

test:   ## Run all tests
	@go clean --testcache && go test ./...

cover:  ## Run test coverage suite
	@go test ./... --coverprofile=cov.out
	@go tool cover --html=cov.out

build:  ## Builds the CLI
	@CGO_ENABLED=1 go build ${GO_FLAGS} \
	-ldflags "-w -s -X main.version=${VERSION} -X main.commit=${GIT_REV} -X main.date=${DATE}" \
	-v -o ${OUTPUT_BIN} main.go

img:    ## Build Docker Image
	@docker build --rm -t ${IMAGE} .

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":[^:]*?## "}; {printf "\033[38;5;69m%-30s\033[38;5;38m %s\033[0m\n", $$1, $$2}'
