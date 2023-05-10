package controllers

import (
	"encoding/json"
	"github.com/Pad-TodoList/todoList-server/src/migrate"
	"github.com/Pad-TodoList/todoList-server/src/models"
	"github.com/Pad-TodoList/todoList-server/src/utils"
	"github.com/gorilla/mux"
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

		utils.HandleResponse(http.StatusCreated, result, w)
	}
}

func UpdateTask(dataAccess migrate.DataAccessObject) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var task models.IdentifiableTask

		err := json.NewDecoder(r.Body).Decode(&task)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		result := dataAccess.UpdateTask(task)

		utils.HandleResponse(http.StatusNoContent, result, w)
	}
}

func DeleteTask(dataAccess migrate.DataAccessObject) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		result := dataAccess.DeleteTask(id)

		utils.HandleResponse(http.StatusOK, result, w)
	}
}

func GetOneTask(dataAccess migrate.DataAccessObject) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		result := dataAccess.GetOneTask(id)

		utils.HandleResponse(http.StatusOK, result, w)
	}
}

func GetUserTask(dataAccess migrate.DataAccessObject) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		result := dataAccess.GetUserTask(id)

		utils.HandleResponse(http.StatusOK, result, w)
	}
}
