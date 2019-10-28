//Given a non-empty string check if it can be constructed by taking a substring of it and appending multiple copies of the substring together. You may assume the given string consists of lowercase English letters only and its length will not exceed 10000. 
//
// 
//
// Example 1: 
//
// 
//Input: "abab"
//Output: True
//Explanation: It's the substring "ab" twice.
// 
//
// Example 2: 
//
// 
//Input: "aba"
//Output: False
// 
//
// Example 3: 
//
// 
//Input: "abcabcabcabc"
//Output: True
//Explanation: It's the substring "abc" four times. (And the substring "abcabc" twice.)
// 
// Related Topics String

package main

//leetcode submit region begin(Prohibit modification and deletion)
func repeatedSubstringPattern(s string) bool {
	for i := 1; i < len(s); i++ {
		if len(s) % i == 0 {
			size := len(s) / i
			cmp := s[0: i]
			valid := true
			for j := 1; j < size; j++ {
				if s[j * i: i * (j + 1)] != cmp {
					valid = false
					break
				}
			}
			if valid {
				return valid
			}
		}
	}
	return false
}
//leetcode submit region end(Prohibit modification and deletion)
