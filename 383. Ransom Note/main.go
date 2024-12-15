package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {

	dict := make(map[rune]int)

	for _, r := range magazine {
		dict[r]++
	}

	for _, r := range ransomNote {
		if v, e := dict[r]; !e || v == 0 {
			return false
		}

		dict[r]--
	}

	return true
}

func main() {
	ransomNote := "aal"
	magazine := "aab"

	fmt.Println(canConstruct(ransomNote, magazine))
}
