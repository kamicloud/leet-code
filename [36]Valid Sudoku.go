//Determine if a 9x9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules: 
//
// 
// Each row must contain the digits 1-9 without repetition. 
// Each column must contain the digits 1-9 without repetition. 
// Each of the 9 3x3 sub-boxes of the grid must contain the digits 1-9 without repetition. 
// 
//
// 
//A partially filled sudoku which is valid. 
//
// The Sudoku board could be partially filled, where empty cells are filled with the character '.'. 
//
// Example 1: 
//
// 
//Input:
//[
//  ["5","3",".",".","7",".",".",".","."],
//  ["6",".",".","1","9","5",".",".","."],
//  [".","9","8",".",".",".",".","6","."],
//  ["8",".",".",".","6",".",".",".","3"],
//  ["4",".",".","8",".","3",".",".","1"],
//  ["7",".",".",".","2",".",".",".","6"],
//  [".","6",".",".",".",".","2","8","."],
//  [".",".",".","4","1","9",".",".","5"],
//  [".",".",".",".","8",".",".","7","9"]
//]
//Output: true
// 
//
// Example 2: 
//
// 
//Input:
//[
//  ["8","3",".",".","7",".",".",".","."],
//  ["6",".",".","1","9","5",".",".","."],
//  [".","9","8",".",".",".",".","6","."],
//  ["8",".",".",".","6",".",".",".","3"],
//  ["4",".",".","8",".","3",".",".","1"],
//  ["7",".",".",".","2",".",".",".","6"],
//  [".","6",".",".",".",".","2","8","."],
//  [".",".",".","4","1","9",".",".","5"],
//  [".",".",".",".","8",".",".","7","9"]
//]
//Output: false
//Explanation: Same as Example 1, except with the 5 in the top left corner being 
//    modified to 8. Since there are two 8's in the top left 3x3 sub-box, it is invalid.
// 
//
// Note: 
//
// 
// A Sudoku board (partially filled) could be valid but is not necessarily solvable. 
// Only the filled cells need to be validated according to the mentioned rules. 
// The given board contain only digits 1-9 and the character '.'. 
// The given board size is always 9x9. 
// 
//
package main

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < len(board); i++ {
		var m = make(map[byte]byte)
		for j := 0; j < len(board[i]); j++ {
			_, err := m[board[i][j]]
			if err != false {
				return false
			}

			if board[i][j] != '.' {
				m[board[i][j]] = board[i][j]
			}
		}
	}

	for j := 0; j < len(board[0]); j++ {
		var m = make(map[byte]byte)
		for i := 0; i < len(board); i++ {
			_, err := m[board[i][j]]
			if err != false {
				return false
			}

			if board[i][j] != '.' {
				m[board[i][j]] = board[i][j]
			}
		}
	}

	for i := 0; i < len(board); i+= 3 {
		for j := 0; j < len(board[i]); j+= 3 {
			k, l := i, j

			var m = make(map[byte]byte)
			for {
				_, err := m[board[k][l]]
				if err != false {
					return false
				}


				if board[k][l] != '.' {
					m[board[k][l]] = board[k][l]
				}

				if k%3 == 2 {
					if l%3 < 2 {
						l++
						k = i
					} else {
						break
					}
				} else {
					k++
				}
			}

		}
	}

	return true
}