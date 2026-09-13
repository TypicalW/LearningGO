package main

import (
	"errors"
	"fmt"
	"time"
)

// Created a data-type holding bluerprint for To-do objects
type Todo struct {
	Title     string
	Status    bool
	StartTime time.Time
	EndTime   *time.Time
}

// created a data-type with type being []Todo which is a slice
// Todos is a datatype holding Todo structs
type Todos []Todo

func (todos *Todos) add(title string) {
	todo := Todo{
		Title:     title,
		Status:    false,
		StartTime: time.Now(),
		EndTime:   nil,
	}

	*todos = append(*todos, todo)
}

func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("Invalid index bozo")
		fmt.Println(err)
	}
	return nil
}

func (todos *Todos) delete(index int) error {
	t := *todos
	if err := t.validateIndex(index); err != nil {
		return err
	}
	*todos = append(t[:index], t[index+1:]...)
	return nil
}
