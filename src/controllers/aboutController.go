package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Pad-TodoList/todoList-server/src/models"
)

func HandlerAbout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	routes := []models.RouteInfo{
		{Name: "about.json", Status: models.RouteStatus{Path: "/about.json", Status: "OK"}},
		{Name: "Create account", Status: models.RouteStatus{Path: "POST /register", Status: "NOT OK"}},
		{Name: "Login", Status: models.RouteStatus{Path: "POST /login", Status: "NOT OK"}},
		{Name: "Get user data", Status: models.RouteStatus{Path: "GET /user/get/{id}", Status: "NOT OK"}},
		{Name: "Update user data", Status: models.RouteStatus{Path: "PUT /user/update/{id}", Status: "NOT OK"}},
		{Name: "Delete user data", Status: models.RouteStatus{Path: "DELETE /user/delete/{id}", Status: "NOT OK"}},
		{Name: "Create tasks", Status: models.RouteStatus{Path: "POST /task/create", Status: "NOT OK"}},
		{Name: "Update task", Status: models.RouteStatus{Path: "PUT /task/update/{id}", Status: "NOT OK"}},
		{Name: "Get one task", Status: models.RouteStatus{Path: "GET /task/get/{id}", Status: "NOT OK"}},
		{Name: "Get all user tasks", Status: models.RouteStatus{Path: "GET /task/getUser/{id}", Status: "NOT OK"}},
		{Name: "Delete task", Status: models.RouteStatus{Path: "DELETE /task/delete/{id}", Status: "NOT OK"}},
	}
	err := json.NewEncoder(w).Encode(routes)
	if err != nil {
		log.Fatalln("There was an error encoding the initialized struct")
	}
}
