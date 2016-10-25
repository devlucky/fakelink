package main

import (
	"github.com/devlucky/fakelink/src/api"
	"github.com/devlucky/fakelink/src/links"
	"github.com/devlucky/fakelink/src/templates"
	"log"
	"net/http"
	"os"
	"github.com/devlucky/fakelink/src/images"
)

func main() {
	config := &api.Config{
		Template: templates.Get(),
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
		ImageMaxWidth: 512,
		ImageMaxHeight: 512,
	}
	router := api.NewRouter(config)

	log.Println("Listening on 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
