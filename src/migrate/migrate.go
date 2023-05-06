package migrate

import (
	"github.com/Pad-TodoList/todoList-server/src/migrate/postgreSql"
	"github.com/Pad-TodoList/todoList-server/src/models"
)

type DataAccessObject interface {
	connection() models.DataAccessMessage
	createUser(user models.User) models.DataAccessMessage
	updateUser(user models.IdentifiableUser) models.DataAccessMessage
	getUser(id string) models.DataAccessMessage
	deleteUser(id string) models.DataAccessMessage
	createTask(task models.Task) models.DataAccessMessage
	updateTask(task models.IdentifiableTask) models.DataAccessMessage
	getOneTask(id string) models.DataAccessMessage
	getUserTask(id string) models.DataAccessMessage
	deleteTask(id string) models.DataAccessMessage
}

type database struct {
	host, port, user, password, databaseName string
}

func (p database) connection() models.DataAccessMessage {
	return postgreSql.Connection()
}

func (p database) createUser(user models.User) models.DataAccessMessage {
	return postgreSql.CreateUser(user)
}

func (p database) updateUser(user models.IdentifiableUser) models.DataAccessMessage {
	return postgreSql.UpdateUser(user)
}

func (p database) getUser(id string) models.DataAccessMessage {
	return postgreSql.GetUser(id)
}

func (p database) deleteUser(id string) models.DataAccessMessage {
	return postgreSql.DeleteUser(id)
}

func (p database) createTask(task models.Task) models.DataAccessMessage {
	return postgreSql.CreateTask(task)
}

func (p database) updateTask(task models.IdentifiableTask) models.DataAccessMessage {
	return postgreSql.UpdateTask(task)
}

func (p database) getOneTask(id string) models.DataAccessMessage {
	return postgreSql.GetOneTask(id)
}

func (p database) getUserTask(id string) models.DataAccessMessage {
	return postgreSql.GetUserTask(id)
}

func (p database) deleteTask(id string) models.DataAccessMessage {
	return postgreSql.DeleteTask(id)
}

func getDataAccess() DataAccessObject {
	dataAccess := database{host: "", port: "", user: "", password: "", databaseName: ""}
	return dataAccess
}
