package models

type UserToken struct {
	UserId      string `json:"userId"`
	AccessToken string `json:"accessToken"`
}

type User struct {
	Nickname  string `json:"name"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type IdentifiableUser struct {
	User
	Uuid string `json:"uuid"`
}
