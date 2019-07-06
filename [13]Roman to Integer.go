//Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M. 
//
// 
//Symbol       Value
//I             1
//V             5
//X             10
//L             50
//C             100
//D             500
//M             1000 
//
// For example, two is written as II in Roman numeral, just two one's added together. Twelve is written as, XII, which is simply X + II. The number twenty seven is written as XXVII, which is XX + V + II. 
//
// Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used: 
//
// 
// I can be placed before V (5) and X (10) to make 4 and 9. 
// X can be placed before L (50) and C (100) to make 40 and 90. 
// C can be placed before D (500) and M (1000) to make 400 and 900. 
// 
//
// Given a roman numeral, convert it to an integer. Input is guaranteed to be within the range from 1 to 3999. 
//
// Example 1: 
//
// 
//Input: "III"
//Output: 3 
//
// Example 2: 
//
// 
//Input: "IV"
//Output: 4 
//
// Example 3: 
//
// 
//Input: "IX"
//Output: 9 
//
// Example 4: 
//
// 
//Input: "LVIII"
//Output: 58
//Explanation: L = 50, V= 5, III = 3.
// 
//
// Example 5: 
//
// 
//Input: "MCMXCIV"
//Output: 1994
//Explanation: M = 1000, CM = 900, XC = 90 and IV = 4. 
//


func romanToInt(s string) int {
	i := 0
	var res = 0
	str := []rune(s)
	for {
		if i >= len(str) {
			break
		}
		if str[i] == rune('I') {
			if i + 1 >= len(str) {
				res += 1
			} else if str[i + 1] == rune('V') {
				res += 4
				i += 2
				continue
			} else if str[i + 1] == rune('X') {
				res += 9
				i += 2
				continue
			} else {
				res += 1
			}
		}
		if str[i] == rune('V') {
			res += 5
		}
		if str[i] == rune('X') {
			if i + 1 >= len(str) {
				res += 10
			} else if str[i + 1] == rune('L') {
				res += 40
				i += 2
				continue
			} else if str[i + 1] == rune('C') {
				res += 90
				i += 2
				continue
			} else {
				res += 10
			}
		}
		if str[i] == rune('L') {
			res += 50
		}
		if str[i] == rune('C') {
			if i + 1 >= len(str) {
				res += 100
			} else if str[i + 1] == rune('D') {
				res += 400
				i += 2
				continue
			} else if str[i + 1] == rune('M') {
				res += 900
				i += 2
				continue
			} else {
				res += 100
			}
		}
		if str[i] == rune('D') {
			res += 500
		}
		if str[i] == rune('M') {
			res += 1000
		}


		i++
	}
	return res
}