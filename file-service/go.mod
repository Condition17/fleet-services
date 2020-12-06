module github.com/Condition17/fleet-services/file-service

go 1.13

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace (
	github.com/Condition17/fleet-services/file-service => ./.
	github.com/coreos/etcd => github.com/ozonru/etcd/v3 v3.3.0-rc.0-grpc1.30.0
)

require (
	github.com/Condition17/fleet-services/lib v0.0.0-20201206215626-f08cb717fda8
	github.com/Condition17/fleet-services/run-controller-service v0.0.0-20201206215626-f08cb717fda8
	github.com/Condition17/fleet-services/user-service v0.0.0-20201206215626-f08cb717fda8 // indirect
	github.com/Microsoft/go-winio v0.4.15 // indirect
	github.com/coreos/etcd v3.3.25+incompatible // indirect
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/go-delve/delve v1.5.0 // indirect
	github.com/go-git/go-git/v5 v5.2.0 // indirect
	github.com/gofrs/uuid v3.3.0+incompatible
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/mock v1.4.4 // indirect
	github.com/golang/protobuf v1.4.3
	github.com/gomodule/redigo v1.8.3
	github.com/google/martian/v3 v3.1.0 // indirect
	github.com/google/uuid v1.1.2 // indirect
	github.com/imdario/mergo v0.3.11 // indirect
	github.com/jackc/pgx/v4 v4.9.2 // indirect
	github.com/kevinburke/ssh_config v0.0.0-20201106050909-4977a11b4351 // indirect
	github.com/lib/pq v1.8.0 // indirect
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/broker/googlepubsub/v2 v2.9.1
	github.com/micro/go-plugins/registry/kubernetes/v2 v2.9.1
	github.com/miekg/dns v1.1.35 // indirect
	github.com/mitchellh/hashstructure v1.1.0 // indirect
	github.com/nats-io/nats.go v1.10.0 // indirect
	github.com/ozonru/etcd/v3 v3.3.0-rc.0-grpc1.30.0 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/tkanos/gonfig v0.0.0-20181112185242-896f3d81fadf // indirect
	github.com/xanzy/ssh-agent v0.3.0 // indirect
	go.opencensus.io v0.22.5 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.16.0 // indirect
	golang.org/x/exp v0.0.0-20201008143054-e3b2a7f2fdc7 // indirect
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b // indirect
	golang.org/x/sys v0.0.0-20201204225414-ed752295db88 // indirect
	golang.org/x/text v0.3.4 // indirect
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e // indirect
	golang.org/x/tools v0.0.0-20201204222352-654352759326 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/grpc v1.34.0
	google.golang.org/protobuf v1.25.0
	gorm.io/driver/postgres v1.0.5 // indirect
	honnef.co/go/tools v0.0.1-2020.1.6 // indirect
	rsc.io/quote/v3 v3.1.0 // indirect
)
