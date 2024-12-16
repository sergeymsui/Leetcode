package main

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	// Создаем мапу для отслеживания клонированных узлов
	visited := make(map[*Node]*Node)

	// Вспомогательная функция для клонирования узла
	var clone func(*Node) *Node
    
	clone = func(n *Node) *Node {
		if n == nil {
			return nil
		}

		// Если узел уже клонирован, возвращаем его копию
		if clonedNode, found := visited[n]; found {
			return clonedNode
		}

		// Создаем новый узел
		clonedNode := &Node{Val: n.Val}
		visited[n] = clonedNode

		// Клонируем соседей
		for _, neighbor := range n.Neighbors {
			clonedNode.Neighbors = append(clonedNode.Neighbors, clone(neighbor))
		}

		return clonedNode
	}

	return clone(node)
}

func main() {

	n1 := &Node{
		Val: 1,
	}

	n2 := &Node{
		Val: 2,
	}

	n3 := &Node{
		Val: 3,
	}

	n4 := &Node{
		Val: 4,
	}

	n1.Neighbors = append(n1.Neighbors, n2)
	n2.Neighbors = append(n2.Neighbors, n3)
	n3.Neighbors = append(n3.Neighbors, n4)
	n4.Neighbors = append(n4.Neighbors, nil)

	cloneGraph(nil)
}
