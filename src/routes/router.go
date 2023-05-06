package route

import (
	"net/http"

	"github.com/Pad-TodoList/todoList-server/src/controllers"
	"github.com/gorilla/mux"
)

func Router() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/about.json", controllers.HandlerAbout).Methods("GET")
	LoginRoutes(r)
	UserRoutes(r)
	TaskRoutes(r)
	return r
}
