//Given an integer, return its base 7 string representation. 
//
// Example 1: 
// 
//Input: 100
//Output: "202"
// 
// 
//
// Example 2: 
// 
//Input: -7
//Output: "-10"
// 
// 
//
// Note:
//The input will be in range of [-1e7, 1e7].
//

package main

//leetcode submit region begin(Prohibit modification and deletion)
func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	var res string
	neg := false
	if num < 0 {
		neg = true
		num = -num
	}

	for num != 0 {

		res = strconv.Itoa(num % 7) + res
		num /= 7
	}


	if neg {
		res = "-" + res
	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)
