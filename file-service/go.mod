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
	cloud.google.com/go v0.71.0 // indirect
	cloud.google.com/go/pubsub v1.8.2 // indirect
	github.com/Condition17/fleet-services/lib v0.0.0-20201103210223-0cf24d84e514
	github.com/Condition17/fleet-services/run-controller-service v0.0.0-20201103210223-0cf24d84e514 // indirect
	github.com/Condition17/fleet-services/user-service v0.0.0-20201103210223-0cf24d84e514 // indirect
	github.com/gofrs/uuid v3.3.0+incompatible
	github.com/golang/protobuf v1.4.3
	github.com/gomodule/redigo v1.8.2
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/broker/googlepubsub/v2 v2.9.1
	github.com/micro/go-plugins/registry/kubernetes/v2 v2.9.1
	github.com/tkanos/gonfig v0.0.0-20181112185242-896f3d81fadf
	golang.org/x/tools v0.0.0-20201103190053-ac612affd56b // indirect
	google.golang.org/genproto v0.0.0-20201103154000-415bd0cd5df6 // indirect
	google.golang.org/protobuf v1.25.0
)
