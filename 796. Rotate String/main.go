package main

import "fmt"

func rotate(s string) string {

	t := ""

	for _, v := range s {
		t = string(v) + t
	}

	return t
}

func rotateString(s string, goal string) bool {

	if goal == rotate(s) {
		return true
	}

	if len(goal) != len(s) {

		return false
	}

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {

			if s[i] == goal[j] {

				if s[i:]+s[0:i] == goal {
					return true
				}
			}
		}

	}

	return false
}

func main() {

	s := "abcde"
	goal := "cdeab"

	fmt.Println(rotateString(s, goal))

}
