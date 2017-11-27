package gm

import "time"

type Todo struct {
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}
type User  struct{
	Name string  `json:"name"`
	Sex  string  `json:"sex"`
	Level int     `json:"level"`
}

type Todos []Todo
type Users []User
