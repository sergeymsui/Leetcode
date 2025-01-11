package main

import (
	"fmt"
	"math"
)

func Log(base, x float64) float64 {
	return math.Log(x) / math.Log(base)
}

func smallestNumber(n int) int {
	power := math.Ceil(Log(2, float64(n)))
	return int(math.Pow(2, power)) - 1
}

func main() {
	fmt.Println(smallestNumber(3))
}
