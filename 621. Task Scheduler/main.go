!!!

package main

import "fmt"

func leastInterval(tasks []byte, n int) int {
	indexes := make(map[byte]int)

	for _, v := range tasks {
		indexes[v]++
	}

	return 0
}

func main() {
	tasks := []byte{'A', 'B', 'C', 'D', 'E', 'A', 'B', 'C', 'D', 'E'}
	// tasks := []byte{'A', 'A', 'A', 'B', 'B', 'B'}
	n := 4

	fmt.Println(leastInterval(tasks, n))
}
