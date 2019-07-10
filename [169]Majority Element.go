//Given an array of size n, find the majority element. The majority element is the element that appears more than ⌊ n/2 ⌋ times. 
//
// You may assume that the array is non-empty and the majority element always exist in the array. 
//
// Example 1: 
//
// 
//Input: [3,2,3]
//Output: 3 
//
// Example 2: 
//
// 
//Input: [2,2,1,1,1,2,2]
//Output: 2
// 
//


func majorityElement(nums []int) int {
	var m = make(map[int]int)
	res, c := 0, 0

	for _, v := range nums {
		count, exists := m[v]
		if exists {
			m[v] = count + 1
		} else {
			m[v] = 1
		}
	}

	for k, v := range m {
		if v > c {
			res = k
			c = v
		}
	}

	return res
}