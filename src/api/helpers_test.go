package api

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func expectStatus(t *testing.T, rr *httptest.ResponseRecorder, status int) {
	if status := rr.Code; status != status {
		t.Errorf("Expected status to be %s. Instead, it was %s", status, rr.Code)
	}
}

func expectHeaderToContain(t *testing.T, rr *httptest.ResponseRecorder, header string, values []string) {
	for _, value := range values {
		if !strings.Contains(rr.HeaderMap.Get(header), value) {
			t.Errorf("Expected %s header to contain value %s", header, value)
		}
	}
}

func expectBodyToContain(t *testing.T, rr *httptest.ResponseRecorder, values []string) {
	bodyString := rr.Body.String()

	for _, value := range values {
		if !strings.Contains(bodyString, value) {
			t.Errorf("Expected body %s to contain value %s", bodyString, value)
		}
	}
}
