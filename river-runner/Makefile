.PHONY: proto
proto:
	 protoc --proto_path=proto/river-runner --go_out=proto/river-runner --go_opt=paths=source_relative \
	 	--go-grpc_out=proto/river-runner --go-grpc_opt=paths=source_relative \
	 	proto/river-runner/river-runner.proto

build:
	go build -o service -v .

build_linux:
	GOOS=linux GOARCH=amd64	go build -o service -v .