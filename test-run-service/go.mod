module github.com/Condition17/fleet-services/test-run-service

go 1.13

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace (
	github.com/Condition17/fleet-services/test-run-service => ./.
	github.com/coreos/etcd => github.com/ozonru/etcd v3.3.20-grpc1.27-origmodule+incompatible
	google.golang.org/grpc => google.golang.org/grpc v1.27.0
)

require (
	cloud.google.com/go v0.71.0 // indirect
	github.com/Condition17/fleet-services/lib v0.0.0-20201103081339-2d11f130c838
	github.com/Condition17/fleet-services/run-controller-service v0.0.0-20201103081339-2d11f130c838
	github.com/Condition17/fleet-services/user-service v0.0.0-20201103081339-2d11f130c838
	github.com/ghodss/yaml v1.0.0
	github.com/golang/protobuf v1.4.3
	github.com/jackc/pgx/v4 v4.9.1 // indirect
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/broker/googlepubsub/v2 v2.9.1
	github.com/nats-io/jwt v1.2.0 // indirect
	github.com/tkanos/gonfig v0.0.0-20181112185242-896f3d81fadf
	golang.org/x/net v0.0.0-20201031054903-ff519b6c9102 // indirect
	golang.org/x/text v0.3.4 // indirect
	golang.org/x/tools v0.0.0-20201102212025-f46e4245211d // indirect
	google.golang.org/genproto v0.0.0-20201102152239-715cce707fb0 // indirect
	google.golang.org/protobuf v1.25.0
	gopkg.in/yaml.v2 v2.3.0
	gorm.io/driver/postgres v1.0.5
	gorm.io/gorm v1.20.5
)
