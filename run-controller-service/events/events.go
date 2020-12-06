package events

const (
	TestRunInitiated                         = "test-run.initiated"
	FileChunksUploaded                       = "file.chunksUploaded"
	FileSystemProvisioned                    = "filesystem.provisioned"
	ExecutorInstanceProvisioned              = "executorInstance.provisioned"
	WssTestRunStateChanged                   = "wss.test_run_state_changed"
	TEST_RUN_FINISHED                        = "test-run.finished"
	FILE_SYSTEM_CREATED                      = "filesystem.created"
	EXECUTOR_INSTANCE_CREATED                = "executorinstance.created"
	FILE_ASSEMBLY_SUCCEEDED                  = "file.assemblySucceeded"
	WSS_FILE_SYSTEM_CREATION_COMPLETED       = "fileSystemCreationCompleted"
	WSS_EXECUTOR_INSTANCE_CREATION_COMPLETED = "executorInstanceCreationCompleted"
	WSS_FILE_SUCCESSFULLY_ASSEMBLED          = "fileSuccessfullyAssembled"
	WSS_TEST_RUN_FINSHED                     = "testRunFinished"
	WSS_ERROR                                = "error"
)
