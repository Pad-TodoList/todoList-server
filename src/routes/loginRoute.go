package route

import (
	"github.com/Pad-TodoList/todoList-server/src/controllers"
	"github.com/Pad-TodoList/todoList-server/src/migrate"
	"github.com/Pad-TodoList/todoList-server/src/utils"
	"github.com/gorilla/mux"
)

func LoginRoutes(router *mux.Router, dataAccess migrate.DataAccessObject) {
	router.HandleFunc("/login", controllers.Login(dataAccess)).Methods("POST")
	router.HandleFunc("/register", controllers.Register(dataAccess)).Methods("POST")
	router.HandleFunc("/register", utils.HandleOptions).Methods("OPTIONS")
	router.HandleFunc("/login", utils.HandleOptions).Methods("OPTIONS")
}
