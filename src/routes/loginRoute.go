package route

import (
	"github.com/Pad-TodoList/todoList-server/src/controllers"
	"github.com/Pad-TodoList/todoList-server/src/migrate"
	"github.com/gorilla/mux"
	"net/http"
)

func handleOptions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Optionally, you can set additional headers such as 'Access-Control-Max-Age' for caching

	w.WriteHeader(http.StatusOK)
}

func LoginRoutes(router *mux.Router, dataAccess migrate.DataAccessObject) {
	router.HandleFunc("/login", controllers.Login(dataAccess)).Methods("POST")
	router.HandleFunc("/register", controllers.Register(dataAccess)).Methods("POST")
	router.HandleFunc("/register", handleOptions).Methods("OPTIONS")
	router.HandleFunc("/login", handleOptions).Methods("OPTIONS")
}
