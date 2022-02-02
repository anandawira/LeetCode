package graph

type Node struct {
	Val       int
	Neighbors []*Node
}

var oldToNew map[*Node]*Node = make(map[*Node]*Node)

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	if val, ok := oldToNew[node]; ok {
		return val
	}

	newNode := &Node{Val: node.Val}
	oldToNew[node] = newNode

	for _, neighbor := range node.Neighbors {
		newNode.Neighbors = append(newNode.Neighbors, cloneGraph(neighbor))
	}

	return newNode
}
