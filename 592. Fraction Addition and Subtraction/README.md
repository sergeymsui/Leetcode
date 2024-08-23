# Intuition

<!-- Describe your first thoughts on how to solve this problem. -->

The main problem, as it seemed, is the correct parsing of the source string. The problem can be represented as the sum of fractions `1/2-1/2 = 1/2+(-1/2)` etc. To find a common divisor for all fractions, you need to multiply the denominators.

# Approach

<!-- Describe your approach to solving the problem. -->

First, we parse the original string, selecting the numerators and denominators. At the same time, we find a divisor, albeit not the largest, but a common one. Then we add and reduce if necessary.

# Complexity

- Time complexity: O(2N)

- Space complexity: O(2N + const)

# Code

```golang []
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
```
