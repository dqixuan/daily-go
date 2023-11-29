package stack

/*
*

有n条鱼每条鱼的位置及大小均不同，他们沿着

	x轴游动，有的向左，有的向右。游动的速度是一样的，两条鱼相遇大鱼会吃掉小鱼。从左到右给出每条鱼的大小和游动的方向（
	0表示向左，1表示向右）。问足够长的时间之后，能剩下多少条鱼？
*/
func eatFish(val, direct []int) int {
	stack := make([]int, 0, len(val))
	for i := range val {
		hasEat := false
		for len(stack) > 0 && direct[stack[len(stack)-1]] == 1 && direct[i] == 0 {
			if val[stack[len(stack)-1]] > val[i] {
				hasEat = true
				break
			}
			stack = stack[:len(stack)-1]
		}

		if !hasEat {
			stack = append(stack, i)
		}
	}
	return len(stack)
}
