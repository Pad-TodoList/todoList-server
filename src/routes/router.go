package route

import (
	"github.com/Pad-TodoList/todoList-server/src/migrate"
	"net/http"

	"github.com/Pad-TodoList/todoList-server/src/controllers"
	"github.com/gorilla/mux"
)

func Router(dataAccess migrate.DataAccessObject) http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/about.json", controllers.HandlerAbout(dataAccess)).Methods("GET")
	UserRoutes(r, dataAccess)
	TaskRoutes(r, dataAccess)
	LoginRoutes(r, dataAccess)
	return r
}
