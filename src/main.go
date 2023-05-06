package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"

	route "github.com/Pad-TodoList/todoList-server/src/routes"
)

func getGoDotEnvVariable(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file", err.Error())
	}
	return os.Getenv(key)
}

func main() {
	port := getGoDotEnvVariable("PORT")
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
