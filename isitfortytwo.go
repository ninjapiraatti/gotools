package gotools

// IsItFortytwo tells you if the parameter is divisible by 42 and if it is, it tells the multiplier. Else it returns 0.
// The function exists really just to demonstrate shorthand/ternary.
func IsItFortytwo(nbr int) int {
	if res := nbr / 42; res%1 == 0 {
		return res
	}
	return 0
}
