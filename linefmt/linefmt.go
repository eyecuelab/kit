package linefmt

import "strings"
import "fmt"

var lastLine string

const consoleWidth = 120

//Clear n characters from the current line.
func Clear(n int) bool {
	if n < 0 {
		return false
	}
	fmt.Print(strings.Repeat("\b", n))
	return true
}

//ClearAll clears the current line of consoleWidth (default 120) characters.
func ClearAll() {
	fmt.Print("\r", strings.Repeat(" ", consoleWidth), "\r")
}

//Printf clears the current line, then prints using fmt.Printf
func Printf(format string, args ...interface{}) (int, error) {
	ClearAll()
	return fmt.Printf(format, args...)
}

//Print clears the current line, then prints using fmt.Print
func Print(args ...interface{}) (int, error) {
	ClearAll()
	return fmt.Print(args...)
}
