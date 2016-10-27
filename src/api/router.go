package api

import (
	"fmt"
	"github.com/devlucky/fakelink/src/images"
	"github.com/devlucky/fakelink/src/links"
	"github.com/devlucky/fakelink/src/templates"
	"github.com/julienschmidt/httprouter"
	"os"
)

// NewRouter creates the router for the main API.
func NewRouter(config *Config) *httprouter.Router {
	router := httprouter.New()
	router.OPTIONS("/*path", injectConfig(config, cors))
	router.GET("/random", injectConfig(config, getRandom))
	router.GET("/links/:slug", injectConfig(config, getLink))
	router.POST("/links", injectConfig(config, postLink))

	return router
}

func inMemoryConf() *Config {
	return &Config{
		RootPath:       fmt.Sprintf("%s/src/github.com/devlucky/fakelink", os.Getenv("GOPATH")),
		DebugMode:      true,
		Template:       templates.Get(),
		LinkStore:      links.NewInMemoryStore(),
		ImageStore:     images.NewInMemoryStore(),
		ImageMaxWidth:  64,
		ImageMaxHeight: 64,
	}
}
