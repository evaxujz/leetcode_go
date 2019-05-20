package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

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
