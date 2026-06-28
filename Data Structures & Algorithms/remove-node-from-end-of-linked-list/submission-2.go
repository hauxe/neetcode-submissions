func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var current *ListNode = nil
	for head != nil {
		current, head, head.Next = head, head.Next, current
	}
	for i := 1; current != nil; i++ {
		if i == n {
			current = current.Next
			// fmt.Println("current")
			// printNodes(current)
			continue
		}
		head, current, current.Next = current, current.Next, head
		// fmt.Println("head")
		// printNodes(head)
	}
	return head
}