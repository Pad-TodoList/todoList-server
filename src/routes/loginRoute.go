package route

import (
	"github.com/Pad-TodoList/todoList-server/src/controllers"
	"github.com/Pad-TodoList/todoList-server/src/migrate"
	"github.com/gorilla/mux"
)

func LoginRoutes(router *mux.Router, dataAccess migrate.DataAccessObject) {
	router.HandleFunc("/login", controllers.HandlerLogin(dataAccess)).Methods("POST")
	router.HandleFunc("/register", controllers.HandlerLogin(dataAccess)).Methods("POST")
}
