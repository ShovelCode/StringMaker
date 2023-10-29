package main

import (
	"fmt"
	"strings"
)

// RepeatString takes a string `str` and an integer `n`
// and returns a new string which is `str` repeated `n` times.
func RepeatString(str string, n int) string {
	var result strings.Builder
	for i := 0; i < n; i++ {
		result.WriteString(str)
	}
	return result.String()
}

func main() {
	str := "Go"
	n := 3
	result := RepeatString(str, n)
	fmt.Println("Result:", result) // Output: Result: GoGoGo
}

