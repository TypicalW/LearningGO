package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var todos Todos

	if len(os.Args) < 2 {
		fmt.Println("Usage: todo <command> [arguments]")
		return
	}

	command := os.Args[1]

	switch command {

	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a task title")
			return
		}

		title := os.Args[2]
		todos.add(title)

	case "list":
		todos.list()

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a task index")
			return
		}

		index, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Index must be a number")
			return
		}

		if err := todos.delete(index); err != nil {
			fmt.Println(err)
			return
		}

	case "edit":
		if len(os.Args) < 4 {
			fmt.Println("Usage: todo edit <index> <new title>")
			return
		}

		index, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Index must be a number")
			return
		}

		title := os.Args[3]

		if err := todos.edit(index, title); err != nil {
			fmt.Println(err)
			return
		}

	case "toggle":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a task index")
			return
		}

		index, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Index must be a number")
			return
		}

		if err := todos.toggle(index); err != nil {
			fmt.Println(err)
			return
		}

	default:
		fmt.Println("Unknown command:", command)
	}
}
