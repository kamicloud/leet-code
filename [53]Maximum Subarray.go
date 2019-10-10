//Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum. 
//
// Example: 
//
// 
//Input: [-2,1,-3,4,-1,2,1,-5,4],
//Output: 6
//Explanation: [4,-1,2,1] has the largest sum = 6.
// 
//
// Follow up: 
//
// If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle. 
//
package main


func maxSubArray(nums []int) int {
	max := math.MinInt64
	curSum := 0

	for i := 0; i < len(nums); i++ {
		curSum = int(math.Max(float64(curSum + nums[i]), float64(nums[i])))
		max = int(math.Max(float64(max), float64(curSum)))
	}

	return max
}
