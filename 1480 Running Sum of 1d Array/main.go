package main

import "fmt"

func runningSum(nums []int) []int {

	ans := []int{}

	s := 0
	for i := 0; i < len(nums); i++ {

		s += nums[i]
		ans = append(ans, s)

	}

	return ans
}

func main() {

	a := []int{1, 2, 3, 4}

	fmt.Println(runningSum(a))
}
