/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	var current *ListNode = nil
	for head != nil {
		head, head.Next, current = head.Next, current, head
	}
	return current
}
