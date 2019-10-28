//Your task is to calculate ab mod 1337 where a is a positive integer and b is an extremely large positive integer given in the form of an array. 
//
// Example 1: 
//
// 
// 
//Input: a = 2, b = [3]
//Output: 8
// 
//
// 
// Example 2: 
//
// 
//Input: a = 2, b = [1,0]
//Output: 1024
// 
// 
// Related Topics Math

package main

//leetcode submit region begin(Prohibit modification and deletion)
func superPow(a int, b []int) int {
	if a == 1 || (len(b) == 1 && b[0] == 0) {
		return 1
	}
	a = a % 1337
	ten := pow10(a)
	res := 0
	for i := 0; i < len(b); i++ {
		if b[i] == 0 {
			continue
		}

		ic := powN(pow10N(ten, len(b) - 1 - i), b[i])

		if i == len(b) - 1 {
			ic *= powN(a, b[i])
		}

		if res == 0 {
			res = ic
		} else {
			res = res * ic % 1337
		}
	}
	res = res % 1337
	return res
}

func powN(x int, n int) int {
	res := 1
	for i := 0; i < n; i++ {
		res = res * x % 1337
	}
	return res
}

func pow10(x int) int {
	return powN(x, 10)
}

func pow10N(ten int, c int) int {
	if c == 0 {
		return 1
	}
	for i := 1; i < c; i++ {
		ten = pow10(ten) % 1337
	}
	return ten
}


//leetcode submit region end(Prohibit modification and deletion)
