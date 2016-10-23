package api

import (
	"encoding/json"
	"fmt"
	"github.com/devlucky/fakelink/src/links"
	"github.com/devlucky/fakelink/src/templates"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type PostLinkResponse struct {
	Slug string `json:"slug"`
}

func PostLink(w http.ResponseWriter, r *http.Request, ps httprouter.Params, c *Config) {
	var input *templates.Values

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request body. Could not be parsed into JSON"))
		return
	}
	defer r.Body.Close()

	link, err := links.NewLink("v1", input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		msg := fmt.Sprintf("Invalid values. Error was: %s", err)
		w.Write([]byte(msg))
		return
	}

	slug := c.LinkStore.Create(link)

	jsonResp, err := json.Marshal(&PostLinkResponse{slug})
	if err != nil {
		log.Printf("Unexpected error %s when marshaling the response into JSON", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(jsonResp))
}
