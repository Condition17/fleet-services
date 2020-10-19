
GOPATH:=$(shell go env GOPATH)
MODIFY=Mproto/imports/api.proto=github.com/micro/go-micro/v2/api/proto

.PHONY: proto
proto:
    
	protoc --plugin=protoc-gen-go=${GOPATH}/bin/protoc-gen-go --plugin=protoc-gen-micro=${GOPATH}/bin/protoc-gen-micro --proto_path=. --micro_out=${MODIFY}:. --go_out=${MODIFY}:. proto/run-controller-service/run-controller-service.proto  
    

.PHONY: build
build: proto

	go build -o run-controller-service *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t run-controller-service:latest
