module upload

go 1.13

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace (
	github.com/coreos/etcd => github.com/ozonru/etcd v3.3.20-grpc1.27-origmodule+incompatible
	google.golang.org/grpc => google.golang.org/grpc v1.27.0
)

require (
	github.com/gofrs/uuid v3.2.0+incompatible
	github.com/gogo/protobuf v1.2.2-0.20190723190241-65acae22fc9d // indirect
	github.com/golang/protobuf v1.4.2
	github.com/gomodule/redigo v1.8.2
	github.com/joho/godotenv v1.3.0
	github.com/kr/pretty v0.2.0 // indirect
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/registry/mdns/v2 v2.9.1
	github.com/stretchr/testify v1.6.1 // indirect
	golang.org/x/crypto v0.0.0-20200709230013-948cd5f35899 // indirect
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b // indirect
	golang.org/x/net v0.0.0-20200707034311-ab3426394381 // indirect
	golang.org/x/sys v0.0.0-20200727154430-2d971f7391a4 // indirect
	golang.org/x/text v0.3.3 // indirect
	golang.org/x/tools v0.0.0-20200727154626-eaaaedc6afa7 // indirect
	google.golang.org/genproto v0.0.0-20200729003335-053ba62fc06f // indirect
	google.golang.org/grpc v1.30.0 // indirect
	google.golang.org/protobuf v1.25.0
	gopkg.in/yaml.v2 v2.2.8 // indirect
	honnef.co/go/tools v0.0.1-2020.1.4 // indirect
)
