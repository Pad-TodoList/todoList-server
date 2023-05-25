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
		//w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		routes := []models.RouteInfo{
			{Name: "about.json", Status: models.RouteStatus{Path: "/about.json", Status: "OK"}},
			{Name: "Create account", Status: models.RouteStatus{Path: "POST /register", Status: "OK"}},
			{Name: "Login", Status: models.RouteStatus{Path: "POST /login", Status: "OK"}},
			{Name: "Get user data", Status: models.RouteStatus{Path: "GET /getUser/{id}", Status: "OK"}},
			{Name: "Update user data", Status: models.RouteStatus{Path: "PUT /updateUser/{id}", Status: "OK"}},
			{Name: "Delete user data", Status: models.RouteStatus{Path: "DELETE /deleteUser/{id}", Status: "OK"}},
			{Name: "Create tasks", Status: models.RouteStatus{Path: "POST /createTask", Status: "OK"}},
			{Name: "Update task", Status: models.RouteStatus{Path: "PUT /updateTask/{id}", Status: "OK"}},
			{Name: "Get one task", Status: models.RouteStatus{Path: "GET /getTask/{id}", Status: "OK"}},
			{Name: "Get all user tasks", Status: models.RouteStatus{Path: "GET /getUserTasks/{id}", Status: "OK"}},
			{Name: "Delete task", Status: models.RouteStatus{Path: "DELETE /deleteTask/{id}", Status: "OK"}},
		}
		err := json.NewEncoder(w).Encode(routes)
		if err != nil {
			log.Fatalln("There was an error encoding the initialized struct")
		}
	}
}
