module github.com/Condition17/fleet-services/file-builder

go 1.14

replace github.com/Condition17/fleet-services/file-builder => ./.

require (
	cloud.google.com/go/pubsub v1.9.0
	cloud.google.com/go/storage v1.12.0
	github.com/Condition17/fleet-services/binary-builder v0.0.0-20201124203551-1694b744d0ab
	github.com/Condition17/fleet-services/file-service v0.0.0-20201207230256-fd5e120ba798
	github.com/Condition17/fleet-services/lib v0.0.0-20201207230256-fd5e120ba798
	github.com/Condition17/fleet-services/resource-manager-service v0.0.0-20201207230256-fd5e120ba798
	github.com/Condition17/fleet-services/run-controller-service v0.0.0-20201207230256-fd5e120ba798
	github.com/golang/protobuf v1.4.3
	github.com/google/go-cmp v0.5.4 // indirect
	golang.org/x/net v0.0.0-20201207224615-747e23833adb // indirect
	golang.org/x/oauth2 v0.0.0-20201207163604-931764155e3f // indirect
	golang.org/x/sys v0.0.0-20201207223542-d4d67f95c62d // indirect
	google.golang.org/genproto v0.0.0-20201207150747-9ee31aac76e7 // indirect
	google.golang.org/grpc v1.34.0
	google.golang.org/protobuf v1.25.0
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
