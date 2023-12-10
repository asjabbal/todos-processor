package model

import "fmt"

type Todo struct {
	Id        int    `json:"id"`
	UserId    int    `json:"userId"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func (todo Todo) PrintValues() bool {
	if todo.Id == 0 || len(todo.Title) == 0 {
		return false
	}
	fmt.Printf("for todo with id=%v:\n\ttitle=%v\n\tcompleted=%v\n", todo.Id, todo.Title, todo.Completed)
	return true
}
