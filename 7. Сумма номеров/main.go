package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findArrays(narrays []int, n int, k int) int {

	i := 0
	cnums := 0

	for i < n {

		s := 0
		j := i

		for j < n && s < k {

			s += narrays[j]
			j++
		}

		if s == k {
			cnums++
		}

		i++
	}

	return cnums
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	const maxCapacity = 2 * 100000
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	scanner.Scan()
	dsets := strings.Split(scanner.Text(), " ")

	n, _ := strconv.Atoi(dsets[0])
	k, _ := strconv.Atoi(dsets[1])

	scanner.Scan()
	nsets := strings.Split(scanner.Text(), " ")

	narrays := []int{}

	for _, v := range nsets {
		d, _ := strconv.Atoi(v)
		narrays = append(narrays, d)
	}

	fmt.Println(findArrays(narrays, n, k))
}
