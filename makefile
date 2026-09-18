BINARY_NAME=nothing-postrenderer

DEFAULT_GOAL := build

.PHONY: build
build:
	go build -o build/$(BINARY_NAME) main.go

.PHONY: test
test:
	go test ./internal/postrenderer