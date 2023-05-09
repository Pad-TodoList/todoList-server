package controllers

import (
	"encoding/json"
	"github.com/Pad-TodoList/todoList-server/src/migrate"
	"github.com/Pad-TodoList/todoList-server/src/models"
	"github.com/Pad-TodoList/todoList-server/src/utils"
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
			userSave := dataAccess.GetUser(result.Data.(models.UserToken).UserId)
			hashedPassword := utils.HashString(user.Password)
			if hashedPassword != userSave.Data.(models.IdentifiableUser).Password {
				utils.BadResponse(w, models.ErrorMessage{Code: models.Forbidden, Message: "Bad password"})
				return
			}
		}

		utils.HandleResponse(http.StatusOK, result, w)
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
		hashedPassword := utils.HashString(user.Password)
		result := dataAccess.CreateUser(models.User{
			Nickname:  user.Nickname,
			Firstname: user.Firstname,
			Lastname:  user.Lastname,
			Email:     user.Email,
			Password:  hashedPassword,
		})

		utils.HandleResponse(http.StatusCreated, result, w)
	}
}
