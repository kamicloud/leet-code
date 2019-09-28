//Given an array of integers and an integer k, find out whether there are two distinct indices i and j in the array such that nums[i] = nums[j] and the absolute difference between i and j is at most k. 
//
// 
// Example 1: 
//
// 
//Input: nums = [1,2,3,1], k = 3
//Output: true
// 
//
// 
// Example 2: 
//
// 
//Input: nums = [1,0,1,1], k = 1
//Output: true
// 
//
// 
// Example 3: 
//
// 
//Input: nums = [1,2,3,1,2,3], k = 2
//Output: false
// 
// 
// 
// 
//

package main

func containsNearbyDuplicate(nums []int, k int) bool {
	mp := map[int][]int{}
	for i := 0; i < len(nums); i++ {
		poses, exists := mp[nums[i]]

		if exists {
			for j := 0; j < len(poses); j++ {
				if i - poses[j] <= k {
					return true
				}
			}
			mp[nums[i]] = append(mp[nums[i]], i)
		} else {
			mp[nums[i]] = []int{i}

		}
	}

	return false
}
