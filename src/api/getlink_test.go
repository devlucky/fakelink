package api

import (
	"fmt"
	"github.com/devlucky/fakelink/src/links"
	"github.com/devlucky/fakelink/src/templates"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetExistingLink(t *testing.T) {
	title := "the-great-api-test"

	config := inMemoryConf()
	slug := config.LinkStore.Create(
		&links.Link{
			Values: templates.Values{
				Title: title,
			},
		},
	)

	req, err := http.NewRequest("GET", fmt.Sprintf("/links/%s", slug), nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	NewRouter(config).ServeHTTP(rr, req)

	expectStatus(t, rr, http.StatusOK)
	expectBodyToContain(t, rr, []string{title})
}

func TestGetMissingLink(t *testing.T) {
	req, err := http.NewRequest("GET", "/links/missing", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	NewRouter(inMemoryConf()).ServeHTTP(rr, req)

	expectStatus(t, rr, http.StatusNotFound)
}
