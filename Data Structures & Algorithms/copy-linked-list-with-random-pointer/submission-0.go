/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	m := make(map[*Node]*Node)
	var newHead *Node = nil
	var newCurrent *Node = nil
	current := head

	for current != nil {
		node := &Node{
			Val:    current.Val,
			Next:   nil,
			Random: nil,
		}
		if newCurrent == nil {
			newHead, newCurrent = node, node
		} else {
			newCurrent, newCurrent.Next = node, node
		}
		m[current] = node
		current = current.Next
	}
	current = head
	for current != nil {
		if current.Random != nil {
			m[current].Random = m[current.Random]
		}
		current = current.Next
	}

	return newHead
}
