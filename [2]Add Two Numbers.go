//You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list. 
//
// You may assume the two numbers do not contain any leading zero, except the number 0 itself. 
//
// Example: 
//
// 
//Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
//Output: 7 -> 0 -> 8
//Explanation: 342 + 465 = 807.
// 
//

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNumbersWithCarry(l1, l2, 0)
}

func addTwoNumbersWithCarry(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
	temp := ListNode{
		Val:  carry,
		Next: nil,
	}
	if l1 != nil {
		temp.Val += l1.Val
		l1 = l1.Next
	} else {
		l1 = nil
	}
	if l2 != nil {
		temp.Val += l2.Val
		l2 = l2.Next
	} else {
		l2 = nil
	}

	carry = temp.Val / 10

	temp.Val = temp.Val % 10

	if l1 != nil || l2 != nil || carry != 0 {
		temp.Next = addTwoNumbersWithCarry(l1, l2, carry)
	}

	return &temp
}
