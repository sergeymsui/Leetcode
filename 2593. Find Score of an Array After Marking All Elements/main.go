package main

import "fmt"

func min(nums []int) ([]int, int) {

	m := nums[0]
	pos := 0

	for i, v := range nums {
		if v < m && v > 0 {
			m = v
			pos = []int{i}
		} else if v == m {
			pos = append(pos, i)
		}
	}

	return pos, m
}

func findScore(nums []int) int64 {
	var score int64

	marked_num := 0

	for marked_num < len(nums) {
		pos, min_value := min(nums)

		score += int64(nums[min_index])

		marked_num++

		indexes[min_value] = append(indexes[0, min_pos])
	}

	return score
}

func main() {
	nums := []int{2, 1, 3, 4, 5, 2}

	fmt.Println(findScore(nums))
}
