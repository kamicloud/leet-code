//Given a string s consists of upper/lower-case alphabets and empty space characters ' ', return the length of last word in the string. 
//
// If the last word does not exist, return 0. 
//
// Note: A word is defined as a character sequence consists of non-space characters only. 
//
// Example: 
//
// 
//Input: "Hello World"
//Output: 5
// 
//
// 
//



func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	arr := strings.Split(s, " ")

	if len(arr) == 0 {
		return 0
	}
	return len(strings.Split(arr[len(arr) - 1], ""))
}