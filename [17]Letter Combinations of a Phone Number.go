//Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent. 
//
// A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters. 
//
// 
//
// Example: 
//
// 
//Input: "23"
//Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
// 
//
// Note: 
//
// Although the above answer is in lexicographical order, your answer could be in any order you want. 
//

package main


func letterCombinations(digits string) []string {
	dig := strings.Split(digits, "")
	mp := map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}

	var x []string

	for _, v := range dig {
		if x == nil {
			x = mp[v]
		} else {
			x = crossJoin(mp[v], x)
		}
	}

	return x
}
func crossJoin(left []string, right []string) []string {
	var res []string

	for _, v1 := range left {
		for _, v2 := range right {
			res = append(res, strings.Join([]string{v2,v1, }, ""))
		}
	}
	return res
}