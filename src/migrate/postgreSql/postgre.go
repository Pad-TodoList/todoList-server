package postgreSql

import (
	"database/sql"
	"fmt"
	"github.com/Pad-TodoList/todoList-server/src/models"
	"github.com/Pad-TodoList/todoList-server/src/utils"
)

type Database struct {
	Host, Port, User, Password, DatabaseName string
}

func (d Database) CreateAccessToken(id string) models.DataAccessMessage {
	token := utils.GenerateAccessToken(50)
	insertData := `insert into "usertoken"("userid", "token") values($1, $2)`
	db := connectToDatabase(d)
	if db == nil {
		return models.DataAccessMessage{Status: false, Data: models.ErrorMessage{Code: models.ConnectionDatabase, Message: "Error connection DB"}}
	}
	_, err := db.Exec(insertData, id, token)

	if err != nil {
		return models.DataAccessMessage{Status: false, Data: models.ErrorMessage{Code: models.DatabaseAction, Message: err.Error()}}
	}
	return models.DataAccessMessage{Status: true, Data: models.UserToken{UserId: id, AccessToken: token}}
}

func (d Database) IsGoodAccessToken(accessToken string) models.DataAccessMessage {
	db := connectToDatabase(d)
	if db == nil {
		return models.DataAccessMessage{Status: false, Data: models.ErrorMessage{Code: models.ConnectionDatabase, Message: "Error connection DB"}}
	}
	rows, err := db.Query("SELECT * from usertoken WHERE token = $1", accessToken)
	var length int

	if err != nil {
		return models.DataAccessMessage{Status: false, Data: models.ErrorMessage{Code: models.DatabaseAction, Message: err.Error()}}
	}
	defer rows.Close()
	for rows.Next() {
		length++
	}
	if length == 0 {
		return models.DataAccessMessage{Status: false, Data: models.ErrorMessage{Code: models.NotFound, Message: "Access token not found"}}
	} else {
		return models.DataAccessMessage{Status: true, Data: "Good access token"}
	}
}

func (d Database) DeleteAccessToken(id string) models.DataAccessMessage {
	db := connectToDatabase(d)
	if db == nil {
		return models.DataAccessMessage{Status: false, Data: models.ErrorMessage{Code: models.ConnectionDatabase, Message: "Error connection DB"}}
	}
	removeData := "delete from usertoken where userid = $1"
	_, err := db.Exec(removeData, id)

	if err != nil {
		return models.DataAccessMessage{Status: false, Data: models.ErrorMessage{Code: models.DatabaseAction, Message: err.Error()}}
	}
	return models.DataAccessMessage{Status: true, Data: "Delete OK"}
}

func connectToDatabase(d Database) *sql.DB {
	psqlConn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", d.User, d.Password, d.Host, d.Port, d.DatabaseName)
	db, err := sql.Open("postgres", psqlConn)

	if err != nil {
		return nil
	}
	return db
}

func (d Database) CanConnectToDatabase() models.DataAccessMessage {
	if connectToDatabase(d) == nil {
		return models.DataAccessMessage{Status: false, Data: models.ErrorMessage{Code: models.ConnectionDatabase, Message: "Error connection DB"}}
	}
	fmt.Println("Connected to PostGreSQL !")
	return models.DataAccessMessage{Status: true, Data: "Database connected"}
}

func (d Database) CreateUser(user models.User) models.DataAccessMessage {
	var id string
	insertData := `insert into "todolistuser"("nickname", "firstname", "lastname", "email", "password") values($1, $2, $3, $4, $5)`
	db := connectToDatabase(d)
	if db == nil {
		return models.DataAccessMessage{Status: false, Data: models.ErrorMessage{Code: models.ConnectionDatabase, Message: "Error connection DB"}}
	}
	_, err := db.Exec(insertData, user.Nickname, user.Firstname, user.Lastname, user.Email, user.Password)
	if err != nil {
		return models.DataAccessMessage{Status: false, Data: models.ErrorMessage{Code: models.Forbidden, Message: err.Error()}}
	}
	rows, err := db.Query("SELECT id from todolistuser WHERE email = $1", user.Email)
	if err != nil {
		return models.DataAccessMessage{Status: false, Data: models.ErrorMessage{Code: models.DatabaseAction, Message: err.Error()}}
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			return models.DataAccessMessage{Status: false, Data: models.ErrorMessage{Code: models.DatabaseAction, Message: err.Error()}}
		}
	}
	return d.CreateAccessToken(id)
}

func (d Database) UpdateUser(user models.IdentifiableUser) models.DataAccessMessage {
	updateData := `update "todolistuser" set "nickname"=$1, "firstname"=$2, "lastname"=$3, "email"=$4, "password"=$5 where "id"=$6`
	db := connectToDatabase(d)
	if db == nil {
		return models.DataAccessMessage{Status: false, Data: models.ErrorMessage{Code: models.ConnectionDatabase, Message: "Error connection DB"}}
	}
	_, err := db.Exec(updateData, user.Nickname, user.Firstname, user.Lastname, user.Email, user.Password, user.Id)
	if err != nil {
		return models.DataAccessMessage{Status: false, Data: models.ErrorMessage{Code: models.DatabaseAction, Message: err.Error()}}
	}
	return d.GetUser(user.Id)
}

func (d Database) GetUser(id string) models.DataAccessMessage {
	db := connectToDatabase(d)
	var user models.IdentifiableUser
	if db == nil {
		return models.DataAccessMessage{Status: false, Data: models.ErrorMessage{Code: models.ConnectionDatabase, Message: "Error connection DB"}}
	}
	rows, err := db.Query("SELECT id, nickname, firstname, lastname, email, password from todolistuser WHERE id = $1", id)

	if err != nil {
		return models.DataAccessMessage{Status: false, Data: models.ErrorMessage{Code: models.DatabaseAction, Message: err.Error()}}
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Nickname, &user.Firstname, &user.Lastname, &user.Email, &user.Password)
		if err != nil {
			return models.DataAccessMessage{Status: false, Data: models.ErrorMessage{Code: models.DatabaseAction, Message: err.Error()}}
		}
	}
	if len(user.Nickname) > 0 {
		return models.DataAccessMessage{Status: true, Data: user}
	} else {
		return models.DataAccessMessage{Status: false, Data: models.ErrorMessage{Code: models.NotFound, Message: "User not found"}}
	}
}

func (d Database) FindUserToken(user models.User) models.DataAccessMessage {
	db := connectToDatabase(d)
	if db == nil {
		return models.DataAccessMessage{Status: false, Data: models.ErrorMessage{Code: models.ConnectionDatabase, Message: "Error connection DB"}}
	}
	rows, err := db.Query("SELECT id from todolistuser WHERE email = $1", user.Email)

	if err != nil {
		return models.DataAccessMessage{Status: false, Data: models.ErrorMessage{Code: models.DatabaseAction, Message: err.Error()}}
	}
	defer rows.Close()
	var id string
	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			return models.DataAccessMessage{Status: false, Data: models.ErrorMessage{Code: models.DatabaseAction, Message: err.Error()}}
		}
	}
	rows, err = db.Query("SELECT token from usertoken WHERE userid = $1", id)

	if err != nil {
		return models.DataAccessMessage{Status: false, Data: models.ErrorMessage{Code: models.DatabaseAction, Message: err.Error()}}
	}
	defer rows.Close()
	var token string
	for rows.Next() {
		err := rows.Scan(&token)
		if err != nil {
			return models.DataAccessMessage{Status: false, Data: models.ErrorMessage{Code: models.DatabaseAction, Message: err.Error()}}
		}
	}
	if len(token) == 0 {
		return models.DataAccessMessage{Status: false, Data: models.ErrorMessage{Code: models.NotFound, Message: "User not found"}}
	}
	return models.DataAccessMessage{Status: true, Data: models.UserToken{UserId: id, AccessToken: token}}
}

func (d Database) DeleteUser(id string) models.DataAccessMessage {
	deleteData := `delete from "todolistuser" where id=$1`
	db := connectToDatabase(d)
	if db == nil {
		return models.DataAccessMessage{Status: false, Data: models.ErrorMessage{Code: models.ConnectionDatabase, Message: "Error connection DB"}}
	}
	_, err := db.Exec(deleteData, id)

	if err != nil {
		return models.DataAccessMessage{Status: false, Data: models.ErrorMessage{Code: models.DatabaseAction, Message: err.Error()}}
	}
	return d.DeleteAccessToken(id)
}

func (d Database) CreateTask(task models.Task) models.DataAccessMessage {
	var createdTask models.IdentifiableTask
	insertData := `insert into "task"("name", "description", "startdate", "enddate", "status", "userid") values($1, $2, $3, $4, $5, $6)`
	db := connectToDatabase(d)
	if db == nil {
		return models.DataAccessMessage{Status: false, Data: models.ErrorMessage{Code: models.ConnectionDatabase, Message: "Error connection DB"}}
	}
	_, err := db.Exec(insertData, task.Name, task.Description, task.StartDate, task.EndDate, task.Status, task.UserId)
	if err != nil {
		return models.DataAccessMessage{Status: false, Data: err.Error()}
	}
	rows, err := db.Query("SELECT id, name, description, startDate, endDate, status, userId from task WHERE name = $1", task.Name)

	if err != nil {
		return models.DataAccessMessage{Status: false, Data: models.ErrorMessage{Code: models.DatabaseAction, Message: err.Error()}}
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&createdTask.Id, &createdTask.Name, &createdTask.Description, &createdTask.StartDate, &createdTask.EndDate, &createdTask.Status, &createdTask.UserId)
		if err != nil {
			return models.DataAccessMessage{Status: false, Data: models.ErrorMessage{Code: models.DatabaseAction, Message: err.Error()}}
		}
	}
	return models.DataAccessMessage{Status: true, Data: createdTask}
}

func (d Database) UpdateTask(task models.IdentifiableTask) models.DataAccessMessage {
	updateData := `update "task" set "name"=$1, "description"=$2, "startdate"=$3, "enddate"=$4, "status"=$5, "userid"=$6 where "id"=$7`
	db := connectToDatabase(d)
	if db == nil {
		return models.DataAccessMessage{Status: false, Data: models.ErrorMessage{Code: models.ConnectionDatabase, Message: "Error connection DB"}}
	}
	_, err := db.Exec(updateData, task.Name, task.Description, task.StartDate, task.EndDate, task.Status, task.UserId, task.Id)

	if err != nil {
		return models.DataAccessMessage{Status: false, Data: models.ErrorMessage{Code: models.DatabaseAction, Message: err.Error()}}
	}
	return d.GetOneTask(task.Id)
}

func (d Database) GetOneTask(id string) models.DataAccessMessage {
	var task models.IdentifiableTask
	total := 0
	db := connectToDatabase(d)
	if db == nil {
		return models.DataAccessMessage{Status: false, Data: models.ErrorMessage{Code: models.ConnectionDatabase, Message: "Error connection DB"}}
	}
	rows, err := db.Query("SELECT id, name, description, startDate, endDate, status, userId from task WHERE id = $1", id)

	if err != nil {
		return models.DataAccessMessage{Status: false, Data: models.ErrorMessage{Code: models.DatabaseAction, Message: err.Error()}}
	}
	defer rows.Close()
	for rows.Next() {
		total++
		err := rows.Scan(&task.Id, &task.Name, &task.Description, &task.StartDate, &task.EndDate, &task.Status, &task.UserId)
		if err != nil {
			return models.DataAccessMessage{Status: false, Data: models.ErrorMessage{Code: models.DatabaseAction, Message: err.Error()}}
		}
	}
	if total == 0 {
		return models.DataAccessMessage{Status: false, Data: models.ErrorMessage{Code: models.NotFound, Message: "Task not found"}}
	} else {
		return models.DataAccessMessage{Status: true, Data: task}
	}
}

func (d Database) GetUserTask(id string) models.DataAccessMessage {
	var tasks []models.IdentifiableTask
	db := connectToDatabase(d)
	if db == nil {
		return models.DataAccessMessage{Status: false, Data: models.ErrorMessage{Code: models.ConnectionDatabase, Message: "Error connection DB"}}
	}
	rows, err := db.Query("SELECT id, name, description, startDate, endDate, status, userId from task WHERE userId = $1", id)

	if err != nil {
		return models.DataAccessMessage{Status: false, Data: models.ErrorMessage{Code: models.DatabaseAction, Message: err.Error()}}
	}
	defer rows.Close()

	for rows.Next() {
		var task models.IdentifiableTask
		err := rows.Scan(&task.Id, &task.Name, &task.Description, &task.StartDate, &task.EndDate, &task.Status, &task.UserId)
		tasks = append(tasks, task)
		if err != nil {
			return models.DataAccessMessage{Status: false, Data: models.ErrorMessage{Code: models.DatabaseAction, Message: err.Error()}}
		}
	}
	return models.DataAccessMessage{Status: true, Data: tasks}
}

func (d Database) DeleteTask(id string) models.DataAccessMessage {
	deleteData := `delete from "task" where id=$1`
	db := connectToDatabase(d)
	if db == nil {
		return models.DataAccessMessage{Status: false, Data: models.ErrorMessage{Code: models.ConnectionDatabase, Message: "Error connection DB"}}
	}
	_, err := db.Exec(deleteData, id)

	if err != nil {
		return models.DataAccessMessage{Status: false, Data: models.ErrorMessage{Code: models.DatabaseAction, Message: err.Error()}}
	}
	return models.DataAccessMessage{Status: true, Data: "Delete task OK"}
}
