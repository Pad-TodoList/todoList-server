package migrate

import (
	"github.com/Pad-TodoList/todoList-server/src/migrate/postgreSql"
	"github.com/Pad-TodoList/todoList-server/src/models"
	"github.com/Pad-TodoList/todoList-server/src/utils"
)

type DataAccessObject interface {
	CanConnectToDatabase() models.DataAccessMessage
	CreateUser(user models.User) models.DataAccessMessage
	UpdateUser(user models.IdentifiableUser) models.DataAccessMessage
	GetUser(id string) models.DataAccessMessage
	FindUserToken(user models.User) models.DataAccessMessage
	DeleteUser(id string) models.DataAccessMessage
	CreateTask(task models.Task) models.DataAccessMessage
	UpdateTask(task models.IdentifiableTask) models.DataAccessMessage
	GetOneTask(id string) models.DataAccessMessage
	GetUserTask(id string) models.DataAccessMessage
	DeleteTask(id string) models.DataAccessMessage
	CreateAccessToken(id string) models.DataAccessMessage
	IsGoodAccessToken(accessToken string) models.DataAccessMessage
	DeleteAccessToken(id string) models.DataAccessMessage
}

func GetDataAccess() DataAccessObject {
	dataAccess := postgreSql.Database{
		Host:         utils.GetGoDotEnvVariable("DATABASE_HOST"),
		Port:         utils.GetGoDotEnvVariable("DATABASE_PORT"),
		User:         utils.GetGoDotEnvVariable("DATABASE_USER"),
		Password:     utils.GetGoDotEnvVariable("DATABASE_PASSWORD"),
		DatabaseName: utils.GetGoDotEnvVariable("DATABASE_NAME"),
	}
	return dataAccess
}
