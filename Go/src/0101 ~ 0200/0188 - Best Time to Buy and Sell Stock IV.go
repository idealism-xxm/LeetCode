// 链接：https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iv/
// 题意：给定一只股票连续几天的价格，求最多只买卖 k 次的情况下最大获利？

// 输入： [2,4,1], k = 2
// 输出： 2
// 解释： 第一天买 (price = 2) ，第二天卖 (price = 4) ，获利 2
//		 总共进行 1 次买卖，共获利 2

// 输入： [3,2,6,5,0,3], k = 2
// 输出： 7
// 解释： 第二天买 (price = 2) ，第三天卖 (price = 6) ，获利 4
//		 第五天买 (price = 0) ，第六天卖 (price = 3) ，获利 3
//		 总共进行 2 次买卖，共获利 7

// 思路： DP
//
//		0123 的加强版，可以直接复用 0123 的 DP 思路
//		设 dp[i][j] 表示前 i 天经过了 j 次买/卖后的最大获利
//		其中： j = 0 表示未进行交易
//			  j = 1 表示买入了第 1 支股票
//			  j = 2 表示卖出了第 1 支股票
//			  ...
//			  j 为奇数时：表示买入了第 (j + 1) >> 1 支股票
//			  j 为偶数时：表示卖出了第 j >> 1 支股票
//		j 最大为 k << 1
//
//		初始化： dp[0][0] = 0
//		状态转移：
//			j == 0: dp[i][0] = dp[i - 1][0]
//			j >  0: dp[i][j] = max(dp[i][j], dp[i][j - 1] + sign(j) * prices[i])
//		其中：
//			j & 1 == 1 时， sign(j) 返回 -1
//			j & 1 == 0 时， sign(j) 返回 1
//
//		由于状态转移时只会用到它上边和左上的数据，所以可以优化为 O(k) ，直接用 dp[(k << 1) | 1]，从后往前更新即可
//		（本题不改变顺序也行，因为此时从前往后相当于今天可能进行了两次操作，
//		但这不影响，因为这样无任何盈利，还会减少可交易的次数）
//
//		时间复杂度： O(n * k)
//		空间复杂度： O(min(k, n))

func maxProfit(k int, prices []int) int {
	// 由于 k 可能非常大，但买卖次数最多为 len(prices) >> 1
	if k > (len(prices) >> 1) {
		k = len(prices) >> 1
	}
	// 最大下标为 k << 1 ，所以要分配 (k << 1) | 1 的空间
	dp := make([]int, (k << 1) | 1)
	// 设为一个不可能达到的最小值，方便简化后续逻辑（如果不支持，则只能对每一位标记是否合法）
	for i := k << 1; i > 0; i-- {
		dp[i] = math.MinInt32
	}
	// 状态转移
	for _, price := range prices {
		for j, sign := k << 1, 1; j > 0; j, sign = j - 1, sign * -1 {
			// 交易 j 次，
			// 要么今天不交易，直接用前一天交易 j 次的结果
			// 要么今天交易，用前一天交易 j - 1 次的结果 + 今天交易后获得的钱
			dp[j] = max(dp[j], dp[j - 1] + sign * price)
		}
	}
	// 返回所有情况中的最大值即可
	return max(dp[0], dp[1:]...)
}

func max(firstNum int, remainNums ...int) int {
	for _, num := range remainNums {
		if firstNum < num {
			firstNum = num
		}
	}
	return firstNum
}
