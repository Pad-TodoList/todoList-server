package main

import (
	"fmt"
	"github.com/Pad-TodoList/todoList-server/src/migrate"
	route "github.com/Pad-TodoList/todoList-server/src/routes"
	"github.com/Pad-TodoList/todoList-server/src/utils"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"time"
)

func main() {
	dataAccess := migrate.GetDataAccess()
	output := dataAccess.CanConnectToDatabase()
	if !output.Status {
		fmt.Printf("Error with database : %s\n", output.Data)
	}
	port := utils.GetGoDotEnvVariable("PORT")
	fmt.Printf("server running on port %s\n", port)
	router := route.Router(dataAccess)
	srv := &http.Server{
		Handler:      router,
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
