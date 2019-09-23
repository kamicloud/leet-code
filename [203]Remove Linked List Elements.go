//Remove all elements from a linked list of integers that have value val. 
//
// Example: 
//
// 
//Input:  1->2->6->3->4->5->6, val = 6
//Output: 1->2->3->4->5
// 
//

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	for {
		if head == nil || head.Val != val {
			break
		}
		head = head.Next
	}

	if head == nil || head.Next == nil {
		return head
	}

	temp := head

	for {
		if temp == nil || temp.Next == nil {
			break
		}

		if temp.Next.Val == val {
			temp.Next = temp.Next.Next
		} else {
			temp = temp.Next
		}
	}

	return head
}
