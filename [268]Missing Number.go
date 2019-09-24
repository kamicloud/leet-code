//Given an array containing n distinct numbers taken from 0, 1, 2, ..., n, find the one that is missing from the array.
//
// Example 1:
//
//
//Input: [3,0,1]
//Output: 2
//
//
// Example 2:
//
//
//Input: [9,6,4,2,3,5,7,0,1]
//Output: 8
//
//
// Note:
//Your algorithm should run in linear runtime complexity. Could you implement it using only constant extra space complexity?


package main


func missingNumber(nums []int) int {
	if len(nums) == 1 {
		if nums[0] == 1 {
			return 0
		}
		return 1
	}

	res := 0
	max := int(math.MinInt64)
	min := int(math.MaxInt64)
	for i := 0; i < len(nums); i++ {
		res += nums[i]
		if max < nums[i] {
			max = nums[i]
		}
		if min > nums[i] {
			min = nums[i]
		}
	}

	ans := ((max + min) * (max - min + 1)) / 2

	if ans == res {
		if min == 0 {
			return max + 1
		}
		return 0
	}

	return ans - res
}