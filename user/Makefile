
GOPATH:=$(shell go env GOPATH)
.PHONY: init
init:
	go install github.com/golang/protobuf/protoc-gen-go@latest
	go install github.com/micro/micro/v3/cmd/protoc-gen-micro@latest
	go install github.com/micro/micro/v3/cmd/protoc-gen-openapi@latest

.PHONY: proto
proto: # makefile中需要使用（shell pwd）才能识别
	docker run --rm -v $(shell pwd):$(shell pwd) -w $(shell pwd) zhugeqing/protoc \
--go_out=./ --micro_out=./ -I ./ ./proto/user/*.proto

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o user *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: dockerBuild
dockerBuild:
	docker image  build -t user:latest .
