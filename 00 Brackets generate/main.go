package main

import "fmt"

func bracketsGenerate(br, l, r int, s string) {

	if l == 0 && r == 0 {
		fmt.Println(s)
	}

	if l > 0 {
		bracketsGenerate(br+1, l-1, r, s+"(")
	}

	if br > 0 && r > 0 {
		bracketsGenerate(br-1, l, r-1, s+")")
	}

}

func main() {
	n := 0

	fmt.Scan(&n)

	bracketsGenerate(0, n, n, "")
}
