package model

type Task struct {
	Id        int    `json:"id"`
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}
