//Given a string, find the length of the longest substring without repeating characters. 
//
// 
// Example 1: 
//
// 
//Input: "abcabcbb"
//Output: 3 
//Explanation: The answer is "abc", with the length of 3. 
// 
//
// 
// Example 2: 
//
// 
//Input: "bbbbb"
//Output: 1
//Explanation: The answer is "b", with the length of 1.
// 
//
// 
// Example 3: 
//
// 
//Input: "pwwkew"
//Output: 3
//Explanation: The answer is "wke", with the length of 3. 
//             Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
// 
// 
// 
// 
//

func lengthOfLongestSubstring(s string) int {
	var m = make(map[rune]int)
	var max, start = 0, 0

	for i := 0; i < len(s); i++ {
		temp, isPresent := m[rune(s[i])]

		if isPresent && temp >= start {
			if max <= i-temp {
				max = i - temp
			}
			start = temp + 1
		} else {
			if max < i-start+1 {
				max = i - start + 1
			}
		}
		m[rune(s[i])] = i
	}

	return max
}
