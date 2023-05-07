package postgreSql

import (
	"database/sql"
	"fmt"
	"github.com/Pad-TodoList/todoList-server/src/models"
)

type Database struct {
	Host, Port, User, Password, DatabaseName string
	Database                                 *sql.DB
}

func (d Database) Connection() models.DataAccessMessage {
	psqlConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", d.Host, d.Port, d.User, d.Password, d.DatabaseName)
	db, err := sql.Open("postgres", psqlConn)
	if err != nil {
		return models.DataAccessMessage{Status: false, Data: err.Error()}
	}
	defer db.Close()
	d.Database = db
	err = d.Database.Ping()
	if err != nil {
		return models.DataAccessMessage{Status: false, Data: err.Error()}
	}
	fmt.Println("Connected!")
	return models.DataAccessMessage{Status: true, Data: "Database connected"}
}

func (d Database) CreateUser(user models.User) models.DataAccessMessage {
	return models.DataAccessMessage{Status: true, Data: "1"}
}

func (d Database) UpdateUser(user models.IdentifiableUser) models.DataAccessMessage {
	return models.DataAccessMessage{Status: true, Data: "1"}
}

func (d Database) GetUser(id string) models.DataAccessMessage {
	return models.DataAccessMessage{Status: true, Data: "1"}
}

func (d Database) DeleteUser(id string) models.DataAccessMessage {
	return models.DataAccessMessage{Status: true, Data: "1"}
}

func (d Database) CreateTask(task models.Task) models.DataAccessMessage {
	return models.DataAccessMessage{Status: true, Data: "1"}
}

func (d Database) UpdateTask(task models.IdentifiableTask) models.DataAccessMessage {
	return models.DataAccessMessage{Status: true, Data: "1"}
}

func (d Database) GetOneTask(id string) models.DataAccessMessage {
	return models.DataAccessMessage{Status: true, Data: "1"}
}

func (d Database) GetUserTask(id string) models.DataAccessMessage {
	return models.DataAccessMessage{Status: true, Data: "1"}
}

func (d Database) DeleteTask(id string) models.DataAccessMessage {
	return models.DataAccessMessage{Status: true, Data: "1"}
}
