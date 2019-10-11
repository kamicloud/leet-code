//Given a string, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases. 
//
// Note: For the purpose of this problem, we define empty string as valid palindrome. 
//
// Example 1: 
//
// 
//Input: "A man, a plan, a canal: Panama"
//Output: true
// 
//
// Example 2: 
//
// 
//Input: "race a car"
//Output: false
// 
// Related Topics Two Pointers String

package main

//leetcode submit region begin(Prohibit modification and deletion)
func isPalindrome(s string) bool {
	sx := strings.ToUpper(s)

	s = ""
	for i := 0; i < len(sx); i ++ {
		if (sx[i] <= 'Z' && sx[i] >= 'A') || (sx[i] <= '9' && sx[i] >= '0') {
			s += string(sx[i])
		}
	}

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s) - i - 1] {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
