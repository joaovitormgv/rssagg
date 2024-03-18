package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWihtJSON(w, 200, struct{}{})
}
