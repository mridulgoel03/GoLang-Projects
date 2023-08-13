package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	ID   int
	Name string
}

var tasks []Task
var lastID int

func addTask(name string) {
	lastID++
	tasks = append(tasks, Task{ID: lastID, Name: name})
	fmt.Println("Task added:", name)
}

func listTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks.")
		return
	}
	fmt.Println("Tasks:")
	for _, task := range tasks {
		fmt.Printf("[%d] %s\n", task.ID, task.Name)
	}
}

func removeTask(id int) {
	found := false
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			found = true
			break
		}
	}
	if found {
		fmt.Println("Task removed.")
	} else {
		fmt.Println("Task not found.")
	}
}

func main() {
	fmt.Println("Simple Task Manager")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			fmt.Println("Goodbye!")
			break
		}

		if input == "list" {
			listTasks()
			continue
		}

		if strings.HasPrefix(input, "add ") {
			taskName := strings.TrimPrefix(input, "add ")
			addTask(taskName)
			continue
		}

		if strings.HasPrefix(input, "remove ") {
			taskIDStr := strings.TrimPrefix(input, "remove ")
			taskID, err := strconv.Atoi(taskIDStr)
			if err != nil {
				fmt.Println("Invalid task ID.")
			} else {
				removeTask(taskID)
			}
			continue
		}

		fmt.Println("Invalid command. Available commands: list, add [task], remove [taskID], exit")
	}
}
