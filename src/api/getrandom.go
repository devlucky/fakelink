package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"fmt"
)

func GetRandom(w http.ResponseWriter, r *http.Request, ps httprouter.Params, c *Config) {
	slug := c.LinkStore.FindRandom()
	if slug == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/links/%s", slug), http.StatusTemporaryRedirect)
}
