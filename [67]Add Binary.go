//Given two binary strings, return their sum (also a binary string). 
//
// The input strings are both non-empty and contains only characters 1 or 0. 
//
// Example 1: 
//
// 
//Input: a = "11", b = "1"
//Output: "100" 
//
// Example 2: 
//
// 
//Input: a = "1010", b = "1011"
//Output: "10101" 
//

func addBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}
	if len(b) == 0 {
		return a
	}
	var ar = []rune(a)
	var br = []rune(b)
	var apos, bpos = len(a) - 1, len(b) - 1
	var zero = rune('0')
	var one = rune('1')
	var carry = false
	var currentRune rune

	for {
		currentRune = ar[apos] + br[bpos] - zero
		if carry {
			currentRune += 1
		}

		carry = currentRune > one
		currentRune = (currentRune - zero) % 2 + zero
		ar[apos] = currentRune

		apos--
		bpos--

		if bpos < 0 {
			if carry {
				if apos < 0 {
					return "1" + string(ar)
				}
				return string(append([]rune(addBinary(string(ar[:apos + 1]), "1")), ar[apos + 1:]...))
			}
			return string(ar)
		}

	}
}