package utils

import "net/http"

func HandleOptions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Optionally, you can set additional headers such as 'Access-Control-Max-Age' for caching

	w.WriteHeader(http.StatusOK)
}
