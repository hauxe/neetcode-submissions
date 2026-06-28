/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseKGroup(head *ListNode, k int) *ListNode {
	var begin *ListNode = nil
	var end *ListNode = nil
	var prev *ListNode = nil
	current := head
	head = nil
	for current != nil {
		begin, current = current, move(current, k)
		if current == nil {
			return head
		}
		prev, end, current = end, current, current.Next
		reverse(begin, end)
		begin, end = end, begin
		if head == nil {
			head = begin
			prev = end
			prev.Next = current
		} else {
			prev, prev.Next, end.Next = end, begin, current
		}
	}
	return head
}

func move(head *ListNode, k int) *ListNode {
	for ; head != nil && k > 1; k-- {
		head = head.Next
	}
	return head
}

func reverse(begin, end *ListNode) {
	var current *ListNode = nil
	last := false
	for begin != nil {
		if begin == end {
			last = true
		}
		current, begin, begin.Next = begin, begin.Next, current
		if last {
			break
		}
	}
}
