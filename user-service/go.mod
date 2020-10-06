module github.com/Condition17/fleet-services/user-service/proto

go 1.13

replace (
	github.com/Condition17/fleet-services/user-service => ./.
	github.com/coreos/etcd => github.com/ozonru/etcd v3.3.20-grpc1.27-origmodule+incompatible
	google.golang.org/grpc => google.golang.org/grpc v1.27.0
)

require (
	github.com/Condition17/fleet-services/common v0.0.0-20201006201526-a3474ae68cb6
	github.com/Condition17/fleet-services/user-service v0.0.0-20201004170504-061a0231c0d2
	github.com/coreos/etcd v3.3.25+incompatible // indirect
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/protobuf v1.4.2
	github.com/google/uuid v1.1.2 // indirect
	github.com/imdario/mergo v0.3.11 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/lib/pq v1.8.0 // indirect
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.9.1
	github.com/miekg/dns v1.1.31 // indirect
	github.com/nats-io/jwt v1.0.1 // indirect
	github.com/nats-io/nats.go v1.10.0 // indirect
	github.com/nats-io/nkeys v0.2.0 // indirect
	github.com/satori/go.uuid v1.2.0
	github.com/tkanos/gonfig v0.0.0-20181112185242-896f3d81fadf
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.16.0 // indirect
	golang.org/x/crypto v0.0.0-20201002170205-7f63de1d35b0
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b // indirect
	golang.org/x/net v0.0.0-20201002202402-0a1ea396d57c // indirect
	golang.org/x/tools v0.0.0-20201002184944-ecd9fd270d5d // indirect
	google.golang.org/genproto v0.0.0-20201002142447-3860012362da // indirect
	google.golang.org/grpc v1.32.0 // indirect
	google.golang.org/protobuf v1.25.0
	gorm.io/driver/postgres v1.0.2
	gorm.io/gorm v1.20.2
)
