package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Uso: go run main.go <command> [argumentos]")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Uso: add <tarefa>")
			return
		}
		AddTask(os.Args[2])

	case "list":
		ListTasks()

	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Uso: done <id>")
			return
		}
		MarkDone(os.Args[2])

	case "undone":
		if len(os.Args) < 3 {
			fmt.Println("Uso: undone <id>")
			return
		}
		MarkUndone(os.Args[2])

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Uso: delete <id>")
			return
		}

		DeleteTask(os.Args[2])

	default:
		fmt.Println("Invalid command. Use: add, list, done, delete")
	}
}
