.PHONY: build
build: 
	go build cmd/apiserver/main.go

run:
	go run cmd/apiserver/main.go

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build