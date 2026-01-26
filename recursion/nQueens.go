package main

import "strings"

func solveNQueens(n int) [][]string {
	grid := make([][]int, n)
	for i := range grid {
		grid[i] = make([]int, n)
	}
	ans := [][]string{}
	generateGrid(&grid, 0, n, &ans)
	return ans
}

func generateGrid(grid *[][]int, row, n int, ans *[][]string) {
	if row == n {
		generateStringAnswer(grid, n, ans)
		return
	}
	for col := 0; col < n; col++ {
		if isSafe(row, col, n, grid) {
			(*grid)[row][col] = 1
			generateGrid(grid, row+1, n, ans)
			(*grid)[row][col] = 0
		}
	}
}

func isSafe(row, col, n int, grid *[][]int) bool {
	for i := 0; i < row; i++ {
		if (*grid)[i][col] == 1 {
			return false
		}
	}
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if (*grid)[i][j] == 1 {
			return false
		}
	}
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if (*grid)[i][j] == 1 {
			return false
		}
	}
	return true
}

func generateStringAnswer(grid *[][]int, n int, ans *[][]string) {
	board := []string{}
	for i := 0; i < n; i++ {
		var sb strings.Builder
		for j := 0; j < n; j++ {
			if (*grid)[i][j] == 1 {
				sb.WriteByte('Q')
			} else {
				sb.WriteByte('.')
			}
		}
		board = append(board, sb.String())
	}
	*ans = append(*ans, board)
}