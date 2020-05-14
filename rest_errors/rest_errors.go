package rest_errors

import (
	"errors"
	"net/http"
)

type RestErr struct {
	Message string        `json:"message"`
	Status  int           `json:"status"`
	Error   string        `json:"error"`
	Causes  []interface{} `json:"causes"`
}

func NewRestError(message string, status int, err string, causes []interface{}) RestErr {
	return RestErr{
		Message: message,
		Status:  status,
		Error:   err,
		Causes:  causes,
	}
}

func NewError(msg string) error {
	return errors.New(msg)
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   http.StatusText(http.StatusBadRequest),
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   http.StatusText(http.StatusNotFound),
	}
}

func NewInternalServerError(message string, err error) *RestErr {
	result := &RestErr{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   http.StatusText(http.StatusInternalServerError),
		//Causes: []interface{}{err},
	}
	if err != nil {
		result.Causes = append(result.Causes, err.Error())
	}

	return result
}
