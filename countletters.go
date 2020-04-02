package gotools

// Countletters counts the occurrences of rune c
func Countletters(str string, c byte) int {
	count := 0
	for i := range str {
		if str[i] == c {
			count++
		}
	}
	return count
}
