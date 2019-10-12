//Write a function that takes a string as input and reverse only the vowels of a string.
//
// Example 1:
//
//
//Input: "hello"
//Output: "holle"
//
//
//
// Example 2:
//
//
//Input: "leetcode"
//Output: "leotcede"
//
//
// Note:
//The vowels does not include the letter "y".
//
//
// Related Topics Two Pointers String

package main


//leetcode submit region begin(Prohibit modification and deletion)
func reverseVowels(s string) string {
	if len(s) == 0 || len(s) == 1 {
		return s
	}
	sa := []byte(s)
	i := -1
	j := len(sa)
	var temp byte
	for {
		if i == j || i == len(s) - 1 {
			break
		}

		if temp == 0 {
			i++
			if isVowel(sa[i]) {
				temp = sa[i]
			}
		} else if temp > 0 {
			j--
			if isVowel(sa[j]) {
				sa[i], sa[j], temp = sa[j], temp, 0
			}

		}
	}

	return string(sa)
}

func isVowel(c byte) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' ||
		c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U'
}


//leetcode submit region end(Prohibit modification and deletion)
