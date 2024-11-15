package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func jsonError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Responding with 500 errors: " + msg)
	}

	type errorResponse = struct {
		Error string `json:"error"`
	}

	jsonResponse(w, code, errorResponse{
		Error: msg,
	})
}

func jsonResponse(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to Marshal json response: %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
