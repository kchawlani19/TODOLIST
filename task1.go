package main

import (
	"fmt"
	"strings"
)

type Task struct {
	Description string
}

var tasks []Task

func main() {
	var command string

	fmt.Println("Welcome to my todie-godie todo list")
	fmt.Println("kya kya milega: add <task>, list, delete <index>, dafahojao")

	for {
		fmt.Printf("$ ")
		fmt.Scanln(&command)

		parts := strings.SplitN(command, " ", 2)

		switch parts[0] {
		case "add":
			if len(parts) < 2 {
				fmt.Println("Please provide task desription")
				continue
			}
			tasks = append(tasks, Task{Description: parts[1]})
			fmt.Println("Task added")

		case "List":
			if len(tasks) == 0 {
				fmt.Println("No tasks available.")

			} else {
				for i, task := range tasks {
					fmt.Printf("%d. %s\n", i, &task.Description)
				}
			}

		case "delete":
			if len(parts) < 2 {
				fmt.Println("Please provide index to delete")
				continue
			}
			var index int
			_, err := fmt.Scanf(parts[1], "%d", &index)
			if err != nil || index < 0 || index >= len(tasks) {
				fmt.Println("invalid index. ")
				continue
			}

			tasks = append(tasks[:index], tasks[index+1:]...)
			fmt.Println("Task deleted")

		case "exit":
			fmt.Println("lalallalalbybyebye")
			return

		default:
			fmt.Println("unknown command. Please use add, list, delete, or exit.")
		}
	}
}
