module github.com/Condition17/fleet-services/river-runner

go 1.14

replace github.com/Condition17/fleet-services/river-runner => ./

require (
	cloud.google.com/go/pubsub v1.9.0
	github.com/Condition17/fleet-services/file-service v0.0.0-20201207205346-13b7968f3516
	github.com/Condition17/fleet-services/lib v0.0.0-20201207205346-13b7968f3516
	github.com/Condition17/fleet-services/resource-manager-service v0.0.0-20201207205346-13b7968f3516
	github.com/Condition17/fleet-services/river v0.0.0-20201207205346-13b7968f3516
	github.com/Condition17/fleet-services/run-controller-service v0.0.0-20201207205346-13b7968f3516
	github.com/golang/protobuf v1.4.3
	golang.org/x/oauth2 v0.0.0-20201207163604-931764155e3f // indirect
	golang.org/x/tools v0.0.0-20201207191902-7bb39e4ca9ac // indirect
	google.golang.org/genproto v0.0.0-20201207150747-9ee31aac76e7 // indirect
	google.golang.org/grpc v1.34.0
	google.golang.org/protobuf v1.25.0
)
