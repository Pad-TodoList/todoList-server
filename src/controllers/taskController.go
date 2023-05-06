package controllers

import (
	"encoding/json"
	"log"
	"net/http"
)

func HandlerTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode("task route")
	if err != nil {
		log.Fatalln("There was an error encoding the initialized struct")
	}
}
