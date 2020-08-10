module test-run

go 1.13

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace (
	github.com/Condition17/fleet-services/upload => ../upload
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)

require (
	github.com/Condition17/fleet-services/upload v0.0.0-00010101000000-000000000000
	github.com/golang/protobuf v1.4.2
	github.com/jinzhu/gorm v1.9.15
	github.com/joho/godotenv v1.3.0
	github.com/lib/pq v1.3.0
	github.com/micro/go-micro/v2 v2.9.1
	google.golang.org/protobuf v1.25.0
)
