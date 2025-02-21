package leetcode

func IsValidSudoku(board [][]byte) bool {
	var rows [9]map[byte]bool
	var columns [9]map[byte]bool
	var boxes [9]map[byte]bool

	for i := range rows {
		rows[i] = make(map[byte]bool)
		columns[i] = make(map[byte]bool)
		boxes[i] = make(map[byte]bool)
	}

	for row := range board {
		for col := range board[row] {
			if board[row][col] != '.' {
				boxIdx := (row/3)*3 + (col / 3)

				if rows[row][board[row][col]] || columns[col][board[row][col]] || boxes[boxIdx][board[row][col]] {
					return false
				}

				rows[row][board[row][col]] = true
				columns[col][board[row][col]] = true
				boxes[boxIdx][board[row][col]] = true
			}
		}
	}

	return true
}

// TODO: write the bitwise shit
func IsValidSudokuBitMasking(board [][]byte) bool {
	rows := make([]int, 9)
	cols := make([]int, 9)
	squares := make([]int, 9)

	for r, row := range board {
		for c, cell := range row {
			if cell == '.' {
				continue
			}

			val := cell - '1'
			bit := 1 << val
			squareIdx := (r/3)*3 + c/3

			if rows[r]&bit != 0 || cols[c]&bit != 0 ||
				squares[squareIdx]&bit != 0 {
				return false
			}

			rows[r] |= bit
			cols[c] |= bit
			squares[squareIdx] |= bit
		}
	}

	return true
}
