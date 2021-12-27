package offer

/**

 */

func movingCount( threshold int ,  rows int ,  cols int ) int {
	// write code here
	return 0
}

func helpMoveCount(threshold, i, j, row, cols int, visit[][]int, count *int) {
	if visit[i][j] == 1 {
		return
	}
	if getSum(i, j) > threshold {
		return
	}
	visit[i][j] = 1
}

func getSum(a, b int) int {
	ans := 0
	for a != 0 {
		ans += a%10
		a /= 10
	}
	for b != 0 {
		ans += b%10
		b /= 10
	}
	return  ans
}