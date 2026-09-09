package main

import (
	"time"
)

type Todo struct {
	Title      string
	Status     bool
	StartTime  time.Time
	EndTime    *time.Time
	Completion uint8
}

type Todos []Todo

func (todos *Todos) add(title string) {
	todo := Todo{
		Title:      title,
		Status:     false,
		StartTime:  time.Now(),
		EndTime:    nil,
		Completion: 0,
	}

	*todos = append(*todos, todo)
}
