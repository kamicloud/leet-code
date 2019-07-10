//Given a sorted linked list, delete all duplicates such that each element appear only once. 
//
// Example 1: 
//
// 
//Input: 1->1->2
//Output: 1->2
// 
//
// Example 2: 
//
// 
//Input: 1->1->2->3->3
//Output: 1->2->3
// 
//

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var current = head

	for {
		if current.Next != nil {
			if current.Next.Val == current.Val {
				current.Next = current.Next.Next
			} else {
				current = current.Next
			}
		} else {
			break
		}
	}


	return head
}