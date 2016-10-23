package api

import (
	"github.com/devlucky/fakelink/src/links"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
)

// Config is a container for all the interfaces and configuration options the API uses.
// It will be injected to the endpoints in order to allow them to access these options in a DI way
type Config struct {
	Template  *template.Template
	LinkStore links.Store
}

// Wraps an endpoint handler with a function that has access to a Config
func InjectConfig(c *Config, f func(http.ResponseWriter, *http.Request, httprouter.Params, *Config)) func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		f(w, r, ps, c)
	}
}
