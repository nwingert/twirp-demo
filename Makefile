# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

# Files
SERVER_NAME=server
CLIENT_NAME=client

GOPACKAGES = $(shell go list ./...  | grep -v /vendor/)
PROTOS = $(shell find ./rpc -name '*.proto')

default: build

workdir:
	mkdir -p workdir

protobuf:
	protoc --proto_path=$(GOPATH)/src:. --twirp_out=. --go_out=. $(PROTOS)

build: protobuf server client

server:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o workdir/$(SERVER_NAME) ./cmd/server

client:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o workdir/$(CLIENT_NAME) ./cmd/client

test: test-all

test-all:
	@go test -v $(GOPACKAGES)