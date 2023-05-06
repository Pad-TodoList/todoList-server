package models

type DataAccessMessage struct {
	Status bool `json:"status"`
	Data   any  `json:"data"`
}
