package main

import (
	"bufio"
	"fmt"
	"os"
)

// Task represents a todo item
type Task struct {
	Description string
	Completed   bool
}

// TodoList manages a collection of tasks
type TodoList struct {
	Tasks []Task
}

// AddTask adds a new task to the list
func (t *TodoList) AddTask(description string) {
	t.Tasks = append(t.Tasks, Task{Description: description, Completed: false})
	fmt.Println("Task added successfully!")
}

// ListTasks displays all tasks
func (t *TodoList) ListTasks() {
	if len(t.Tasks) == 0 {
		fmt.Println("No tasks in your list!")
		return
	}

	fmt.Println("Your Todo List:")
	for i, task := range t.Tasks {
		status := " "
		if task.Completed {
			status = "✓"
		}
		fmt.Printf("%d. [%s] %s\n", i+1, status, task.Description)
	}
}

// CompleteTask marks a task as completed
func (t *TodoList) CompleteTask(index int) {
	if index < 0 || index >= len(t.Tasks) {
		fmt.Println("Invalid task number!")
		return
	}
	t.Tasks[index].Completed = true
	fmt.Printf("Task '%s' marked as completed!\n", t.Tasks[index].Description)
}

// RemoveTask removes a task from the list
func (t *TodoList) RemoveTask(index int) {
	if index < 0 || index >= len(t.Tasks) {
		fmt.Println("Invalid task number!")
		return
	}
	removedTask := t.Tasks[index]
	t.Tasks = append(t.Tasks[:index], t.Tasks[index+1:]...)
	fmt.Printf("Task '%s' removed from the list!\n", removedTask.Description)
}

func main() {
	todoList := TodoList{}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n==== Todo List App ====")
		fmt.Println("1. Add Task")
		fmt.Println("2. List Tasks")
		fmt.Println("3. Complete Task")
		fmt.Println("4. Remove Task")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice (1-5): ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Print("Enter task description: ")
			scanner.Scan()
			description := scanner.Text()
			todoList.AddTask(description)

		case "2":
			todoList.ListTasks()

		case "3":
			todoList.ListTasks()
			fmt.Print("Enter task number to mark as completed: ")
			scanner.Scan()
			var taskNum int
			fmt.Sscanf(scanner.Text(), "%d", &taskNum)
			todoList.CompleteTask(taskNum - 1)

		case "4":
			todoList.ListTasks()
			fmt.Print("Enter task number to remove: ")
			scanner.Scan()
			var taskNum int
			fmt.Sscanf(scanner.Text(), "%d", &taskNum)
			todoList.RemoveTask(taskNum - 1)

		case "5":
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

