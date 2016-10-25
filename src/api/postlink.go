package api

import (
	"encoding/json"
	"github.com/devlucky/fakelink/src/images"
	"github.com/devlucky/fakelink/src/links"
	"github.com/julienschmidt/httprouter"
	"github.com/satori/go.uuid"
	"image"
	"net/http"
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
		errorResponse(w, http.StatusBadRequest, "Format is not multipart/form-data", err, c)
		return
	}

	input := &PostLinkInput{}
	err = json.Unmarshal([]byte(r.FormValue("json")), &input)
	if err != nil {
		errorResponse(w, http.StatusBadRequest, "Invalid request body. Multipart form needs a 'json' key", err, c)
		return
	}

	// We pass the new link through the creator in order to validate the raw input
	link, err := links.NewLink(input.Link.Values, input.Link.Private)
	if err != nil {
		errorResponse(w, http.StatusBadRequest, "The link's structure or values are invalid", err, c)
		return
	}

	// If a custom image was uploaded, we store it and point the values to the image's URL
	file, _, err := r.FormFile("image")
	if err == nil {
		img, _, err := image.Decode(file)
		if err != nil {
			errorResponse(w, http.StatusBadRequest, "The image could not be decoded", err, c)
			return
		}

		thumbnail := images.Thumbnail(img, c.ImageMaxWidth, c.ImageMaxHeight)
		imageUrl, err := c.ImageStore.Put(uuid.NewV4().String(), thumbnail)
		if err != nil {
			errorResponse(w, http.StatusInternalServerError, "Could upload image", err, c)
			return
		}

		link.Values.Image = imageUrl
	}

	slug := c.LinkStore.Create(link)

	jsonResp, err := json.Marshal(&PostLinkOutput{slug})
	if err != nil {
		errorResponse(w, http.StatusInternalServerError, "Unexpected error when marshaling the response into JSON", err, c)
		return
	}

	response(w, http.StatusCreated, jsonResp)
}
