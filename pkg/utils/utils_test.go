package utils

import (
	"net/http"
	"testing"
)

type ErrorBody struct {
	ErrorMsg *string `json:"error,omitempty"`
}

func Test_APIResponse_StatusOK(t *testing.T) {
	value := "BodyValue"
	expected := value
	result, _ := APIResponse(http.StatusOK, value)

	if result.Body != expected && result.StatusCode != 200 {
		t.Errorf("Test_APIResponse_StatusOK FAILED, expected `result.body` -> %s, got -> %s", expected, result.Body)
	} else {
		t.Logf("Test_APIResponse_StatusOK SUCCEED, expected `result.body` -> %s, got -> %s", expected, result.Body)
	}
}

func Test_APIResponse_StatusBadRequest(t *testing.T) {
	msg := "sample message"
	expected := `{"error":"sample message"}`
	result, _ := APIResponse(http.StatusBadRequest, ErrorBody{ErrorMsg: &msg})

	if result.Body != expected && result.StatusCode != 401 {
		t.Errorf("Test_APIResponse_StatusBadRequest FAILED, expected `result.body` -> %s, got -> %s", expected, result.Body)
	} else {
		t.Logf("Test_APIResponse_StatusBadRequest SUCCEED, expected `result.body` -> %s, got -> %s", expected, result.Body)
	}
}
