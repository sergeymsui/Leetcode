// 2022. Convert 1D Array Into 2D Array
// You are given a 0-indexed 1-dimensional (1D) integer array original, and two integers, m and n. You are tasked with creating a 2-dimensional (2D) array with  m rows and n columns using all the elements from original.
// The elements from indices 0 to n - 1 (inclusive) of original should form the first row of the constructed 2D array, the elements from indices n to 2 * n - 1 (inclusive) should form the second row of the constructed 2D array, and so on.
// Return an m x n 2D array constructed according to the above procedure, or an empty 2D array if it is impossible.

package main

import "fmt"

func construct2DArray(original []int, m int, n int) [][]int {

	if n*m != len(original) {
		return [][]int{}
	}

	ans := [][]int{}

	s := 0
	for i := 0; i < m; i++ {

		row := []int{}
		for j := 0; j < n; j++ {
			row = append(row, original[s+j])
		}

		ans = append(ans, row)
		s += n

	}

	return ans
}

func main() {
	original := []int{1, 2, 3, 4}
	m := 2
	n := 2

	fmt.Println(construct2DArray(original, m, n))
}

