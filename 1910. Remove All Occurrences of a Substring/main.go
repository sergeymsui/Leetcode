package main

import "fmt"

func removeOccurrences(s string, part string) string {

	i := 0
	for i < len(s) {
		j := i
		r := 0

		for j < len(s) && r < len(part) && s[j] == part[r] {
			r++
			j++
		}

		if len(part) == r {
			s = s[:i] + s[j:]
			i -= 2
		} else {
			i++
		}

	}

	return s
}

func main() {
	// s = "daabcbaabcbc", part = "abc"
	fmt.Println(removeOccurrences("daabcbaabcbc", "abc"))
	fmt.Println(removeOccurrences("ixcupqoixcupqokevnpokevnpoknqywmlhevgc", "ixcupqokevnpo"))

}
