// 476. Number Complement

// The complement of an integer is the integer you get when you flip all the 0's to 1's and all the 1's to 0's in its binary representation.

// For example, The integer 5 is "101" in binary and its complement is "010" which is the integer 2.
// Given an integer num, return its complement.

package main

import "fmt"

func findComplement(num int) int {
	mask := 0
	temp := num

	for temp != 0 {
		mask = (mask << 1) | 1
		temp >>= 1
	}

	return num ^ mask
}

func main() {
	fmt.Println(findComplement(15))
}
