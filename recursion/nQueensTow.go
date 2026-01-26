package main

func solveNQueenss(n int) [][]string {
	res := [][]string{}
	grid := make([][]int, n)
	for i := range grid {
		grid[i] = make([]int, n)
	}
	helperQueens(0, n, &grid, &res)
	return res
}

func helperQueens(row, n int, grid *[][]int, res *[][]string) {
	if row == n {
		addAnswer(grid, res)
		return
	}

	for i := 0; i < n; i++ {
		if isSafeAgain(row, i, n, grid) {
			(*grid)[row][i] = 1
			helperQueens(row+1, n, grid, res)
			(*grid)[row][i] = 0
		}
	}
}

func isSafeAgain(row, col, n int, grid *[][]int) bool {
	// Check column
	for i := 0; i < row; i++ {
		if (*grid)[i][col] == 1 {
			return false
		}
	}

	// Check left diagonal
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if (*grid)[i][j] == 1 {
			return false
		}
	}

	// Check right diagonal
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if (*grid)[i][j] == 1 {
			return false
		}
	}

	return true
}

func addAnswer(grid *[][]int, res *[][]string) {
	arr := []string{}
	for i := 0; i < len(*grid); i++ {
		str := ""
		for j := 0; j < len((*grid)[i]); j++ {
			if (*grid)[i][j] == 1 {
				str += "Q"
			} else {
				str += "."
			}
		}
		arr = append(arr, str)
	}
	*res = append(*res, arr)
}
