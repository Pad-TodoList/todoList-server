package models

import "github.com/google/uuid"

type UserToken struct {
	UserId      string `json:"userId"`
	AccessToken string `json:"acessToken"`
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
	UUID uuid.UUID `json:"uuid"`
}
