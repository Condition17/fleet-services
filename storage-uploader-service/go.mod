module github.com/Condition17/fleet-services/storage-uploader-service

go 1.13

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace (
	github.com/Condition17/fleet-services/storage-uploader-service => ./.
	github.com/coreos/etcd => github.com/ozonru/etcd/v3 v3.3.0-rc.0-grpc1.30.0
)

require (
	cloud.google.com/go/storage v1.12.0
	github.com/Condition17/fleet-services/file-service v0.0.0-20201206212458-e7f5a325a3a5
	github.com/Condition17/fleet-services/lib v0.0.0-20201206212458-e7f5a325a3a5
	github.com/Condition17/fleet-services/run-controller-service v0.0.0-20201206212458-e7f5a325a3a5 // indirect
	github.com/Condition17/fleet-services/user-service v0.0.0-20201206212458-e7f5a325a3a5 // indirect
	github.com/Microsoft/go-winio v0.4.15 // indirect
	github.com/golang/protobuf v1.4.3
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/broker/googlepubsub/v2 v2.9.1
	github.com/micro/go-plugins/registry/kubernetes/v2 v2.9.1
	github.com/ozonru/etcd/v3 v3.3.0-rc.0-grpc1.30.0 // indirect
	golang.org/x/sys v0.0.0-20201204225414-ed752295db88 // indirect
	golang.org/x/tools v0.0.0-20201204222352-654352759326 // indirect
	google.golang.org/protobuf v1.25.0
)
