//Given a 32-bit signed integer, reverse digits of an integer. 
//
// Example 1: 
//
// 
//Input: 123
//Output: 321
// 
//
// Example 2: 
//
// 
//Input: -123
//Output: -321
// 
//
// Example 3: 
//
// 
//Input: 120
//Output: 21
// 
//
// Note: 
//Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231, 231 − 1]. For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows. 
//

func reverse(x int) int {
	return int(reverse32(int32(x)))
}
func reverse32(x int32) int32 {
	var y int32 = 0

	var oldy int32 = 0
	var egtZero = true

	if x < 0 {
		egtZero = false
		x = -x
	}

	for {
		oldy = y
		y = y*10 + x%10
		x = x / 10

		if oldy != y / 10 {
			return 0
		}

		if x == 0 {
			break
		}
	}

	if !egtZero {
		y = -y
	}

	return y
}