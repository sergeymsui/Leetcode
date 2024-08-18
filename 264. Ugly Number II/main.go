// 264. Ugly Number II
// An ugly number is a positive integer whose prime factors are limited to 2, 3, and 5.

// Given an integer n, return the nth ugly number.

package main

import (
	"fmt"
)

func nthUglyNumber(n int) int {

	minors := []int{2, 3, 5}
	indices := []int{0, 0, 0}

	ugly := []int{1}

	for i := 1; i < n; i++ {

		nextUgly := []int{
			ugly[indices[0]] * minors[0],
			ugly[indices[1]] * minors[1],
			ugly[indices[2]] * minors[2],
		}

		minUgly := min(nextUgly[0], nextUgly[1], nextUgly[2])
		ugly = append(ugly, minUgly)

		for j := 0; j < 3; j++ {
			if nextUgly[j] == minUgly {
				indices[j]++
			}
		}
	}

	return ugly[n-1]
}

func main() {
	fmt.Println(nthUglyNumber(19))
}
