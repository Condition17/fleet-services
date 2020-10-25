module github.com/Condition17/fleet-services/lib

go 1.14

// this replaces help testing the implemented funcitonality
replace (
	github.com/Condition17/fleet-services/lib => ./.
	github.com/coreos/etcd => github.com/ozonru/etcd v3.3.20-grpc1.27-origmodule+incompatible
	google.golang.org/grpc => google.golang.org/grpc v1.27.0
)

require (
	github.com/Condition17/fleet-services/common v0.0.0-20201010123938-c45c2dee5d72 // indirect
	github.com/Condition17/fleet-services/run-controller-service v0.0.0-20201025103716-bc8e68e1010b
	github.com/Condition17/fleet-services/user-service v0.0.0-20201025103716-bc8e68e1010b
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.9.1
	github.com/miekg/dns v1.1.35 // indirect
	github.com/xanzy/ssh-agent v0.3.0 // indirect
	golang.org/x/net v0.0.0-20201024042810-be3efd7ff127 // indirect
	golang.org/x/sys v0.0.0-20201024232916-9f70ab9862d5 // indirect
	google.golang.org/genproto v0.0.0-20201022181438-0ff5f38871d5 // indirect
	gorm.io/gorm v1.20.5
)
