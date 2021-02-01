package runStates

type TestRunStateType string

type runStatesEnumStruct struct {
	Initiated                     TestRunStateType
	InitiatedDone                 TestRunStateType
	FileUpload                    TestRunStateType
	FileUploadDone                TestRunStateType
	ProvisionFs                   TestRunStateType
	ProvisionFsDone               TestRunStateType
	ProvisionExecutorInstance     TestRunStateType
	ProvisionExecutorInstanceDone TestRunStateType
	AssembleFile                  TestRunStateType
	AssembleFileDone              TestRunStateType
	Evaluating                    TestRunStateType
	EvaluationDone                TestRunStateType
	Finished                      TestRunStateType
	Error                         TestRunStateType
}

var TestRunState = runStatesEnumStruct{
	Initiated:                     "initiated",
	InitiatedDone:                 "initiated_done",
	FileUpload:                    "file_upload",
	FileUploadDone:                "file_upload_done",
	ProvisionFs:                   "provision_fs",
	ProvisionFsDone:               "provision_fs_done",
	ProvisionExecutorInstance:     "provision_executor_instance",
	ProvisionExecutorInstanceDone: "provision_executor_instance_done",
	AssembleFile:                  "assemble_file",
	AssembleFileDone:              "assemble_file_done",
	Evaluating:                    "evaluating",
	EvaluationDone:                "evaluationDone",
	Finished:                      "finished",
	Error:                         "error",
}
