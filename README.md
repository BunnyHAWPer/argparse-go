# ArgParse Go

A powerful command-line argument parser for Go applications.

[![Go Version](https://img.shields.io/badge/Go-1.18+-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)
[![Version](https://img.shields.io/badge/Version-1.0.0-orange.svg)](https://github.com/bunnyhawper/argparse-go)

ArgParse Go is a flexible and feature-rich command-line argument parsing library for Go applications. It provides intuitive APIs for defining and handling various argument types, subcommands, and validation rules.

## Features

- 📦 Easy to use, lightweight, with zero dependencies
- 💪 Support for multiple argument types (string, int, float, bool, list, etc.)
- 🔍 Automatic help text generation
- 🧩 Support for subcommands
- ✅ Required arguments and validation
- 🔄 Default values and choices
- 🕰️ Date/time parsing

## Important Note

This library is hosted at `github.com/bunnyhawper/argparse-go` but uses the package name `argparse` for better usability. When importing the library, use:

```go
import "github.com/bunnyhawper/argparse-go"
```

Then in your code, you'll refer to it simply as `argparse`:

```go
parser := argparse.NewParser("myapp", "My application")
```

## Installation

```bash
go get github.com/bunnyhawper/argparse-go
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/bunnyhawper/argparse-go"
)

func main() {
    // Create a new parser
    parser := argparse.NewParser("example", "Example command-line application")
    
    // Add help and version flags
    parser.AddHelp()
    parser.AddVersion()
    
    // Add a required string argument
    parser.String("n", "name", &argparse.Argument{
        Description: "Your name",
        IsRequired:  true,
    })
    
    // Add a boolean flag
    parser.Bool("v", "verbose", &argparse.Argument{
        Description: "Enable verbose output",
        DefaultVal:  false,
    })
    
    // Add a positional argument
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
    }
    
    fmt.Printf("Processing file: %s\n", parser.GetString("file"))
}
```

Run with:
```bash
go run main.go -n "John Doe" -v example.txt
```

## API Reference

### Creating a Parser

```go
parser := argparse.NewParser(name, description)
```

#### Parameters:
- `name` (string): The name of your application
- `description` (string): A short description of your application

### Parser Methods

#### Setting Parser Information

```go
parser.SetEpilog(epilog)    // Sets text to display after help message
parser.SetVersion(version)  // Sets version string
parser.AddHelp()            // Adds -h/--help option
parser.AddVersion()         // Adds -V/--version option
```

#### Adding Arguments

```go
// Add a flag argument (general method)
parser.Flag(shortName, longName, options)

// Type-specific methods
parser.String(shortName, longName, options)    // String argument
parser.Int(shortName, longName, options)       // Integer argument
parser.Float(shortName, longName, options)     // Float argument
parser.Bool(shortName, longName, options)      // Boolean flag
parser.List(shortName, longName, options)      // List of values
parser.Counter(shortName, longName, options)   // Counter (increments with each occurrence)
parser.DateTime(shortName, longName, options)  // Date/time value

// Positional arguments
parser.Positional(name, options)
```

#### Parameters:
- `shortName` (string): Short name for the argument (e.g., "v" for -v)
- `longName` (string): Long name for the argument (e.g., "verbose" for --verbose)
- `name` (string): Name for positional arguments
- `options` (*Argument): Argument options (see below)

### Argument Options

When creating arguments, you can provide an `Argument` struct with the following options:

```go
&argparse.Argument{
    Description: "Description of the argument",
    IsRequired:  true,                    // Whether the argument is required
    ArgType:     argparse.String,         // Argument type (automatically set by type-specific methods)
    DefaultVal:  "default value",         // Default value if argument is not provided
    ValidChoices: []string{"opt1", "opt2"}, // Valid choices for the argument
}
```

### Argument Modifiers

After creating an argument, you can add modifiers:

```go
arg := parser.String("n", "name", &argparse.Argument{...})

arg.Required()              // Make the argument required
arg.Default("John Doe")     // Set a default value
arg.Help("Help text")       // Set help text
arg.Choices([]string{...})  // Set valid choices
```

### Subcommands

```go
// Create a subcommand
cmd := parser.NewCommand(name, description)

// Add arguments to the subcommand
cmd.Parser.String(...)
cmd.Parser.Int(...)
// etc.
```

### Parsing Arguments

```go
// Parse arguments and exit on error
parser.ParseOrExit()

// Parse arguments and handle errors manually
args, err := parser.Parse(os.Args[1:])
if err != nil {
    // Handle error
}
```

### Getting Argument Values

```go
// Get values using type-specific methods
s := parser.GetString("name")     // Get string value
i := parser.GetInt("count")       // Get integer value
f := parser.GetFloat("amount")    // Get float value
b := parser.GetBool("verbose")    // Get boolean value
l := parser.GetList("tags")       // Get list value
dt := parser.GetDateTime("date")  // Get datetime value

// Generic method (returns interface{})
val := parser.Get("name")
```

## Argument Types

| Type     | Description                          | Example                            |
|----------|--------------------------------------|-----------------------------------|
| String   | Text value                           | `-s "hello"` or `--string "hello"` |
| Int      | Integer value                        | `-i 42` or `--int 42`              |
| Float    | Floating-point value                 | `-f 3.14` or `--float 3.14`        |
| Bool     | Boolean flag                         | `-b` or `--bool`                   |
| List     | List of values                       | `-l "one,two,three"` or `--list "one,two,three"` |
| Counter  | Increments with each occurrence      | `-c -c -c` (value would be 3)      |
| DateTime | Date and time value                  | `--date "2023-01-01"` or `--date "2023-01-01 15:30:00"` |

## Examples

### Basic Example

See the [basic example](examples/basic/main.go) for a simple application.

### Subcommands Example

See the [advanced example](examples/advanced/main.go) for an application with subcommands.

### All Argument Types

See the [complete example](examples/complete/main.go) for an application using all argument types.

## Command-line Usage

When you run your application with `-h` or `--help`, it will display automatically generated help text:

```
usage: example [-h] [-V] -n NAME [-b] file

Example command-line application

positional arguments:
  file                  File to process

optional arguments:
  -h, --help            Show this help message and exit
  -V, --version         Show program's version and exit
  -n, --name NAME       Your name
  -b, --verbose         Enable verbose output
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

Created by [Dhruv Rawat](https://github.com/bunnyhawper) with ❤️
