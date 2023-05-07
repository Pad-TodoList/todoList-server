package route

import (
	"github.com/Pad-TodoList/todoList-server/src/controllers"
	"github.com/Pad-TodoList/todoList-server/src/middleware"
	"github.com/Pad-TodoList/todoList-server/src/migrate"
	"github.com/gorilla/mux"
)

func TaskRoutes(router *mux.Router, dataAccess migrate.DataAccessObject) {
	taskRouter := router.PathPrefix("/task").Subrouter()
	taskRouter.Use(middleware.TaskMiddleware)
	taskRouter.HandleFunc("/create", controllers.HandlerTask(dataAccess)).Methods("POST")
	taskRouter.HandleFunc("/update/{id}", controllers.HandlerTask(dataAccess)).Methods("PUT")
	taskRouter.HandleFunc("/get/{id}", controllers.HandlerTask(dataAccess)).Methods("GET")
	taskRouter.HandleFunc("/getUser/{id}", controllers.HandlerTask(dataAccess)).Methods("GET")
	taskRouter.HandleFunc("/delete/{id}", controllers.HandlerTask(dataAccess)).Methods("DELETE")
}
