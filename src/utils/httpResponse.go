package utils

import (
	"encoding/json"
	"github.com/Pad-TodoList/todoList-server/src/models"
	"log"
	"net/http"
)

func GoodResponse(w http.ResponseWriter, code int, data any) {
	w.Header().Set("Content-Type", "application/json")
	//w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, "There was an error encoding the initialized struct", http.StatusInternalServerError)
	}
}

func BadResponse(w http.ResponseWriter, error models.ErrorMessage) {
	w.Header().Set("Content-Type", "application/json")
	//w.Header().Set("Access-Control-Allow-Origin", "*")
	switch error.Code {
	case models.NotFound:
		w.WriteHeader(http.StatusNotFound)
	case models.DatabaseAction:
		w.WriteHeader(http.StatusInternalServerError)
	case models.ConnectionDatabase:
		w.WriteHeader(http.StatusInternalServerError)
	case models.Forbidden:
		w.WriteHeader(http.StatusForbidden)
	}
	err := json.NewEncoder(w).Encode(error.Message)
	if err != nil {
		log.Fatalln("There was an error encoding the initialized struct")
	}
}

func HandleResponse(goodStatusCode int, result models.DataAccessMessage, w http.ResponseWriter) {
	if result.Status {
		GoodResponse(w, goodStatusCode, result.Data)
	} else {
		_, ok := result.Data.(models.ErrorMessage)
		if !ok {
			log.Fatalln("Error with error message")
		} else {
			BadResponse(w, result.Data.(models.ErrorMessage))
		}
	}
}
