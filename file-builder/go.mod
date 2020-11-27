module github.com/Condition17/fleet-services/file-builder

go 1.14

replace github.com/Condition17/fleet-services/file-builder => ./.

require (
	cloud.google.com/go/pubsub v1.8.3
	cloud.google.com/go/storage v1.12.0
	github.com/Condition17/fleet-services/file-service v0.0.0-20201127145314-6f7078d558b3
	github.com/Condition17/fleet-services/lib v0.0.0-20201127145314-6f7078d558b3
	github.com/Condition17/fleet-services/resource-manager-service v0.0.0-20201127145314-6f7078d558b3
	github.com/Condition17/fleet-services/run-controller-service v0.0.0-20201127145314-6f7078d558b3
	github.com/golang/protobuf v1.4.3
	github.com/google/go-cmp v0.5.4 // indirect
	golang.org/x/sys v0.0.0-20201126233918-771906719818 // indirect
	golang.org/x/tools v0.0.0-20201125231158-b5590deeca9b // indirect
	google.golang.org/grpc v1.33.2
	google.golang.org/protobuf v1.25.0
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
