// 3163. String Compression III
// Given a string word, compress it using the following algorithm:

// Begin with an empty string comp. While word is not empty, use the following operation:
// Remove a maximum length prefix of word made of a single character c repeating at most 9 times.
// Append the length of the prefix followed by c to comp.
// Return the string comp.

package main

import (
	"fmt"
	"strconv"
)

func compressedString(word string) string {

	t := 0
	n := len(word)
	ans := []byte{}

	for t < n {

		ptr := t

		for ptr < len(word) && word[t] == word[ptr] && (ptr-t) < 9 {
			ptr++
		}

		ans = append(ans, []byte(strconv.Itoa(ptr-t))...)
		ans = append(ans, word[t])

		t = ptr
	}

	return string(ans)
}

func main() {

	s := "aaaaaaaaaaaaaabb"

	fmt.Println(compressedString(s))
}
