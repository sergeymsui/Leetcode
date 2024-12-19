package main

import "fmt"

func maxChunksToSorted(arr []int) int {

	prefixElements := make([]int, len(arr))
	copy(prefixElements, arr)

	suffixElements := make([]int, len(arr))
	copy(suffixElements, arr)

	for i := 1; i < len(arr); i++ {
		prefixElements[i] = max(prefixElements[i], prefixElements[i-1])
	}

	for i := len(arr) - 2; i >= 0; i-- {
		suffixElements[i] = min(suffixElements[i], suffixElements[i+1])
	}

	chunks := 1
	for i := 1; i < len(arr); i++ {
		if prefixElements[i-1] < suffixElements[i] {
			chunks++
		}
	}

	return chunks
}

func main() {
	arr := []int{1}

	fmt.Println(maxChunksToSorted(arr))
}
