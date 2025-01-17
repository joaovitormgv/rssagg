package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWihtError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Respondind with 5XX error:", msg)
	}
	type errResponse struct {
		Error string `json:"error"`
	}

	respondWihtJSON(w, code, errResponse{
		Error: msg,
	})
}

func respondWihtJSON(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Fail to marshal JSON response: %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)

}
