GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GORUN=$(GOCMD) run
PROJECT=circa10a/witchonstephendrive.com
BINARY=witch
VERSION=0.1.0
GOBUILDFLAGS=-ldflags="-s -w"

build:
	$(GOBUILD) $(GOBUILDFLAGS) -o $(BINARY)

run:
	$(GORUN) .

compile:
	GOOS=linux GOARCH=amd64 go build $(GOBUILDFLAGS) -o bin/$(BINARY)-linux-amd64
	GOOS=linux GOARCH=arm go build $(GOBUILDFLAGS) -o bin/$(BINARY)-linux-arm
	GOOS=linux GOARCH=arm64 go build $(GOBUILDFLAGS) -o bin/$(BINARY)-linux-arm64
	GOOS=darwin GOARCH=amd64 go build $(GOBUILDFLAGS) -o bin/$(BINARY)-darwin-amd64

docs:
	# Swagger
	# https://github.com/swaggo/echo-swagger
	swag init -o ./api

lint:
	golangci-lint run -v