//Given a linked list, remove the n-th node from the end of list and return its head. 
//
// Example: 
//
// 
//Given linked list: 1->2->3->4->5, and n = 2.
//
//After removing the second node from the end, the linked list becomes 1->2->3->5.
// 
//
// Note: 
//
// Given n will always be valid. 
//
// Follow up: 
//
// Could you do this in one pass? 
//

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	l := 0
	temp := head
	for {
		if temp == nil {
			break
		}
		temp = temp.Next
		l++
	}

	if n == l {
		return head.Next
	}

	temp = head

	for i := l - n; i > 1; i-- {
		temp = temp.Next
	}

	if n == 1 {
		temp.Next = nil
	} else {
		temp.Next = temp.Next.Next
	}

	return head
}
