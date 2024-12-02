package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// createTask prompts the user for a task description and adds it to the task list.
func createTask(tasks []string) []string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter task description: ")
	scanner.Scan()
	// Trimming leading/trailing spaces
	taskDescription := strings.TrimSpace(scanner.Text())

	// Validate that input is not empty
	if taskDescription == "" {
		fmt.Println("Task description cannot be empty.")
		return tasks
	}

	// Append the new task to the list and confirm creation
	tasks = append(tasks, taskDescription)
	fmt.Println("Task created successfully!")
	return tasks
}

// showTasks displays the list of tasks, formatted with task numbers.
func showTasks(tasks []string) {
	if len(tasks) == 0 {
		fmt.Println("No tasks available!")
		return
	}

	// Print all tasks with their index
	fmt.Println("Your tasks:")
	for idx, task := range tasks {
		fmt.Printf("[%d]: %s\n", idx+1, task)
	}
}

// main is the entry point where the CLI interacts with the user.
func main() {
	// Initialize an empty slice for tasks
	var tasks []string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to the CLI To-Do List!")

	// Main loop for accepting commands
	for {
		fmt.Print("Enter a command (add, list, exit): ")
		scanner.Scan()
		command := strings.ToLower(strings.TrimSpace(scanner.Text()))

		switch command {
		case "add":
			tasks = createTask(tasks)
		case "list":
			showTasks(tasks)
		case "exit":
			// Exit the loop and program
			fmt.Println("Exiting the application.")
			return
		default:
			fmt.Println("Invalid command. Please use 'add', 'list', or 'exit'.")
		}
	}
}
