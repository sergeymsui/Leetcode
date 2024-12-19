package main

import "fmt"

func isSubsequenc(s1 string, s2 string) bool {

	dict := make(map[rune]int)

	for _, v := range s2 {
		dict[v]++
	}

	for _, v := range s1 {

		if _, e := dict[v]; !e || dict[v] == 0 {
			return false
		}

		dict[v]--
	}

	return true
}

func main() {
	s1 := ""
	s2 := ""

	fmt.Scan(&s1)
	fmt.Scan(&s2)

	fmt.Println(isSubsequenc(s1, s2))
}
