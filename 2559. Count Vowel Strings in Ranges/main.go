package main

import "fmt"

func isVowel(r byte) bool {
	vowel := []byte{'a', 'e', 'i', 'o', 'u'}

	for _, v := range vowel {
		if v == r {
			return true
		}
	}
	return false
}

func vowelStrings(words []string, queries [][]int) []int {

	res := []int{}

	lrValues := make([]int, len(words)+1)
	lvalue := 0

	for i, v := range words {
		if isVowel(v[0]) && isVowel(v[len(v)-1]) {
			lvalue++
		}
		lrValues[i+1] = lvalue
	}

	for i := 0; i < len(queries); i++ {

		s := queries[i][0]
		f := queries[i][1]

		res = append(res, lrValues[f+1]-lrValues[s])
	}

	return res
}

func main() {

	words := []string{"aba", "bcb", "ece", "aa", "e"}
	queries := [][]int{{0, 2}, {1, 4}, {1, 1}}

	fmt.Println(vowelStrings(words, queries))
}
