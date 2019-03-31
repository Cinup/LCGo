package code

func maxProfit1(prices []int) int {
	max := 0
	min := 0xffffff
	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else {
			if prices[i]-min > max {
				max = prices[i] - min
			}
		}
	}
	return max
}
func maxProfit2(prices []int) int {
	max := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			max += prices[i] - prices[i-1]
		}
	}
	return max
}
