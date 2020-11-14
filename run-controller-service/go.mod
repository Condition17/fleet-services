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
	github.com/Condition17/fleet-services/file-service v0.0.0-20201114193928-f881acfb0b4f
	github.com/Condition17/fleet-services/lib v0.0.0-20201114193928-f881acfb0b4f
	github.com/Condition17/fleet-services/resource-manager-service v0.0.0-20201114193928-f881acfb0b4f
	github.com/Condition17/fleet-services/test-run-service v0.0.0-20201114193928-f881acfb0b4f
	github.com/Condition17/fleet-services/user-service v0.0.0-20201114193928-f881acfb0b4f // indirect
	github.com/go-git/go-git/v5 v5.2.0 // indirect
	github.com/golang/protobuf v1.4.3
	github.com/google/go-cmp v0.5.3 // indirect
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/broker/googlepubsub/v2 v2.9.1
	github.com/micro/go-plugins/registry/kubernetes/v2 v2.9.1
	github.com/nats-io/jwt v1.2.0 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	golang.org/x/crypto v0.0.0-20201112155050-0c6587e931a9 // indirect
	golang.org/x/sys v0.0.0-20201113233024-12cec1faf1ba // indirect
	golang.org/x/text v0.3.4 // indirect
	golang.org/x/tools v0.0.0-20201113202037-1643af1435f3 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20201113130914-ce600e9a6f9e // indirect
	google.golang.org/protobuf v1.25.0
)
