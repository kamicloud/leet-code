//Given a valid (IPv4) IP address, return a defanged version of that IP address. 
//
// A defanged IP address replaces every period "." with "[.]". 
//
// 
// Example 1: 
// Input: address = "1.1.1.1"
//Output: "1[.]1[.]1[.]1"
// Example 2: 
// Input: address = "255.100.50.0"
//Output: "255[.]100[.]50[.]0"
// 
// 
// Constraints: 
//
// 
// The given address is a valid IPv4 address. 
// Related Topics String

package main

//leetcode submit region begin(Prohibit modification and deletion)
func defangIPaddr(address string) string {
    str := ""
	for i := 0; i < len(address); i++ {
		if address[i] == '.' {
			str += "[.]"
		} else {
			str += string(address[i])
		}
	}
	return str
}
//leetcode submit region end(Prohibit modification and deletion)
