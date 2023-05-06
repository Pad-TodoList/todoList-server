package controllers

import (
	"encoding/json"
	"log"
	"net/http"
)

func HandlerLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode("login route")
	if err != nil {
		log.Fatalln("There was an error encoding the initialized struct")
	}
}
