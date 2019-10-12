//Given two arrays, write a function to compute their intersection. 
//
// Example 1: 
//
// 
//Input: nums1 = [1,2,2,1], nums2 = [2,2]
//Output: [2,2]
// 
//
// 
// Example 2: 
//
// 
//Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
//Output: [4,9] 
// 
//
// Note: 
//
// 
// Each element in the result should appear as many times as it shows in both arrays. 
// The result can be in any order. 
// 
//
// Follow up: 
//
// 
// What if the given array is already sorted? How would you optimize your algorithm? 
// What if nums1's size is small compared to nums2's size? Which algorithm is better? 
// What if elements of nums2 are stored on disk, and the memory is limited such that you cannot load all elements into the memory at once? 
// 
// Related Topics Hash Table Two Pointers Binary Search Sort

package main

//leetcode submit region begin(Prohibit modification and deletion)
func intersect(nums1 []int, nums2 []int) []int {
	mp := map[int]int{}
	var res []int

	for i := 0; i < len(nums1); i++ {
		count := mp[nums1[i]]

		mp[nums1[i]] = count + 1
	}

	for i := 0; i < len(nums2); i++ {
		count, exists := mp[nums2[i]]

		if exists && count > 0 {
			res = append(res, nums2[i])
			mp[nums2[i]] = count - 1
		}
	}

	return res
}
//leetcode submit region end(Prohibit modification and deletion)
