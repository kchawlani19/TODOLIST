package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/kchawlani19/todo/todo"
)

func main() {

	tasks, err := todo.LoadTasks()
	if err != nil {
		fmt.Printf("Error loading tasks: %v\n", err)
		tasks = todo.Tasks{}
	}

	if len(os.Args) < 2 {
		fmt.Println("Please provide a command: add, list, delete, done, or exit")
		return
	}
	command, args := os.Args[1], os.Args[2:]
	arg := strings.Join(args, " ")

	switch command {
	case "add":
		if arg == "" {
			fmt.Println("Please provide task description")
		}

		todo.AddTask(tasks, arg)

	case "list":
		todo.ListTask(tasks)

	case "delete":

		todo.DeleteTask(tasks, arg)

	case "done":

		todo.Done(tasks, arg)

	case "exit":
		fmt.Println("Byebyeee")
		return

	default:
		fmt.Println("Unknown command. Please use add, list, delete, done, or exit.")
	}
}
