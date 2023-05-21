package route

import (
	"github.com/Pad-TodoList/todoList-server/src/controllers"
	"github.com/Pad-TodoList/todoList-server/src/middleware"
	"github.com/Pad-TodoList/todoList-server/src/migrate"
	"github.com/Pad-TodoList/todoList-server/src/utils"
	"github.com/gorilla/mux"
)

func UserRoutes(router *mux.Router, dataAccess migrate.DataAccessObject) {
	userRouter := router.PathPrefix("/user").Subrouter()
	userRouter.Use(middleware.UserMiddleware(dataAccess))
	userRouter.HandleFunc("/update", controllers.UpdateUser(dataAccess)).Methods("PUT")
	userRouter.HandleFunc("/get/{id}", controllers.GetUser(dataAccess)).Methods("GET")
	userRouter.HandleFunc("/delete/{id}", controllers.DeleteUser(dataAccess)).Methods("DELETE")
	userRouter.HandleFunc("/update", utils.HandleOptions).Methods("OPTIONS")
	userRouter.HandleFunc("/get/{id}", utils.HandleOptions).Methods("OPTIONS")
	userRouter.HandleFunc("/delete/{id}", utils.HandleOptions).Methods("OPTIONS")
}
