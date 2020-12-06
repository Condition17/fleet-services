module github.com/Condition17/fleet-services/resource-manager-service

go 1.13

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace (
	github.com/Condition17/fleet-services/resource-manager-service => ./.
	github.com/coreos/etcd => github.com/ozonru/etcd/v3 v3.3.0-rc.0-grpc1.30.0
)

require (
	github.com/Condition17/fleet-services/lib v0.0.0-20201206221626-e5b9674ec1ac
	github.com/Condition17/fleet-services/run-controller-service v0.0.0-20201206221626-e5b9674ec1ac
	github.com/Condition17/fleet-services/test-run-service v0.0.0-20201206221626-e5b9674ec1ac
	github.com/Condition17/fleet-services/user-service v0.0.0-20201206221626-e5b9674ec1ac
	github.com/golang/protobuf v1.4.3
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/broker/googlepubsub/v2 v2.9.1
	github.com/ozonru/etcd/v3 v3.3.0-rc.0-grpc1.30.0 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	google.golang.org/api v0.36.0
	google.golang.org/protobuf v1.25.0
	gorm.io/driver/postgres v1.0.5
	gorm.io/gorm v1.20.8
)
