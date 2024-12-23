package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	s := strings.Builder{}
	Solution(os.Stdin, &s)

	fmt.Println(s.String())
}

func dfs(graph map[int][]int, colors []string, v int, stack *[]int) {
	colors[v] = "gray"

	for _, c := range graph[v] {
		if colors[c] == "white" {
			dfs(graph, colors, c, stack)
		}
	}

	colors[v] = "black"
	*stack = append(*stack, v)
}

func Solution(r io.Reader, s *strings.Builder) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()

	initData := strings.Fields(scanner.Text())
	n, _ := strconv.Atoi(initData[0])
	m, _ := strconv.Atoi(initData[1])

	graph := make(map[int][]int)

	for i := 0; i < m; i++ {
		scanner.Scan()
		data := strings.Fields(scanner.Text())
		k, _ := strconv.Atoi(data[0])
		l, _ := strconv.Atoi(data[1])

		graph[k] = append(graph[k], l)
	}

	colors := []string{}
	stack := []int{}

	for i := 0; i < n+1; i++ {
		colors = append(colors, "white")
	}

	for i := n; i > 0; i-- {
		if colors[i] == "white" {
			dfs(graph, colors, i, &stack)
		}
	}
	
	for len(stack) > 0 {
		fmt.Printf("%d ", stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}
}
