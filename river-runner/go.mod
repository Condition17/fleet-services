module github.com/Condition17/fleet-services/river-runner

go 1.14

replace github.com/Condition17/fleet-services/river-runner => ./

require (
	cloud.google.com/go/pubsub v1.9.1
	github.com/Condition17/fleet-services/file-service v0.0.0-20210130234853-9ee45ad119e1
	github.com/Condition17/fleet-services/lib v0.0.0-20210130234853-9ee45ad119e1
	github.com/Condition17/fleet-services/resource-manager-service v0.0.0-20210130234853-9ee45ad119e1
	github.com/Condition17/fleet-services/river v0.0.0-20210130234853-9ee45ad119e1
	github.com/Condition17/fleet-services/run-controller-service v0.0.0-20210130234853-9ee45ad119e1
	github.com/golang/protobuf v1.4.3
	github.com/miekg/dns v1.1.37 // indirect
	google.golang.org/api v0.38.0 // indirect
	google.golang.org/grpc v1.35.0
	google.golang.org/protobuf v1.25.0
)
