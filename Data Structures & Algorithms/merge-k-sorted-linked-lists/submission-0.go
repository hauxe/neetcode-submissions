/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeKLists(lists []*ListNode) *ListNode {
	var result *ListNode
	var current *ListNode
	var lastNode *ListNode
	nums := len(lists)
	for nums > 0 {
		nums = 0
		minIdx := -1
		for i, list := range lists {
			if list == nil {
				continue
			}
			nums++
			if minIdx == -1 {
				lastNode = list
				minIdx = i
				continue
			}
			if list.Val < lastNode.Val {
				lastNode = list
				minIdx = i
			}
		}
		if nums == 1 {
			if current == nil {
				return lastNode
			}
			current.Next = lastNode
			return result
		}
		if minIdx >= 0 {
			lists[minIdx] = lists[minIdx].Next
			if lists[minIdx] == nil {
				nums--
			}
			if result == nil {
				result, current = lastNode, lastNode
				continue
			}
			current, current.Next = lastNode, lastNode
		}

	}

	return result
}
