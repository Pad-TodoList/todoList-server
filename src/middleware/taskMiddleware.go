package middleware

import (
	"fmt"
	"github.com/Pad-TodoList/todoList-server/src/migrate"
	"net/http"
)

func TaskMiddleware(dataAccess migrate.DataAccessObject) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			//if r.URL.Path == "/about.json" {
			//	fmt.Println("Request received for /about.json")
			//}
			fmt.Println(r.URL.Path, "task middleware")
			next.ServeHTTP(w, r)
		})
	}
}
