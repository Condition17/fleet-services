module github.com/Condition17/fleet-services/binary-builder

go 1.14

replace github.com/Condition17/fleet-services/binary-builder => ./.

require (
	github.com/Condition17/fleet-services/file-service v0.0.0-20201126204404-b8ea3fa6a45c
	github.com/Condition17/fleet-services/lib v0.0.0-20201126204404-b8ea3fa6a45c
	github.com/Condition17/fleet-services/resource-manager-service v0.0.0-20201126204404-b8ea3fa6a45c
	github.com/golang/protobuf v1.4.3
	github.com/gomodule/redigo v1.8.2
	golang.org/x/sys v0.0.0-20201126144705-a4b67b81d3d2 // indirect
	google.golang.org/grpc v1.33.2
	google.golang.org/protobuf v1.25.0
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
