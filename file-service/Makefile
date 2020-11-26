
GOPATH:=$(shell go env GOPATH)
MODIFY=Mproto/imports/api.proto=github.com/micro/go-micro/v2/api/proto

.PHONY: proto
proto:
    
	protoc --plugin=protoc-gen-go=${GOPATH}/bin/protoc-gen-go --plugin=protoc-gen-micro=${GOPATH}/bin/protoc-gen-micro --proto_path=. --micro_out=${MODIFY}:. --go_out=${MODIFY}:. proto/file-service/file-service.proto \
	&& protoc --proto_path=proto/file-service --go_out=proto/file-service/grpc --go_opt=paths=source_relative \
                                             		--go-grpc_out=proto/file-service/grpc --go-grpc_opt=paths=source_relative \
                                             		proto/file-service/file-service.proto

.PHONY: build
build: proto

	go build -o service *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t file-service:latest