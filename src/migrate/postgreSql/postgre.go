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
	return models.DataAccessMessage{Status: true, Data: "OK"}
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
	fmt.Println("Connected!")
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
	return models.DataAccessMessage{Status: true, Data: "1"}
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
