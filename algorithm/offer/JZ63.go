package offer


func maxProfit( prices []int ) int {
	// write code here
	if len(prices) == 0 {
		return 0
	}
	minIndex := 0
	ans := 0
	for i, v := range prices {
		ans = max(ans, v-prices[minIndex])
		if v < prices[minIndex] {
			minIndex = i
		}
	}
	return ans
}
