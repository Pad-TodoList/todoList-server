package middleware

import (
	"github.com/Pad-TodoList/todoList-server/src/migrate"
	"net/http"
)

func TaskMiddleware(dataAccess migrate.DataAccessObject) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			result := dataAccess.IsGoodAccessToken(r.Header.Get("accessToken"))
			if !result.Status {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
