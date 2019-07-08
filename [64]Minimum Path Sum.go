//Given a m x n grid filled with non-negative numbers, find a path from top left to bottom right which minimizes the sum of all numbers along its path. 
//
// Note: You can only move either down or right at any point in time. 
//
// Example: 
//
// 
//Input:
//[
//  [1,3,1],
//  [1,5,1],
//  [4,2,1]
//]
//Output: 7
//Explanation: Because the path 1→3→1→1→1 minimizes the sum.
// 
//

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			} else if i != 0 && j == 0 {
				grid[i][j] += grid[i - 1][j]
			} else if j != 0 && i == 0 {
				grid[i][j] += grid[i][j - 1]
			} else {
				if grid[i][j - 1] < grid[i - 1][j] {
					grid[i][j] += grid[i][j - 1]
				} else {
					grid[i][j] += grid[i - 1][j]
				}
			}
		}
	}


	return grid[m - 1][n - 1]
}