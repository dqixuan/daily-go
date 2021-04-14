package main

import "fmt"

func main() {
	grid := make([][]int, 8)
	grid[0] = []int{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}
	grid[1] = []int{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}
	grid[2] = []int{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0}
	grid[3] = []int{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0}
	grid[4] = []int{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0}
	grid[5] = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0}
	grid[6] = []int{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}
	grid[7] = []int{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}

	fmt.Println(MaxAreaOfIsland(grid))
}

func dfs(i, j int, grid [][]int) int {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] != 1 {
		return 0
	}
	grid[i][j] = 0
	num := 1
	num += dfs(i-1, j, grid)
	num += dfs(i+1, j, grid)
	num += dfs(i, j-1, grid)
	num += dfs(i, j+1, grid)
	return  num
}

func MaxAreaOfIsland(grid [][]int) int {
	ans := 0
	for i:=0; i < len(grid); i++ {
		for j:=0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				ans = max(ans, dfs(i, j, grid))
			}
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
