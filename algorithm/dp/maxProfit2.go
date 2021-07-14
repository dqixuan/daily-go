package dp

/**
	给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
	设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。
	注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）
 */

func maxProfit2(prices []int) int {
	ans := 0

	for i, v := range prices {
		if i > 1 {
			profit := v-prices[i-1]
			if profit > 0 {
				ans += profit
			}
		}
	}

	return ans
}

func maxScoreSightSeeingPair(nums []int) int {
	max := nums[0]+0
	min := nums[0]-0
	for i, v := range nums {
		if max < i+v {
			max = i+v
		}
		if min < v-i {
			min = v-i
		}
	}
	return min + max
}
func maxScoreSightseeingPair(values []int) int {
	max := values[0]
	idx := 0
	ans := 0
	for i, v := range values {
		if max + v -i > ans && i > idx {
			ans = max + v - i
		}
		if i + v > max {
			max = i + v
			idx = i
		}
	}
	return ans
}
