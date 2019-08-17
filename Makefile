# Go parameters
GOCMD=go
GOPACKR=packr
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=privoxy-front
ROOT_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

all: test build


build: gen
		$(GOPACKR) build -o $(BINARY_NAME) -v

.PHONY: arm5
arm5: gen
	mkdir -p release
	GOOS=linux GOARCH=arm GOARM=5 $(GOPACKR) build -o release/$(BINARY_NAME)-v1.0.0-arm5

.PHONY: linux
linux: gen
	mkdir -p release
	GOOS=linux GOARCH=amd64 $(GOPACKR) build -o release/$(BINARY_NAME)-v1.0.0-linux-amd64

.PHONY: darwin
darwin: gen
	mkdir -p release
	GOOS=darwin GOARCH=amd64 $(GOPACKR) build -o release/$(BINARY_NAME)-v1.0.0-darwin-amd64

.PHONY: build
build: gen arm5 linux darwin


gen:
	go get -u github.com/rakyll/statik
	go generate

test:
		$(GOTEST) -v ./...
clean:
		$(GOCLEAN)
		rm -f release/*
run:
		$(GOPACKR) build -o $(BINARY_NAME) -v ./main.go
		./$(BINARY_NAME)
