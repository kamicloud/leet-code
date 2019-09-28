//Given a linked list, rotate the list to the right by k places, where k is non-negative. 
//
// Example 1: 
//
// 
//Input: 1->2->3->4->5->NULL, k = 2
//Output: 4->5->1->2->3->NULL
//Explanation:
//rotate 1 steps to the right: 5->1->2->3->4->NULL
//rotate 2 steps to the right: 4->5->1->2->3->NULL
// 
//
// Example 2: 
//
// 
//Input: 0->1->2->NULL, k = 4
//Output: 2->0->1->NULL
//Explanation:
//rotate 1 steps to the right: 2->0->1->NULL
//rotate 2 steps to the right: 1->2->0->NULL
//rotate 3 steps to the right: 0->1->2->NULL
//rotate 4 steps to the right: 2->0->1->NULL 
//

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */




package main


func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	temp := head
	l := 0
	for {
		if temp == nil {
			break
		}
		l++
		temp = temp.Next
	}

	k = k % l
	temp = head

	if k == 0 {
		return head
	}

	var end1 *ListNode = nil

	for i := 0; i <= (l - k) - 1; i++ {
		if i == (l - k) - 1 {
			end1 = temp
		}

		temp = temp.Next
	}
	head1 := temp


	for {
		if temp.Next == nil {
			temp.Next = head
			break
		}
		temp = temp.Next
	}
	end1.Next = nil

	return head1
}