/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
	slow, fast := head, head
	var current *ListNode = nil
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		current, slow, slow.Next = slow, slow.Next, current
	}
	prev := slow
	for slow.Next != nil {
		if fast != nil { // odd
			prev, current, current.Next, slow.Next, slow.Next.Next = current, current.Next, slow.Next, slow.Next.Next, prev
		} else { // even
			prev, current, current.Next, slow.Next, slow.Next.Next = slow.Next, current.Next, prev, slow.Next.Next, current
		}
	}
	if current != nil {
		current, current.Next = nil, prev
	}
}
