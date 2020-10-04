module github.com/Condition17/fleet-services/user-service/proto

go 1.13

replace (
	github.com/Condition17/fleet-services/user-service => ./.
	github.com/coreos/etcd => github.com/ozonru/etcd v3.3.20-grpc1.27-origmodule+incompatible
	google.golang.org/grpc => google.golang.org/grpc v1.27.0
)

require (
	github.com/Condition17/fleet-services/user-service v0.0.0-00010101000000-000000000000
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/golang/protobuf v1.4.0
	github.com/micro/go-micro/v2 v2.9.1
	google.golang.org/protobuf v1.22.0
)
