package api

import (
	"bytes"
	"encoding/json"
	"github.com/devlucky/fakelink/src/links"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestPostLinkWithWrongInput(t *testing.T) {
	req, err := http.NewRequest("POST", "/links", bytes.NewReader([]byte{}))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	NewRouter(inMemoryConf()).ServeHTTP(rr, req)

	expectStatus(t, rr, http.StatusBadRequest)
}

func TestPostLink(t *testing.T) {
	input := &PostLinkInput{
		Link: links.RandomLink(),
	}

	inputBytes, err := json.Marshal(input)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/links", bytes.NewReader(inputBytes))
	if err != nil {
		t.Fatal(err)
	}

	config := inMemoryConf()
	rr := httptest.NewRecorder()
	NewRouter(config).ServeHTTP(rr, req)

	expectStatus(t, rr, http.StatusCreated)
	expectHeaderToContain(t, rr, "Access-Control-Allow-Origin", []string{"*"})
	expectHeaderToContain(t, rr, "Content-Type", []string{"application/json"})

	output := &PostLinkOutput{}
	json.Unmarshal(rr.Body.Bytes(), output)

	link := config.LinkStore.Find(output.Slug)
	if link == nil {
		t.Error("Expected POST /links to return the slug that identifies the links")
	}

	if !reflect.DeepEqual(link, input.Link) {
		t.Error("Expected input and saved links to be the same")
	}
}
