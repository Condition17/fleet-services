module github.com/Condition17/fleet-services/user-service

go 1.13

replace (
	github.com/Condition17/fleet-services/user-service => ./.
	github.com/coreos/etcd => github.com/ozonru/etcd v3.3.20-grpc1.27-origmodule+incompatible
	google.golang.org/grpc => google.golang.org/grpc v1.27.0
)

require (
	cloud.google.com/go v0.71.0 // indirect
	github.com/Condition17/fleet-services/lib v0.0.0-20201107200319-06e7c1fec256
	github.com/Condition17/fleet-services/run-controller-service v0.0.0-20201107200319-06e7c1fec256 // indirect
	github.com/Microsoft/go-winio v0.4.15 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/golang/protobuf v1.4.3
	github.com/jackc/pgx/v4 v4.9.2 // indirect
	github.com/kevinburke/ssh_config v0.0.0-20201106050909-4977a11b4351 // indirect
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/broker/googlepubsub/v2 v2.9.1
	github.com/miekg/dns v1.1.35 // indirect
	github.com/tkanos/gonfig v0.0.0-20181112185242-896f3d81fadf
	github.com/xanzy/ssh-agent v0.3.0 // indirect
	golang.org/x/crypto v0.0.0-20201016220609-9e8e0b390897
	golang.org/x/sys v0.0.0-20201107080550-4d91cf3a1aaf // indirect
	golang.org/x/tools v0.0.0-20201105220310-78b158585360 // indirect
	google.golang.org/api v0.35.0 // indirect
	google.golang.org/genproto v0.0.0-20201106154455-f9bfe239b0ba // indirect
	google.golang.org/grpc v1.33.2 // indirect
	google.golang.org/protobuf v1.25.0
	gorm.io/driver/postgres v1.0.5
	gorm.io/gorm v1.20.5
)
