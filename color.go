package gotools

import "fmt"

const (
	Cyan    = "\033[0;96m%s\033[0m"
	Red     = "\033[0;91m%s\033[0m"
	Yellow  = "\033[0;93m%s\033[0m"
	Green   = "\033[0;92m%s\033[0m"
	Magenta = "\033[0;95m%s\033[0m"
)

// Color colors the input string
func Color(color string, text string) {
	fmt.Printf(color, text)
}
