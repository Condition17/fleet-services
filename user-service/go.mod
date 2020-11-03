module github.com/Condition17/fleet-services/user-service

go 1.13

replace (
	github.com/Condition17/fleet-services/user-service => ./.
	github.com/coreos/etcd => github.com/ozonru/etcd v3.3.20-grpc1.27-origmodule+incompatible
	google.golang.org/grpc => google.golang.org/grpc v1.27.0
)

require (
	cloud.google.com/go v0.71.0 // indirect
	github.com/Condition17/fleet-services/lib v0.0.0-20201103081339-2d11f130c838
	github.com/Condition17/fleet-services/run-controller-service v0.0.0-20201103081339-2d11f130c838 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/golang/protobuf v1.4.3
	github.com/jackc/pgx/v4 v4.9.1 // indirect
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/broker/googlepubsub/v2 v2.9.1
	github.com/miekg/dns v1.1.35 // indirect
	github.com/tkanos/gonfig v0.0.0-20181112185242-896f3d81fadf
	github.com/xanzy/ssh-agent v0.3.0 // indirect
	golang.org/x/crypto v0.0.0-20201016220609-9e8e0b390897
	golang.org/x/tools v0.0.0-20201102212025-f46e4245211d // indirect
	google.golang.org/genproto v0.0.0-20201102152239-715cce707fb0 // indirect
	google.golang.org/protobuf v1.25.0
	gorm.io/driver/postgres v1.0.5
	gorm.io/gorm v1.20.5
)
