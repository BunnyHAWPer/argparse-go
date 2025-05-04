// Package argparse provides a powerful command-line argument parser for Go applications.
//
// Version: 1.0.0
// Author: Dhruv Rawat
// GitHub: bunnyhawper
package argparse

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Version information
const (
	Version = "1.0.0"
	Author  = "Dhruv Rawat"
	GitHub  = "bunnyhawper"
)

// ArgumentType defines the type of an argument
type ArgumentType int

const (
	// String argument type
	String ArgumentType = iota
	// Int argument type
	Int
	// Float argument type
	Float
	// Bool argument type
	Bool
	// List argument type (multiple values)
	List
	// Counter argument type (counts occurrences)
	Counter
	// DateTime argument type
	DateTime
)

// Argument represents a command-line argument
type Argument struct {
	Name         string
	ShortName    string
	Description  string
	IsRequired   bool
	ArgType      ArgumentType
	DefaultVal   interface{}
	ValidChoices []string
	value        interface{}
	isSet        bool
	isPositional bool
	parent       *Parser
}

// Parser represents the argument parser
type Parser struct {
	name        string
	description string
	epilog      string
	version     string
	args        []*Argument
	positional  []*Argument
	subparsers  map[string]*Parser
	parent      *Parser
	subparser   string
}

// Command represents a subcommand in the parser
type Command struct {
	Parser *Parser
}

// NewParser creates a new argument parser
func NewParser(name string, description string) *Parser {
	return &Parser{
		name:        name,
		description: description,
		version:     Version,
		args:        make([]*Argument, 0),
		positional:  make([]*Argument, 0),
		subparsers:  make(map[string]*Parser),
	}
}

// SetEpilog sets the epilog text for the parser
func (p *Parser) SetEpilog(epilog string) *Parser {
	p.epilog = epilog
	return p
}

// SetVersion sets the version for the parser
func (p *Parser) SetVersion(version string) *Parser {
	p.version = version
	return p
}

// AddHelp adds a help argument to the parser
func (p *Parser) AddHelp() *Argument {
	help := p.Flag("h", "help", &Argument{
		Description: "Show this help message and exit",
		ArgType:     Bool,
		DefaultVal:  false,
	})
	return help
}

// AddVersion adds a version argument to the parser
func (p *Parser) AddVersion() *Argument {
	version := p.Flag("V", "version", &Argument{
		Description: "Show program's version and exit",
		ArgType:     Bool,
		DefaultVal:  false,
	})
	return version
}

// NewCommand creates a new subcommand
func (p *Parser) NewCommand(name string, description string) *Command {
	subparser := &Parser{
		name:        name,
		description: description,
		version:     p.version,
		args:        make([]*Argument, 0),
		positional:  make([]*Argument, 0),
		subparsers:  make(map[string]*Parser),
		parent:      p,
	}

	p.subparsers[name] = subparser

	return &Command{
		Parser: subparser,
	}
}

// Flag adds a new flag argument
func (p *Parser) Flag(shortName, longName string, options *Argument) *Argument {
	if options == nil {
		options = &Argument{}
	}

	options.ShortName = shortName
	options.Name = longName
	options.isPositional = false
	options.parent = p

	p.args = append(p.args, options)
	return options
}

// String adds a string argument
func (p *Parser) String(shortName, longName string, options *Argument) *Argument {
	if options == nil {
		options = &Argument{}
	}
	options.ArgType = String

	return p.Flag(shortName, longName, options)
}

// Int adds an integer argument
func (p *Parser) Int(shortName, longName string, options *Argument) *Argument {
	if options == nil {
		options = &Argument{}
	}
	options.ArgType = Int

	return p.Flag(shortName, longName, options)
}

// Float adds a float argument
func (p *Parser) Float(shortName, longName string, options *Argument) *Argument {
	if options == nil {
		options = &Argument{}
	}
	options.ArgType = Float

	return p.Flag(shortName, longName, options)
}

// Bool adds a boolean argument
func (p *Parser) Bool(shortName, longName string, options *Argument) *Argument {
	if options == nil {
		options = &Argument{}
	}
	options.ArgType = Bool
	options.DefaultVal = false

	return p.Flag(shortName, longName, options)
}

// List adds a list argument
func (p *Parser) List(shortName, longName string, options *Argument) *Argument {
	if options == nil {
		options = &Argument{}
	}
	options.ArgType = List
	if options.DefaultVal == nil {
		options.DefaultVal = make([]string, 0)
	}

	return p.Flag(shortName, longName, options)
}

// Counter adds a counter argument
func (p *Parser) Counter(shortName, longName string, options *Argument) *Argument {
	if options == nil {
		options = &Argument{}
	}
	options.ArgType = Counter
	if options.DefaultVal == nil {
		options.DefaultVal = 0
	}

	return p.Flag(shortName, longName, options)
}

// DateTime adds a datetime argument
func (p *Parser) DateTime(shortName, longName string, options *Argument) *Argument {
	if options == nil {
		options = &Argument{}
	}
	options.ArgType = DateTime

	return p.Flag(shortName, longName, options)
}

// Positional adds a positional argument
func (p *Parser) Positional(name string, options *Argument) *Argument {
	if options == nil {
		options = &Argument{}
	}

	options.Name = name
	options.isPositional = true
	options.parent = p

	p.positional = append(p.positional, options)
	return options
}

// Required sets the argument as required
func (a *Argument) Required() *Argument {
	a.IsRequired = true
	return a
}

// Default sets the default value for the argument
func (a *Argument) Default(value interface{}) *Argument {
	a.DefaultVal = value
	return a
}

// Help sets the help text for the argument
func (a *Argument) Help(helpText string) *Argument {
	a.Description = helpText
	return a
}

// Choices sets the valid choices for the argument
func (a *Argument) Choices(choices []string) *Argument {
	a.ValidChoices = choices
	return a
}

// Parse parses the command line arguments
func (p *Parser) Parse(args []string) (map[string]interface{}, error) {
	if args == nil {
		args = os.Args[1:]
	}

	// Initialize result map
	result := make(map[string]interface{})

	// Add default values
	for _, arg := range p.args {
		if arg.DefaultVal != nil {
			result[arg.Name] = arg.DefaultVal
		}
	}

	for _, arg := range p.positional {
		if arg.DefaultVal != nil {
			result[arg.Name] = arg.DefaultVal
		}
	}

	// Process arguments
	positionalIndex := 0
	hasVersionFlag := false
	hasHelpFlag := false

	for i := 0; i < len(args); i++ {
		arg := args[i]

		// Check for subcommand
		if len(p.subparsers) > 0 && i == 0 {
			if subparser, ok := p.subparsers[arg]; ok {
				p.subparser = arg
				subResult, err := subparser.Parse(args[1:])
				if err != nil {
					return nil, err
				}

				result["subcommand"] = arg
				for k, v := range subResult {
					result[k] = v
				}
				return result, nil
			}
		}

		// Process flags
		if strings.HasPrefix(arg, "-") {
			var name string
			var value string
			hasValue := false

			if strings.HasPrefix(arg, "--") {
				// Long option
				parts := strings.SplitN(arg[2:], "=", 2)
				name = parts[0]
				if len(parts) > 1 {
					value = parts[1]
					hasValue = true
				}

				// Check for help and version flags
				if name == "help" {
					hasHelpFlag = true
					result["help"] = true
				} else if name == "version" {
					hasVersionFlag = true
					result["version"] = true
				}

				// Find matching argument
				found := false
				for _, option := range p.args {
					if option.Name == name {
						found = true

						switch option.ArgType {
						case Bool:
							result[option.Name] = true
							option.isSet = true

						case Counter:
							count, _ := result[option.Name].(int)
							result[option.Name] = count + 1
							option.isSet = true

						default:
							if hasValue {
								parsedValue, err := parseValue(option.ArgType, value)
								if err != nil {
									return nil, fmt.Errorf("invalid value for --%s: %v", name, err)
								}
								result[option.Name] = parsedValue
								option.isSet = true
							} else {
								if i+1 < len(args) && !strings.HasPrefix(args[i+1], "-") {
									i++
									parsedValue, err := parseValue(option.ArgType, args[i])
									if err != nil {
										return nil, fmt.Errorf("invalid value for --%s: %v", name, err)
									}
									result[option.Name] = parsedValue
									option.isSet = true
								} else {
									return nil, fmt.Errorf("argument --%s requires a value", name)
								}
							}
						}
						break
					}
				}

				if !found {
					return nil, fmt.Errorf("unknown argument: --%s", name)
				}

			} else {
				// Short option
				shortName := arg[1:]

				// Check for help and version flags
				if shortName == "h" {
					hasHelpFlag = true
					result["help"] = true
				} else if shortName == "V" {
					hasVersionFlag = true
					result["version"] = true
				}

				// Handle multiple short options (e.g., -abc)
				for j, shortOpt := range shortName {
					found := false
					for _, option := range p.args {
						if option.ShortName == string(shortOpt) {
							found = true

							switch option.ArgType {
							case Bool:
								result[option.Name] = true
								option.isSet = true

							case Counter:
								count, _ := result[option.Name].(int)
								result[option.Name] = count + 1
								option.isSet = true

							default:
								if j < len(shortName)-1 {
									// If not the last character, use the rest as value
									value = shortName[j+1:]
									parsedValue, err := parseValue(option.ArgType, value)
									if err != nil {
										return nil, fmt.Errorf("invalid value for -%c: %v", shortOpt, err)
									}
									result[option.Name] = parsedValue
									option.isSet = true
									j = len(shortName) // Stop processing
								} else {
									if i+1 < len(args) && !strings.HasPrefix(args[i+1], "-") {
										i++
										parsedValue, err := parseValue(option.ArgType, args[i])
										if err != nil {
											return nil, fmt.Errorf("invalid value for -%c: %v", shortOpt, err)
										}
										result[option.Name] = parsedValue
										option.isSet = true
									} else {
										return nil, fmt.Errorf("argument -%c requires a value", shortOpt)
									}
								}
							}
							break
						}
					}

					if !found {
						return nil, fmt.Errorf("unknown argument: -%c", shortOpt)
					}

					// If we used the rest of the string as a value, stop processing
					if j+1 < len(shortName) {
						break
					}
				}
			}
		} else {
			// Positional argument
			if positionalIndex < len(p.positional) {
				pos := p.positional[positionalIndex]
				parsedValue, err := parseValue(pos.ArgType, arg)
				if err != nil {
					return nil, fmt.Errorf("invalid value for %s: %v", pos.Name, err)
				}
				result[pos.Name] = parsedValue
				pos.isSet = true
				positionalIndex++
			} else {
				return nil, fmt.Errorf("unrecognized positional argument: %s", arg)
			}
		}
	}

	// Special case: --help or -h
	if hasHelpFlag {
		p.PrintHelp()
		os.Exit(0)
	}

	// Special case: --version or -V
	if hasVersionFlag {
		fmt.Printf("%s %s\n", p.name, p.version)
		os.Exit(0)
	}

	// Check required arguments (only if help/version not specified)
	for _, arg := range p.args {
		if arg.IsRequired && !arg.isSet {
			if arg.isPositional {
				return nil, fmt.Errorf("required argument missing: %s", arg.Name)
			} else {
				if arg.ShortName != "" {
					return nil, fmt.Errorf("required argument missing: --%s/-%s", arg.Name, arg.ShortName)
				} else {
					return nil, fmt.Errorf("required argument missing: --%s", arg.Name)
				}
			}
		}
	}

	for _, arg := range p.positional {
		if arg.IsRequired && !arg.isSet {
			return nil, fmt.Errorf("required positional argument missing: %s", arg.Name)
		}
	}

	return result, nil
}

// ParseOrExit parses command line arguments or exits on error
func (p *Parser) ParseOrExit() map[string]interface{} {
	result, err := p.Parse(nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n\n", err)
		p.PrintHelp()
		os.Exit(1)
	}
	return result
}

// PrintHelp prints the help message
func (p *Parser) PrintHelp() {
	fmt.Printf("Usage: %s", p.name)

	if len(p.args) > 0 {
		fmt.Printf(" [options]")
	}

	for _, pos := range p.positional {
		if pos.IsRequired {
			fmt.Printf(" %s", pos.Name)
		} else {
			fmt.Printf(" [%s]", pos.Name)
		}
	}

	if len(p.subparsers) > 0 {
		fmt.Printf(" {")
		first := true
		for name := range p.subparsers {
			if !first {
				fmt.Printf(",")
			}
			fmt.Printf("%s", name)
			first = false
		}
		fmt.Printf("}")
	}

	fmt.Printf("\n\n%s\n\n", p.description)

	if len(p.positional) > 0 {
		fmt.Printf("Positional arguments:\n")
		for _, pos := range p.positional {
			fmt.Printf("  %-20s %s\n", pos.Name, pos.Description)
		}
		fmt.Printf("\n")
	}

	if len(p.args) > 0 {
		fmt.Printf("Optional arguments:\n")
		for _, arg := range p.args {
			if arg.ShortName != "" {
				fmt.Printf("  -%s, --%-15s %s\n", arg.ShortName, arg.Name, arg.Description)
			} else {
				fmt.Printf("      --%-15s %s\n", arg.Name, arg.Description)
			}
		}
		fmt.Printf("\n")
	}

	if len(p.subparsers) > 0 {
		fmt.Printf("Commands:\n")
		for name, subparser := range p.subparsers {
			fmt.Printf("  %-20s %s\n", name, subparser.description)
		}
		fmt.Printf("\n")
	}

	if p.epilog != "" {
		fmt.Printf("%s\n", p.epilog)
	}
}

// Helper function to parse values based on type
func parseValue(argType ArgumentType, value string) (interface{}, error) {
	switch argType {
	case String:
		return value, nil

	case Int:
		return strconv.Atoi(value)

	case Float:
		return strconv.ParseFloat(value, 64)

	case Bool:
		return strconv.ParseBool(value)

	case List:
		return strings.Split(value, ","), nil

	case DateTime:
		// Try common date formats
		formats := []string{
			time.RFC3339,
			"2006-01-02",
			"2006-01-02 15:04:05",
			"01/02/2006",
			"01/02/2006 15:04:05",
		}

		for _, format := range formats {
			t, err := time.Parse(format, value)
			if err == nil {
				return t, nil
			}
		}
		return nil, errors.New("invalid datetime format")

	default:
		return value, nil
	}
}

// Get retrieves the value of an argument by name
func (p *Parser) Get(name string) interface{} {
	args, _ := p.Parse(nil)
	return args[name]
}

// GetString retrieves the string value of an argument
func (p *Parser) GetString(name string) string {
	val := p.Get(name)
	if val == nil {
		return ""
	}
	if str, ok := val.(string); ok {
		return str
	}
	return fmt.Sprintf("%v", val)
}

// GetInt retrieves the int value of an argument
func (p *Parser) GetInt(name string) int {
	val := p.Get(name)
	if val == nil {
		return 0
	}
	if i, ok := val.(int); ok {
		return i
	}
	return 0
}

// GetFloat retrieves the float value of an argument
func (p *Parser) GetFloat(name string) float64 {
	val := p.Get(name)
	if val == nil {
		return 0
	}
	if f, ok := val.(float64); ok {
		return f
	}
	return 0
}

// GetBool retrieves the bool value of an argument
func (p *Parser) GetBool(name string) bool {
	val := p.Get(name)
	if val == nil {
		return false
	}
	if b, ok := val.(bool); ok {
		return b
	}
	return false
}

// GetList retrieves the list value of an argument
func (p *Parser) GetList(name string) []string {
	val := p.Get(name)
	if val == nil {
		return []string{}
	}
	if list, ok := val.([]string); ok {
		return list
	}
	return []string{}
}

// GetDateTime retrieves the datetime value of an argument
func (p *Parser) GetDateTime(name string) time.Time {
	val := p.Get(name)
	if val == nil {
		return time.Time{}
	}
	if t, ok := val.(time.Time); ok {
		return t
	}
	return time.Time{}
}
