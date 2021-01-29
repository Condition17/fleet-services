module github.com/Condition17/fleet-services/test-run-service

go 1.13

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.

replace (
	github.com/Condition17/fleet-services/test-run-service => ./.
	github.com/coreos/etcd => github.com/ozonru/etcd/v3 v3.3.0-rc.0-grpc1.30.0
)

require (
	cloud.google.com/go v0.75.0 // indirect
	cloud.google.com/go/pubsub v1.9.1 // indirect
	cloud.google.com/go/storage v1.12.0
	github.com/Condition17/fleet-services/file-service v0.0.0-20201213183648-f069e2830265 // indirect
	github.com/Condition17/fleet-services/lib v0.0.0-20210129205841-3118c743e89b
	github.com/Condition17/fleet-services/run-controller-service v0.0.0-20210129205841-3118c743e89b
	github.com/Condition17/fleet-services/user-service v0.0.0-20210129205841-3118c743e89b
	github.com/Microsoft/go-winio v0.4.16 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.4.3
	github.com/google/uuid v1.2.0 // indirect
	github.com/jackc/pgproto3/v2 v2.0.7 // indirect
	github.com/micro/go-micro v1.18.0 // indirect
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/broker/googlepubsub/v2 v2.9.1
	github.com/ozonru/etcd/v3 v3.3.0-rc.0-grpc1.30.0 // indirect
	github.com/tkanos/gonfig v0.0.0-20181112185242-896f3d81fadf // indirect
	go.opencensus.io v0.22.6 // indirect
	golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad // indirect
	golang.org/x/mod v0.4.1 // indirect
	golang.org/x/net v0.0.0-20210119194325-5f4716e94777 // indirect
	golang.org/x/oauth2 v0.0.0-20210126194326-f9ce19ea3013 // indirect
	golang.org/x/sys v0.0.0-20210124154548-22da62e12c0c // indirect
	golang.org/x/text v0.3.5 // indirect
	golang.org/x/tools v0.1.0 // indirect
	google.golang.org/api v0.37.0 // indirect
	google.golang.org/genproto v0.0.0-20210126160654-44e461bb6506 // indirect
	google.golang.org/grpc v1.35.0 // indirect
	google.golang.org/protobuf v1.25.0
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gorm.io/driver/postgres v1.0.7
	gorm.io/gorm v1.20.12
)
