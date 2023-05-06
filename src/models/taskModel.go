package models

type Task struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
	Status      string `json:"status"`
	UserId      string `json:"userId"`
}

type IdentifiableTask struct {
	Task
	Uuid string `json:"uuid"`
}
