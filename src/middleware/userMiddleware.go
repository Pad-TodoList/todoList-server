package middleware

import (
	"fmt"
	"net/http"
)

func UserMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//if r.URL.Path == "/about.json" {
		//	fmt.Println("Request received for /about.json")
		//}
		fmt.Println(r.URL.Path, "user middleware")
		next.ServeHTTP(w, r)
	})
}
