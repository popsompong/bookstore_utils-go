package rest_errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNewInternalServerError(t *testing.T) {
	err := NewInternalServerError("this is the message", errors.New("database error"))
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "this is the message", err.Message)
	assert.EqualValues(t, http.StatusText(http.StatusInternalServerError), err.Error)
	assert.NotNil(t, err.Causes)
	assert.EqualValues(t, 1, len(err.Causes()))
	assert.EqualValues(t, "database error", err.Causes()[0])

	errBytes, _ := json.Marshal(err)
	fmt.Println(string(errBytes))
}

func TestNewBadRequestError(t *testing.T) {
	//TODO: Test!
}

func TestNewNotFoundError(t *testing.T) {
	//TODO: Test!
}

func TestNewError(t *testing.T) {
	//TODO: Test!
}
