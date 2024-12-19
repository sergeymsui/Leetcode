package main

import "fmt"

func insertionSort(a []int) {

	for i := 1; i < len(a); i++ {

		for j := i; j > 0 && a[j-1] > a[j]; j-- {
			a[j], a[j-1] = a[j-1], a[j]
		}

	}

	fmt.Println(a)
}

func main() {

	a := []int{
		5, 3, 6, 4, 1, 5, 1, 1, 0, 0, 0, 1000, -8, -99, 1,
	}

	insertionSort(a)

}
