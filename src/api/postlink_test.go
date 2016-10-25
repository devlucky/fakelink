package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/devlucky/fakelink/src/links"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

func TestPostLinkWithWrongFormat(t *testing.T) {
	req, err := http.NewRequest("POST", "/links", bytes.NewReader([]byte{}))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	NewRouter(inMemoryConf()).ServeHTTP(rr, req)

	expectStatus(t, rr, http.StatusBadRequest)
}

func TestPostInvalidLink(t *testing.T) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	err := bodyWriter.WriteField("json", "{}")
	if err != nil {
		t.Fatalf("Unexpected error writing multipart/form-data: %s", err)
	}

	req, err := http.NewRequest("POST", "/links", bodyBuf)
	if err != nil {
		t.Fatalf("Unexpected error creating a request: %s", err)
	}
	req.Header.Set("Content-Type", bodyWriter.FormDataContentType())
	bodyWriter.Close()

	config := inMemoryConf()
	rr := httptest.NewRecorder()
	NewRouter(config).ServeHTTP(rr, req)

	expectStatus(t, rr, http.StatusBadRequest)
}

func TestPostLinkWithoutImage(t *testing.T) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	/*
		JSON field
	*/
	input := &PostLinkInput{
		Link: *links.RandomLink(),
	}
	inputBytes, err := json.Marshal(input)
	if err != nil {
		t.Fatalf("Unexpected error marshaling input to JSON: %s", err)
	}
	err = bodyWriter.WriteField("json", string(inputBytes))
	if err != nil {
		t.Fatalf("Unexpected error writing multipart/form-data: %s", err)
	}

	/*
		Request & Response
	*/
	req, err := http.NewRequest("POST", "/links", bodyBuf)
	if err != nil {
		t.Fatalf("Unexpected error creating a request: %s", err)
	}
	req.Header.Set("Content-Type", bodyWriter.FormDataContentType())
	bodyWriter.Close()

	config := inMemoryConf()
	rr := httptest.NewRecorder()
	NewRouter(config).ServeHTTP(rr, req)

	expectStatus(t, rr, http.StatusCreated)
	expectHeaderToContain(t, rr, "Access-Control-Allow-Origin", []string{"*"})
	expectHeaderToContain(t, rr, "Content-Type", []string{"application/json"})

	/*
		Output
	*/
	output := &PostLinkOutput{}
	json.Unmarshal(rr.Body.Bytes(), output)

	link := config.LinkStore.Find(output.Slug)
	if link == nil {
		t.Error("Expected POST /links to return the slug that identifies the links")
	}

	if !reflect.DeepEqual(*link, input.Link) {
		t.Error("Expected input and saved links to be the same")
	}
}

func TestPostLinkWithImage(t *testing.T) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	/*
		JSON field
	*/
	input := &PostLinkInput{
		Link: *links.RandomLink(),
	}
	inputBytes, err := json.Marshal(input)
	if err != nil {
		t.Fatalf("Unexpected error marshaling input to JSON: %s", err)
	}
	err = bodyWriter.WriteField("json", string(inputBytes))
	if err != nil {
		t.Fatalf("Unexpected error writing multipart/form-data: %s", err)
	}

	/*
		IMAGE field
	*/
	filename := "sharknado.jpg"
	fileWriter, err := bodyWriter.CreateFormFile("image", filename)
	if err != nil {
		t.Fatalf("Unexpected error writing multipart/form-data: %s", err)
	}
	imageFile, err := os.Open(fmt.Sprintf("../../assets/images/%s", filename))
	if err != nil {
		t.Fatalf("Unexpected error opening file %s: %s", filename, err)
	}
	_, err = io.Copy(fileWriter, imageFile)
	if err != nil {
		t.Fatalf("Unexpected error writing image file to multipart/form-data: %s", err)
	}

	/*
		Request & Response
	*/
	req, err := http.NewRequest("POST", "/links", bodyBuf)
	if err != nil {
		t.Fatalf("Unexpected error creating a request: %s", err)
	}
	req.Header.Set("Content-Type", bodyWriter.FormDataContentType())
	bodyWriter.Close()

	config := inMemoryConf()
	rr := httptest.NewRecorder()
	NewRouter(config).ServeHTTP(rr, req)

	expectStatus(t, rr, http.StatusCreated)
	expectHeaderToContain(t, rr, "Access-Control-Allow-Origin", []string{"*"})
	expectHeaderToContain(t, rr, "Content-Type", []string{"application/json"})

	/*
		Output
	*/
	output := &PostLinkOutput{}
	json.Unmarshal(rr.Body.Bytes(), output)

	link := config.LinkStore.Find(output.Slug)
	if link == nil {
		t.Error("Expected POST /links to return the slug that identifies the links")
	}

	if link.Values.Image == input.Link.Values.Image {
		t.Error("Expected the link's Image to point to the uploaded file")
	}
}
