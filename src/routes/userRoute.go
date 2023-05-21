package route

import (
	"github.com/Pad-TodoList/todoList-server/src/controllers"
	"github.com/Pad-TodoList/todoList-server/src/migrate"
	"github.com/Pad-TodoList/todoList-server/src/utils"
	"github.com/gorilla/mux"
)

func UserRoutes(router *mux.Router, dataAccess migrate.DataAccessObject) {
	//userRouter.Use(middleware.UserMiddleware(dataAccess))
	router.HandleFunc("/updateUser", controllers.UpdateUser(dataAccess)).Methods("PUT")
	router.HandleFunc("/getUser/{id}", controllers.GetUser(dataAccess)).Methods("GET")
	router.HandleFunc("/deleteUser/{id}", controllers.DeleteUser(dataAccess)).Methods("DELETE")
	router.HandleFunc("/updateUser", utils.HandleOptions).Methods("OPTIONS")
	router.HandleFunc("/getUser/{id}", utils.HandleOptions).Methods("OPTIONS")
	router.HandleFunc("/deleteUser/{id}", utils.HandleOptions).Methods("OPTIONS")
}
