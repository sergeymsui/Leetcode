package main

import "fmt"

func OddCount(n int) int {

	if n%2 > 0 {
		n--
	}

	res := int(n / 2)

	return res
}

func main() {
	fmt.Println(OddCount(7))
}
