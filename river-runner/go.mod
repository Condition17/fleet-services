module github.com/Condition17/fleet-services/river-runner

go 1.14

replace github.com/Condition17/fleet-services/river-runner => ./

require (
	cloud.google.com/go/pubsub v1.9.0
	github.com/Condition17/fleet-services/file-service v0.0.0-20201203221802-2d8665f2d498
	github.com/Condition17/fleet-services/lib v0.0.0-20201203221802-2d8665f2d498
	github.com/Condition17/fleet-services/resource-manager-service v0.0.0-20201203221802-2d8665f2d498
	github.com/golang/protobuf v1.4.3
	golang.org/x/net v0.0.0-20201202161906-c7110b5ffcbb // indirect
	golang.org/x/oauth2 v0.0.0-20201203001011-0b49973bad19 // indirect
	golang.org/x/sys v0.0.0-20201202213521-69691e467435 // indirect
	golang.org/x/tools v0.0.0-20201204062850-545788942d5f // indirect
	google.golang.org/grpc v1.34.0
	google.golang.org/protobuf v1.25.0
)
