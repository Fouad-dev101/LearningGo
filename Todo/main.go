package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	list := NewTodoList()
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("📝 In-Memory To-Do List")
	fmt.Println("Commands: add <title> | list | done <id> | delete <id> | quit")

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}

		parts := strings.SplitN(input, " ", 2)
		cmd := strings.ToLower(parts[0])

		switch cmd {
		case "add":
			if len(parts) < 2 {
				fmt.Println("Usage: add <title>")
				continue
			}
			task, err := list.Add(parts[1])
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			fmt.Printf("Added task #%d: %s\n", task.ID, task.Title)

		case "list":
			tasks := list.List()
			if len(tasks) == 0 {
				fmt.Println("(no tasks yet)")
				continue
			}
			for _, t := range tasks {
				status := "[ ]"
				if t.Done {
					status = "[x]"
				}
				fmt.Printf("%s %d. %s\n", status, t.ID, t.Title)
			}

		case "done":
			id, err := parseID(parts)
			if err != nil {
				fmt.Println(err)
				continue
			}
			if err := list.Complete(id); err != nil {
				fmt.Println("Error:", err)
				continue
			}
			fmt.Printf("Task #%d marked as done ✅\n", id)

		case "delete":
			id, err := parseID(parts)
			if err != nil {
				fmt.Println(err)
				continue
			}
			if err := list.Delete(id); err != nil {
				fmt.Println("Error:", err)
				continue
			}
			fmt.Printf("Task #%d deleted 🗑️\n", id)

		case "quit", "exit":
			fmt.Println("Bye 👋")
			return

		default:
			fmt.Println("Unknown command. Try: add, list, done, delete, quit")
		}
	}
}

func parseID(parts []string) (int, error) {
	if len(parts) < 2 {
		return 0, fmt.Errorf("Usage: %s <id>", parts[0])
	}
	id, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, fmt.Errorf("invalid ID: %s", parts[1])
	}
	return id, nil
}