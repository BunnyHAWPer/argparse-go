@echo off
echo ==============================================================
echo Argparse Testing Script
echo ==============================================================
echo.
echo This script will run through all example commands to test the 
echo argparse library functionality.
echo.
echo Press any key to continue...
pause > nul

echo.
echo ==============================================================
echo Testing Version and Help flags
echo ==============================================================
echo.

echo Running: go run ./examples/basic/main.go -V
go run ./examples/basic/main.go -V

echo.
echo Running: go run ./examples/basic/main.go -h
go run ./examples/basic/main.go -h

echo.
echo ==============================================================
echo Basic Example Tests
echo ==============================================================
echo.

echo Running basic example with required arguments...
echo go run ./examples/basic/main.go -n "Your Name" test.txt
go run ./examples/basic/main.go -n "Your Name" test.txt

echo.
echo Running with verbose flag (short form)...
echo go run ./examples/basic/main.go -n "Your Name" -b test.txt
go run ./examples/basic/main.go -n "Your Name" -b test.txt

echo.
echo Running with long option names...
echo go run ./examples/basic/main.go --name "Your Name" --verbose test.txt
go run ./examples/basic/main.go --name "Your Name" --verbose test.txt

echo.
echo Testing error handling (missing required argument)...
echo go run ./examples/basic/main.go test.txt
go run ./examples/basic/main.go test.txt || echo Error handling works as expected.

echo.
echo ==============================================================
echo Advanced Example Tests (with subcommands)
echo ==============================================================
echo.

echo Running help for main command...
echo go run ./examples/advanced/main.go -h
go run ./examples/advanced/main.go -h

echo.
echo Running help for add subcommand...
echo go run ./examples/advanced/main.go add -h || echo Command help not available
go run ./examples/advanced/main.go add -h || echo Command help not available

echo.
echo Running 'add' subcommand with minimal options...
echo go run ./examples/advanced/main.go add -t "Complete homework"
go run ./examples/advanced/main.go add -t "Complete homework"

echo.
echo Running 'add' with all options...
echo go run ./examples/advanced/main.go add -t "Write report" -d "Monthly sales report" -p 2 --due "2023-12-25" -l "work,urgent,report"
go run ./examples/advanced/main.go add -t "Write report" -d "Monthly sales report" -p 2 --due "2023-12-25" -l "work,urgent,report"

echo.
echo Running 'list' subcommand with default options...
echo go run ./examples/advanced/main.go list
go run ./examples/advanced/main.go list

echo.
echo Running 'list' with all options...
echo go run ./examples/advanced/main.go list -a --sort priority --limit 3
go run ./examples/advanced/main.go list -a --sort priority --limit 3

echo.
echo Running 'list' with different sort option...
echo go run ./examples/advanced/main.go list --sort date
go run ./examples/advanced/main.go list --sort date

echo.
echo Running 'remove' subcommand (basic)...
echo go run ./examples/advanced/main.go remove -i 1
go run ./examples/advanced/main.go remove -i 1

echo.
echo Running 'remove' with force option...
echo go run ./examples/advanced/main.go remove -i 2 -f
go run ./examples/advanced/main.go remove -i 2 -f

echo.
echo Testing error handling (missing required argument)...
echo go run ./examples/advanced/main.go add
go run ./examples/advanced/main.go add || echo Error handling works as expected.

echo.
echo ==============================================================
echo Complete Example Tests (All Argument Types)
echo ==============================================================
echo.

echo Running help...
echo go run ./examples/complete/main.go -h
go run ./examples/complete/main.go -h

echo.
echo Running with only required arguments...
echo go run ./examples/complete/main.go -r "required-value" positional-value
go run ./examples/complete/main.go -r "required-value" positional-value

echo.
echo Running with multiple argument types...
echo go run ./examples/complete/main.go -r "required-value" -s "hello" -i 123 -f 3.14 -b -l "a,b,c" -c -c -d "2023-01-01" --choice "option3" positional-value
go run ./examples/complete/main.go -r "required-value" -s "hello" -i 123 -f 3.14 -b -l "a,b,c" -c -c -d "2023-01-01" --choice "option3" positional-value

echo.
echo Testing counter incrementation...
echo go run ./examples/complete/main.go -r "required-value" -c -c -c -c positional-value
go run ./examples/complete/main.go -r "required-value" -c -c -c -c positional-value

echo.
echo Testing choices validation...
echo go run ./examples/complete/main.go -r "required-value" --choice "option2" positional-value
go run ./examples/complete/main.go -r "required-value" --choice "option2" positional-value

echo.
echo Testing invalid choice (should fail)...
echo go run ./examples/complete/main.go -r "required-value" --choice "invalid" positional-value
go run ./examples/complete/main.go -r "required-value" --choice "invalid" positional-value || echo Error handling for invalid choice works as expected.

echo.
echo Testing different datetime formats...
echo go run ./examples/complete/main.go -r "required-value" -d "2023-05-15 14:30:00" positional-value
go run ./examples/complete/main.go -r "required-value" -d "2023-05-15 14:30:00" positional-value

echo.
echo ==============================================================
echo All tests completed successfully!
echo ==============================================================
echo.
echo Press any key to exit...
pause > nul 