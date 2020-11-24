module github.com/Condition17/fleet-services/run-controller-service

go 1.13

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace (
	github.com/Condition17/fleet-services/run-controller-service => ./.
	github.com/coreos/etcd => github.com/ozonru/etcd/v3 v3.3.0-rc.0-grpc1.30.0
)

require (
	github.com/Condition17/fleet-services/binary-builder v0.0.0-20201124203551-1694b744d0ab
	github.com/Condition17/fleet-services/file-service v0.0.0-20201124203551-1694b744d0ab
	github.com/Condition17/fleet-services/lib v0.0.0-20201124203551-1694b744d0ab
	github.com/Condition17/fleet-services/resource-manager-service v0.0.0-20201124203551-1694b744d0ab
	github.com/Condition17/fleet-services/test-run-service v0.0.0-20201124203551-1694b744d0ab
	github.com/Condition17/fleet-services/user-service v0.0.0-20201124203551-1694b744d0ab // indirect
	github.com/go-git/go-git/v5 v5.2.0 // indirect
	github.com/golang/protobuf v1.4.3
	github.com/google/go-cmp v0.5.4 // indirect
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/broker/googlepubsub/v2 v2.9.1
	github.com/micro/go-plugins/registry/kubernetes/v2 v2.9.1
	github.com/nats-io/jwt v1.2.0 // indirect
	github.com/ozonru/etcd/v3 v3.3.0-rc.0-grpc1.30.0 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	golang.org/x/crypto v0.0.0-20201124201722-c8d3bf9c5392 // indirect
	golang.org/x/text v0.3.4 // indirect
	golang.org/x/tools v0.0.0-20201124202034-299f270db459 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/grpc v1.33.2
	google.golang.org/protobuf v1.25.0
)
