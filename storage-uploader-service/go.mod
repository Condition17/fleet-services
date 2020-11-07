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
	cloud.google.com/go v0.71.0 // indirect
	cloud.google.com/go/storage v1.12.0
	github.com/Condition17/fleet-services/file-service v0.0.0-20201107200319-06e7c1fec256
	github.com/Condition17/fleet-services/lib v0.0.0-20201107200319-06e7c1fec256
	github.com/Condition17/fleet-services/run-controller-service v0.0.0-20201107200319-06e7c1fec256 // indirect
	github.com/Condition17/fleet-services/user-service v0.0.0-20201107200319-06e7c1fec256 // indirect
	github.com/Microsoft/go-winio v0.4.15 // indirect
	github.com/golang/protobuf v1.4.3
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/broker/googlepubsub/v2 v2.9.1
	github.com/micro/go-plugins/registry/kubernetes/v2 v2.9.1
	golang.org/x/sys v0.0.0-20201107080550-4d91cf3a1aaf // indirect
	golang.org/x/tools v0.0.0-20201105220310-78b158585360 // indirect
	google.golang.org/protobuf v1.25.0
)
