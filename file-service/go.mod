module github.com/Condition17/fleet-services/file-service

go 1.13

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace (
	file-service => ./.
	github.com/coreos/etcd => github.com/ozonru/etcd v3.3.20-grpc1.27-origmodule+incompatible
	google.golang.org/grpc => google.golang.org/grpc v1.27.0
)

require (
	file-service v0.0.0-00010101000000-000000000000
	github.com/Condition17/fleet-services/common v0.0.0-20200919205711-d2a60ae78150
	github.com/gofrs/uuid v3.2.0+incompatible
	github.com/golang/protobuf v1.4.2
	github.com/gomodule/redigo v1.8.2
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/registry/kubernetes/v2 v2.9.1
	github.com/tkanos/gonfig v0.0.0-20181112185242-896f3d81fadf
	google.golang.org/protobuf v1.25.0
)
