package handler

import (
	"context"
	b64 "encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Condition17/fleet-services/lib/auth"
	runStateEvents "github.com/Condition17/fleet-services/run-controller-service/events"
	runControllerProto "github.com/Condition17/fleet-services/run-controller-service/proto/run-controller-service"
	"github.com/Condition17/fleet-services/test-run-service/model"
	proto "github.com/Condition17/fleet-services/test-run-service/proto/test-run-service"
	runStates "github.com/Condition17/fleet-services/test-run-service/run-states"
	microErrors "github.com/micro/go-micro/v2/errors"
	"gorm.io/gorm"
	"log"
	"reflect"
	"time"
)

func (h *Handler) Create(ctx context.Context, req *proto.CreateTestRunRequest, res *proto.TestRunDetails) error {
	user, _ := auth.GetUserFromDecodedToken(ctx)
	newTestRun := model.MarshalTestRun(req.TestRun)
	newTestRun.UserID = user.Id
	createdTestRun, err := h.TestRunRepository.Create(newTestRun)
	if err != nil {
		return microErrors.BadRequest(h.Service.Name(), fmt.Sprintf("%v", err))
	}
	res.TestRun = model.UnmarshalTestRun(createdTestRun)

	// send test run initiated event
	eventData, _ := json.Marshal(
		&runControllerProto.TestRunInitiatedEventData{
			TestRunSpec: &runControllerProto.TestRunSpec{
				Id:   uint32(createdTestRun.ID),
				Name: createdTestRun.Name,
			},
			FileSpec: &runControllerProto.FileSpec{
				Name:         req.FileSpec.Name,
				Size:         req.FileSpec.Size,
				MaxChunkSize: req.FileSpec.MaxChunkSize,
			},
		},
	)
	h.SendRunStateEvent(ctx, runStateEvents.TestRunInitiated, eventData)

	return nil
}

func (h *Handler) AssignFile(ctx context.Context, req *proto.AssignRequest, res *proto.EmptyResponse) error {
	user, _ := auth.GetUserFromDecodedToken(ctx)
	testRun, err := h.TestRunRepository.GetUserTestRun(user.Id, req.TestRunId)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return microErrors.NotFound(h.Service.Name(), "Test run not found")
		}
		return microErrors.InternalServerError(h.Service.Name(), fmt.Sprintf("%v", err))
	}

	if testRun.FileID != "" {
		return microErrors.Conflict(h.Service.Name(), "Test run already has a file assigned to it.")
	}

	testRun.FileID = req.FileId
	if err := h.TestRunRepository.Update(testRun); err != nil {
		return microErrors.InternalServerError(h.Service.Name(), fmt.Sprintf("%v", err))
	}

	return nil
}

func (h *Handler) Get(ctx context.Context, req *proto.TestRun, res *proto.TestRunDetails) error {
	user, _ := auth.GetUserFromDecodedToken(ctx)
	result, err := h.TestRunRepository.GetUserTestRun(user.Id, req.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return microErrors.NotFound(h.Service.Name(), "Test run with this id not found")
		}
		return microErrors.InternalServerError(h.Service.Name(), fmt.Sprintf("%v", err))
	}
	res.TestRun = model.UnmarshalTestRun(result)

	return nil
}

func (h *Handler) GetById(ctx context.Context, req *proto.TestRun, res *proto.TestRunDetails) error {
	var isServiceCaller bool = ctx.Value("serviceCaller").(bool)

	if !isServiceCaller {
		return microErrors.Unauthorized(h.Service.Name(), "Caller not authorized for this operation")
	}

	result, err := h.TestRunRepository.GetTestRunById(req.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return microErrors.NotFound(h.Service.Name(), "No test run associated with the specified id was found")
		}
		return microErrors.InternalServerError(h.Service.Name(), fmt.Sprintf("%v", err))
	}
	res.TestRun = model.UnmarshalTestRun(result)

	return nil
}

func (h *Handler) List(ctx context.Context, req *proto.EmptyRequest, res *proto.ListResponse) error {
	user, _ := auth.GetUserFromDecodedToken(ctx)
	results, err := h.TestRunRepository.GetAll(user.Id)
	if err != nil {
		return microErrors.InternalServerError(h.Service.Name(), fmt.Sprintf("%v", err))
	}
	res.TestRuns = model.UnmarshalTestRunsCollection(results)

	return nil
}

func (h *Handler) Delete(ctx context.Context, req *proto.TestRun, res *proto.EmptyResponse) error {
	user, _ := auth.GetUserFromDecodedToken(ctx)
	if err := h.TestRunRepository.Delete(user.Id, req.Id); err != nil {
		return microErrors.InternalServerError(h.Service.Name(), fmt.Sprintf("%v", err))
	}

	return nil
}

func (h *Handler) ChangeState(ctx context.Context, req *proto.TestRunStateSpec, res *proto.TestRunDetails) error {
	//var isServiceCaller bool = ctx.Value("serviceCaller").(bool)
	//
	//if !isServiceCaller {
	//	return microErrors.Unauthorized(h.Service.Name(), "Caller not authorized for this operation")
	//}

	testRun, err := h.TestRunRepository.GetTestRunById(req.TestRunId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return microErrors.NotFound(h.Service.Name(), "Test run not found")
		}
		return microErrors.InternalServerError(h.Service.Name(), fmt.Sprintf("%v", err))
	}

	if !isValidTestRunState(req.State) {
		return microErrors.BadRequest(h.Service.Name(), fmt.Sprintf("State '%v' not recognized as a valid state", req.State))
	}

	if err := h.updateState(testRun, runStates.TestRunStateType(req.State), req.StateMetadata); err != nil {
		return microErrors.Forbidden(h.Service.Name(), fmt.Sprintf("%v", err))
	}

	res.TestRun = model.UnmarshalTestRun(testRun)

	return nil
}

func (h *Handler) ForceStop(ctx context.Context, req *proto.ForceStopRequest, res *proto.EmptyResponse) error {
	log.Println("Force stop test run:", req.TestRunId)

	testRun, err := h.TestRunRepository.GetTestRunById(req.TestRunId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return microErrors.NotFound(h.Service.Name(), "Test run not found")
		}
		return microErrors.InternalServerError(h.Service.Name(), fmt.Sprintf("%v", err))
	}

	if testRun.State != runStates.TestRunState.Initiated && testRun.State != runStates.TestRunState.FileUpload {
		return nil
	}

	if err := h.updateState(testRun, runStates.TestRunState.Error, "Force-stopped"); err != nil {
		return microErrors.InternalServerError(h.Service.Name(), fmt.Sprintf("%v", err))
	}
	return nil
}


func (h *Handler) RegisterRunIssue(ctx context.Context, req *proto.RunIssue, res *proto.EmptyResponse) error {
	// Process request async
	go func() {
		var err error
		var issueFileUrl string
		// Upload as a binary file, the input bytes that determined the issue
		bucketFileName := fmt.Sprintf("%v/%v_%v.bin", req.TestRunId, req.Issue, time.Now().Unix())
		if issueFileUrl, err = h.cloudStorageClient.UploadBytes(runIssuesBucketName, bucketFileName, req.InputBytes); err != nil {
			log.Fatalf("Error encountered uploading issue binary to bucket '%v' in file '%v': %v\n", runIssuesBucketName, bucketFileName, err)
		}

		newRunIssue := model.MarshalRunIssue(req)
		newRunIssue.InputBinUrl = issueFileUrl
		if err = h.RunIssueRepository.Create(newRunIssue); err != nil {
			log.Fatalln("Error encountered while creating run issue:", err)
		}
	}()
	return nil
}

func isValidTestRunState(state string) bool {
	e := reflect.ValueOf(&runStates.TestRunState).Elem()
	for i:=0; i < e.NumField(); i++ {
		if e.Field(i).Interface() == runStates.TestRunStateType(state) {
			return true
		}
	}

	return false
}

func (h *Handler) updateState(testRun *model.TestRun, newState runStates.TestRunStateType, newStateMetadata string) error {
	if testRun.State == runStates.TestRunState.Error {
		return errors.New(fmt.Sprintf("Could not change state from '%s' to '%s': testrun %d has a fixed error state.", testRun.State, newState, testRun.ID))
	}

	// if the new state is an error one, store the last valid state in it's associated metadata
	var stateMetadataBytes []byte
	if runStates.TestRunState.Error != newState {
		stateMetadataBytes = []byte(newStateMetadata)
		testRun.StateMetadata = b64.StdEncoding.EncodeToString([]byte(newStateMetadata))
	} else {
		stateMetadataBytes, _ = json.Marshal(map[string]string{"lastValidState": string(testRun.State), "metadata": newStateMetadata})
	}

	if runStates.TestRunState.Error == newState || runStates.TestRunState.Finished == newState {
		testRun.FinishedAt = time.Now()
	}

	testRun.StateMetadata = b64.StdEncoding.EncodeToString(stateMetadataBytes)
	testRun.State = newState

	if err := h.TestRunRepository.Update(testRun); err != nil {
		return err
	}

	return nil
}