package main

import (
	"fmt"
	"github.com/Pad-TodoList/todoList-server/src/utils"
	"log"
	"net/http"
	"time"

	route "github.com/Pad-TodoList/todoList-server/src/routes"
)

func main() {
	port := utils.GetGoDotEnvVariable("PORT")
	fmt.Printf("server running on port %s\n", port)
	router := route.Router()
	srv := &http.Server{
		Handler:      router,
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
