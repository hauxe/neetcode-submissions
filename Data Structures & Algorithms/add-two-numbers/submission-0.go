/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var result *ListNode = nil
	var current *ListNode = nil
	remain := 0
	for l1 != nil || l2 != nil {
		if l2 == nil {
			l1, l2 = l2, l1
		}
		if l1 == nil {
			if remain == 0 {
				if result == nil {
					return l2
				}
				current.Next = l2
				return result
			}
			l2.Val += remain
			l2.Val, remain = l2.Val%10, l2.Val/10
			current, current.Next, l2 = l2, l2, l2.Next
			continue
		}
		v := l1.Val + l2.Val + remain
		l1.Val, remain = v%10, v/10
		if result == nil {
			result = l1
			current = l1
		} else {
			current, current.Next = l1, l1
		}
		l1, l2 = l1.Next, l2.Next
	}
	if remain > 0 {
		current.Next = &ListNode{
			Val: remain,
		}
	}
	return result
}
