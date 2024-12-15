package main

import "fmt"

func isAnagram(s string, t string) bool {

	dict := make(map[rune]int)

	for _, r := range s {
		dict[r]++
	}

	for _, r := range t {

		if v, e := dict[r]; !e || v == 0 {
			return false
		} else {
			dict[r]--
		}
	}

	for _, v := range dict {
		if v > 0 {
			return false
		}
	}

	return true
}

func main() {
	s := "care"
	t := "car"
	fmt.Println(isAnagram(s, t))
}
