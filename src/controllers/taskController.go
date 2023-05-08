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

func CreateTask(dataAccess migrate.DataAccessObject) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var task models.Task

		err := json.NewDecoder(r.Body).Decode(&task)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		result := dataAccess.CreateTask(task)
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

func UpdateTask(dataAccess migrate.DataAccessObject) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode("task route")
		if err != nil {
			log.Fatalln("There was an error encoding the initialized struct")
		}
	}
}

func DeleteTask(dataAccess migrate.DataAccessObject) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		result := dataAccess.DeleteTask(id)

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

func GetOneTask(dataAccess migrate.DataAccessObject) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode("task route")
		if err != nil {
			log.Fatalln("There was an error encoding the initialized struct")
		}
	}
}

func GetUserTask(dataAccess migrate.DataAccessObject) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode("task route")
		if err != nil {
			log.Fatalln("There was an error encoding the initialized struct")
		}
	}
}
