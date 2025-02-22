package main

import "fmt"

type ProductOfNumbers struct {
	arr []int
}

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{
		arr: []int{},
	}
}

func (this *ProductOfNumbers) Add(num int) {
	this.arr = append(this.arr, num)
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	n := len(this.arr)

	if n == 0 {
		return 0
	}

	m := 1

	for i := n - k; i < n; i++ {
		m *= this.arr[i]
	}

	return m
}

func main() {
	// ["ProductOfNumbers","add","add","add","add","add","getProduct","getProduct","getProduct","add","getProduct"]
	// [[],[3],[0],[2],[5],[4],[2],[3],[4],[8],[2]]
	p := Constructor()
	p.Add(3)
	p.Add(0)
	p.Add(2)
	p.Add(5)
	p.Add(4)

	fmt.Println(p.GetProduct(2))
	fmt.Println(p.GetProduct(3))
	fmt.Println(p.GetProduct(4))
	p.Add(8)
	fmt.Println(p.GetProduct(2))
}
