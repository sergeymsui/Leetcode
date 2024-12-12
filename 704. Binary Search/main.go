package main

import "fmt"

func find(nums []int, target int, pos int) int {
	if len(nums) == 1 {

		if nums[0] == target {
			return pos
		}

		return -1
	}

	d := int(len(nums) / 2)

	if v := find(nums[:d], target, pos); v != -1 {
		return v
	} else {
		return find(nums[d:], target, d+pos)
	}
}

func search(nums []int, target int) int {
	return find(nums, target, 0)
}

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9

	fmt.Println(search(nums, target))
}
