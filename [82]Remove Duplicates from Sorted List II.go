//Given a sorted linked list, delete all nodes that have duplicate numbers, leaving only distinct numbers from the original list. 
//
// Example 1: 
//
// 
//Input: 1->2->3->3->4->4->5
//Output: 1->2->5
// 
//
// Example 2: 
//
// 
//Input: 1->1->1->2->3
//Output: 2->3
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

//type ListNode struct {
//	Val int
//	Next *ListNode
//}



func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	temp := map[int]int{}
	temp2 := head

	for {

		if temp2 == nil {
			break
		} else {
			//temp[temp2.Val] = 0
			_, exists := temp[temp2.Val]

			if exists {
				temp[temp2.Val] = 1
				//temp2.Next = temp2.Next.Next
			} else {
				temp[temp2.Val] = 0
			}
			temp2 = temp2.Next

		}




	}

	for {
		remove, exists := temp[head.Val]

		if exists && remove == 1{
			head = head.Next
		} else {
			break
		}
		if head == nil {
			return head
		}
	}
	if head == nil || head.Next == nil {
		return head
	}

	temp2 = head

	for {
		if temp2 == nil || temp2.Next == nil {
			break
		}
		remove, exists := temp[temp2.Next.Val]

		if exists && remove == 1{
			temp2.Next = temp2.Next.Next
		} else {
			temp2 = temp2.Next
		}
	}




	return head
}