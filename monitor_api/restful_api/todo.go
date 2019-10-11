package restful_api

import "time"

//Todo ...
type Todo struct {
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

//Todos is ...
type Todos []Todo
