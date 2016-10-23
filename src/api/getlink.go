package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func GetLink(w http.ResponseWriter, r *http.Request, ps httprouter.Params, c *Config) {
	slug := ps.ByName("slug")

	link := c.LinkStore.Find(slug)
	if link == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	c.Template.Execute(w, link.Values)
}
