package errors

import (
	"encoding/json"
	"fmt"
)

type Error struct {
	Code           int         `json:"code"`
	DisplayMessage string      `json:"displayMessage"`
	Details        interface{} `json:"details"`
}

func (e *Error) Error() string {
	b, _ := json.Marshal(e)
	return string(b)
}

func EventUnmarshalError(target interface{}, errorDetails interface{}) error {
	return &Error{
		Code:           999,
		DisplayMessage: "Could not unmarshal given event object.",
		Details: map[string]string{
			"target": fmt.Sprintf("%v", target),
			"error":  fmt.Sprintf("%v", errorDetails),
		},
	}
}

func FileCreationError(target interface{}, errorDetails interface{}) error {
	return &Error{
		Code:           1000,
		DisplayMessage: "Error while creating a file for test run.",
		Details: map[string]string{
			"target": fmt.Sprintf("%v", target),
			"error":  fmt.Sprintf("%v", errorDetails),
		},
	}
}

func FileAssignError(target interface{}, errorDetails interface{}) error {
	return &Error{
		Code:           1001,
		DisplayMessage: "Error encountered while assigning the file to the created test run.",
		Details: map[string]string{
			"target": fmt.Sprintf("%v", target),
			"error":  fmt.Sprintf("%v", errorDetails),
		},
	}
}

func TestRunRetrievalError(target interface{}, errorDetails interface{}) error {
	return &Error{
		Code:           1002,
		DisplayMessage: "Error encountered while retrieving test run.",
		Details: map[string]string{
			"target": fmt.Sprintf("%v", target),
			"error":  fmt.Sprintf("%v", errorDetails),
		},
	}
}

func FileSystemCreationError(target interface{}, errorDetails interface{}) error {
	return &Error{
		Code:           1003,
		DisplayMessage: "Error encountered while creating file system.",
		Details: map[string]string{
			"target": fmt.Sprintf("%v", target),
			"error":  fmt.Sprintf("%v", errorDetails),
		},
	}
}


func AssembleFileRequestError(target interface{}, errorDetails interface{}) error {
	return &Error{
		Code:           1004,
		DisplayMessage: "Error encountered while requesting file assembly.",
		Details: map[string]string{
			"target": fmt.Sprintf("%v", target),
			"error":  fmt.Sprintf("%v", errorDetails),
		},
	}
}

func RiverRunRequestError(target interface{}, errorDetails interface{}) error {
	return &Error{
		Code:           1005,
		DisplayMessage: "Error encountered while requesting file River run.",
		Details: map[string]string{
			"target": fmt.Sprintf("%v", target),
			"error":  fmt.Sprintf("%v", errorDetails),
		},
	}
}

func TestRunUserBytesContextAppendError(target interface{}, errorDetails interface{}) error {
	return &Error{
		Code:           1006,
		DisplayMessage: "Error appending test run's user bytes to context.",
		Details: map[string]string{
			"target": fmt.Sprintf("%v", target),
			"error":  fmt.Sprintf("%v", errorDetails),
		},
	}
}