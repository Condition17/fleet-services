GOPATH:=$(shell go env GOPATH)

compile_proto:
	protoc --proto_path=proto/binary-builder --go_out=proto/binary-builder --go_opt=paths=source_relative \
		--go-grpc_out=proto/binary-builder --go-grpc_opt=paths=source_relative \
		proto/binary-builder/binary-builder.proto

build:
	GOOS=linux GOARCH=amd64 go build -o service -v . 
