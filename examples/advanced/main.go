package main

import (
	"fmt"
	"time"

	"github.com/bunnyhawper/argparse-go"
)

func main() {
	// Create a new parser
	parser := argparse.NewParser("taskmgr", "Task management application")

	// Add help and version
	parser.AddHelp()
	parser.AddVersion()

	// Create subcommands
	addCmd := parser.NewCommand("add", "Add a new task")
	listCmd := parser.NewCommand("list", "List all tasks")
	removeCmd := parser.NewCommand("remove", "Remove a task")

	// Add arguments to "add" command
	addCmd.Parser.String("t", "title", &argparse.Argument{
		Description: "Task title",
		IsRequired:  true,
	})

	addCmd.Parser.String("d", "description", &argparse.Argument{
		Description: "Task description",
	})

	addCmd.Parser.Int("p", "priority", &argparse.Argument{
		Description: "Task priority (1-5)",
		DefaultVal:  3,
	}).Choices([]string{"1", "2", "3", "4", "5"})

	addCmd.Parser.DateTime("", "due", &argparse.Argument{
		Description: "Due date (format: YYYY-MM-DD)",
	})

	addCmd.Parser.List("l", "labels", &argparse.Argument{
		Description: "Task labels (comma-separated)",
	})

	// Add arguments to "list" command
	listCmd.Parser.Bool("a", "all", &argparse.Argument{
		Description: "Show all tasks including completed",
	})

	listCmd.Parser.String("s", "sort", &argparse.Argument{
		Description: "Sort order",
		DefaultVal:  "priority",
	}).Choices([]string{"priority", "date", "title"})

	listCmd.Parser.Int("", "limit", &argparse.Argument{
		Description: "Maximum number of tasks to show",
		DefaultVal:  10,
	})

	// Add arguments to "remove" command
	removeCmd.Parser.Int("i", "id", &argparse.Argument{
		Description: "Task ID to remove",
		IsRequired:  true,
	})

	removeCmd.Parser.Bool("f", "force", &argparse.Argument{
		Description: "Force removal without confirmation",
	})

	// Parse arguments
	parser.ParseOrExit()

	// Handle subcommands
	switch parser.GetString("subcommand") {
	case "add":
		handleAddCommand(addCmd.Parser)

	case "list":
		handleListCommand(listCmd.Parser)

	case "remove":
		handleRemoveCommand(removeCmd.Parser)

	default:
		// No subcommand specified, show general help
		parser.PrintHelp()
	}
}

func handleAddCommand(parser *argparse.Parser) {
	title := parser.GetString("title")
	description := parser.GetString("description")
	priority := parser.GetInt("priority")

	fmt.Printf("Adding task: %s\n", title)
	fmt.Printf("Priority: %d\n", priority)

	if description != "" {
		fmt.Printf("Description: %s\n", description)
	}

	dueDate := parser.GetDateTime("due")
	if !dueDate.IsZero() {
		fmt.Printf("Due date: %s\n", dueDate.Format("2006-01-02"))
	}

	labels := parser.GetList("labels")
	if len(labels) > 0 {
		fmt.Println("Labels:")
		for _, label := range labels {
			fmt.Printf("  - %s\n", label)
		}
	}

	fmt.Println("Task added successfully!")
}

func handleListCommand(parser *argparse.Parser) {
	showAll := parser.GetBool("all")
	sortBy := parser.GetString("sort")
	limit := parser.GetInt("limit")

	fmt.Printf("Listing tasks (limit: %d)\n", limit)
	fmt.Printf("Sort by: %s\n", sortBy)

	if showAll {
		fmt.Println("Including completed tasks")
	} else {
		fmt.Println("Only showing pending tasks")
	}

	// Simulate listing tasks
	for i := 1; i <= 5; i++ {
		fmt.Printf("[%d] Task %d - Priority: %d, Due: %s\n",
			i,
			i,
			(i%3)+1,
			time.Now().AddDate(0, 0, i).Format("2006-01-02"))
	}
}

func handleRemoveCommand(parser *argparse.Parser) {
	id := parser.GetInt("id")
	force := parser.GetBool("force")

	if force {
		fmt.Printf("Removing task %d without confirmation...\n", id)
	} else {
		fmt.Printf("Are you sure you want to remove task %d? (simulating 'yes')\n", id)
	}

	fmt.Printf("Task %d removed successfully!\n", id)
}
