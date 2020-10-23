module test-run-service

go 1.13

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace (
	github.com/Condition17/fleet-services/test-run-service => ./.
	github.com/coreos/etcd => github.com/ozonru/etcd v3.3.20-grpc1.27-origmodule+incompatible
	google.golang.org/grpc => google.golang.org/grpc v1.27.0
)

require (
	github.com/Condition17/fleet-services/common v0.0.0-20201023195539-96e4832eeb3f
	github.com/Condition17/fleet-services/file-service v0.0.0-20201018213019-6857638ba994
	github.com/Condition17/fleet-services/run-controller-service v0.0.0-20201023195539-96e4832eeb3f
	github.com/Condition17/fleet-services/test-run-service v0.0.0-20201023195539-96e4832eeb3f
	github.com/Condition17/fleet-services/user-service v0.0.0-20201023195539-96e4832eeb3f
	github.com/go-delve/delve v1.5.0
	github.com/golang/protobuf v1.4.3
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.9.1
	github.com/tkanos/gonfig v0.0.0-20181112185242-896f3d81fadf
	github.com/xanzy/ssh-agent v0.3.0 // indirect
	golang.org/x/net v0.0.0-20201022231255-08b38378de70 // indirect
	golang.org/x/sys v0.0.0-20201022201747-fb209a7c41cd // indirect
	google.golang.org/genproto v0.0.0-20201022181438-0ff5f38871d5 // indirect
	google.golang.org/protobuf v1.25.0
	gorm.io/driver/postgres v1.0.5
	gorm.io/gorm v1.20.5
)
