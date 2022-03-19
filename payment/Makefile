

.PHONY: proto
proto:
	sudo docker run --rm  -v $(shell pwd):$(shell pwd) -w $(shell pwd) cap1573/protoc -I ./ --micro_out=./ --go_out=./ ./proto/payment/*.proto

.PHONY: build
build: proto

	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o payment-service *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t payment-service:latest
