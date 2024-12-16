package main

import "fmt"

func sortColors(nums []int) {

	i := 1

	for i < len(nums) {

		for j := i - 1; j >= 0 && nums[j] > nums[j+1]; j-- {
			nums[j], nums[j+1] = nums[j+1], nums[j]
		}

		i++
	}

}

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}

	sortColors(nums)

	fmt.Println(nums)
}
