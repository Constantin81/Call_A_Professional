.PHONY: build
build: 
	go build cmd/apiserver/main.go

run:
	go run cmd/apiserver/main.go

.DEFAULT_GOAL := build