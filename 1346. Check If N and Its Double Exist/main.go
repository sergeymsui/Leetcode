package main

import "fmt"

func checkIfExist(arr []int) bool {

	storage := make(map[int][]int)

	for i := 0; i < len(arr); i++ {
		storage[arr[i]] = append(storage[arr[i]], i)
	}

	for i := 0; i < len(arr); i++ {

		if pos, e := storage[2*arr[i]]; e {

			for _, v := range pos {
				if v != i {
					return true
				}
			}

		}
	}

	return false
}

func main() {

	arr := []int{
		-2, 0, 10, -19, 4, 6, -8,
	}

	fmt.Println(checkIfExist(arr))
}
