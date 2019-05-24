//Given a non-empty array of integers, every element appears twice except for one. Find that single one. 
//
// Note: 
//
// Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory? 
//
// Example 1: 
//
// 
//Input: [2,2,1]
//Output: 1
// 
//
// Example 2: 
//
// 
//Input: [4,1,2,1,2]
//Output: 4
// 
//

import java.util.*;

class Solution {
    public int singleNumber(int[] nums) {
        HashMap<Integer, Integer> map = new HashMap();
        for (int i = 0; i < nums.length; i++) {

            if (null != map.get(nums[i])) {
                map.remove(nums[i]);
            } else {
                map.put(nums[i], 1);

            }
        }

		Integer[] keys = map.keySet().toArray(new Integer[0]);
		return keys[0].intValue();
    }
}