package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
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

func (todos *Todos) toggle(index int) error {
	if err := todos.validateIndex(index); err != nil {
		return err
	}
	t := *todos
	todo := &t[index]

	if !todo.Status {
		completedTime := time.Now()
		todo.EndTime = &completedTime
	} else {
		todo.EndTime = nil
	}
	todo.Status = !todo.Status
	return nil

}

func (todos *Todos) edit(index int, title string) error {
	if err := todos.validateIndex(index); err != nil {
		return err
	}
	(*todos)[index].Title = title
	return nil
}

func (todos *Todos) list() {
	table := table.New(os.Stdout)
	table.SetRowLines(true)
	table.SetHeaders("#", "Task", "Status", "Started at", "Finished at")

	for index, t := range *todos {
		completed := "X"
		completedAt := ""

		if t.Status {
			completed = "DONE"
			if t.EndTime != nil {
				completedAt = t.EndTime.Format(time.RFC1123)
			}
		}
		table.AddRow(strconv.Itoa(index), t.Title, completed, t.StartTime.Format(time.RFC1123), completedAt)

	}
	table.Render()
}
