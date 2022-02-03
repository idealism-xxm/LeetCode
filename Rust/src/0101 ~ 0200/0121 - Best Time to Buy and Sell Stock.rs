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

impl Solution {
    pub fn max_profit(prices: Vec<i32>) -> i32 {
        prices
            // 转成迭代器
            .iter()
            // 跳过第一天，同一天买卖不亏不赚
            // （这里是否跳过不影响结果，
            //  是为了语义上更符合我们后续使用的 pre 变量）
            .skip(1)
            // 使用 fold 积累结果
            .fold(
                // 维护两个值 ans 和 pre ，
                // ans 表示目前进行一次买卖时的最大获利
                // pre 表示当前天之前的股票最低价
                (0, prices[0]), 
                // 假设当天 cur 买入，则前几天选择最低价 pre 买入，则获利 cur - pre ，
                // 更新最大获利 ans = ans.max(cur - pre) 。
                // 同时更新后一天之前的股票最低价 pre = min(pre, cu)
                |(ans, pre), &cur| (ans.max(cur - pre), pre.min(cur)),
            )
            // 最后返回最大获利即可
            .0
    }
}
