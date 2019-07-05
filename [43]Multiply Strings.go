//Given two non-negative integers num1 and num2 represented as strings, return the product of num1 and num2, also represented as a string. 
//
// Example 1: 
//
// 
//Input: num1 = "2", num2 = "3"
//Output: "6" 
//
// Example 2: 
//
// 
//Input: num1 = "123", num2 = "456"
//Output: "56088"
// 
//
// Note: 
//
// 
// The length of both num1 and num2 is < 110. 
// Both num1 and num2 contain only digits 0-9. 
// Both num1 and num2 do not contain any leading zero, except the number 0 itself. 
// You must not use any built-in BigInteger library or convert the inputs to integer directly. 
// 
//

func multiply(num1 string, num2 string) string {
	var res = "0"
	for i := len(num2) - 1; i >= 0; i-- {
		for j := rune('1'); j <= rune(num2[i]); j++ {
			res = addTwo(res, num1)
		}
		num1 += "0"
	}

	for i := 0; i < len(res); {
		if len(res) > 1 && rune(res[i]) == rune('0') {
			res = string([]rune(res[i + 1:]))
		} else {
			break
		}
	}

	return res
}

func addTwo(a string, b string) string {
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
	var one = rune('9')
	var carry = false
	var currentRune rune

	for {
		currentRune = ar[apos] + br[bpos] - zero
		if carry {
			currentRune += 1
		}

		carry = currentRune > one
		currentRune = (currentRune - zero) % 10 + zero
		ar[apos] = currentRune

		apos--
		bpos--

		if bpos < 0 {
			if carry {
				if apos < 0 {
					return "1" + string(ar)
				}
				return string(append([]rune(addTwo(string(ar[:apos + 1]), "1")), ar[apos + 1:]...))
			}
			return string(ar)
		}

	}
}