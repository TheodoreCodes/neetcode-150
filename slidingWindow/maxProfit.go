package slidingWindow

// maxProfit solves NeetCode "Buy and Sell Crypto"
// https://neetcode.io/problems/buy-and-sell-crypto/question
//
// prices[i] is the price of a coin on day i. Choose one day to buy and a later
// day to sell, and return the maximum profit. Return 0 if no profit is possible.
//
// Constraints:
//
//	1 <= prices.length <= 100
//	0 <= prices[i] <= 100
func maxProfit(prices []int) int {
	maxProfit := 0
	prevSell := 0
	for sellIdx := len(prices) - 1; sellIdx > 0; sellIdx-- {
		if prices[sellIdx] < prevSell {
			continue
		}

		for buyIdx := range sellIdx {
			if prices[sellIdx]-prices[buyIdx] > maxProfit {
				maxProfit = prices[sellIdx] - prices[buyIdx]
			}
		}

		prevSell = prices[sellIdx]
	}

	return maxProfit
}
