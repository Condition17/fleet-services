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
	cloud.google.com/go v0.70.0 // indirect
	github.com/Condition17/fleet-services/file-service v0.0.0-20201025124251-378b9f3827a1
	github.com/Condition17/fleet-services/lib v0.0.0-20201025124251-378b9f3827a1
	github.com/coreos/etcd v3.3.25+incompatible // indirect
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/go-delve/delve v1.5.0
	github.com/go-git/go-git/v5 v5.2.0 // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/protobuf v1.4.3
	github.com/google/uuid v1.1.2 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/imdario/mergo v0.3.11 // indirect
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/broker/googlepubsub/v2 v2.9.1
	github.com/micro/go-plugins/registry/kubernetes/v2 v2.9.1
	github.com/miekg/dns v1.1.35 // indirect
	github.com/nats-io/jwt v1.1.0 // indirect
	github.com/nats-io/nats.go v1.10.0 // indirect
	github.com/nats-io/nkeys v0.2.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.16.0 // indirect
	golang.org/x/crypto v0.0.0-20201016220609-9e8e0b390897 // indirect
	golang.org/x/exp v0.0.0-20201008143054-e3b2a7f2fdc7 // indirect
	golang.org/x/net v0.0.0-20201024042810-be3efd7ff127 // indirect
	golang.org/x/sync v0.0.0-20201020160332-67f06af15bc9 // indirect
	golang.org/x/sys v0.0.0-20201024232916-9f70ab9862d5 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/grpc v1.33.1 // indirect
	google.golang.org/protobuf v1.25.0
	honnef.co/go/tools v0.0.1-2020.1.6 // indirect
)
