module github.com/Condition17/fleet-services/test-run-service

go 1.13

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.

replace (
	github.com/Condition17/fleet-services/test-run-service => ./.
	github.com/coreos/etcd => github.com/ozonru/etcd/v3 v3.3.0-rc.0-grpc1.30.0
)

require (
	github.com/Condition17/fleet-services/file-service v0.0.0-20201204085701-818ebf0007af
	github.com/Condition17/fleet-services/lib v0.0.0-20201206190129-e30184dea4ed
	github.com/Condition17/fleet-services/run-controller-service v0.0.0-20201206190129-e30184dea4ed
	github.com/Condition17/fleet-services/user-service v0.0.0-20201206190129-e30184dea4ed
	github.com/ghodss/yaml v1.0.0
	github.com/golang/protobuf v1.4.3
	github.com/jackc/pgx/v4 v4.10.0 // indirect
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/broker/googlepubsub/v2 v2.9.1
	github.com/ozonru/etcd/v3 v3.3.0-rc.0-grpc1.30.0 // indirect
	github.com/tkanos/gonfig v0.0.0-20181112185242-896f3d81fadf
	golang.org/x/sys v0.0.0-20201204225414-ed752295db88 // indirect
	golang.org/x/text v0.3.4 // indirect
	golang.org/x/tools v0.0.0-20201204222352-654352759326 // indirect
	google.golang.org/protobuf v1.25.0
	gopkg.in/yaml.v2 v2.4.0
	gorm.io/driver/postgres v1.0.5
	gorm.io/gorm v1.20.8
)
