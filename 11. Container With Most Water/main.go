package main

import (
	"fmt"
)

func maxArea(height []int) int {

	max_square := 0

	i := 0
	j := len(height) - 1

	for i < j {

		s := min(height[i], height[j]) * (j - i)

		if s > max_square {
			max_square = s
		}

		if height[j] > height[i] {
			i++
		} else {
			j--
		}

	}

	return max_square
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}

	fmt.Println(maxArea(height))
}
