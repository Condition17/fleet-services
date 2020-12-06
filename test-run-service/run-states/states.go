package runStates

type TestRunStateType string

type runStatesEnumStruct struct {
	Initiated                     TestRunStateType
	FileUpload                    TestRunStateType
	FileUploadDone                TestRunStateType
	ProvisionFs                   TestRunStateType
	ProvisionFsDone               TestRunStateType
	ProvisionExecutorInstance     TestRunStateType
	ProvisionExecutorInstanceDone TestRunStateType
	AssembleFile                  TestRunStateType
	AssembleFileDone              TestRunStateType
	Evaluating                    TestRunStateType
	Failed                        TestRunStateType
	Succeeded                     TestRunStateType
	Error                         TestRunStateType
}

var TestRunState = runStatesEnumStruct{
	Initiated:                     "initiated",
	FileUpload:                    "file_upload",
	FileUploadDone:                "file_upload_done",
	ProvisionFs:                   "provision_fs",
	ProvisionFsDone:               "provision_fs_done",
	ProvisionExecutorInstance:     "provision_executor_instance",
	ProvisionExecutorInstanceDone: "provision_executor_instance_done",
	AssembleFile:                  "assemble_file",
	AssembleFileDone:              "assemble_file_done",
	Evaluating:                    "evaluating",
	Failed:                        "failed",
	Succeeded:                     "succeeded",
	Error:                         "error",
}
