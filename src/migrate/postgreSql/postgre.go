package postgreSql

import (
	"database/sql"
	"fmt"
	"github.com/Pad-TodoList/todoList-server/src/models"
	"github.com/Pad-TodoList/todoList-server/src/utils"
)

type Database struct {
	Host, Port, User, Password, DatabaseName string
	Database                                 *sql.DB
}

func (d Database) CreateAccessToken(id string) models.DataAccessMessage {
	token := utils.GenerateAccessToken(50)
	insertData := `insert into "usertoken"("userid", "token") values($1, $2)`
	psqlConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", d.Host, d.Port, d.User, d.Password, d.DatabaseName)
	db, _ := sql.Open("postgres", psqlConn)
	_, err := db.Exec(insertData, id, token)

	if err != nil {
		return models.DataAccessMessage{Status: false, Data: err.Error()}
	}
	return models.DataAccessMessage{Status: true, Data: models.UserToken{UserId: id, AccessToken: token}}
}

func (d Database) GetAccessToken(accessToken string) models.DataAccessMessage {
	psqlConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", d.Host, d.Port, d.User, d.Password, d.DatabaseName)
	db, _ := sql.Open("postgres", psqlConn)
	rows, err := db.Query("SELECT * from usertoken WHERE token = $1", accessToken)
	var length int

	if err != nil {
		return models.DataAccessMessage{Status: false, Data: err.Error()}
	}
	defer rows.Close()
	for rows.Next() {
		length++
	}
	if length == 0 {
		return models.DataAccessMessage{Status: false, Data: "Bad access token"}
	} else {
		return models.DataAccessMessage{Status: true, Data: "Good access token"}
	}
}

func (d Database) DeleteAccessToken(id string) models.DataAccessMessage {
	psqlConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", d.Host, d.Port, d.User, d.Password, d.DatabaseName)
	db, _ := sql.Open("postgres", psqlConn)
	removeData := "delete from usertoken where userid = $1"
	_, err := db.Exec(removeData, id)

	if err != nil {
		return models.DataAccessMessage{Status: false, Data: err.Error()}
	}
	return models.DataAccessMessage{Status: true, Data: "Delete OK"}
}

func (d Database) Connection() models.DataAccessMessage {
	psqlConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", d.Host, d.Port, d.User, d.Password, d.DatabaseName)
	db, err := sql.Open("postgres", psqlConn)

	d.Database = db
	if err != nil {
		return models.DataAccessMessage{Status: false, Data: err.Error()}
	}
	err = d.Database.Ping()
	if err != nil {
		return models.DataAccessMessage{Status: false, Data: err.Error()}
	}
	fmt.Println("Connected to PostGreSQL !")
	return models.DataAccessMessage{Status: true, Data: "Database connected"}
}

func (d Database) CreateUser(user models.User) models.DataAccessMessage {
	var id string
	insertData := `insert into "todolistuser"("nickname", "firstname", "lastname", "email", "password") values($1, $2, $3, $4, $5)`
	psqlConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", d.Host, d.Port, d.User, d.Password, d.DatabaseName)
	db, _ := sql.Open("postgres", psqlConn)
	_, err := db.Exec(insertData, user.Nickname, user.Firstname, user.Lastname, user.Email, user.Password)
	if err != nil {
		return models.DataAccessMessage{Status: false, Data: err.Error()}
	}
	rows, err := db.Query("SELECT id from todolistuser WHERE email = $1", user.Email)
	if err != nil {
		return models.DataAccessMessage{Status: false, Data: err.Error()}
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			return models.DataAccessMessage{Status: false, Data: err.Error()}
		}
	}
	return d.CreateAccessToken(id)
}

func (d Database) UpdateUser(user models.IdentifiableUser) models.DataAccessMessage {
	updateData := `update "todolistuser" set "nickname"=$1, "firstname"=$2, "lastname"=$3, "email"=$4, "password"=$5 where "id"=$6`
	psqlConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", d.Host, d.Port, d.User, d.Password, d.DatabaseName)
	db, _ := sql.Open("postgres", psqlConn)
	_, err := db.Exec(updateData, user.Nickname, user.Firstname, user.Lastname, user.Email, user.Password, user.Id)
	if err != nil {
		return models.DataAccessMessage{Status: false, Data: err.Error()}
	}
	return d.GetUser(user.Id)
}

func (d Database) GetUser(id string) models.DataAccessMessage {
	psqlConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", d.Host, d.Port, d.User, d.Password, d.DatabaseName)
	db, _ := sql.Open("postgres", psqlConn)
	rows, err := db.Query("SELECT id, nickname, firstname, lastname, email, password from todolistuser WHERE id = $1", id)

	if err != nil {
		return models.DataAccessMessage{Status: false, Data: err.Error()}
	}
	defer rows.Close()
	var nickname string
	var firstname string
	var lastname string
	var email string
	var password string
	for rows.Next() {
		err := rows.Scan(&id, &nickname, &firstname, &lastname, &email, &password)
		if err != nil {
			fmt.Printf("error: %s\n", err.Error())
		}
	}
	if len(nickname) > 0 {
		return models.DataAccessMessage{Status: true, Data: models.IdentifiableUser{
			User: models.User{Nickname: nickname, Firstname: firstname, Lastname: lastname, Email: email, Password: password},
			Id:   id,
		}}
	} else {
		return models.DataAccessMessage{Status: false, Data: "Not found"}
	}
}

func (d Database) FindUser(user models.User) models.DataAccessMessage {
	psqlConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", d.Host, d.Port, d.User, d.Password, d.DatabaseName)
	db, _ := sql.Open("postgres", psqlConn)
	rows, err := db.Query("SELECT id from todolistuser WHERE email = $1", user.Email)

	if err != nil {
		return models.DataAccessMessage{Status: false, Data: err.Error()}
	}
	defer rows.Close()
	var id string
	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			return models.DataAccessMessage{Status: false, Data: err.Error()}
		}
	}
	rows, err = db.Query("SELECT token from usertoken WHERE userid = $1", id)

	if err != nil {
		return models.DataAccessMessage{Status: false, Data: err.Error()}
	}
	defer rows.Close()
	var token string
	for rows.Next() {
		err := rows.Scan(&token)
		if err != nil {
			return models.DataAccessMessage{Status: false, Data: err.Error()}
		}
	}
	return models.DataAccessMessage{Status: false, Data: models.UserToken{UserId: id, AccessToken: token}}
}

func (d Database) DeleteUser(id string) models.DataAccessMessage {
	deleteData := `delete from "todolistuser" where id=$1`
	psqlConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", d.Host, d.Port, d.User, d.Password, d.DatabaseName)
	db, _ := sql.Open("postgres", psqlConn)
	_, err := db.Exec(deleteData, id)
	if err != nil {
		return models.DataAccessMessage{Status: false, Data: err.Error()}
	}
	return d.DeleteAccessToken(id)
}

func (d Database) CreateTask(task models.Task) models.DataAccessMessage {
	var createdTask models.IdentifiableTask
	insertData := `insert into "task"("name", "description", "startdate", "enddate", "status", "userid") values($1, $2, $3, $4, $5, $6)`
	psqlConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", d.Host, d.Port, d.User, d.Password, d.DatabaseName)
	db, _ := sql.Open("postgres", psqlConn)
	_, err := db.Exec(insertData, task.Name, task.Description, task.StartDate, task.EndDate, task.Status, task.UserId)
	if err != nil {
		return models.DataAccessMessage{Status: false, Data: err.Error()}
	}
	rows, err := db.Query("SELECT id, name, description, startDate, endDate, status, userId from task WHERE name = $1", task.Name)
	if err != nil {
		return models.DataAccessMessage{Status: false, Data: err.Error()}
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&createdTask.Id, &createdTask.Name, &createdTask.Description, &createdTask.StartDate, &createdTask.EndDate, &createdTask.Status, &createdTask.UserId)
		if err != nil {
			return models.DataAccessMessage{Status: false, Data: err.Error()}
		}
	}
	return models.DataAccessMessage{Status: true, Data: createdTask}
}

func (d Database) UpdateTask(task models.IdentifiableTask) models.DataAccessMessage {
	updateData := `update "task" set "name"=$1, "description"=$2, "startdate"=$3, "enddate"=$4, "status"=$5, "userid"=$6 where "id"=$7`
	psqlConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", d.Host, d.Port, d.User, d.Password, d.DatabaseName)
	db, _ := sql.Open("postgres", psqlConn)
	_, err := db.Exec(updateData, task.Name, task.Description, task.StartDate, task.EndDate, task.Status, task.UserId, task.Id)
	if err != nil {
		return models.DataAccessMessage{Status: false, Data: err.Error()}
	}
	return d.GetOneTask(task.Id)
}

func (d Database) GetOneTask(id string) models.DataAccessMessage {
	var task models.IdentifiableTask
	psqlConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", d.Host, d.Port, d.User, d.Password, d.DatabaseName)
	db, _ := sql.Open("postgres", psqlConn)
	rows, err := db.Query("SELECT id, name, description, startDate, endDate, status, userId from task WHERE id = $1", id)
	if err != nil {
		return models.DataAccessMessage{Status: false, Data: err.Error()}
	}
	defer rows.Close()
	total := 0
	for rows.Next() {
		total++
		err := rows.Scan(&task.Id, &task.Name, &task.Description, &task.StartDate, &task.EndDate, &task.Status, &task.UserId)
		if err != nil {
			return models.DataAccessMessage{Status: false, Data: err.Error()}
		}
	}
	if total == 0 {
		return models.DataAccessMessage{Status: false, Data: "Not found"}
	} else {
		return models.DataAccessMessage{Status: true, Data: task}
	}
}

func (d Database) GetUserTask(id string) models.DataAccessMessage {
	var tasks []models.IdentifiableTask
	psqlConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", d.Host, d.Port, d.User, d.Password, d.DatabaseName)
	db, _ := sql.Open("postgres", psqlConn)
	rows, err := db.Query("SELECT id, name, description, startDate, endDate, status, userId from task WHERE userId = $1", id)
	if err != nil {
		return models.DataAccessMessage{Status: false, Data: err.Error()}
	}
	defer rows.Close()
	total := 0
	for rows.Next() {
		total++
		var task models.IdentifiableTask
		err := rows.Scan(&task.Id, &task.Name, &task.Description, &task.StartDate, &task.EndDate, &task.Status, &task.UserId)
		tasks = append(tasks, task)
		if err != nil {
			return models.DataAccessMessage{Status: false, Data: err.Error()}
		}
	}
	if total == 0 {
		return models.DataAccessMessage{Status: false, Data: "Not found"}
	} else {
		return models.DataAccessMessage{Status: true, Data: tasks}
	}
}

func (d Database) DeleteTask(id string) models.DataAccessMessage {
	deleteData := `delete from "task" where id=$1`
	psqlConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", d.Host, d.Port, d.User, d.Password, d.DatabaseName)
	db, _ := sql.Open("postgres", psqlConn)
	_, err := db.Exec(deleteData, id)
	if err != nil {
		return models.DataAccessMessage{Status: false, Data: err.Error()}
	}
	return models.DataAccessMessage{Status: true, Data: "Delete task OK"}
}
