// Link: https://leetcode.com/problems/best-time-to-buy-and-sell-stock
package besttimetobuyandsellstock

// Complexity: O(n)
func maxProfit(prices []int) int {
	max := 0
	for left, right := 0, 0; right < len(prices); right++ {
		if prices[left] > prices[right] {
			left = right
		} else {
			diff := prices[right] - prices[left]
			if diff > max {
				max = diff
			}
		}
	}
	return max
}
