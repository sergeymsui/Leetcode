// 592. Fraction Addition and Subtraction

// Given a string expression representing an expression of fraction addition and subtraction, return the calculation result in string format.

// The final result should be an irreducible fraction. If your final result is an integer, change it to the format of a fraction that has a denominator 1. So in this case, 2 should be converted to 2/1.

package main

import (
	"fmt"
	"strconv"
)

func fractionAddition(expression string) string {

	// Greatest common divisor
	gcd := 1

	numerators := []int{}
	denominators := []int{}

	// last symbol position
	l := 0

	// is numerator or not
	f := false

	for i, v := range expression {
		if v == '/' {
			numerator := expression[l:i]
			l = i + 1

			if d, e := strconv.Atoi(numerator); e == nil {
				numerators = append(numerators, d)
			}

			f = true
		}

		if f && (v == '+' || v == '-') {
			denominator := expression[l:i]
			l = i

			if d, e := strconv.Atoi(denominator); e == nil {
				denominators = append(denominators, d)
				gcd *= d
			}

			f = false
		} else if i == len(expression)-1 {
			denominator := expression[l : i+1]

			if d, e := strconv.Atoi(denominator); e == nil {
				denominators = append(denominators, d)
				gcd *= d
			}
		}
	}

	summ := 0
	for i := 0; i < len(denominators); i++ {
		summ += numerators[i] * (gcd / denominators[i])
	}

	if summ == 0 {
		return "0/1"
	} else {

		for i := 1; i <= gcd; i++ {
			if (summ%i == 0) && (gcd%i == 0) {
				summ = summ / i
				gcd = gcd / i

				i = 1
			}
		}

	}

	return fmt.Sprintf("%d/%d", summ, gcd)
}

func main() {
	fmt.Println(fractionAddition("-1/2+1/2"))
	fmt.Println(fractionAddition("1/2-1/2+1/3"))
	fmt.Println(fractionAddition("1/3-1/2"))
}
