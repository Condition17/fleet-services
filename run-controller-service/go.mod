module github.com/Condition17/fleet-services/run-controller-service

go 1.13

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace (
	github.com/Condition17/fleet-services/run-controller-service => ./.
	github.com/coreos/etcd => github.com/ozonru/etcd v3.3.20-grpc1.27-origmodule+incompatible
	google.golang.org/grpc => google.golang.org/grpc v1.27.0
)

require (
	cloud.google.com/go v0.71.0 // indirect
	github.com/Condition17/fleet-services/file-service v0.0.0-20201107223016-b8381b8f372a
	github.com/Condition17/fleet-services/lib v0.0.0-20201107223016-b8381b8f372a
	github.com/Condition17/fleet-services/test-run-service v0.0.0-20201107223016-b8381b8f372a
	github.com/Condition17/fleet-services/user-service v0.0.0-20201107223016-b8381b8f372a // indirect
	github.com/go-git/go-git/v5 v5.2.0 // indirect
	github.com/golang/protobuf v1.4.3
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/broker/googlepubsub/v2 v2.9.1
	github.com/micro/go-plugins/registry/kubernetes/v2 v2.9.1
	github.com/nats-io/jwt v1.2.0 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	golang.org/x/net v0.0.0-20201031054903-ff519b6c9102 // indirect
	golang.org/x/sys v0.0.0-20201107080550-4d91cf3a1aaf // indirect
	golang.org/x/text v0.3.4 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.25.0
)
