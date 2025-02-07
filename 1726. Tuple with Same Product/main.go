package main

import (
	"fmt"
	"slices"
)

func tupleSameProduct(nums []int) int {

	slices.Sort(nums)
	n := len(nums)

	result := 0

	i := 0
	for i < n {
		j := i + 1
		for j < n {
			p := nums[i] * nums[j]

			range_set := make(map[int]bool)

			l := i + 1
			for l < j {

				k := nums[l]

				if p%k == 0 {
					if _, e := range_set[int(p/k)]; e {
						result += 8
					}
					range_set[k] = true
				}

				l++
			}

			j++
		}
		i++
	}

	return result
}

func main() {
	nums := []int{1, 4, 2, 5, 10}
	fmt.Println(tupleSameProduct(nums))
}
