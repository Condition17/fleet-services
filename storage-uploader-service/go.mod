module github.com/Condition17/fleet-services/storage-uploader-service

go 1.13

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace (
	github.com/Condition17/fleet-services/storage-uploader-service => ./.
	github.com/coreos/etcd => github.com/ozonru/etcd v3.3.20-grpc1.27-origmodule+incompatible
	google.golang.org/grpc => google.golang.org/grpc v1.27.0
)

require (
	cloud.google.com/go/storage v1.10.0
	github.com/Condition17/fleet-services/file-service v0.0.0-20200928220040-6ae6c4f52080
	github.com/Condition17/fleet-services/lib v0.0.0-20201031183420-a99aa12d16c6
	github.com/golang/protobuf v1.4.3
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/broker/googlepubsub/v2 v2.9.1
	github.com/micro/go-plugins/registry/kubernetes/v2 v2.9.1
	google.golang.org/protobuf v1.25.0
)
