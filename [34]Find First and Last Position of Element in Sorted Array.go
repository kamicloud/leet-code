//Given an array of integers nums sorted in ascending order, find the starting and ending position of a given target value. 
//
// Your algorithm's runtime complexity must be in the order of O(log n). 
//
// If the target is not found in the array, return [-1, -1]. 
//
// Example 1: 
//
// 
//Input: nums = [5,7,7,8,8,10], target = 8
//Output: [3,4] 
//
// Example 2: 
//
// 
//Input: nums = [5,7,7,8,8,10], target = 6
//Output: [-1,-1] 
// Related Topics Array Binary Search

package main

//leetcode submit region begin(Prohibit modification and deletion)
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return []int{0, 0}
		} else {
			return []int{-1, -1}
		}

	}


	midIndex := len(nums) / 2

	if nums[midIndex] < target {
		res := searchRange(nums[midIndex:], target)

		if res[0] == -1 {
			return res
		} else {
			return []int{res[0] + midIndex , res[1] + midIndex}
		}

	} else if nums[midIndex] > target {
		return searchRange(nums[0: midIndex], target)
	} else {
		x := nums[0: midIndex]
		y := nums[midIndex:]
		left := searchRange(x, target)
		right := searchRange(y, target)

		if left[0] == -1 && right[0] == -1 {
			return []int{-1, -1}
		} else if left[0] == -1 {
			return []int{right[0] + midIndex, right[1] + midIndex}
		} else if right[0] == -1 {
			return left
		} else {

			return []int{left[0], right[1] + midIndex}
		}
	}

}
//leetcode submit region end(Prohibit modification and deletion)
