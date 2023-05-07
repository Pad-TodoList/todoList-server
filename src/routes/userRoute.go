package route

import (
	"github.com/Pad-TodoList/todoList-server/src/controllers"
	"github.com/Pad-TodoList/todoList-server/src/middleware"
	"github.com/Pad-TodoList/todoList-server/src/migrate"
	"github.com/gorilla/mux"
)

func UserRoutes(router *mux.Router, dataAccess migrate.DataAccessObject) {
	userRouter := router.PathPrefix("/user").Subrouter()
	userRouter.Use(middleware.UserMiddleware)
	userRouter.HandleFunc("/update/{id}", controllers.HandlerUser(dataAccess)).Methods("PUT")
	userRouter.HandleFunc("/get/{id}", controllers.HandlerUser(dataAccess)).Methods("GET")
	userRouter.HandleFunc("/delete/{id}", controllers.HandlerUser(dataAccess)).Methods("DELETE")
}
