package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Task struct {
	ID          int        `json:"id"`
	Description string     `json:"description"`
	Status      bool       `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	CompletedAt *time.Time `json:"completed_at"`
}

func (t Task) String() string {
	return fmt.Sprintf("ID: %d, Desc: %v, Status: %v, Created: %v, Completed: %v",
		t.ID, t.Description, t.Status, t.CreatedAt, t.CompletedAt)
}

type Tasks []Task

const filename = "newtasks.json"

func main() {
	fmt.Println("Welcome to my todie-godie todo list")
	fmt.Println("kya kya milega: add <task>, list, delete <id>, done <id>, exit")

	tasks, err := loadTasks()
	if err != nil {
		fmt.Printf("Error loading tasks: %v\n", err)
		tasks = Tasks{}
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		line = strings.TrimSpace(line)
		parts := strings.SplitN(line, " ", 2)
		if len(parts) > 1 {
			parts[1] = strings.TrimSpace(parts[1])
		}
		command := strings.TrimSpace(parts[0])

		switch command {
		case "add":
			if len(parts) < 2 || parts[1] == "" {
				fmt.Println("Please provide task description")
				continue
			}

			newID := 1
			if len(tasks) > 0 {
				newID = tasks[len(tasks)-1].ID + 1
			}

			t := Task{ID: newID, Description: parts[1], CreatedAt: time.Now()}
			tasks = append(tasks, t)
			fmt.Println("Task added")

			// Print JSON representation
			jsonBytes, err := json.Marshal(t)
			if err != nil {
				fmt.Printf("Error marshaling task: %v\n", err)
				continue
			}
			fmt.Println("Task JSON:", string(jsonBytes))

			// Save to file
			if err := saveTasks(tasks); err != nil {
				fmt.Printf("Error saving tasks: %v\n", err)
			}

		case "list":
			if len(tasks) == 0 {
				fmt.Println("No tasks available.")
			} else {
				for _, task := range tasks {
					jsonBytes, err := json.Marshal(task)
					if err != nil {
						fmt.Printf("Error marshaling task %d: %v\n", task.ID, err)
						continue
					}
					fmt.Printf(" %s (JSON: %s)\n", task.String(), string(jsonBytes))
				}
			}

		case "delete":
			if len(parts) < 2 {
				fmt.Println("Please provide ID to delete")
				continue
			}
			id, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("Invalid ID")
				continue
			}

			found := false
			for i, task := range tasks {
				if task.ID == id {
					tasks = append(tasks[:i], tasks[i+1:]...)
					found = true
					break
				}
			}

			if !found {
				fmt.Println("Task with this ID not found")
				continue
			}

			fmt.Println("Task deleted")

			if err := saveTasks(tasks); err != nil {
				fmt.Printf("Error saving tasks: %v\n", err)
			}

		case "done":
			if len(parts) < 2 {
				fmt.Println("Please provide valid ID")
				continue
			}
			id, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("Invalid ID")
				continue
			}

			found := false
			for i, task := range tasks {
				if task.ID == id {
					completedTime := time.Now()
					tasks[i].Status = true
					tasks[i].CompletedAt = &completedTime
					found = true

					// Print updated task JSON
					jsonBytes, err := json.Marshal(tasks[i])
					if err != nil {
						fmt.Printf("Error marshaling task: %v\n", err)
						continue
					}
					fmt.Printf("Task %d marked as done\n", id)
					fmt.Println("Updated task JSON:", string(jsonBytes))
					break
				}
			}

			if !found {
				fmt.Println("Task with this ID not found")
				continue
			}

			if err := saveTasks(tasks); err != nil {
				fmt.Printf("Error saving tasks: %v\n", err)
			}

		case "exit":
			fmt.Println("lalallalalbybyebye")
			return

		default:
			fmt.Println("Unknown command. Please use add, list, delete, done, or exit.")
		}
	}
}

// loadTasks reads tasks from the JSON file
func loadTasks() (Tasks, error) {
	file, err := os.Open(filename)
	if os.IsNotExist(err) {
		return Tasks{}, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	var tasks Tasks
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tasks); err != nil {
		return nil, fmt.Errorf("failed to decode JSON: %v", err)
	}

	return tasks, nil
}

// saveTasks writes tasks to the JSON file
func saveTasks(tasks Tasks) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Pretty print JSON
	if err := encoder.Encode(tasks); err != nil {
		return fmt.Errorf("failed to encode JSON: %v", err)
	}

	return nil
}
