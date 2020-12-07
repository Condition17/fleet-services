package events

const (
	// Events sent between microservices
	TestRunInitiated            = "test-run.initiated"
	FileChunksUploaded          = "file.chunksUploaded"
	FileSystemProvisioned       = "filesystem.provisioned"
	ExecutorInstanceProvisioned = "executorInstance.provisioned"
	FileAssemblySuccess         = "file.assemblySuccess"
	FileEvaluationFinished      = "file.evaluationFinished"
	// Events that are sent using the WSS server to it's client (frontend application in our case)
	WssTestRunStateChanged = "wss.test_run_state_changed"
)
