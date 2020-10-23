module github.com/Condition17/fleet-services/file-service

go 1.13

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace (
	github.com/Condition17/fleet-services/file-service => ./.
	github.com/coreos/etcd => github.com/ozonru/etcd v3.3.20-grpc1.27-origmodule+incompatible
	google.golang.org/grpc => google.golang.org/grpc v1.27.0
)

require (
	cloud.google.com/go v0.70.0 // indirect
	cloud.google.com/go/pubsub v1.8.1 // indirect
	github.com/Condition17/fleet-services/common v0.0.0-20201021191733-63365d839dc7
	github.com/Condition17/fleet-services/user-service v0.0.0-20201021191733-63365d839dc7 // indirect
	github.com/gofrs/uuid v3.3.0+incompatible
	github.com/golang/protobuf v1.4.3
	github.com/gomodule/redigo v1.8.2
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/broker/googlepubsub/v2 v2.9.1
	github.com/micro/go-plugins/registry/kubernetes/v2 v2.9.1
	github.com/nats-io/jwt v1.1.0 // indirect
	github.com/tkanos/gonfig v0.0.0-20181112185242-896f3d81fadf
	golang.org/x/exp v0.0.0-20201008143054-e3b2a7f2fdc7 // indirect
	golang.org/x/net v0.0.0-20201021035429-f5854403a974 // indirect
	golang.org/x/sync v0.0.0-20201020160332-67f06af15bc9 // indirect
	golang.org/x/sys v0.0.0-20201020230747-6e5568b54d1a // indirect
	golang.org/x/tools v0.0.0-20201021171030-d105bfabbdbe // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20201021134325-0d71844de594 // indirect
	google.golang.org/grpc v1.33.1 // indirect
	google.golang.org/protobuf v1.25.0
	honnef.co/go/tools v0.0.1-2020.1.6 // indirect
)
