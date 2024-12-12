package main

import (
	"fmt"
	"strings"
	"unicode"
)


func isPalindrome(s string) bool {

	normalize := ""

	for _, s := range strings.ToLower(s) {

		if unicode.IsLetter(s) || unicode.IsDigit(s) {
			normalize += string(s)
		}

	}

	if len(normalize) == 0 {
        return true;
    }

	i := 0
	j := len(normalize) - 1
	f := true

	for i > j {
		f = f && (normalize[i] == normalize[j])

		i++
		j--
	}

	return f
}

func main() {
	s := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(s))
}
