package arrays

// isValidSudoku solves NeetCode "Valid Sudoku"
// https://neetcode.io/problems/valid-sudoku/question
//
// Given a 9x9 Sudoku board, determine whether it is valid. Only filled cells
// need to be validated: each row, each column, and each of the nine 3x3 boxes
// must contain the digits '1'-'9' at most once. Empty cells are '.'.
//
// Constraints:
//
//	board.length == 9
//	board[i].length == 9
//	board[i][j] is a digit '1'-'9' or '.'
func isValidSudoku(board [][]byte) bool {
	var rows [9][10]bool
	var cols [9][10]bool
	var boxes [9][10]bool

	for i := range board {
		for j := range board[i] {
			if board[i][j] == '.' {
				continue
			}
			digit := int(board[i][j] - '0')
			if rows[i][digit] ||
				cols[j][digit] ||
				boxes[(i/3)*3+(j/3)][digit] {
				return false
			}
			rows[i][digit] = true
			cols[j][digit] = true
			boxes[(i/3)*3+(j/3)][digit] = true
		}
	}

	return true
}
