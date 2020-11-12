module github.com/Condition17/fleet-services/resource-manager-service

go 1.13

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace (
	github.com/Condition17/fleet-services/resource-manager-service => ./.
	github.com/coreos/etcd => github.com/ozonru/etcd v3.3.20-grpc1.27-origmodule+incompatible
	google.golang.org/grpc => google.golang.org/grpc v1.27.0
)

require (
	cloud.google.com/go v0.72.0 // indirect
	cloud.google.com/go/pubsub v1.8.3 // indirect
	github.com/Condition17/fleet-services/lib v0.0.0-20201107220710-33139c8c7d11
	github.com/Condition17/fleet-services/test-run-service v0.0.0-20201111212302-bc6c8bb28f29
	github.com/Condition17/fleet-services/user-service v0.0.0-20201111212302-bc6c8bb28f29
	github.com/golang/protobuf v1.4.3
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/broker/googlepubsub/v2 v2.9.1
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	golang.org/x/net v0.0.0-20201110031124-69a78807bb2b // indirect
	golang.org/x/oauth2 v0.0.0-20201109201403-9fd604954f58 // indirect
	golang.org/x/sys v0.0.0-20201110211018-35f3e6cf4a65 // indirect
	golang.org/x/tools v0.0.0-20201111133315-69daaf961d65 // indirect
	google.golang.org/api v0.35.0
	google.golang.org/genproto v0.0.0-20201111145450-ac7456db90a6 // indirect
	google.golang.org/protobuf v1.25.0
	gorm.io/driver/postgres v1.0.5
	gorm.io/gorm v1.20.6
)
