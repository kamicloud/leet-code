//Given two strings s and t , write a function to determine if t is an anagram of s. 
//
// Example 1: 
//
// 
//Input: s = "anagram", t = "nagaram"
//Output: true
// 
//
// Example 2: 
//
// 
//Input: s = "rat", t = "car"
//Output: false
// 
//
// Note: 
//You may assume the string contains only lowercase alphabets. 
//
// Follow up: 
//What if the inputs contain unicode characters? How would you adapt your solution to such case? 
//
package main

func isAnagram(s string, t string) bool {
    sa := map[byte]int{}
	ta := map[byte]int{}

	if len(s) != len(t) {
		return false
	}

	for i := 0; i < len(s); i++ {
		scount, _ := sa[s[i]]

		scount++
		sa[s[i]] = scount


		tcount, _ := ta[t[i]]

		tcount++
		ta[t[i]] = tcount
	}

	if len(sa) != len(ta) {
		return false
	}

	for k, v := range sa {
		count, exists := ta[k]

		if !exists || count != v {
			return false
		}
	}

	return true
}