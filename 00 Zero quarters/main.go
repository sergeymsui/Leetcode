package main

import "fmt"

func solution(n int, p string) {
	if n == 0 {
		fmt.Println(p)

		return
	}

	solution(n-1, p+"0")
	solution(n-1, p+"1")
}

func main() {
	n := 0

	fmt.Scan(&n)

	solution(n, "")
}
