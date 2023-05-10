package controllers

import (
	"encoding/json"
	"github.com/Pad-TodoList/todoList-server/src/migrate"
	"github.com/Pad-TodoList/todoList-server/src/models"
	"github.com/Pad-TodoList/todoList-server/src/utils"
	"github.com/gorilla/mux"
	"net/http"
)

func GetUser(dataAccess migrate.DataAccessObject) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		result := dataAccess.GetUser(id)

		utils.HandleResponse(http.StatusOK, result, w)
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

		utils.HandleResponse(http.StatusNoContent, result, w)
	}
}

func DeleteUser(dataAccess migrate.DataAccessObject) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		result := dataAccess.DeleteUser(id)

		utils.HandleResponse(http.StatusOK, result, w)
	}
}
