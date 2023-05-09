package models

const (
	ConnectionDatabase = 0
	DatabaseAction     = 1
	NotFound           = 2
	Forbidden          = 3
)

type ErrorMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
