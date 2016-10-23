package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"fmt"
	"html/template"
	"errors"
)

type V1 struct {
	Title string
	Type string
	Url string
	Image string
}

const templateV1 = `
<!DOCTYPE html>
<html prefix="og: http://ogp.me/ns#">
<head>
    <title>{{.Title}}</title>
    <meta property="og:title" content="{{.Title}}" />
    <meta property="og:type" content="{{.Type}}" />
    <meta property="og:url" content="{{.Url}}" />
    <meta property="og:image" content="{{.Image}}" />
</head>
</html>
`

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

	values := &V1{
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
	TemplateStore *TemplateStore
}

type TemplateStore struct {
	templates map[string]*template.Template
}

func NewTemplateStore() (*TemplateStore) {
	templates := make(map[string]*template.Template)

	// TODO: Make this lazy and based on a map constant
	v1, err := template.New("v1").Parse(templateV1)
	if err != nil {
		log.Fatalf("Unexpected error parsing the template: %s", err)
	}
	templates["v1"] = v1

	return &TemplateStore{templates}
}

func (store *TemplateStore) GetTemplate(version string) (*template.Template, error) {
	if t, ok := store.templates[version]; ok {
		return t, nil
	}

	return nil, errors.New(fmt.Sprintf("Could not find template %s", version))
}

func main() {
	router := httprouter.New()

	// TODO: Take this piece of code to a template store that loads lazily and stores the templates in the cache

	conf := &Config{
		TemplateStore: NewTemplateStore(),
	}

	router.GET("/links/:slug", InjectConfig(conf, GetLink))

	log.Println("Listening on 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

