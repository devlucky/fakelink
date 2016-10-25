package api

import (
	"encoding/json"
	"fmt"
	"github.com/devlucky/fakelink/src/links"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"image"
	"github.com/devlucky/fakelink/src/images"
	"github.com/satori/go.uuid"
)

type PostLinkInput struct {
	Link links.Link `json:"link"`
}

type PostLinkOutput struct {
	Slug string `json:"slug"`
}

// We expect a multipart/form-data request containing:
// 	- an optional "image"
// 	- a "json" with the expected input as values
func PostLink(w http.ResponseWriter, r *http.Request, ps httprouter.Params, c *Config) {
	err := r.ParseMultipartForm(1024)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	input := &PostLinkInput{}
	err = json.Unmarshal([]byte(r.FormValue("json")), &input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request body. Multipart form needs a 'json' key"))
		return
	}

	// We pass the new link through the creator in order to validate the raw input
	link, err := links.NewLink(input.Link.Values, input.Link.Private)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Invalid link: %s", err)))
		return
	}

	// If a custom image was uploaded, we store it and point the values to the image's URL
	file, _, err := r.FormFile("image")
	if err == nil {
		img, _, err := image.Decode(file)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("Could not decode image: %s", err)))
			return
		}

		thumbnail := images.Thumbnail(img, c.ImageMaxWidth, c.ImageMaxHeight)
		imageUrl, err := c.ImageStore.Put(uuid.NewV4().String(), thumbnail)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("Could upload image: %s", err)))
			return
		}

		input.Link.Values.Image = imageUrl
	}

	slug := c.LinkStore.Create(link)

	jsonResp, err := json.Marshal(&PostLinkOutput{slug})
	if err != nil {
		log.Printf("Unexpected error %s when marshaling the response into JSON", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(jsonResp))
}
