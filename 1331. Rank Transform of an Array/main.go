// 1331. Rank Transform of an Array

package main

import (
	"fmt"
	"sort"
)

func arrayRankTransform(arr []int) []int {

	orig := []int{}
	orig = append(orig, arr...)

	rankMap := make(map[int]int)

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	i := 1
	for _, v := range arr {
		if _, exists := rankMap[v]; !exists {
			rankMap[v] = i
			i++
		}
	}

	for i, v := range orig {
		arr[i] = rankMap[v]
	}

	return arr
}

func main() {

	a := []int{
		27, 46, -3, -36, 31, -14, -7, -36, 27, -14, 41, -40, 23,
	}

	res := arrayRankTransform(a)

	fmt.Println(res)

}
