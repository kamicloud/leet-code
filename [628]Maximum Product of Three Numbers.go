//Given an integer array, find three numbers whose product is maximum and output the maximum product. 
//
// Example 1: 
//
// 
//Input: [1,2,3]
//Output: 6
// 
//
// 
//
// Example 2: 
//
// 
//Input: [1,2,3,4]
//Output: 24
// 
//
// 
//
// Note: 
//
// 
// The length of the given array will be in range [3,104] and all elements are in the range [-1000, 1000]. 
// Multiplication of any three numbers in the input won't exceed the range of 32-bit signed integer. 
// 
//
// 
//




func maximumProduct(nums []int) int {
	min1, min2 := math.MaxInt64,math.MaxInt64
	max1, max2, max3 := math.MinInt64,math.MinInt64,math.MinInt64

	for i := 0; i < len(nums); i++ {
		val := nums[i]

		if val < min2 {
			if val < min1 {
				min2 = min1
				min1 = val
			} else {
				min2 = val
			}
		}
		if val > max3 {
			if val > max2 {
				max3 = max2
				if val > max1 {
					max2 = max1
					max1 = val
				} else {
					max2 = val
				}
			} else {
				max3 = val
			}
		}
	}

	res1 := min1 * min2 * max1
	res2 := max1 * max2 * max3

	if res1 < res2 {
		return res2
	}

	return res1
}