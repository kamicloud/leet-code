//Write a function to find the longest common prefix string amongst an array of strings. 
//
// If there is no common prefix, return an empty string "". 
//
// Example 1: 
//
// 
//Input: ["flower","flow","flight"]
//Output: "fl"
// 
//
// Example 2: 
//
// 
//Input: ["dog","racecar","car"]
//Output: ""
//Explanation: There is no common prefix among the input strings.
// 
//
// Note: 
//
// All given inputs are in lowercase letters a-z. 
//

func longestCommonPrefix(strs []string) string {
	strslen := len(strs)

	if strslen == 0 {
		return ""
	}

	firstlen := len(strs[0])

	if strslen == 1 {
		return strs[0]
	}

	var temp = ""

	for i := 0; i < firstlen; i++ {
		var brk = false
		for j := 1; j < strslen; j++ {
			if (i >= len(strs[j]) && len(strs[j]) < firstlen) || strs[j][i] != strs[0][i] {
				brk = true
				break
			}
		}

		if brk {
			break
		}
		temp += string(strs[0][i])
	}

	return temp
}
