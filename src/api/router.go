package api

import (
	"fmt"
	"github.com/devlucky/fakelink/src/images"
	"github.com/devlucky/fakelink/src/links"
	"github.com/devlucky/fakelink/src/templates"
	"github.com/julienschmidt/httprouter"
	"os"
)

func NewRouter(config *Config) *httprouter.Router {
	router := httprouter.New()
	router.OPTIONS("/*path", injectConfig(config, CORS))
	router.GET("/random", injectConfig(config, GetRandom))
	router.GET("/links/:slug", injectConfig(config, GetLink))
	router.POST("/links", injectConfig(config, PostLink))

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
