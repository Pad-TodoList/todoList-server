package controllers

import (
	"encoding/json"
	"github.com/Pad-TodoList/todoList-server/src/migrate"
	"log"
	"net/http"

	"github.com/Pad-TodoList/todoList-server/src/models"
)

func HandlerAbout(_ migrate.DataAccessObject) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		routes := []models.RouteInfo{
			{Name: "about.json", Status: models.RouteStatus{Path: "GET /about.json", Status: "OK"}},
			{Name: "Create account", Status: models.RouteStatus{Path: "POST /register", Status: "OK"}},
			{Name: "Login", Status: models.RouteStatus{Path: "POST /login", Status: "OK"}},
			{Name: "Get user data", Status: models.RouteStatus{Path: "GET /user/get/{id}", Status: "OK"}},
			{Name: "Update user data", Status: models.RouteStatus{Path: "PUT /user/update", Status: "OK"}},
			{Name: "Delete user data", Status: models.RouteStatus{Path: "DELETE /user/delete/{id}", Status: "OK"}},
			{Name: "Create tasks", Status: models.RouteStatus{Path: "POST /task/create", Status: "OK"}},
			{Name: "Update task", Status: models.RouteStatus{Path: "PUT /task/update", Status: "OK"}},
			{Name: "Get one task", Status: models.RouteStatus{Path: "GET /task/get/{id}", Status: "OK"}},
			{Name: "Get all user tasks", Status: models.RouteStatus{Path: "GET /task/getUser/{id}", Status: "OK"}},
			{Name: "Delete task", Status: models.RouteStatus{Path: "DELETE /task/delete/{id}", Status: "OK"}},
		}
		err := json.NewEncoder(w).Encode(routes)
		if err != nil {
			log.Fatalln("There was an error encoding the initialized struct")
		}
	}
}
