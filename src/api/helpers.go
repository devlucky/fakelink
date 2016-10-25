package api

import (
	"encoding/json"
	"net/http"
)

func response(w http.ResponseWriter, status int, json []byte) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(json))
}

type ErrorResponseOutput struct {
	Message      string `json:"message"`
	DebugMessage string `json:"debug_mesage"`
}

func errorResponse(w http.ResponseWriter, status int, message string, err error, c *Config) {
	output := &ErrorResponseOutput{Message: message}
	if c.DebugMode {
		output.DebugMessage = err.Error()
	}

	jsonResp, err := json.Marshal(output)
	if err != nil {
		jsonResp = []byte("{}")
	}

	response(w, status, jsonResp)
}
