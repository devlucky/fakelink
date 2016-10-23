package main

import (
	"github.com/devlucky/fakelink/src/api"
	"github.com/devlucky/fakelink/src/links"
	"github.com/devlucky/fakelink/src/templates"
	"log"
	"net/http"
)

func main() {
	config := &api.Config{
		Template: templates.Get(),
		LinkStore: links.NewInMemoryStore(),
	}
	router := api.NewRouter(config)

	log.Println("Listening on 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
