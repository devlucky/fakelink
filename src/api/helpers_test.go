package api

import (
	"encoding/json"
	"errors"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

/*
	GENERIC TEST HELPERS
*/
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

/*
	TESTS FOR SPECIFIC HELPER FUNCTIONS
*/

func TestResponse(t *testing.T) {
	endpoint := "/test_success"

	router := NewRouter(inMemoryConf())
	router.GET(endpoint, func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		response(w, http.StatusTeapot, []byte("{\"some\": \"json\"}"))
	})

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		t.Fatalf("Unexpected error creating a request: %s", err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	expectStatus(t, rr, http.StatusTeapot)
	expectHeaderToContain(t, rr, "Content-Type", []string{"application/json"})
	expectHeaderToContain(t, rr, "Access-Control-Allow-Origin", []string{"*"})
	expectBodyToContain(t, rr, []string{"some", "json"})
}

func TestErrorResponse(t *testing.T) {
	testErrorResponseDebugMode(t, true)
	testErrorResponseDebugMode(t, false)
}

func testErrorResponseDebugMode(t *testing.T, debugMode bool) {
	endpoint := "/test_error"
	apiMessage := "API-level message"
	internalMessage := errors.New("Debug-only message")

	conf := inMemoryConf()
	conf.DebugMode = debugMode

	router := NewRouter(conf)
	router.GET(endpoint, injectConfig(conf, func(w http.ResponseWriter, r *http.Request, ps httprouter.Params, c *Config) {
		errorResponse(w, http.StatusNotAcceptable, apiMessage, internalMessage, c)
	}))

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		t.Fatalf("Unexpected error creating a request: %s", err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	expectStatus(t, rr, http.StatusNotAcceptable)
	expectHeaderToContain(t, rr, "Content-Type", []string{"application/json"})
	expectHeaderToContain(t, rr, "Access-Control-Allow-Origin", []string{"*"})

	resp := &ErrorResponseOutput{}
	err = json.Unmarshal(rr.Body.Bytes(), resp)
	if err != nil {
		t.Fatalf("Unexpected error unmarshaling JSON response: %s", err)
	}

	if resp.Message != apiMessage {
		t.Errorf("Expected error's API message to be %s. Instead, it was %s", apiMessage, resp.Message)
	}

	if debugMode && resp.DebugMessage != internalMessage.Error() {
		t.Errorf("Expected error's debug message to be %s when debug mode is on. Instead, it was %s", internalMessage, resp.DebugMessage)
	}

	if !debugMode && resp.DebugMessage != "" {
		t.Errorf("Expected error's debug message to be empty when debug mode is not on. Instead, it was %s", resp.DebugMessage)
	}
}
