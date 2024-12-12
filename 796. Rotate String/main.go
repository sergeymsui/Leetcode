package main

import "fmt"

// rotate reverses the input string s and returns the result
func rotate(s string) string {
	t := "" // Initialize an empty string to store the reversed version of s

	// Loop through each character in s, prepending it to t
	for _, v := range s {
		t = string(v) + t
	}

	return t // Return the reversed string
}

// rotateString checks if the string goal can be obtained by rotating the string s
func rotateString(s string, goal string) bool {
	// If goal is the reverse of s, return true immediately
	if goal == rotate(s) {
		return true
	}

	// If s and goal have different lengths, they cannot be rotations of each other
	if len(goal) != len(s) {
		return false
	}

	// Loop through each character in s to attempt rotation matching with goal
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			// Check if the characters at positions i in s and j in goal match
			if s[i] == goal[j] {
				// If a rotation of s starting from i matches goal, return true
				if s[i:]+s[0:i] == goal {
					return true
				}
			}
		}
	}

	// If no rotation of s matches goal, return false
	return false
}

func main() {

	s := "abcde"
	goal := "cdeab"

	fmt.Println(rotateString(s, goal))

}
