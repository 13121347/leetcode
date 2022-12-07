package Algo_Dynamic

import "math"

//执行超时了
func maxProfit(prices []int) int {
	//1 最大利润
	maxProfit := 0

	//2 遍历数组
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			diffProfix := prices[j] - prices[i]
			if maxProfit < diffProfix {
				maxProfit = diffProfix
			}
		}
	}

	//3 返回结果
	return maxProfit
}

func maxProfit2(prices []int) int {
	minprice, maxprofit := math.MaxInt32, 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minprice {
			minprice = prices[i]
		} else if prices[i]-minprice > maxprofit {
			maxprofit = prices[i] - minprice
		}
	}
	return maxprofit
}
