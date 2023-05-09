package route

import (
	"github.com/Pad-TodoList/todoList-server/src/controllers"
	"github.com/Pad-TodoList/todoList-server/src/middleware"
	"github.com/Pad-TodoList/todoList-server/src/migrate"
	"github.com/gorilla/mux"
)

func TaskRoutes(router *mux.Router, dataAccess migrate.DataAccessObject) {
	taskRouter := router.PathPrefix("/task").Subrouter()
	taskRouter.Use(middleware.TaskMiddleware(dataAccess))
	taskRouter.HandleFunc("/create", controllers.CreateTask(dataAccess)).Methods("POST")
	taskRouter.HandleFunc("/update", controllers.UpdateTask(dataAccess)).Methods("PUT")
	taskRouter.HandleFunc("/get/{id}", controllers.GetOneTask(dataAccess)).Methods("GET")
	taskRouter.HandleFunc("/getUser/{id}", controllers.GetUserTask(dataAccess)).Methods("GET")
	taskRouter.HandleFunc("/delete/{id}", controllers.DeleteTask(dataAccess)).Methods("DELETE")
}
