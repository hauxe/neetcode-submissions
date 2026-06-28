/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
	m := make(map[*ListNode]struct{})
	for head != nil {
		if _, exist := m[head]; exist {
			return true
		}
		m[head] = struct{}{}
		head = head.Next
	}
	return false
}
