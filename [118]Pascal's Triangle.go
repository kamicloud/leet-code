//Given a non-negative integer numRows, generate the first numRows of Pascal's triangle. 
//
// 
//In Pascal's triangle, each number is the sum of the two numbers directly above it. 
//
// Example: 
//
// 
//Input: 5
//Output:
//[
//     [1],
//    [1,1],
//   [1,2,1],
//  [1,3,3,1],
// [1,4,6,4,1]
//]
// 
//


package main

func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}
	res := [][]int{{1}}
	if numRows == 1 {
		return res
	}
	res = append(res, []int{1, 1})

	if numRows == 2 {
		return res
	}



	for i := 2; i < numRows; i++ {
		rown := []int{1}
		for j := 1; j < len(res[i - 1]); j++ {
			rown = append(rown, res[i - 1][j] + res[i - 1][j - 1])
		}

		rown = append(rown, res[i - 1][len(res[i - 1]) - 1])

		res = append(res, rown)
	}

	return res
}