package main

import (
	"fmt"
	"math"
	"sort"
)

func side(x1, x2, y1, y2 int) float64 {
	return math.Sqrt(float64((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1)))
}

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {

	points := [][]int{p1, p2, p3, p4}
	lengths := []float64{}

	for i := 0; i < len(points); i++ {
		for j := 0; j < len(points); j++ {
			if i == j {
				continue
			}

			a := points[i]
			b := points[j]

			lengths = append(lengths, side(a[0], b[0], a[1], b[1]))
		}
	}

	sort.Float64s(lengths)

	sideLenght := lengths[0]

	for _, v := range lengths[0:8] {
		if sideLenght != v {
			return false
		}
	}

	sideLenght = lengths[9]
	for _, v := range lengths[8:] {
		if sideLenght != v {
			return false
		}
	}

	return true
}

func main() {

	p1 := []int{0, 0}
	p2 := []int{1, 1}
	p3 := []int{1, 0}
	p4 := []int{0, 1}

	fmt.Println(validSquare(p1, p2, p3, p4))
}
