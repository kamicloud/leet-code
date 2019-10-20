//Given two strings s and t, determine if they are isomorphic. 
//
// Two strings are isomorphic if the characters in s can be replaced to get t. 
//
// All occurrences of a character must be replaced with another character while preserving the order of characters. No two characters may map to the same character but a character may map to itself. 
//
// Example 1: 
//
// 
//Input: s = "egg", t = "add"
//Output: true
// 
//
// Example 2: 
//
// 
//Input: s = "foo", t = "bar"
//Output: false 
//
// Example 3: 
//
// 
//Input: s = "paper", t = "title"
//Output: true 
//
// Note: 
//You may assume both s and t have the same length. 
// Related Topics Hash Table

package main

//leetcode submit region begin(Prohibit modification and deletion)
func isIsomorphic(s string, t string) bool {
	mp := make(map[byte]byte)
	tp := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		c, ifPresent := mp[s[i]]
		_, ifPresent2 := tp[t[i]]

		if ifPresent && c != t[i] {
			return false
		}
		if !ifPresent {
			if ifPresent2 {
				return false
			}
			mp[s[i]] = t[i]
			tp[t[i]] = s[i]
		}
	}
	return true
}
//leetcode submit region end(Prohibit modification and deletion)
