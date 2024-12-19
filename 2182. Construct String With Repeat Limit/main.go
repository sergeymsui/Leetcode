// !!!

package main

import (
	"fmt"
	"strings"
)

func repeatLimitedString(s string, repeatLimit int) string {
	dict := make(map[rune]int)

	for _, v := range s {
		dict[v]++
	}

	var builder strings.Builder

	current_char_index := rune('z')

	for current_char_index >= 'a' { // Предполагаем, что работа с английскими буквами
		if dict[current_char_index] == 0 {
			current_char_index--
			continue
		}

		use := min(repeatLimit, dict[current_char_index])
		dict[current_char_index] -= use

		for i := 0; i < use; i++ {
			builder.WriteRune(current_char_index)
		}

		if dict[current_char_index] > 0 {
			smaller_char_index := current_char_index - 1

			for smaller_char_index >= 'a' && dict[smaller_char_index] == 0 {
				smaller_char_index--
			}

			if smaller_char_index < 'a' {
				break
			}

			builder.WriteRune(smaller_char_index)
			dict[smaller_char_index]--
		}
	}

	return builder.String()
}

func main() {
	s := "cczazcc"
	repeatLimit := 3

	fmt.Println(repeatLimitedString(s, repeatLimit))
}
