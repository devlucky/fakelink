package api

import (
	"github.com/devlucky/fakelink/src/links"
	"github.com/devlucky/fakelink/src/templates"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetRandom(t *testing.T) {
	config := inMemoryConf()
	slug := config.LinkStore.Create(
		&links.Link{
			Values: &templates.Values{},
		},
	)

	req, err := http.NewRequest("GET", "/random", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	NewRouter(config).ServeHTTP(rr, req)

	expectStatus(t, rr, http.StatusTemporaryRedirect)
	expectHeaderToContain(t, rr, "Location", []string{slug})
}

func TestGetRandomWhenStoreIsEmpty(t *testing.T) {
	req, err := http.NewRequest("GET", "/random", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	NewRouter(inMemoryConf()).ServeHTTP(rr, req)

	expectStatus(t, rr, http.StatusNotFound)
}
