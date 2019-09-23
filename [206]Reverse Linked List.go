//Reverse a singly linked list. 
//
// Example: 
//
// 
//Input: 1->2->3->4->5->NULL
//Output: 5->4->3->2->1->NULL
// 
//
// Follow up: 
//
// A linked list can be reversed either iteratively or recursively. Could you implement both? 
//

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


func reverseList(head *ListNode) *ListNode {

	if head == nil  || head.Next == nil {
		return head
	}

	rl, tail := doReverseList(head.Next)
	head.Next = nil
	tail.Next = head

	return rl
}

func doReverseList(head *ListNode) (*ListNode, *ListNode) {
	if head == nil  || head.Next == nil {
		return head, head
	} else if head.Next.Next == nil {
		head.Next.Next = head
		return head.Next, head
	}

	rl, tail := doReverseList(head.Next)
	head.Next = nil
	tail.Next = head

	return rl, tail.Next
}