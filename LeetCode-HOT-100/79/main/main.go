package main

import "fmt"

type pair struct{ x, y int }

var directions = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	used := make([][]bool, m)
	for i := range used {
		used[i] = make([]bool, n)
	}
	var dfs func(index, row, col int) bool
	dfs = func(index, row, col int) bool {
		if board[row][col] != word[index] {
			return false
		}
		if index == len(word)-1 {
			return true
		}
		used[row][col] = true
		defer func() { used[row][col] = false }()
		for _, dir := range directions {
			if newRow, newCol := row+dir.x, col+dir.y; 0 <= newRow && newRow < m && 0 <= newCol && newCol < n && !used[newRow][newCol] {
				if dfs(index+1, newRow, newCol) {
					return true
				}
			}
		}
		return false
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(0, i, j) {
				return true
			}
		}
	}
	return false
}

func main() {
	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	word := "SEE"
	flag := exist(board, word)
	fmt.Println(flag)
}
