//Given a column title as appear in an Excel sheet, return its corresponding column number. 
//
// For example: 
//
// 
//    A -> 1
//    B -> 2
//    C -> 3
//    ...
//    Z -> 26
//    AA -> 27
//    AB -> 28 
//    ...
// 
//
// Example 1: 
//
// 
//Input: "A"
//Output: 1
// 
//
// Example 2: 
//
// 
//Input: "AB"
//Output: 28
// 
//
// Example 3: 
//
// 
//Input: "ZY"
//Output: 701
// Related Topics Math

package main

//leetcode submit region begin(Prohibit modification and deletion)
func titleToNumber(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		res = res * 26 + int(rune(s[i]) - rune('A')) + 1
	}

	return res
}
//leetcode submit region end(Prohibit modification and deletion)
