package gotools

// Use var when outside of a function.
// If you don't set a value, zero value of the type is set.
// answer is therefore 0.

var answer int

// Add adds two numbers
func Add(x int, y int) int {
	return x + y
}
