//Given a positive integer, return its corresponding column title as appear in an Excel sheet. 
//
// For example: 
//
// 
//    1 -> A
//    2 -> B
//    3 -> C
//    ...
//    26 -> Z
//    27 -> AA
//    28 -> AB 
//    ...
// 
//
// Example 1: 
//
// 
//Input: 1
//Output: "A"
// 
//
// Example 2: 
//
// 
//Input: 28
//Output: "AB"
// 
//
// Example 3: 
//
// 
//Input: 701
//Output: "ZY"
//


func convertToTitle(n int) string {
	var res string
	for {
		if n == 0 {
			break
		}

		temp := n % 26
		if temp == 0 {
			temp = 26
			n -= 1
		}

		temp += int(rune('A')) - 1

		res = string([]rune{rune(temp)}) + res

		n /= 26
	}

	return res
}