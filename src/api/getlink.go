package api

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func GetLink(w http.ResponseWriter, r *http.Request, ps httprouter.Params, c *Config) {
	slug := ps.ByName("slug")

	link := c.LinkStore.Find(slug)
	if link == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	template, err := c.TemplateStore.GetTemplate(link.Version)
	if err != nil {
		log.Printf("Unexpected template %s could not be found. Error: %s", "v1", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	template.Execute(w, link.Values)
}
