package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCORSOptions(t *testing.T) {
	req, err := http.NewRequest("OPTIONS", "/anything", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	NewRouter(inMemoryConf()).ServeHTTP(rr, req)

	expectStatus(t, rr, http.StatusOK)
	expectHeaderToContain(t, rr, "Access-Control-Allow-Headers", []string{"Content-Type"})
	expectHeaderToContain(t, rr, "Access-Control-Allow-Methods", []string{"GET", "POST", "OPTIONS", "PUT", "PATCH", "DELETE"})
	expectHeaderToContain(t, rr, "Access-Control-Allow-Origin", []string{"*"})
}
