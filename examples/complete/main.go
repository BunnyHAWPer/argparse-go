package main

import (
	"fmt"
	"time"

	"github.com/bunnyhawper/argparse-go"
)

func main() {
	// Create a new parser
	parser := argparse.NewParser("complete", "Demonstration of all argument types")

	// Add help and version
	parser.AddHelp()
	parser.AddVersion()

	// String arguments
	parser.String("s", "string", &argparse.Argument{
		Description: "String argument example",
		DefaultVal:  "default string",
	})

	// Integer arguments
	parser.Int("i", "integer", &argparse.Argument{
		Description: "Integer argument example",
		DefaultVal:  42,
	})

	// Float arguments
	parser.Float("f", "float", &argparse.Argument{
		Description: "Float argument example",
		DefaultVal:  3.14,
	})

	// Boolean arguments
	parser.Bool("b", "boolean", &argparse.Argument{
		Description: "Boolean argument example",
		DefaultVal:  false,
	})

	// List arguments
	parser.List("l", "list", &argparse.Argument{
		Description: "List argument example (comma-separated values)",
		DefaultVal:  []string{"default", "values"},
	})

	// Counter arguments
	parser.Counter("c", "counter", &argparse.Argument{
		Description: "Counter argument example (use multiple times to increase)",
		DefaultVal:  0,
	})

	// DateTime arguments
	parser.DateTime("d", "datetime", &argparse.Argument{
		Description: "DateTime argument example (YYYY-MM-DD or YYYY-MM-DD HH:MM:SS)",
	})

	// Arguments with choices
	parser.String("", "choice", &argparse.Argument{
		Description: "Argument with limited choices",
		DefaultVal:  "option1",
	}).Choices([]string{"option1", "option2", "option3"})

	// Required arguments
	parser.String("r", "required", &argparse.Argument{
		Description: "Required argument example",
		IsRequired:  true,
	})

	// Positional argument
	parser.Positional("positional", &argparse.Argument{
		Description: "Positional argument example",
		IsRequired:  true,
	})

	// Parse arguments
	parser.ParseOrExit()

	// Display all values
	fmt.Println("Arguments:")
	fmt.Printf("  String:    %s\n", parser.GetString("string"))
	fmt.Printf("  Integer:   %d\n", parser.GetInt("integer"))
	fmt.Printf("  Float:     %.2f\n", parser.GetFloat("float"))
	fmt.Printf("  Boolean:   %t\n", parser.GetBool("boolean"))

	listVal := parser.GetList("list")
	fmt.Printf("  List:      %v (length: %d)\n", listVal, len(listVal))

	fmt.Printf("  Counter:   %d\n", parser.GetInt("counter"))

	dateVal := parser.GetDateTime("datetime")
	if !dateVal.IsZero() {
		fmt.Printf("  DateTime:  %s\n", dateVal.Format(time.RFC1123))
	} else {
		fmt.Printf("  DateTime:  not provided\n")
	}

	fmt.Printf("  Choice:    %s\n", parser.GetString("choice"))
	fmt.Printf("  Required:  %s\n", parser.GetString("required"))
	fmt.Printf("  Positional: %s\n", parser.GetString("positional"))
}
