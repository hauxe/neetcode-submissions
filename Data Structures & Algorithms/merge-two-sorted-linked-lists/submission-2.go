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
		if list1 == nil {
			if current == nil {
				return list2
			}
			current.Next = list2
			return result
		}
		if list2 == nil {
			if current == nil {
				return list1
			}
			current.Next = list1
			return result
		}
		next := list1
		if list2.Val < next.Val {
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
