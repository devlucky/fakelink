package main

import (
	"github.com/devlucky/fakelink/src/api"
	"github.com/devlucky/fakelink/src/links"
	"github.com/devlucky/fakelink/src/templates"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	router := httprouter.New()

	// TODO: Take this piece of code to a template store that loads lazily and stores the templates in the cache

	conf := &api.Config{
		TemplateStore: templates.NewInMemoryStore(),
		LinkStore:     links.NewInMemoryStore(),
	}

	router.OPTIONS("/*", api.InjectConfig(conf, api.CORS))
	router.GET("/links/:slug", api.InjectConfig(conf, api.GetLink))
	router.POST("/links/:version", api.InjectConfig(conf, api.PostLink))

	log.Println("Listening on 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
