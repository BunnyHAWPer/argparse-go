# 🚩 argparse - A Powerful Go Command-Line Argument Parser

[![Go Version](https://img.shields.io/badge/Go-1.13+-00ADD8?style=flat-square&logo=go)](https://golang.org)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Stability: Beta](https://img.shields.io/badge/stability-beta-33bbff.svg)](https://github.com/bunnyhawper/argparse)

A comprehensive and feature-rich command-line argument parser for Go applications, providing a simple, intuitive API while offering powerful functionality.

## ✨ Features

- 🔧 Easy-to-use API with fluent interface
- 🧰 Support for different argument types:
  - `String`, `Int`, `Float`, `Bool`
  - `List` (comma-separated values)
  - `Counter` (incremented with repeated flags)
  - `DateTime` (with various format support)
- ⚠️ Required arguments with validation
- 🔍 Default values for arguments
- 📍 Positional arguments
- 🌲 Subcommands/nested commands
- ✅ Choices validation (allowed values)
- 📚 Automatic help and version flags
- 🎨 Customizable help messages with description and epilog

## 📦 Installation

```bash
go get github.com/bunnyhawper/argparse
```

## 🚀 Quick Start

```go
package main

import (
    "fmt"
    "github.com/bunnyhawper/argparse"
)

func main() {
    // Create a new parser
    parser := argparse.NewParser("myapp", "My awesome command-line application")
    
    // Add help and version flags
    parser.AddHelp()
    parser.AddVersion()
    
    // Add a string argument with a short name (-n) and long name (--name)
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
        Description: "Input file",
        IsRequired:  true,
    })
    
    // Parse the arguments
    args := parser.ParseOrExit()
    
    // Use the arguments
    fmt.Printf("Hello, %s!\n", parser.GetString("name"))
    
    if parser.GetBool("verbose") {
        fmt.Println("Verbose mode enabled")
    }
    
    fmt.Printf("Processing file: %s\n", parser.GetString("file"))
}
```

## 📘 API Reference

### Creating a Parser

```go
// Create a new parser with name and description
parser := argparse.NewParser("myapp", "Description of my application")

// Set version (optional, overrides default version)
parser.SetVersion("1.2.3")

// Set epilog text (optional, shown at the end of help message)
parser.SetEpilog("For more information, visit: https://example.com")

// Add automatic help flag (-h, --help)
parser.AddHelp()

// Add automatic version flag (-V, --version)
parser.AddVersion()
```

### Adding Arguments

#### Optional Arguments

```go
// String argument
parser.String("s", "string", &argparse.Argument{
    Description: "String argument description",
    DefaultVal:  "default value",
    IsRequired:  false, // optional
})

// Integer argument
parser.Int("i", "integer", &argparse.Argument{
    Description: "Integer argument description",
    DefaultVal:  42,
    IsRequired:  false,
})

// Float argument
parser.Float("f", "float", &argparse.Argument{
    Description: "Float argument description",
    DefaultVal:  3.14,
    IsRequired:  false,
})

// Boolean argument (flag)
parser.Bool("b", "boolean", &argparse.Argument{
    Description: "Boolean argument description",
    DefaultVal:  false,
})

// List argument (comma-separated values)
parser.List("l", "list", &argparse.Argument{
    Description: "List argument description",
    DefaultVal:  []string{"default", "values"},
})

// Counter argument (increments with each occurrence)
parser.Counter("c", "counter", &argparse.Argument{
    Description: "Counter argument description",
    DefaultVal:  0,
})

// DateTime argument
parser.DateTime("d", "datetime", &argparse.Argument{
    Description: "DateTime argument description",
    // Supports formats: RFC3339, "2006-01-02", "2006-01-02 15:04:05", "01/02/2006", etc.
})
```

#### Positional Arguments

```go
// Add a required positional argument
parser.Positional("name", &argparse.Argument{
    Description: "Positional argument description",
    IsRequired:  true,
})

// Add an optional positional argument with default value
parser.Positional("config", &argparse.Argument{
    Description: "Optional positional argument",
    IsRequired:  false,
    DefaultVal:  "default.conf",
})
```

### Argument Customization

You can customize arguments using a fluent API:

```go
// Add a required argument
parser.String("n", "name", &argparse.Argument{
    Description: "Your name",
}).Required()

// Set default value
parser.Int("p", "port", &argparse.Argument{
    Description: "Server port",
}).Default(8080)

// Set help text
parser.String("o", "output", &argparse.Argument{
    DefaultVal: "output.txt",
}).Help("Specify the output file path")

// Set allowed choices
parser.String("m", "mode", &argparse.Argument{
    Description: "Operation mode",
    DefaultVal:  "normal",
}).Choices([]string{"normal", "fast", "safe"})
```

### Working with Subcommands

```go
// Create main parser
parser := argparse.NewParser("app", "Application description")
parser.AddHelp()
parser.AddVersion()

// Create subcommands
addCmd := parser.NewCommand("add", "Add a resource")
listCmd := parser.NewCommand("list", "List resources")
removeCmd := parser.NewCommand("remove", "Remove a resource")

// Add arguments to subcommands
addCmd.Parser.String("n", "name", &argparse.Argument{
    Description: "Resource name",
    IsRequired:  true,
})

listCmd.Parser.Bool("a", "all", &argparse.Argument{
    Description: "Show all resources",
    DefaultVal:  false,
})

removeCmd.Parser.Int("i", "id", &argparse.Argument{
    Description: "Resource ID to remove",
    IsRequired:  true,
})

// Parse arguments
args := parser.ParseOrExit()

// Handle subcommands
switch parser.GetString("subcommand") {
case "add":
    name := addCmd.Parser.GetString("name")
    // Handle add command
    
case "list":
    showAll := listCmd.Parser.GetBool("all")
    // Handle list command
    
case "remove":
    id := removeCmd.Parser.GetInt("id")
    // Handle remove command
}
```

### Parsing Arguments

```go
// Parse arguments (returns map and error)
args, err := parser.Parse(nil)
if err != nil {
    fmt.Fprintf(os.Stderr, "Error: %v\n", err)
    parser.PrintHelp()
    os.Exit(1)
}

// Or use the convenient ParseOrExit method (exits on error)
args := parser.ParseOrExit()
```

### Retrieving Argument Values

```go
// Get raw value
value := parser.Get("name") // returns interface{}

// Get typed values
strValue := parser.GetString("string")
intValue := parser.GetInt("integer")
floatValue := parser.GetFloat("float")
boolValue := parser.GetBool("boolean")
listValue := parser.GetList("list")
dateTimeValue := parser.GetDateTime("datetime")
```

## 🌲 Advanced Usage with Subcommands

```go
package main

import (
    "fmt"
    "github.com/bunnyhawper/argparse"
)

func main() {
    // Create a new parser
    parser := argparse.NewParser("taskmgr", "Task management application")
    parser.AddHelp()
    parser.AddVersion()
    
    // Create subcommands
    addCmd := parser.NewCommand("add", "Add a new task")
    listCmd := parser.NewCommand("list", "List all tasks")
    
    // Add arguments to "add" subcommand
    addCmd.Parser.String("t", "title", &argparse.Argument{
        Description: "Task title",
        IsRequired:  true,
    })
    
    addCmd.Parser.Int("p", "priority", &argparse.Argument{
        Description: "Task priority (1-5)",
        DefaultVal:  3,
    }).Choices([]string{"1", "2", "3", "4", "5"})
    
    // Add arguments to "list" subcommand
    listCmd.Parser.Bool("a", "all", &argparse.Argument{
        Description: "Show all tasks including completed",
    })
    
    listCmd.Parser.String("s", "sort", &argparse.Argument{
        Description: "Sort order",
        DefaultVal:  "priority",
    }).Choices([]string{"priority", "date", "title"})
    
    // Parse arguments
    parser.ParseOrExit()
    
    // Handle subcommands
    switch parser.GetString("subcommand") {
    case "add":
        title := addCmd.Parser.GetString("title")
        priority := addCmd.Parser.GetInt("priority")
        fmt.Printf("Adding task: %s with priority %d\n", title, priority)
        
    case "list":
        showAll := listCmd.Parser.GetBool("all")
        sortBy := listCmd.Parser.GetString("sort")
        fmt.Printf("Listing tasks (sort by: %s, show all: %t)\n", sortBy, showAll)
    }
}
```

## 🧩 All Argument Types

```go
// String arguments
strArg := parser.String("s", "string", &argparse.Argument{
    Description: "String argument example",
    DefaultVal:  "default value",
})

// Integer arguments
intArg := parser.Int("i", "integer", &argparse.Argument{
    Description: "Integer argument example",
    DefaultVal:  42,
})

// Float arguments
floatArg := parser.Float("f", "float", &argparse.Argument{
    Description: "Float argument example",
    DefaultVal:  3.14,
})

// Boolean arguments
boolArg := parser.Bool("b", "boolean", &argparse.Argument{
    Description: "Boolean argument example",
    DefaultVal:  false,
})

// List arguments (comma-separated values)
listArg := parser.List("l", "list", &argparse.Argument{
    Description: "List argument example",
    DefaultVal:  []string{"default", "values"},
})

// Counter arguments (increment with repeated flags)
counterArg := parser.Counter("c", "counter", &argparse.Argument{
    Description: "Counter argument example",
    DefaultVal:  0,
})

// DateTime arguments
dateArg := parser.DateTime("d", "date", &argparse.Argument{
    Description: "Date argument example (YYYY-MM-DD)",
})

// Argument with limited choices
choiceArg := parser.String("", "choice", &argparse.Argument{
    Description: "Argument with limited choices",
    DefaultVal:  "option1",
}).Choices([]string{"option1", "option2", "option3"})

// Required arguments
requiredArg := parser.String("r", "required", &argparse.Argument{
    Description: "Required argument example",
    IsRequired:  true,
})

// Positional arguments
posArg := parser.Positional("positional", &argparse.Argument{
    Description: "Positional argument example",
    IsRequired:  true,
})
```

## 📋 Command Line Usage Examples

### Basic Usage

```bash
# Display help
myapp -h
myapp --help

# Display version
myapp -V
myapp --version

# Basic usage with positional argument
myapp -n "John" input.txt

# Using long options
myapp --name "John" --verbose input.txt

# Multiple short options combined
myapp -nJohn -v input.txt
```

### Subcommands

```bash
# Show main help
taskmgr -h

# Add command
taskmgr add -t "Complete report" -p 2

# List command
taskmgr list --all --sort date

# Remove command
taskmgr remove -i 123 -f
```

## 📁 Examples

Check out the `examples` directory for complete working examples:

- **Basic**: Simple example with flags and positional arguments
- **Advanced**: More complex example with subcommands
- **Complete**: Demonstrates all argument types

## 🔧 Running the Tests

To run all the examples and verify that the library works correctly:

On Windows:
```
.\run_tests.bat
```

On Unix/Linux:
```
./run_tests.sh
```

## 📄 License

MIT License

## 👨‍💻 Author

Created with ❤️ by Dhruv Rawat (GitHub: [bunnyhawper](https://github.com/bunnyhawper)) 