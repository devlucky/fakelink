package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func CORS(w http.ResponseWriter, r *http.Request, ps httprouter.Params, c *Config) {
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, PATCH, DELETE")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
}
