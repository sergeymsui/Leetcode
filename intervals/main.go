package main

import (
	"fmt"
	"sort"
)

func solution(intervals [][]int) [][]int {

	new_intervals := [][]int{}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] <= intervals[j][0]
	})

	i := 0
	for i < len(intervals) {

		j := i + 1

		start := intervals[i][0]
		end := intervals[i][1]

		for j < len(intervals) && intervals[j][0] <= intervals[i][1] {
			end = intervals[j][1]

			i = j
			j++
		}

		new_interval := []int{
			start, end,
		}

		new_intervals = append(new_intervals, new_interval)

		i++
	}

	return new_intervals

}

func main() {

	intervals := [][]int{
		{2, 6}, {1, 3}, {8, 10}, {15, 18},
	}

	for _, s := range solution(intervals) {
		fmt.Println(s)
	}

}
