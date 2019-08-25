package model

type Todo struct {
	ID       int `json:"id"`
	NAME       string `json:"name"`
	COMPLETED int `json:"completed"`
}

type TodoCreate struct {
	NAME       string `json:"name"`
	COMPLETED	int `json:"completed"`
}