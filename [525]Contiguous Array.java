//Given a binary array, find the maximum length of a contiguous subarray with equal number of 0 and 1. 
//
//
// Example 1: 
// 
//Input: [0,1]
//Output: 2
//Explanation: [0, 1] is the longest contiguous subarray with equal number of 0 and 1.
// 
// 
//
// Example 2: 
// 
//Input: [0,1,0]
//Output: 2
//Explanation: [0, 1] (or [1, 0]) is a longest contiguous subarray with equal number of 0 and 1.
// 
// 
//
// Note:
//The length of the given binary array will not exceed 50,000.
//

import java.util.HashMap;

class Solution {
    public int findMaxLength(int[] nums) {
    	HashMap<Integer, Integer> hash = new HashMap<Integer, Integer>();
    	hash.put(0, -1);
        int i = 0, j = 0;
        int one = 0, two = 0;
        int count = 0;
        int maxCount = 0;
        int sum = 0;
        int now = 0;
        int matchingSum = 0;
        if (nums.length == 1) { return 0; }
        for (i = 0; i < nums.length; i++) {
        	sum += nums[i] == 1 ? 1 : -1;
        	if (hash.containsKey(sum)) {
        		maxCount = Math.max(maxCount, i - hash.get(sum));
        	} else {
        		hash.put(sum, i);
        	}
        }
        return maxCount;
    }
}