//Given two arrays, write a function to compute their intersection. 
//
// Example 1: 
//
// 
//Input: nums1 = [1,2,2,1], nums2 = [2,2]
//Output: [2]
// 
//
// 
// Example 2: 
//
// 
//Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
//Output: [9,4] 
// 
//
// Note: 
//
// 
// Each element in the result must be unique. 
// The result can be in any order. 
// 
//
// 
// Related Topics Hash Table Two Pointers Binary Search Sort

package main

//leetcode submit region begin(Prohibit modification and deletion)
func intersection(nums1 []int, nums2 []int) []int {
    mp := map[int]int{}
    resm := map[int]int{}
	var res []int

    for i := 0; i < len(nums1); i++ {
    	mp[nums1[i]] = 1
	}

    for i := 0; i < len(nums2); i++ {
    	_, exists := mp[nums2[i]]

    	if exists {
			resm[nums2[i]] = 1
		}
	}

    for v := range resm {
    	res = append(res, v)
	}

    return res
}
//leetcode submit region end(Prohibit modification and deletion)
