package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	route "github.com/Pad-TodoList/todoList-server/src/routes"
)

func main() {
	port := 8080
	fmt.Printf("server running on port %d\n", port)
	router := route.Router()
	srv := &http.Server{
		Handler:      router,
		Addr:         ":" + strconv.Itoa(port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
