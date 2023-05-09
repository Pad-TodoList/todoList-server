package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/Pad-TodoList/todoList-server/src/migrate"
	"github.com/Pad-TodoList/todoList-server/src/models"
	"github.com/Pad-TodoList/todoList-server/src/utils"
	"log"
	"net/http"
)

func Login(dataAccess migrate.DataAccessObject) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		result := dataAccess.FindUser(user)
		if !result.Status {
			fmt.Printf("error database : %s\n", result.Data)
		}
		w.Header().Set("Content-Type", "application/json")
		if !result.Status {
			w.WriteHeader(http.StatusForbidden)
		} else {
			w.WriteHeader(http.StatusOK)
		}
		err = json.NewEncoder(w).Encode(result.Data)
		if err != nil {
			log.Fatalln("There was an error encoding the initialized struct")
		}
	}
}

func Register(dataAccess migrate.DataAccessObject) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		hashedPassword, err := utils.HashString(user.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		result := dataAccess.CreateUser(models.User{
			Nickname:  user.Nickname,
			Firstname: user.Firstname,
			Lastname:  user.Lastname,
			Email:     user.Email,
			Password:  hashedPassword,
		})
		if !result.Status {
			fmt.Printf("error database : %s\n", result.Data)
		}
		w.Header().Set("Content-Type", "application/json")
		if !result.Status {
			w.WriteHeader(http.StatusForbidden)
		} else {
			w.WriteHeader(http.StatusCreated)
		}
		err = json.NewEncoder(w).Encode(result.Data)
		if err != nil {
			log.Fatalln("There was an error encoding the initialized struct")
		}
	}
}
