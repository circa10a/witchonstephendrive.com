PROJECT=circa10a/witchonstephendrive.com
BINARY=witch
VERSION=0.1.0
GOBUILDFLAGS=-ldflags="-s -w"
DOCKERBUILDDIR=./build
WTICHDOCKERFILE=$(DOCKERBUILDDIR)/Dockerfile
CADDYDOCKERFILE=$(DOCKERBUILDDIR)/Dockerfile.caddy
ASSISTANTRELAYDOCKERFILE=$(DOCKERBUILDDIR)/Dockerfile.assistant-relay
PROMETHEUSDOCKERFILE=$(DOCKERBUILDDIR)/Dockerfile.prometheus
GRAFANADOCKERFILE=$(DOCKERBUILDDIR)/Dockerfile.grafana
DOCKERBUILD=docker build -t $(PROJECT)
DOCKERBUILDX=docker buildx build
DOCKERBUILDARM64=$(DOCKERBUILDX) --platform linux/arm64 -t $(PROJECT)
DOCKERBUILDARMv7=$(DOCKERBUILDX) --platform linux/arm/v7 -t $(PROJECT)
DOCKERPUSH=docker push $(PROJECT)

build:
	go build $(GOBUILDFLAGS) -o $(BINARY)

build-docker:
	$(DOCKERBUILD):witch -f $(WTICHDOCKERFILE) .
	$(DOCKERBUILD):caddy -f $(CADDYDOCKERFILE) .
	$(DOCKERBUILD):prometheus -f $(PROMETHEUSDOCKERFILE) .
	$(DOCKERBUILD):grafana -f  $(GRAFANADOCKERFILE) .
	$(DOCKERBUILD):assistant-relay -f $(ASSISTANTRELAYDOCKERFILE) .

build-docker-arm64:
	$(DOCKERBUILDARM64):witch -f $(WTICHDOCKERFILE) .
	$(DOCKERBUILDARM64):caddy -f $(CADDYDOCKERFILE) .
	$(DOCKERBUILDARM64):prometheus -f $(PROMETHEUSDOCKERFILE) .
	$(DOCKERBUILDARM64):grafana -f  $(GRAFANADOCKERFILE) .
	$(DOCKERBUILDARM64):assistant-relay -f $(ASSISTANTRELAYDOCKERFILE) .

build-docker-armv7:
	$(DOCKERBUILDARM64):witch -f $(WTICHDOCKERFILE) .
	$(DOCKERBUILDARM64):caddy -f $(CADDYDOCKERFILE).
	$(DOCKERBUILDARM64):prometheus -f $(PROMETHEUSDOCKERFILE) .
	$(DOCKERBUILDARM64):grafana -f $(GRAFANADOCKERFILE) .
	$(DOCKERBUILDARM64):assistant-relay -f $(ASSISTANTRELAYDOCKERFILE) .

push-docker:
	$(DOCKERPUSH):witch
	$(DOCKERPUSH):caddy
	$(DOCKERPUSH):prometheus
	$(DOCKERPUSH):grafana
	$(DOCKERPUSH):assistant-relay

run:
	go run .

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

js-lint:
	eslint web/js/*.js

test:
	go test -v ./...
