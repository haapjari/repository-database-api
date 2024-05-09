MAIN_MODULE             := cmd/main.go
IMAGE_VERSION           := latest
SERVICE_NAME 			?= repository-database-api
OUTPUT_PATH 			?= bin/repository-database-api

##  
## Commands
##

fmt:
	gofmt -w ./...
	goimports -w ./...

run:
	go run cmd/main.go
.PHONY: run

dev:
	go run -race cmd/main.go
.PHONY: dev

test:
	go test -v -count=1 ./...
.PHONY: test

docker-build:
	docker build -t ${SERVICE_NAME}:${IMAGE_VERSION} .
.PHONY: docker-build

compile:
	go build -o ${OUTPUT_PATH} ${MAIN_MODULE}
.PHONY: compile
