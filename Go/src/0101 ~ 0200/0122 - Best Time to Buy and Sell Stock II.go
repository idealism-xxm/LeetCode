// 链接：https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/
// 题意：给定一只股票连续几天的价格，求经过不限数量的买卖的情况后最大获利（手上有股票是不能继续买入）？

// 输入： [7,1,5,3,6,4]
// 输出： 7
// 解释： 第二天买 (price = 1) ，第三天卖 (price = 5) ，获利 4
//		 第四天买 (price = 3) ，第五天卖 (price = 6) ，获利 3
//		 总共进行 2 次买卖，共获利 7

// 输入： [1,2,3,4,5]
// 输出： 4
// 解释： 第一天买 (price = 1) ，第五天卖 (price = 5) ，获利 4
//		 总共进行 1 次买卖，共获利 4

// 输入： [7,6,4,3,1]
// 输出： 0
// 解释： 不做任何买卖，获利 0

// 思路1： DP
//		想法基本还是和 0121 一致的：
//		假设我们第 i 天卖出的股票，那么必须在之前买到股票，且价格必须是这些天中最低的
//		所以我们贪心即可，记录前几天股票最低价买入，然后计算当天的利润（不能为负）
//
//		不过可以进行多次交易后，就需要考虑以前得出的结果，然后在可交易的天中进行上述贪心
//		设 dp[i] 表示 prices[:i] 经过交易后的最大获利
//		初始化： dp[0] = 0 （未进行交易，获利 0）
//		状态转移：
//				如果从 dp[j] 转移而来，就需要加上 prices[j + 1:i] 期间最多一次交易的获利
//				dp[i] = max(dp[j] + maxProfitWithOneTransaction(prices[j + 1:i])) (0 <= j < i)
//
//		时间复杂度： O(n ^ 2)
//		空间复杂度： O(n)

func maxProfit(prices []int) int {
	length := len(prices)
	// 如果不够两天，必定无利润
	if length < 2 {
		return 0
	}

	dp := make([]int, length + 1)
	for i := 1; i <= length; i++ {
		// 假设第最后一天卖出股票
		sell := prices[i - 1]
		for j := i - 1; j >= 0; j-- {
			// 假设买入出股票，计算这些天一次交易的最大获利，并进行状态转移
			dp[i] = max(dp[i], dp[j] + (sell - prices[j]))
			// 计算卖出股票的最大值
			sell = max(sell, prices[j])
		}
	}
	return dp[length]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}


// 思路2： 贪心
//		看到题解还有 O(n) 的算法，才发现自己是真得想多了
//
//		思路还是和 0121 类似，
//		不过每次只要发现涨价了，就立刻卖出，赚取当前的利润，
//		然后假装又以当前价格买入，
//			1. 如果第二天涨价了，就假戏真多，转换成真实交易，继续重复上述操作
//			2. 如果第二天降价了，则假装今天以当前价买入
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
	// 假装第一天买入股票
	for i := 1; i < length; i++ {
		// 如果今天涨价了，则昨天假戏真做，今天卖出
		if prices[i] > prices[i - 1] {
			result += prices[i] - prices[i - 1]
		}
		// 再假装今天买入
	}
	return result
}
