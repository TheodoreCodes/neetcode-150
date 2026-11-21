package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// sudokuBoard builds a board from 9 rows of 9 characters ('1'-'9' or '.').
func sudokuBoard(rows [9]string) [][]byte {
	board := make([][]byte, len(rows))
	for i, row := range rows {
		board[i] = []byte(row)
	}
	return board
}

// emptySudokuRows is a fully empty board; tests copy it and fill in cells.
var emptySudokuRows = [9]string{
	".........",
	".........",
	".........",
	".........",
	".........",
	".........",
	".........",
	".........",
	".........",
}

// withCells returns an empty board with the given digits placed at
// (row, col) coordinates, letting tests isolate a single rule violation.
func withCells(cells ...sudokuCell) [9]string {
	rows := emptySudokuRows
	for _, c := range cells {
		r := []byte(rows[c.row])
		r[c.col] = c.digit
		rows[c.row] = string(r)
	}
	return rows
}

type sudokuCell struct {
	row, col int
	digit    byte
}

func TestIsValidSudoku(t *testing.T) {
	tests := []struct {
		name  string
		board [9]string
		want  bool
	}{
		{
			name: "example 1: valid partially filled board",
			board: [9]string{
				"12..3....",
				"4..5.....",
				".98.....3",
				"5...6...4",
				"...8.3..5",
				"7...2...6",
				"......2..",
				"...419..8",
				"....8..79",
			},
			want: true,
		},
		{
			name: "example 2: duplicate 1 in top-left box",
			board: [9]string{
				"12..3....",
				"4..5.....",
				".91.....3",
				"5...6...4",
				"...8.3..5",
				"7...2...6",
				"......2..",
				"...419..8",
				"....8..79",
			},
			want: false,
		},
		{
			name: "fully solved valid board",
			board: [9]string{
				"534678912",
				"672195348",
				"198342567",
				"859761423",
				"426853791",
				"713924856",
				"961537284",
				"287419635",
				"345286179",
			},
			want: true,
		},
		{
			name:  "empty board",
			board: emptySudokuRows,
			want:  true,
		},
		{
			name:  "single filled cell",
			board: withCells(sudokuCell{4, 4, '5'}),
			want:  true,
		},
		{
			name: "same digit in different rows, columns and boxes",
			board: withCells(
				sudokuCell{0, 0, '1'},
				sudokuCell{3, 1, '1'},
				sudokuCell{6, 2, '1'},
				sudokuCell{1, 3, '1'},
				sudokuCell{4, 4, '1'},
				sudokuCell{7, 5, '1'},
				sudokuCell{2, 6, '1'},
				sudokuCell{5, 7, '1'},
				sudokuCell{8, 8, '1'},
			),
			want: true,
		},
		{
			name: "same digit in adjacent boxes with no shared row or column",
			board: withCells(
				sudokuCell{0, 2, '7'},
				sudokuCell{1, 3, '7'},
			),
			want: true,
		},
		{
			name:  "duplicate in a row, same box",
			board: withCells(sudokuCell{0, 0, '5'}, sudokuCell{0, 2, '5'}),
			want:  false,
		},
		{
			name:  "duplicate in a row, different boxes",
			board: withCells(sudokuCell{0, 0, '5'}, sudokuCell{0, 8, '5'}),
			want:  false,
		},
		{
			name:  "duplicate in a column, same box",
			board: withCells(sudokuCell{0, 0, '5'}, sudokuCell{2, 0, '5'}),
			want:  false,
		},
		{
			name:  "duplicate in a column, different boxes",
			board: withCells(sudokuCell{0, 0, '5'}, sudokuCell{8, 0, '5'}),
			want:  false,
		},
		{
			name:  "duplicate in a box only (diagonal neighbours)",
			board: withCells(sudokuCell{0, 0, '5'}, sudokuCell{1, 1, '5'}),
			want:  false,
		},
		{
			name:  "duplicate in a box only (opposite corners)",
			board: withCells(sudokuCell{0, 0, '5'}, sudokuCell{2, 2, '5'}),
			want:  false,
		},
		{
			name:  "duplicate in bottom-right box only",
			board: withCells(sudokuCell{6, 6, '9'}, sudokuCell{8, 8, '9'}),
			want:  false,
		},
		{
			name:  "duplicate in last row",
			board: withCells(sudokuCell{8, 0, '3'}, sudokuCell{8, 8, '3'}),
			want:  false,
		},
		{
			name:  "duplicate in last column",
			board: withCells(sudokuCell{0, 8, '3'}, sudokuCell{8, 8, '3'}),
			want:  false,
		},
		{
			name: "solved board with one altered cell breaks its row and column",
			board: [9]string{
				"534678912",
				"672195348",
				"198342567",
				"859761423",
				"426853791",
				"713924856",
				"961537284",
				"287419635",
				"345286171",
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			board := sudokuBoard(tt.board)

			got := isValidSudoku(board)

			assert.Equal(t, tt.want, got)
			// The solution must not modify the caller's board.
			assert.Equal(t, sudokuBoard(tt.board), board, "board must not be modified")
		})
	}
}
