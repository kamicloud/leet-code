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
        ListNode first = lists[0];
        ListNode temp = null;
        int i = 0;

        if (lists.length == 1) {
            return first;
        }

        for (i = 0; i < lists.length; i++) {
            if (lists[i] == null) {
                continue;
            }
            if (
                    (first == null) ||
                            first.val > lists[i].val
            ) {
                lists[0] = lists[i];
                lists[i] = first;
                first = lists[0];
            }
        }

        ListNode firstp = lists[0];

        while (firstp != null) {

            for (i = 1; i < lists.length; i++) {
                if (lists[i] == null) {
                    continue;
                }
                if (lists[i].val == firstp.val) {
                    temp = lists[i];
                    lists[i] = lists[i].next;
                    temp.next = firstp.next;
                    firstp.next = temp;
                } else if (lists[i].val > firstp.val) {
                    if (firstp.next == null) {
                        firstp.next = lists[i];
                        lists[i] = lists[i].next;
                        firstp.next.next = null;
                    } else {
                        if (firstp.next.val > lists[i].val) {
                            temp = lists[i];
                            lists[i] = lists[i].next;
                            temp.next = firstp.next;
                            firstp.next = temp;
                        }
                    }
                }
            }
            firstp = firstp.next;
        }

        return first;
    }

//    public class ListNode {
//        int val;
//        ListNode next;
//        ListNode(int x) {
//            val = x;
//        }
//    }
}