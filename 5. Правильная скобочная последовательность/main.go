package main

import (
	"bufio"
	"fmt"
	"os"
)

func solution(sequence string) bool {

	m_stack := []rune{}

	for _, v := range sequence {

		last := len(m_stack) - 1

		if v == ')' && last >= 0 && m_stack[last] == '(' {
			m_stack = m_stack[:last]
		} else if v == ']' && last >= 0 && m_stack[last] == '[' {
			m_stack = m_stack[:last]
		} else if v == '}' && last >= 0 && m_stack[last] == '{' {
			m_stack = m_stack[:last]
		} else {
			m_stack = append(m_stack, v)
		}
	}

	return len(m_stack) == 0
}

func main() {
    const maxCapacity = 512*1024  
    buf := make([]byte, maxCapacity)


	scanner := bufio.NewScanner(os.Stdin)
    scanner.Buffer(buf, maxCapacity)
	scanner.Scan()

	sequence := scanner.Text()

	if solution(sequence) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
