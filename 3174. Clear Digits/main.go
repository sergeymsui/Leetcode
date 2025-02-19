package main

import (
	"fmt"
	"unicode"
)

func clearDigits(s string) string {

	for i := 0; i < len(s); i++ {

		n := i
		if unicode.IsDigit(rune(s[i])) {

			if n-1 >= 0 {
				n--
			}

			// fmt.Println(s[:n], s[i+1:])
			s = s[:n] + s[i+1:]

			i = n - 1
		}
	}

	return s
}

func main() {
	fmt.Println(clearDigits("cqo0"))
	fmt.Println(clearDigits("cb34"))
}
