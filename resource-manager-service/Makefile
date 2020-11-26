
GOPATH:=$(shell go env GOPATH)
MODIFY=Mproto/imports/api.proto=github.com/micro/go-micro/v2/api/proto

.PHONY: proto
proto:
    
	protoc --plugin=protoc-gen-go=${GOPATH}/bin/protoc-gen-go --plugin=protoc-gen-micro=${GOPATH}/bin/protoc-gen-micro --proto_path=. --micro_out=${MODIFY}:. --go_out=${MODIFY}:. proto/resource-manager-service/resource-manager-service.proto \
	&& protoc --proto_path=proto/resource-manager-service --go_out=proto/resource-manager-service/grpc --go_opt=paths=source_relative \
                                                 		--go-grpc_out=proto/resource-manager-service/grpc --go-grpc_opt=paths=source_relative \
                                                 		proto/resource-manager-service/resource-manager-service.proto

.PHONY: build
build: proto

	go build -o service *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t resource-manager-service:latest
