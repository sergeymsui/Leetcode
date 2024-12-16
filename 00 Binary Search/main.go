package main

import "fmt"

func binarySearch(a []int, target int, l int, r int) int {

	if len(a) == 0 {
		return -1
	}

	if r-l == 1 {
		if a[l] == target {
			return a[l]
		}

		return -1
	}

	diff := int((r - l) / 2)

	if lv := binarySearch(a, target, l, l+diff); lv > -1 {
		return lv
	}

	if rv := binarySearch(a, target, l+diff, r); rv > -1 {
		return rv
	}

	return -1
}

func main() {

	a := []int{
		5, 3, 5, 88, 1000, 5, 7, 8, 9, 4, 5, 50, 99, 100,
	}

	fmt.Println(binarySearch(a, 8, 0, len(a)))

}
