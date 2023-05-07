package controllers

import (
	"encoding/json"
	"github.com/Pad-TodoList/todoList-server/src/migrate"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandlerUser(dataAccess migrate.DataAccessObject) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode("user route")
		if err != nil {
			log.Fatalln("There was an error encoding the initialized struct")
		}
	}
}

func GetUser(dataAccess migrate.DataAccessObject) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		result := dataAccess.GetUser(id)

		w.Header().Set("Content-Type", "application/json")
		if result.Status {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
		err := json.NewEncoder(w).Encode(result.Data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
