//Given an array, rotate the array to the right by k steps, where k is non-negative. 
//
// Example 1: 
//
// 
//Input: [1,2,3,4,5,6,7] and k = 3
//Output: [5,6,7,1,2,3,4]
//Explanation:
//rotate 1 steps to the right: [7,1,2,3,4,5,6]
//rotate 2 steps to the right: [6,7,1,2,3,4,5]
//rotate 3 steps to the right: [5,6,7,1,2,3,4]
// 
//
// Example 2: 
//
// 
//Input: [-1,-100,3,99] and k = 2
//Output: [3,99,-1,-100]
//Explanation: 
//rotate 1 steps to the right: [99,-1,-100,3]
//rotate 2 steps to the right: [3,99,-1,-100]
// 
//
// Note: 
//
// 
// Try to come up as many solutions as you can, there are at least 3 different ways to solve this problem. 
// Could you do it in-place with O(1) extra space? 
// Related Topics Array

package main

//leetcode submit region begin(Prohibit modification and deletion)
func rotate(nums []int, k int)  {
    k = k % len(nums)

	for i := 0; i < k; i++ {
		doRotate(nums)
	}
}

func doRotate(nums []int) {
	temp := nums[len(nums) - 1]

	for i := len(nums) - 1; i > 0; i-- {
		nums[i] = nums[i - 1]
	}
	nums[0] = temp
}
//leetcode submit region end(Prohibit modification and deletion)
