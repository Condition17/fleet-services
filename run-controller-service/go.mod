module github.com/Condition17/fleet-services/run-controller-service

go 1.13

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace (
	github.com/Condition17/fleet-services/run-controller-service => ./.
	github.com/coreos/etcd => github.com/ozonru/etcd/v3 v3.3.0-rc.0-grpc1.30.0
)

require (
	cloud.google.com/go v0.73.0 // indirect
	cloud.google.com/go/pubsub v1.9.0 // indirect
	github.com/Condition17/fleet-services/file-builder v0.0.0-20201204085701-818ebf0007af
	github.com/Condition17/fleet-services/file-service v0.0.0-20201204085701-818ebf0007af
	github.com/Condition17/fleet-services/lib v0.0.0-20201204085701-818ebf0007af
	github.com/Condition17/fleet-services/resource-manager-service v0.0.0-20201204085701-818ebf0007af
	github.com/Condition17/fleet-services/river-runner v0.0.0-20201204085701-818ebf0007af
	github.com/Condition17/fleet-services/test-run-service v0.0.0-20201204085701-818ebf0007af
	github.com/Condition17/fleet-services/user-service v0.0.0-20201204085701-818ebf0007af // indirect
	github.com/go-git/go-git/v5 v5.2.0 // indirect
	github.com/golang/protobuf v1.4.3
	github.com/google/go-cmp v0.5.4 // indirect
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/broker/googlepubsub/v2 v2.9.1
	github.com/nats-io/jwt v1.2.2 // indirect
	github.com/ozonru/etcd/v3 v3.3.0-rc.0-grpc1.30.0 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	golang.org/x/crypto v0.0.0-20201203163018-be400aefbc4c // indirect
	golang.org/x/net v0.0.0-20201202161906-c7110b5ffcbb // indirect
	golang.org/x/oauth2 v0.0.0-20201203001011-0b49973bad19 // indirect
	golang.org/x/sys v0.0.0-20201202213521-69691e467435 // indirect
	golang.org/x/text v0.3.4 // indirect
	golang.org/x/tools v0.0.0-20201204162204-73cf035baebf // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20201204160425-06b3db808446 // indirect
	google.golang.org/grpc v1.34.0
	google.golang.org/grpc/examples v0.0.0-20201204164231-c7df457e12e0 // indirect
	google.golang.org/protobuf v1.25.0
)
