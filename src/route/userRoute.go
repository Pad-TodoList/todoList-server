package route

import (
	"github.com/Pad-TodoList/todoList-server/src/controllers"
	"github.com/Pad-TodoList/todoList-server/src/middleware"
	"github.com/gorilla/mux"
)

func UserRoutes(router *mux.Router) {
	userRouter := router.PathPrefix("/user").Subrouter()
	userRouter.Use(middleware.UserMiddleware)
	userRouter.HandleFunc("/update/{id}", controllers.HandlerUser).Methods("PUT")
	userRouter.HandleFunc("/get/{id}", controllers.HandlerUser).Methods("GET")
	userRouter.HandleFunc("/delete/{id}", controllers.HandlerUser).Methods("DELETE")
}
