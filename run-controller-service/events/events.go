package events

const (
	// Events sent between microservices
	TestRunInitiated            = "test-run.initiated"
	FileChunksUploaded          = "file.chunksUploaded"
	FileSystemProvisioned       = "filesystem.provisioned"
	ExecutorInstanceProvisioned = "executorInstance.provisioned"
	FileAssemblySuccess         = "file.assemblySuccess"
	// Events that are sent using the WSS server to it's client (frontend application in our case)
	WssTestRunStateChanged          = "wss.test_run_state_changed"
	TEST_RUN_FINISHED               = "test-run.finished"
	FILE_ASSEMBLY_SUCCEEDED         = "file.assemblySucceeded"
	WSS_FILE_SUCCESSFULLY_ASSEMBLED = "fileSuccessfullyAssembled"
	WSS_TEST_RUN_FINSHED            = "testRunFinished"
	WSS_ERROR                       = "error"
)
