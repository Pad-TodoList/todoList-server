package route

import (
	"github.com/Pad-TodoList/todoList-server/src/controllers"
	"github.com/gorilla/mux"
)

func LoginRoutes(router *mux.Router) {
	router.HandleFunc("/login", controllers.HandlerLogin).Methods("POST")
	router.HandleFunc("/register", controllers.HandlerLogin).Methods("POST")
}
