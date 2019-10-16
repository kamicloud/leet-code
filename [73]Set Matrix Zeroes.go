//Given a m x n matrix, if an element is 0, set its entire row and column to 0. Do it in-place. 
//
// Example 1: 
//
// 
//Input: 
//[
//  [1,1,1],
//  [1,0,1],
//  [1,1,1]
//]
//Output: 
//[
//  [1,0,1],
//  [0,0,0],
//  [1,0,1]
//]
// 
//
// Example 2: 
//
// 
//Input: 
//[
//  [0,1,2,0],
//  [3,4,5,2],
//  [1,3,1,5]
//]
//Output: 
//[
//  [0,0,0,0],
//  [0,4,5,0],
//  [0,3,1,0]
//]
// 
//
// Follow up: 
//
// 
// A straight forward solution using O(mn) space is probably a bad idea. 
// A simple improvement uses O(m + n) space, but still not the best solution. 
// Could you devise a constant space solution? 
// 
// Related Topics Array

package main

//leetcode submit region begin(Prohibit modification and deletion)
func setZeroes(matrix [][]int)  {
	cmap, rmap := map[int]int{}, map[int]int{}
    for i := 0; i < len(matrix); i++ {
    	for j := 0; j < len(matrix[i]); j++ {
    		if matrix[i][j] == 0 {
    			cmap[i] = 1
    			rmap[j] = 1
			}
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			_, ifPresent1 := cmap[i]
			_, ifPresent2 := rmap[j]

			if ifPresent1 || ifPresent2 {
				matrix[i][j] = 0
			}
		}
	}

}
//leetcode submit region end(Prohibit modification and deletion)
