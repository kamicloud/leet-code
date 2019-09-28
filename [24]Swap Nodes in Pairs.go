//Given a linked list, swap every two adjacent nodes and return its head.
//
// You may not modify the values in the list's nodes, only nodes itself may be changed.
//
//
//
// Example:
//
//
//Given 1->2->3->4, you should return the list as 2->1->4->3.
//
//

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

package main

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	temp := head

	if head.Next.Next == nil {

	}

	var left, leftP, right, rightP *ListNode
	count := 0


	for {
		if temp == nil {
			break
		}
		if count % 2 == 0 {
			if left == nil {
				left = temp
				leftP = temp
			} else {
				leftP.Next = temp
				leftP = leftP.Next
			}
			temp = temp.Next
			leftP.Next = nil
		} else {
			if right == nil {
				right = temp
				rightP = temp
			} else {
				rightP.Next = temp
				rightP = rightP.Next
			}
			temp = temp.Next
			rightP.Next = nil
		}
		count++
	}
	var res, resP *ListNode

	for i := 0; i < count; i++ {
		if left == nil || right == nil {
			break
		}
		if i % 2 == 0 {
			if res == nil {
				res = right
				resP = right
			} else {
				resP.Next = right
				resP = resP.Next
			}
			right = right.Next
		} else {
			if res == nil {
				res = left
				resP = left
			} else {
				resP.Next = left
				resP = resP.Next
			}
			left = left.Next
		}
	}

	if left != nil {
		resP.Next = left
	}

	return res
}