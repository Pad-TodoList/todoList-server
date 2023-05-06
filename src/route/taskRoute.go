package route

import (
	"github.com/Pad-TodoList/todoList-server/src/controllers"
	"github.com/Pad-TodoList/todoList-server/src/middleware"
	"github.com/gorilla/mux"
)

func TaskRoutes(router *mux.Router) {
	taskRouter := router.PathPrefix("/task").Subrouter()
	taskRouter.Use(middleware.TaskMiddleware)
	taskRouter.HandleFunc("/create", controllers.HandlerTask).Methods("POST")
	taskRouter.HandleFunc("/update/{id}", controllers.HandlerTask).Methods("PUT")
	taskRouter.HandleFunc("/get/{id}", controllers.HandlerTask).Methods("GET")
	taskRouter.HandleFunc("/getUser/{id}", controllers.HandlerTask).Methods("GET")
	taskRouter.HandleFunc("/delete/{id}", controllers.HandlerTask).Methods("DELETE")
}
