module github.com/Condition17/fleet-services/run-controller-service

go 1.13

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace (
	github.com/Condition17/fleet-services/run-controller-service => ./.
	github.com/coreos/etcd => github.com/ozonru/etcd/v3 v3.3.0-rc.0-grpc1.30.0
)

require (
	cloud.google.com/go v0.75.0 // indirect
	cloud.google.com/go/pubsub v1.9.1
	github.com/Condition17/fleet-services/binary-builder v0.0.0-20201124203551-1694b744d0ab // indirect
	github.com/Condition17/fleet-services/file-builder v0.0.0-20210131003915-3bf9dc271c8c
	github.com/Condition17/fleet-services/file-service v0.0.0-20210131003915-3bf9dc271c8c
	github.com/Condition17/fleet-services/lib v0.0.0-20210131003915-3bf9dc271c8c
	github.com/Condition17/fleet-services/resource-manager-service v0.0.0-20210131003915-3bf9dc271c8c
	github.com/Condition17/fleet-services/river-runner v0.0.0-20210131003915-3bf9dc271c8c
	github.com/Condition17/fleet-services/test-run-service v0.0.0-20210131003915-3bf9dc271c8c
	github.com/Condition17/fleet-services/user-service v0.0.0-20210131003915-3bf9dc271c8c // indirect
	github.com/go-git/go-git/v5 v5.2.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.4.3
	github.com/google/go-cmp v0.5.4 // indirect
	github.com/google/uuid v1.2.0 // indirect
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/broker/googlepubsub/v2 v2.9.1
	github.com/miekg/dns v1.1.37 // indirect
	github.com/nats-io/jwt v1.2.2 // indirect
	github.com/ozonru/etcd/v3 v3.3.0-rc.0-grpc1.30.0 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	go.opencensus.io v0.22.6 // indirect
	golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad // indirect
	golang.org/x/mod v0.4.1 // indirect
	golang.org/x/net v0.0.0-20210119194325-5f4716e94777 // indirect
	golang.org/x/oauth2 v0.0.0-20210126194326-f9ce19ea3013 // indirect
	golang.org/x/sys v0.0.0-20210124154548-22da62e12c0c // indirect
	golang.org/x/text v0.3.5 // indirect
	golang.org/x/tools v0.1.0 // indirect
	google.golang.org/api v0.38.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20210126160654-44e461bb6506 // indirect
	google.golang.org/grpc v1.35.0
	google.golang.org/grpc/examples v0.0.0-20201204164231-c7df457e12e0 // indirect
	google.golang.org/protobuf v1.25.0
)
