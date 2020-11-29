module github.com/Condition17/fleet-services/user-service

go 1.13

replace (
	github.com/Condition17/fleet-services/user-service => ./.
	github.com/coreos/etcd => github.com/ozonru/etcd/v3 v3.3.0-rc.0-grpc1.30.0
)

require (
	cloud.google.com/go v0.72.0 // indirect
	cloud.google.com/go/pubsub v1.8.3 // indirect
	github.com/Condition17/fleet-services/lib v0.0.0-20201129170321-6dff70638039
	github.com/Condition17/fleet-services/run-controller-service v0.0.0-20201129170321-6dff70638039 // indirect
	github.com/Microsoft/go-winio v0.4.15 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/golang/protobuf v1.4.3
	github.com/jackc/pgx/v4 v4.9.2 // indirect
	github.com/kevinburke/ssh_config v0.0.0-20201106050909-4977a11b4351 // indirect
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/broker/googlepubsub/v2 v2.9.1
	github.com/miekg/dns v1.1.35 // indirect
	github.com/ozonru/etcd/v3 v3.3.0-rc.0-grpc1.30.0 // indirect
	github.com/tkanos/gonfig v0.0.0-20181112185242-896f3d81fadf
	github.com/xanzy/ssh-agent v0.3.0 // indirect
	golang.org/x/crypto v0.0.0-20201124201722-c8d3bf9c5392
	golang.org/x/net v0.0.0-20201110031124-69a78807bb2b // indirect
	golang.org/x/oauth2 v0.0.0-20201109201403-9fd604954f58 // indirect
	google.golang.org/api v0.35.0 // indirect
	google.golang.org/grpc v1.33.2 // indirect
	google.golang.org/protobuf v1.25.0
	gorm.io/driver/postgres v1.0.5
	gorm.io/gorm v1.20.7
)
