package main

import (
	"fmt"

	"github.com/bunnyhawper/argparse-go"
)

func main() {
	// Create a new parser
	parser := argparse.NewParser("example", "Example command-line application")

	// Add help and version
	parser.AddHelp()
	parser.AddVersion()

	// Add string argument
	parser.String("n", "name", &argparse.Argument{
		Description: "Your name",
		IsRequired:  true,
	})

	// Add flag argument
	parser.Bool("b", "verbose", &argparse.Argument{
		Description: "Enable verbose output",
		DefaultVal:  false,
	})

	// Add positional argument
	parser.Positional("file", &argparse.Argument{
		Description: "File to process",
		IsRequired:  true,
	})

	// Parse arguments
	parser.ParseOrExit()

	// Use the arguments
	fmt.Printf("Hello, %s!\n", parser.GetString("name"))

	if parser.GetBool("verbose") {
		fmt.Println("Verbose mode enabled")
		fmt.Println("Additional information will be displayed")
	}

	fmt.Printf("Processing file: %s\n", parser.GetString("file"))
}
