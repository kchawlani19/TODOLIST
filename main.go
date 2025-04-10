package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	Description string
}

var tasks []Task

func main() {

	fmt.Println("Welcome to my todie-godie todo list")
	fmt.Println("kya kya milega: add <task>, list, delete <index>, exit")

	reader := bufio.NewReader(os.Stdin)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		parts := strings.SplitN(line, " ", 2)
		last := len(parts) - 1
		parts[last] = strings.TrimSuffix(parts[last], "\n")

		switch parts[0] {
		case "add":
			if len(parts) < 2 {
				fmt.Println("Please provide task desription")
				continue
			}
			tasks = append(tasks, Task{Description: parts[1]})
			fmt.Println("Task added")

		case "list":
			if len(tasks) == 0 {
				fmt.Println("No tasks available.")

			} else {
				for i, task := range tasks {
					fmt.Printf(" %d. %s\n", i, task.Description)
				}
			}

		case "delete":
			if len(parts) < 2 {
				fmt.Println("Please provide index to delete")
				continue
			}
			i := parts[len(parts)-1]
			num, err := strconv.ParseInt(i, 10, 64)

			if err != nil || num < 0 || int(num) >= len(tasks) {
				fmt.Println("invalid index. ")
				continue
			}

			tasks = append(tasks[:num], tasks[num+1:]...)
			fmt.Println("Task deleted")

		case "exit":
			fmt.Println("lalallalalbybyebye")
			return

		default:
			fmt.Println("unknown command. Please use add, list, delete, or exit.")
		}
	}
}
