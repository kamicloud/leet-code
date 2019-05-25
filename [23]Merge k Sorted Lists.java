//Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity. 
//
// Example: 
//
// 
//Input:
//[
//  1->4->5,
//  1->3->4,
//  2->6
//]
//Output: 1->1->2->3->4->4->5->6
// 
//

/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */
class Solution {
    public ListNode mergeKLists(ListNode[] lists) {
        if (lists.length == 0) {
            return null;
        }
        return mergeLists(lists, 0, lists.length - 1);
    }

    public ListNode mergeLists(ListNode[] lists, int start, int end) {
        if (start == end) {
            return lists[start];
        }
        int mid = (start + end) / 2;

        ListNode left = mergeLists(lists, start, mid);
        ListNode right = mergeLists(lists, mid + 1, end);

        return mergeTwoNodes(left, right);
    }

    public ListNode mergeTwoNodes(ListNode left, ListNode right) {
        if (left == null) {
            return right;
        }
        if (right == null) {
            return left;
        }

        if (left.val < right.val) {
            left.next = mergeTwoNodes(left.next, right);
            return left;
        } else {
            right.next = mergeTwoNodes(left, right.next);
            return right;
        }
    }

//    public class ListNode {
//        int val;
//        ListNode next;
//        ListNode(int x) {
//            val = x;
//        }
//    }
}