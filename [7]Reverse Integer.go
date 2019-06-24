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
	var y = 0

	var egtZero = true

	if x < 0 {
		egtZero = false
		x = -x
	}

	for {
		y = y*10 + x%10
		x = x / 10

		if x == 0 {
			break
		}
	}

	if !egtZero {
		y = -y
	}

	if y > math.MaxInt32 || y < math.MinInt32 {
		return 0
	}

	return y
}

func reverseString(s []byte) []byte {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
