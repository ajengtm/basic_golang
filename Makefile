VERSION := $(shell git describe --tags)
BUILD := $(shell git rev-parse --short HEAD)
PROJECTNAME := $(shell basename "$(PWD)")

# Go related variables.
GOFILES := $(wildcard *.go)
STIME := $(shell date +%s)

## install: Install missing dependencies. Runs `go get` internally. e.g; make install get=github.com/foo/bar
install: go-get

## build-public-http: start without docker
build:
	@echo "  >  Building Program..."
	GOPRIVATE=gitlab.warungpintar.co go build -ldflags="-s -w" -o bin/${PROJECTNAME} main.go; 
	@echo "Process took $$(($$(date +%s)-$(STIME))) seconds"

## start-auth-http: start without docker
start-http-auth: build
	@echo "  >  Starting Program..."
	go run main.go http-auth
	@echo "Process took $$(($$(date +%s)-$(STIME))) seconds"

## start-auth-http: start without docker
start-http-fetch: build
	@echo "  >  Starting Program..."
	go run main.go http-fetch
	@echo "Process took $$(($$(date +%s)-$(STIME))) seconds"
