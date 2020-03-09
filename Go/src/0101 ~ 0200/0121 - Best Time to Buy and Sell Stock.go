// 链接：https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
// 题意：给定一只股票连续几天的价格，求最多只买卖一次的情况下最大获利？

// 输入： [7,1,5,3,6,4]
// 输出： 5
// 解释： 第二天买 (price = 1) ，第五天卖 (price = 6) ，获利 5

// 输入： [7,6,4,3,1]
// 输出： 0
// 解释： 不做任何买卖，获利 0

// 思路：贪心
//
//		假设我们第 i 天卖出的股票，那么必须在之前买到股票，且价格必须是这些天中最低的
//		所以我们贪心即可，记录前几天股票最低价买入，然后计算当天的利润（不能为负）
//
//		时间复杂度： O(n)
//		空间复杂度： O(1)

func maxProfit(prices []int) int {
	length := len(prices)
	// 如果不够两天，必定无利润
	if length < 2 {
		return 0
	}

	// 不进行买卖，利润为 0
	result := 0
	// 假设第一天买入股票
	buy := prices[0]
	for i := 1; i < length; i++ {
		// 假设卖出股票，计算最大获利
		result = max(result, prices[i] - buy)
		// 计算买入股票的最小值
		buy = min(buy, prices[i])
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
