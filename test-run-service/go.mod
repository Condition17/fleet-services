module github.com/Condition17/fleet-services/test-run-service

go 1.13

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.

replace (
	github.com/Condition17/fleet-services/test-run-service => ./.
	github.com/coreos/etcd => github.com/ozonru/etcd/v3 v3.3.0-rc.0-grpc1.30.0
)

require (
	cloud.google.com/go/pubsub v1.9.1 // indirect
	cloud.google.com/go/storage v1.12.0
	github.com/Condition17/fleet-services/file-service v0.0.0-20201207231537-2ab1f1cbc0f2
	github.com/Condition17/fleet-services/lib v0.0.0-20201210231051-c1deafb7a724
	github.com/Condition17/fleet-services/run-controller-service v0.0.0-20201210231051-c1deafb7a724
	github.com/Condition17/fleet-services/user-service v0.0.0-20201210231051-c1deafb7a724
	github.com/Microsoft/go-winio v0.4.16 // indirect
	github.com/ghodss/yaml v1.0.0
	github.com/golang/protobuf v1.4.3
	github.com/jackc/pgx/v4 v4.10.0 // indirect
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/broker/googlepubsub/v2 v2.9.1
	github.com/ozonru/etcd/v3 v3.3.0-rc.0-grpc1.30.0 // indirect
	github.com/tkanos/gonfig v0.0.0-20181112185242-896f3d81fadf
	golang.org/x/crypto v0.0.0-20201208171446-5f87f3452ae9 // indirect
	golang.org/x/net v0.0.0-20201209123823-ac852fbbde11 // indirect
	golang.org/x/sys v0.0.0-20201210223839-7e3030f88018 // indirect
	golang.org/x/text v0.3.4 // indirect
	golang.org/x/tools v0.0.0-20201210181237-6d345e82f1d8 // indirect
	google.golang.org/genproto v0.0.0-20201210142538-e3217bee35cc // indirect
	google.golang.org/protobuf v1.25.0
	gopkg.in/yaml.v2 v2.4.0
	gorm.io/driver/postgres v1.0.5
	gorm.io/gorm v1.20.8
)
