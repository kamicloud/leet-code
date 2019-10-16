//Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand. 
//
// (i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]). 
//
// You are given a target value to search. If found in the array return its index, otherwise return -1. 
//
// You may assume no duplicate exists in the array. 
//
// Your algorithm's runtime complexity must be in the order of O(log n). 
//
// Example 1: 
//
// 
//Input: nums = [4,5,6,7,0,1,2], target = 0
//Output: 4
// 
//
// Example 2: 
//
// 
//Input: nums = [4,5,6,7,0,1,2], target = 3
//Output: -1 
// Related Topics Array Binary Search

package main

//leetcode submit region begin(Prohibit modification and deletion)
func search(nums []int, target int) int {
	n := len(nums)
	l := 0
	r := n - 1

	for l <= r {
		if nums[l] == target {
			return l
		}

		if nums[r] == target {
			return r
		}

		mid := (l + r) / 2

		if nums[mid] == target {
			return mid
		}

		if nums[l] > nums[r] {
			l += 1
			r -= 1
		} else {
			if nums[mid] > target {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}

	return -1
}
//leetcode submit region end(Prohibit modification and deletion)
//func search(nums []int, target int) int {
//	if len(nums) == 0 || (len(nums) == 1 && nums[0] != target) {
//		return -1
//	}
//	midIndex := len(nums) / 2
//	mid := nums[midIndex]
//	left := nums[0]
//	right := nums[len(nums) - 1]
//
//	if left == target {
//		return 0
//	} else if mid == target {
//		return midIndex
//	} else if right == target {
//		return len(nums) - 1
//	} else if (target < left && target < mid && target < right) || (target > left && target > mid && target > right) {
//		// 左侧或大右侧
//		leftres := search(nums[0:midIndex], target)
//		rightres := search(nums[midIndex:], target)
//		if leftres == -1 && rightres == -1 {
//			return -1
//		} else if leftres == -1 {
//			return rightres + midIndex
//		} else if rightres == -1 {
//			return leftres
//		} else {
//			return leftres
//		}
//	} else if target > mid {
//		// 右
//		res := search(nums[midIndex:], target)
//		if res == -1 {
//			return res
//		}
//		return res + midIndex
//
//	} else {
//		return search(nums[0:midIndex], target)
//	}
//}