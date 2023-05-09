package controllers

import (
	"encoding/json"
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
		result := dataAccess.FindUserToken(user)
		if result.Status {
			utils.GoodResponse(w, http.StatusOK, result.Data)
		} else {
			_, ok := result.Data.(models.ErrorMessage)
			if !ok {
				log.Fatalln("Error with error message")
			} else {
				utils.BadResponse(w, result.Data.(models.ErrorMessage))
			}
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
		if result.Status {
			utils.GoodResponse(w, http.StatusOK, result.Data)
		} else {
			_, ok := result.Data.(models.ErrorMessage)
			if !ok {
				log.Fatalln("Error with error message")
			} else {
				utils.BadResponse(w, result.Data.(models.ErrorMessage))
			}
		}
	}
}
