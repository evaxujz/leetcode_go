package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

// 买卖股票的最佳时机 II

// 给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
// 设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。
// 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
func maxProfit(prices []int) int {
	if prices == nil || len(prices) == 0 || len(prices) == 0 {
		return 0
	}

	buyPrice := prices[0]
	sellPrice := -1
	isBought := false
	profit := 0
	for _, price := range prices {
		if isBought {
			if price > sellPrice {
				sellPrice = price
			} else if price < sellPrice {
				profit += sellPrice - buyPrice
				buyPrice = price
				isBought = false
			}
		} else {
			if price > buyPrice {
				sellPrice = price
				isBought = true
			} else if price < buyPrice {
				buyPrice = price
			}
		}
	}

	if isBought {
		return profit + (sellPrice - buyPrice)
	} else {
		return profit
	}
}
