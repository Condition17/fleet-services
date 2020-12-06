GOPATH:=$(shell go env GOPATH)

.PHONY: proto
proto:
	 protoc --proto_path=proto/file-builder --go_out=proto/file-builder --go_opt=paths=source_relative \
	 	--go-grpc_out=proto/file-builder --go-grpc_opt=paths=source_relative \
	 	proto/file-builder/file-builder.proto

	

build_linux:
	GOOS=linux GOARCH=amd64 go build -o server -v .
