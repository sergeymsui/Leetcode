package main

import "fmt"

func mergeSort(a []int) []int {

	if len(a) <= 1 {
		return a
	}

	diff := int(len(a) / 2)

	r_arr := mergeSort(a[:diff])
	l_arr := mergeSort(a[diff:])

	res := []int{}

	i := 0
	j := 0
	for i < len(l_arr) || j < len(r_arr) {

		if l_arr[i] < r_arr[j] {
			res = append(res, l_arr[i])
			i++
		} else {
			res = append(res, r_arr[j])
			j++
		}

		if i == len(l_arr) {
			res = append(res, r_arr[j:]...)
			break
		}

		if j == len(r_arr) {
			res = append(res, l_arr[i:]...)
			break
		}
	}

	return res
}

func main() {
	a := []int{
		1, 5, 8, 4, 100, -520, 78, 8, 8, 8, 8, 4, 45, 4, 1, 0,
	}

	fmt.Println(mergeSort(a))
}
