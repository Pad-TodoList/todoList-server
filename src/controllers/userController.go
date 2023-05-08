package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/Pad-TodoList/todoList-server/src/migrate"
	"github.com/Pad-TodoList/todoList-server/src/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

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

func UpdateUser(dataAccess migrate.DataAccessObject) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.IdentifiableUser

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, "Bad request body", http.StatusBadRequest)
			return
		}
		result := dataAccess.UpdateUser(user)
		if !result.Status {
			fmt.Printf("error database : %s\n", result.Data)
		}
		w.Header().Set("Content-Type", "application/json")
		if !result.Status {
			w.WriteHeader(http.StatusForbidden)
		} else {
			w.WriteHeader(http.StatusForbidden)
		}
		err = json.NewEncoder(w).Encode(result.Data)
		if err != nil {
			log.Fatalln("There was an error encoding the initialized struct")
		}
	}
}

func DeleteUser(dataAccess migrate.DataAccessObject) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		result := dataAccess.DeleteUser(id)

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
