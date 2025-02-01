package main

import "fmt"

func countSubarrays(nums []int) int64 {
	var count int64 = 0

	for i := 0; i < len(nums); i++ {

		j := i + 1
		for j < len(nums) && nums[j-1] < nums[j] {
			j++
		}

		n := j - i
		count += int64(n * (n + 1) / 2)

		i = j - 1
	}

	return count
}

func main() {
	fmt.Println(countSubarrays([]int{1, 3, 5, 4, 4, 6}))
	//
	fmt.Println(countSubarrays([]int{1, 2, 3, 4, 5}))
}
