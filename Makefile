PROJECT=circa10a/witchonstephendrive.com
BINARY=witch
VERSION=0.1.0
GOBUILDFLAGS=-ldflags="-s -w"
GOGENERATE=go generate ./...
DOCKERBUILDDIR=./build
WITCHDOCKERFILE=$(DOCKERBUILDDIR)/Dockerfile
CADDYDOCKERFILE=$(DOCKERBUILDDIR)/Dockerfile.caddy
ASSISTANTRELAYDOCKERFILE=$(DOCKERBUILDDIR)/Dockerfile.assistant-relay
PROMETHEUSDOCKERFILE=$(DOCKERBUILDDIR)/Dockerfile.prometheus
GRAFANADOCKERFILE=$(DOCKERBUILDDIR)/Dockerfile.grafana
DOCKERBUILD=docker build -t $(PROJECT)
DOCKERBUILDX=docker buildx build
DOCKERBUILDARM64=$(DOCKERBUILDX) --platform linux/arm64 -t $(PROJECT)
DOCKERBUILDARMv7=$(DOCKERBUILDX) --platform linux/arm/v7 -t $(PROJECT)
DOCKERPUSH=docker push $(PROJECT)

.PHONY: build

build:
	$(GOGENERATE)
	go build $(GOBUILDFLAGS) -o $(BINARY)

build-docker:
	$(DOCKERBUILD):$(BINARY) -f $(WITCHDOCKERFILE) .
	$(DOCKERBUILD):caddy -f $(CADDYDOCKERFILE) .
	$(DOCKERBUILD):prometheus -f $(PROMETHEUSDOCKERFILE) .
	$(DOCKERBUILD):grafana -f  $(GRAFANADOCKERFILE) .
	$(DOCKERBUILD):assistant-relay -f $(ASSISTANTRELAYDOCKERFILE) .

build-docker-arm64:
	$(DOCKERBUILDARM64):$(BINARY) -f $(WITCHDOCKERFILE) .
	$(DOCKERBUILDARM64):caddy -f $(CADDYDOCKERFILE) .
	$(DOCKERBUILDARM64):prometheus -f $(PROMETHEUSDOCKERFILE) .
	$(DOCKERBUILDARM64):grafana -f  $(GRAFANADOCKERFILE) .
	$(DOCKERBUILDARM64):assistant-relay -f $(ASSISTANTRELAYDOCKERFILE) .

build-docker-armv7:
	$(DOCKERBUILDARMv7):$(BINARY) -f $(WITCHDOCKERFILE) .
	$(DOCKERBUILDARMv7):caddy -f $(CADDYDOCKERFILE) .
	$(DOCKERBUILDARMv7):prometheus -f $(PROMETHEUSDOCKERFILE) .
	$(DOCKERBUILDARMv7):grafana -f $(GRAFANADOCKERFILE) .
	$(DOCKERBUILDARMv7):assistant-relay -f $(ASSISTANTRELAYDOCKERFILE) .

push-docker:
	$(DOCKERPUSH):$(BINARY)
	$(DOCKERPUSH):caddy
	$(DOCKERPUSH):prometheus
	$(DOCKERPUSH):grafana
	$(DOCKERPUSH):assistant-relay

run:
	$(GOGENERATE)
	go run .

compile:
	GOOS=linux GOARCH=amd64 $(GOGENERATE) && go build $(GOBUILDFLAGS) -o bin/$(BINARY)-linux-amd64
	GOOS=linux GOARCH=arm $(GOGENERATE) && go build $(GOBUILDFLAGS) -o bin/$(BINARY)-linux-arm
	GOOS=linux GOARCH=arm64$(GOGENERATE) && go build $(GOBUILDFLAGS) -o bin/$(BINARY)-linux-arm64
	GOOS=darwin GOARCH=amd64 $(GOGENERATE) && go build $(GOBUILDFLAGS) -o bin/$(BINARY)-darwin-amd64

docs:
	# Swagger
	# https://github.com/swaggo/echo-swagger
	swag init -o ./api

lint:
	$(GOGENERATE)
	golangci-lint run -v

js-lint:
	eslint web/js/*.js

test:
	$(GOGENERATE)
	go test -v ./...
