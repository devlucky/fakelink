package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"github.com/devlucky/fakelink/src/templates"
)

func GetLink(w http.ResponseWriter, r *http.Request, ps httprouter.Params, c *Config) {
	slug := ps.ByName("slug")
	if slug != "existing" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	t, err := c.TemplateStore.GetTemplate("v1")
	if err != nil {
		log.Printf("Unexpected template %s could not be found. Error: %s", "v1", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// TODO: Change this temporary mock for something that comes from the database

	values := &templates.Values{
		Title: "Sharknado",
		Type: "website",
		Url: "http://www.imdb.com/title/tt2724064/",
		Image: "https://images-na.ssl-images-amazon.com/images/M/MV5BOTE2OTk4MTQzNV5BMl5BanBnXkFtZTcwODUxOTM3OQ@@._V1_SY1000_CR0,0,712,1000_AL_.jpg",
	}

	w.WriteHeader(http.StatusOK)
	t.Execute(w, values)
}

func InjectConfig(a *Config, f func (http.ResponseWriter, *http.Request, httprouter.Params, *Config)) (func (http.ResponseWriter, *http.Request, httprouter.Params)) {
	return func (w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		f(w, r, ps, a)
	}
}

type Config struct {
	TemplateStore *templates.Store
}


func main() {
	router := httprouter.New()

	// TODO: Take this piece of code to a template store that loads lazily and stores the templates in the cache

	conf := &Config{
		TemplateStore: templates.NewStore(),
	}

	router.GET("/links/:slug", InjectConfig(conf, GetLink))

	log.Println("Listening on 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

