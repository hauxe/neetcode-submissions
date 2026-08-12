func cloneGraph(node *Node) *Node {
	cloned := make(map[int]*Node)
	return doCloneGraph(node, cloned)
}

func doCloneGraph(node *Node, cloned map[int]*Node) *Node {
	if node == nil {
		return nil
	}
	if cloned[node.Val] != nil {
		return cloned[node.Val]
	}
	n1 := &Node{
		Val: node.Val,
	}
	cloned[n1.Val] = n1
	var neighbors []*Node
	for i := range node.Neighbors {
		n := doCloneGraph(node.Neighbors[i], cloned)
		neighbors = append(neighbors, n)
	}
	n1.Neighbors = neighbors
	return n1
}