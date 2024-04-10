package list

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

type ToDoList struct {
	Tasks []Task `json:"tasks"`
}

func AddTodo(todo *ToDoList, desc string) {
	taskID := len(todo.Tasks) + 1
	newTask := Task{ID: taskID, Description: desc, Completed: false}
	todo.Tasks = append(todo.Tasks, newTask)
	fmt.Println("Task added successfully")
}

func SaveTask(todo *ToDoList) error {
	file, err := json.MarshalIndent(todo, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile("tasks.json", file, 0644)
	if err != nil {
		return err
	}
	fmt.Println("Tasks saved successfully")
	return nil
}

func LoadTasks(todo *ToDoList) {
	file, err := os.ReadFile("tasks.json")
	if err != nil {
		return
	}
	json.Unmarshal(file, todo)
}

func ViewTasks(todo *ToDoList) {
	fmt.Println("Tasks:")
	for _, task := range todo.Tasks {
		status := "Incomplete"
		if task.Completed {
			status = "Completed"
		}
		fmt.Printf("%d. %s - %s\n", task.ID, task.Description, status)
	}
}

func MarkTaskCompleted(todo *ToDoList, id int) {
	for i, task := range todo.Tasks {
		if task.ID == id {
			todo.Tasks[i].Completed = true
			fmt.Println("Task market as completed.")
			return
		}
	}
	fmt.Println("Task not found")
}

func RemoveTask(todo *ToDoList, id int) {
	for i, task := range todo.Tasks {
		if task.ID == id {
			todo.Tasks = append(todo.Tasks[:i], todo.Tasks[i+1:]...)
			fmt.Println("Task removed successfully.")
			return
		}
	}
	fmt.Println("Task not found.")

}
