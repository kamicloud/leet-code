//Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid. 
//
// An input string is valid if: 
//
// 
// Open brackets must be closed by the same type of brackets. 
// Open brackets must be closed in the correct order. 
// 
//
// Note that an empty string is also considered valid. 
//
// Example 1: 
//
// 
//Input: "()"
//Output: true
// 
//
// Example 2: 
//
// 
//Input: "()[]{}"
//Output: true
// 
//
// Example 3: 
//
// 
//Input: "(]"
//Output: false
// 
//
// Example 4: 
//
// 
//Input: "([)]"
//Output: false
// 
//
// Example 5: 
//
// 
//Input: "{[]}"
//Output: true
// 
//

const LA = '{'
const LB = '('
const LC = '['
const RA = '}'
const RB = ')'
const RC = ']'

func isValid(s string) bool {
	var sarr = []byte(s)
	var stack []byte
	var cmp byte

	var sarrlen = len(sarr)
	if sarrlen%2 == 1 {
		return false
	}

	if sarrlen != 0 {
		for i := 0; i < sarrlen; i++ {
			if sarr[i] == LA || sarr[i] == LB || sarr[i] == LC {
				stack = append(stack, sarr[i])
			} else {
				if len(stack) == 0 {
					return false
				}
				cmp = stack[len(stack)-1]
				if (cmp == LA && sarr[i] == RA) ||
					(cmp == LB && sarr[i] == RB) ||
					(cmp == LC && sarr[i] == RC) {
					stack = stack[:len(stack)-1]
				} else {
					return false
				}
			}
		}
		return len(stack) == 0
	}

	return true
}
