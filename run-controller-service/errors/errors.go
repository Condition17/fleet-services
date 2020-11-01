package errors

import "encoding/json"

type Error struct {
	Code           int         `json:"code"`
	DisplayMessage string      `json:"displayMessage"`
	Details        interface{} `json:"details"`
}

func (e *Error) Error() string {
	b, _ := json.Marshal(e)
	return string(b)
}

func EventUnmarshalError(details interface{}) error {
	return &Error{
		Code:           999,
		DisplayMessage: "Could not unmarshal given event object.",
		Details:        details,
	}
}

func FileCreationError(details interface{}) error {
	return &Error{
		Code:           1000,
		DisplayMessage: "Error while creating a file for test run.",
		Details:        details,
	}
}

func FileAssignError(details interface{}) error {
	return &Error{
		Code:           1001,
		DisplayMessage: "Error encountered while assigning the file to the created test run.",
		Details:        details,
	}
}
