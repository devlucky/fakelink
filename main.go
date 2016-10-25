package main

import (
	"fmt"
	"github.com/devlucky/fakelink/src/api"
	"github.com/devlucky/fakelink/src/images"
	"github.com/devlucky/fakelink/src/links"
	"github.com/devlucky/fakelink/src/templates"
	"log"
	"net/http"
	"os"
)

func importLinkExamples(c *api.Config) {
	for _, link := range links.ExampleLinks {
		c.LinkStore.Create(link)
	}
	log.Println("Successfully imported example links")
}

func main() {
	config := &api.Config{
		RootPath:  fmt.Sprintf("%s/src/github.com/devlucky/fakelink", os.Getenv("GOPATH")),
		DebugMode: os.Getenv("DEBUG") == "true",
		Template:  templates.Get(),
		LinkStore: links.NewRedisStore(
			os.Getenv("REDIS_HOST"),
			os.Getenv("REDIS_PORT"),
			os.Getenv("REDIS_PASS"),
		),
		ImageStore: images.NewS3Store(
			os.Getenv("MINIO_HOST"),
			os.Getenv("MINIO_PORT"),
			os.Getenv("MINIO_ACCESS_KEY"),
			os.Getenv("MINIO_SECRET_KEY"),
			os.Getenv("MINIO_PUBLIC_URL"),
		),
		ImageMaxWidth:  512,
		ImageMaxHeight: 512,
	}
	router := api.NewRouter(config)

	// Make sure we only create example links once
	if config.LinkStore.FindRandom() == "" {
		importLinkExamples(config)
	}

	log.Println("Listening on 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
