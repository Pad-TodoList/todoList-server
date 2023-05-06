package postgreSql

import "github.com/Pad-TodoList/todoList-server/src/models"

func Connection() models.DataAccessMessage {
	return models.DataAccessMessage{Status: true, Data: 1}
}

func CreateUser(user models.User) models.DataAccessMessage {
	return models.DataAccessMessage{Status: true, Data: 1}
}

func UpdateUser(user models.IdentifiableUser) models.DataAccessMessage {
	return models.DataAccessMessage{Status: true, Data: 1}
}

func GetUser(id string) models.DataAccessMessage {
	return models.DataAccessMessage{Status: true, Data: 1}
}

func DeleteUser(id string) models.DataAccessMessage {
	return models.DataAccessMessage{Status: true, Data: 1}
}

func CreateTask(task models.Task) models.DataAccessMessage {
	return models.DataAccessMessage{Status: true, Data: 1}
}

func UpdateTask(task models.IdentifiableTask) models.DataAccessMessage {
	return models.DataAccessMessage{Status: true, Data: 1}
}

func GetOneTask(id string) models.DataAccessMessage {
	return models.DataAccessMessage{Status: true, Data: 1}
}

func GetUserTask(id string) models.DataAccessMessage {
	return models.DataAccessMessage{Status: true, Data: 1}
}

func DeleteTask(id string) models.DataAccessMessage {
	return models.DataAccessMessage{Status: true, Data: 1}
}
