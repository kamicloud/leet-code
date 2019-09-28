//Given an integer (signed 32 bits), write a function to check whether it is a power of 4. 
//
// Example 1: 
//
// 
//Input: 16
//Output: true
// 
//
// 
// Example 2: 
//
// 
//Input: 5
//Output: false 
// 
//
// Follow up: Could you solve it without loops/recursion?
package main

func isPowerOfFour(num int) bool {
	if num == 1 {
		return true
	}

	if num < 4 {
		return false
	}

	if (num / 4) * 4 != num {
		return false
	}

	return isPowerOfFour(num / 4)
}