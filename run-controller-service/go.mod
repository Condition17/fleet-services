module github.com/Condition17/fleet-services/run-controller-service

go 1.13

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace (
	github.com/Condition17/fleet-services/run-controller-service => ./.
	github.com/coreos/etcd => github.com/ozonru/etcd/v3 v3.3.0-rc.0-grpc1.30.0
)

require (
	cloud.google.com/go v0.74.0 // indirect
	github.com/Condition17/fleet-services/file-builder v0.0.0-20201213183648-f069e2830265
	github.com/Condition17/fleet-services/file-service v0.0.0-20201213183648-f069e2830265
	github.com/Condition17/fleet-services/lib v0.0.0-20201213183648-f069e2830265
	github.com/Condition17/fleet-services/resource-manager-service v0.0.0-20201213183648-f069e2830265
	github.com/Condition17/fleet-services/river-runner v0.0.0-20201213183648-f069e2830265
	github.com/Condition17/fleet-services/test-run-service v0.0.0-20201213183648-f069e2830265
	github.com/Condition17/fleet-services/user-service v0.0.0-20201213183648-f069e2830265 // indirect
	github.com/go-git/go-git/v5 v5.2.0 // indirect
	github.com/golang/protobuf v1.4.3
	github.com/google/go-cmp v0.5.4 // indirect
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/broker/googlepubsub/v2 v2.9.1
	github.com/nats-io/jwt v1.2.2 // indirect
	github.com/ozonru/etcd/v3 v3.3.0-rc.0-grpc1.30.0 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	golang.org/x/sys v0.0.0-20201211090839-8ad439b19e0f // indirect
	golang.org/x/text v0.3.4 // indirect
	golang.org/x/tools v0.0.0-20201211185031-d93e913c1a58 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20201211151036-40ec1c210f7a // indirect
	google.golang.org/grpc v1.34.0
	google.golang.org/grpc/examples v0.0.0-20201204164231-c7df457e12e0 // indirect
	google.golang.org/protobuf v1.25.0
)
