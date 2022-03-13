GOPATH:=$(shell go env GOPATH)
.PHONY: init
init:
	go install github.com/golang/protobuf/protoc-gen-go@latest
	go install github.com/micro/micro/v3/cmd/protoc-gen-micro@latest
	go install github.com/micro/micro/v3/cmd/protoc-gen-openapi@latest

.PHONY: proto
proto:
	docker run --rm -v $(shell pwd):$(shell pwd) -w $(shell pwd) zhugeqing/protoc \
--go_out=./proto/category --micro_out=./proto/category -I ./ ./proto/category/*.proto

.PHONY: build
build: proto

	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o category-service *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t category-service:latest