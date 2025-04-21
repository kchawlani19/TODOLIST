package todo

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
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

const filename = "tasks.json"

func AddTask(tasks Tasks, description string) error {
	newID := 1
	if len(tasks) > 0 {
		newID = tasks[len(tasks)-1].ID + 1
	}

	t := Task{ID: newID, Description: description, CreatedAt: time.Now()}
	tasks = append(tasks, t)
	fmt.Println("Task added successfully")

	if err := saveTasks(tasks); err != nil {
		return fmt.Errorf("saveTasks returned error: %w", err)
	}
	return nil
}

func ListTask(tasks Tasks) error {

	if len(tasks) == 0 {
		fmt.Println("No tasks available.")
	} else {
		for _, task := range tasks {
			jsonBytes, err := json.Marshal(task)
			if err != nil {
				fmt.Printf("Error marshaling task %d: %v\n", task.ID, err)

			}
			fmt.Printf(" %s\n", string(jsonBytes))
		}
	}
	return nil
}

func DeleteTask(tasks Tasks, arg string) error {

	id, err := strconv.Atoi(arg)
	if err != nil {
		fmt.Printf("Invalid ID: %w \n", err)
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

	}

	fmt.Println("Task deleted")

	if err := saveTasks(tasks); err != nil {
		fmt.Printf("Error saving tasks: %v\n", err)
	}
	return nil
}

func Done(tasks Tasks, arg string) {
	id, err := strconv.Atoi(arg)
	if err != nil {
		fmt.Printf("Invalid ID: %w \n", err)
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

			}
			fmt.Printf("Task %d marked as done\n", id)
			fmt.Println("Updated task JSON:", string(jsonBytes))
			break
		}
	}

	if !found {
		fmt.Println("Task with this ID not found")

	}

	if err := saveTasks(tasks); err != nil {
		fmt.Printf("Error saving tasks: %v\n", err)
	}
}

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

// LoadTasks loads all the tasks from a configured file
func LoadTasks() (Tasks, error) {
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
