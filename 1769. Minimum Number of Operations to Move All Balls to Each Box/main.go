package main

import (
	"fmt"
)

func minOperations(boxes string) []int {

	res := make([]int, len(boxes))

	mv_to_left := 0
	mv_to_right := 0
	balls_to_left := 0
	balls_to_right := 0

	for i := range boxes {

		res[i] += mv_to_left
		balls_to_left += int(boxes[i] - '0')
		mv_to_left += balls_to_left

		j := len(boxes) - 1 - i
		res[j] += mv_to_right
		balls_to_right += int(boxes[j] - '0')
		mv_to_right += balls_to_right
	}

	return res
}

func main() {
	boxes := "001011"
	fmt.Println(minOperations(boxes))
}
