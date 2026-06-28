/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var result *ListNode = nil
	var current *ListNode = nil
	for list1 != nil || list2 != nil {
		next := list1
		if next == nil || (list2 != nil && list2.Val < next.Val) {
			next, list2 = list2, list2.Next
		} else {
			list1 = list1.Next
		}
		if result == nil {
			result = next
			current = next
			continue
		}
		current, current.Next = next, next

	}
	return result
}
