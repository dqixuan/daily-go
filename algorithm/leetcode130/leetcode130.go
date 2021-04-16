package main

import "sort"

/**
    @auther: 丁其轩
    @date:   2021/4/14
**/

func main() {
	board := make([][]byte, 4)
	board[0] = []byte{'X', 'X', 'X', 'X'}
	board[0] = []byte{'X', 'X', 'X', 'X'}
	board[0] = []byte{'X', 'X', 'X', 'X'}
	board[0] = []byte{'X', 'X', 'X', 'X'}
}


func solve(board [][]byte)  {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	dp := make([][]bool, len(board))
	for i :=0; i < len(board); i++ {
		dp[i] = make([]bool, len(board[0]))
	}
	for i:=0; i<len(board[0]);i++ {
		dfs(0, i, board, dp)
		dfs(len(board)-1, i, board, dp)
	}
	for i:=0; i < len(board); i++ {
		dfs(i, 0, board, dp)
		dfs(i, len(board[0])-1, board, dp)
	}
	for i:=0; i < len(board);i++ {
		for j:=0; j<len(board[0]);j++ {
			if board[i][j] == '0' && dp[i][j] {
				board[i][j] = 'X'
			}
		}
	}
	return
}

func dfs(i, j int, board [][]byte, dp [][]bool) {
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || board[i][j] == 'X' || dp[i][j] {
		return
	}
	dp[i][j] = true
	dfs(i+1, j, board, dp)
	dfs(i-1, j, board, dp)
	dfs(i, j+1, board, dp)
	dfs(i, j-1, board, dp)
	return
}

func merge(intervals [][]int) [][]int {
	length := len(intervals)
	if length <= 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] > intervals[j][1]
		}
		return intervals[i][0] > intervals[j][0]
	})

	for i:=0; i < len(intervals); {}
}