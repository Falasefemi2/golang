package main

import (
	"fmt"
	"os"

	"example.com/todo/list"
)

func main() {
	var todo list.ToDoList

	// load exciting task from JSON file if any
	list.LoadTasks(&todo)

	for {
		fmt.Println("1. Add Task")
		fmt.Println("2. View Tasks")
		fmt.Println("3. Mark Task as Completed")
		fmt.Println("4. Remove Task")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Println("Enter task description: ")
			var desc string
			fmt.Scanln(&desc)
			list.AddTodo(&todo, desc)
		case 2:
			list.ViewTasks(&todo)
		case 3:
			list.ViewTasks(&todo)
			fmt.Print("Enter task ID to mark as completed")
			var id int
			fmt.Scanln(&id)
			list.MarkTaskCompleted(&todo, id)
		case 4:
			list.ViewTasks(&todo)
			fmt.Print("Enter task ID to remove: ")
			var id int
			fmt.Scanln(&id)
			list.RemoveTask(&todo, id)
		case 5:
			fmt.Println("Exiting...")
			list.SaveTask(&todo)
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please try again.")

		}

	}
}
