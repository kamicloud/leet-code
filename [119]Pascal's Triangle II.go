//Given a non-negative index k where k ≤ 33, return the kth index row of the Pascal's triangle. 
//
// Note that the row index starts from 0. 
//
// 
//In Pascal's triangle, each number is the sum of the two numbers directly above it. 
//
// Example: 
//
// 
//Input: 3
//Output: [1,3,3,1]
// 
//
// Follow up: 
//
// Could you optimize your algorithm to use only O(k) extra space? 
//

package main

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}

	if rowIndex == 1 {
		return []int{1,1}
	}
	if rowIndex == 2 {
		return []int{1,2,1}
	}
	res := []int{1, 1}


	for i := 3; i < rowIndex + 2; i++ {
		res = append(res, 1)

		for j := len(res) - 2; j > 0; j-- {
			res[j] = res[j] + res[j - 1]
		}
	}
	return res
}