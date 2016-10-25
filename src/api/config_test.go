package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"testing"
)

func TestInjectConfig(t *testing.T) {
	config := &Config{}
	handler := func(w http.ResponseWriter, r *http.Request, ps httprouter.Params, c *Config) {
		if c != config {
			t.Error("Expected the injected config to be the same the handler receives")
		}
	}

	f := injectConfig(config, handler)
	f(nil, nil, nil)
}
