# argparse-go Test Commands

This document contains a comprehensive list of test commands used to verify the functionality of the `argparse-go` library. You can use these commands as examples for your own applications.

## Basic Example Commands

These commands test a simple application with required arguments, flags, and positional arguments.

```bash
# Show version
go run ./examples/basic/main.go -V

# Show help
go run ./examples/basic/main.go -h

# Basic usage with required arguments
go run ./examples/basic/main.go -n "Your Name" test.txt

# Using verbose flag (short form)
go run ./examples/basic/main.go -n "Your Name" -b test.txt

# Using long option names
go run ./examples/basic/main.go --name "Your Name" --verbose test.txt

# Error handling (missing required argument)
go run ./examples/basic/main.go test.txt
```

## Advanced Example Commands (with subcommands)

These commands test a more complex application with subcommands for a task management system.

```bash
# Show main help
go run ./examples/advanced/main.go -h

# Show subcommand help (note: currently shows error - could be improved)
go run ./examples/advanced/main.go add -h

# 'add' subcommand with minimal options
go run ./examples/advanced/main.go add -t "Complete homework"

# 'add' with all options
go run ./examples/advanced/main.go add -t "Write report" -d "Monthly sales report" -p 2 --due "2023-12-25" -l "work,urgent,report"

# 'list' subcommand with default options
go run ./examples/advanced/main.go list

# 'list' with all options
go run ./examples/advanced/main.go list -a --sort priority --limit 3

# 'list' with different sort option
go run ./examples/advanced/main.go list --sort date

# 'remove' subcommand (basic)
go run ./examples/advanced/main.go remove -i 1

# 'remove' with force option
go run ./examples/advanced/main.go remove -i 2 -f

# Error handling (missing required argument)
go run ./examples/advanced/main.go add
```

## Complete Example Commands (All Argument Types)

These commands test an application that demonstrates all argument types supported by the library.

```bash
# Show help
go run ./examples/complete/main.go -h

# Run with only required arguments
go run ./examples/complete/main.go -r "required-value" positional-value

# Run with multiple argument types
go run ./examples/complete/main.go -r "required-value" -s "hello" -i 123 -f 3.14 -b -l "a,b,c" -c -c -d "2023-01-01" --choice "option3" positional-value

# Test counter incrementation
go run ./examples/complete/main.go -r "required-value" -c -c -c -c positional-value

# Test choices validation
go run ./examples/complete/main.go -r "required-value" --choice "option2" positional-value

# Test invalid choice (should fail)
go run ./examples/complete/main.go -r "required-value" --choice "invalid" positional-value

# Test different datetime formats
go run ./examples/complete/main.go -r "required-value" -d "2023-05-15 14:30:00" positional-value
```

## Running All Tests

To automatically run all these commands and verify the library's functionality, use the provided test scripts:

- **Windows**: `.\run_tests.bat`
- **Unix/Linux**: `./run_tests.sh` 