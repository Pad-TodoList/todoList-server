package route

import (
	"github.com/Pad-TodoList/todoList-server/src/controllers"
	"github.com/Pad-TodoList/todoList-server/src/migrate"
	"github.com/Pad-TodoList/todoList-server/src/utils"
	"github.com/gorilla/mux"
)

func TaskRoutes(router *mux.Router, dataAccess migrate.DataAccessObject) {
	//taskRouter.Use(middleware.TaskMiddleware(dataAccess))
	router.HandleFunc("/createTask", controllers.CreateTask(dataAccess)).Methods("POST")
	router.HandleFunc("/updateTask", controllers.UpdateTask(dataAccess)).Methods("PUT")
	router.HandleFunc("/getTask/{id}", controllers.GetOneTask(dataAccess)).Methods("GET")
	router.HandleFunc("/getUserTasks/{id}", controllers.GetUserTask(dataAccess)).Methods("GET")
	router.HandleFunc("/deleteTask/{id}", controllers.DeleteTask(dataAccess)).Methods("DELETE")
	router.HandleFunc("/createTask", utils.HandleOptions).Methods("OPTIONS")
	router.HandleFunc("/updateTask", utils.HandleOptions).Methods("OPTIONS")
	router.HandleFunc("/getTask/{id}", utils.HandleOptions).Methods("OPTIONS")
	router.HandleFunc("/getUserTasks/{id}", utils.HandleOptions).Methods("OPTIONS")
	router.HandleFunc("/deleteTask/{id}", utils.HandleOptions).Methods("OPTIONS")
}
