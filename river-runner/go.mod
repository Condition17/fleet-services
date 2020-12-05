module github.com/Condition17/fleet-services/river-runner

go 1.14

replace github.com/Condition17/fleet-services/river-runner => ./

require (
	cloud.google.com/go/pubsub v1.9.0
	github.com/Condition17/fleet-services/file-service v0.0.0-20201205190639-bbc864952b08
	github.com/Condition17/fleet-services/lib v0.0.0-20201205190639-bbc864952b08
	github.com/Condition17/fleet-services/resource-manager-service v0.0.0-20201205190639-bbc864952b08
	github.com/Condition17/fleet-services/river v0.0.0-20201205190639-bbc864952b08
	github.com/Condition17/fleet-services/run-controller-service v0.0.0-20201205190639-bbc864952b08
	github.com/golang/protobuf v1.4.3
	google.golang.org/grpc v1.34.0
	google.golang.org/protobuf v1.25.0
)
