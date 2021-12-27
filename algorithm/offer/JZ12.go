package offer

/**
	二维数组中是否包含某个路径组成的字符串
 */

func hasPath( matrix [][]byte ,  word string ) bool {
	row := len(matrix)
	col := len(matrix[0])
	visited := make([][]int, row)
	for i:=0; i < row; i++ {
		visited[i] = make([]int, col)
	}
	for i:=0; i < row; i ++ {
		for j:=0; j < col; j++ {
			if si(matrix, visited, i, j, word, 0) {
				return true
			}
		}
	}
	return false
}

func si(grid [][]byte, visited [][]int, i, j int, s string, idx int) bool {
	row := len(grid)
	col := len(grid[0])
	if visited[i][j] == 1 {
		return false
	}
	if grid[i][j] != s[idx] {
		return false
	}
	if idx == len(s)-1 {
		return true
	}
	exist := false
	visited[i][j] = 1
	if i > 0 {
		exist = si(grid, visited, i-1, j, s, idx+1)
		if exist == true {
			return true
		}
	}
	if i+1 < row {
		exist = si(grid, visited, i+1, j, s, idx+1)
		if exist == true {
			return true
		}
	}
	if j > 0 {
		exist = si(grid, visited, i, j-1, s, idx+1)
		if exist == true {
			return true
		}
	}
	if j+1 < col {
		exist =  si(grid, visited, i, j+1, s, idx+1)
		if exist == true {
			return true
		}
	}
	visited[i][j] = 0
	return false
}
