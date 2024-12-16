package main

import (
	"fmt"
	"strconv"
)

func combine(digits []int, keyboard map[int]string, prefix string) {

	if len(digits) == 0 {
		fmt.Printf("%s ", prefix)
		return
	}

	words := keyboard[digits[0]]

	for _, w := range words {
		n_prefix := prefix + string(w)
		combine(digits[1:], keyboard, n_prefix)
	}
}

func main() {

	keyboard := make(map[int]string)

	keyboard[2] = "abc"
	keyboard[3] = "def"
	keyboard[4] = "ghi"
	keyboard[5] = "jkl"
	keyboard[6] = "mno"
	keyboard[7] = "pqrs"
	keyboard[8] = "tuv"
	keyboard[9] = "wxyz"

	input := ""

	fmt.Scan(&input)

	digits := []int{}

	for _, s := range input {
		d, _ := strconv.Atoi(string(s))
		digits = append(digits, d)
	}

	combine(digits, keyboard, "")
}
