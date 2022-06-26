// 链接：https://leetcode.com/problems/maximum-points-you-can-obtain-from-cards/
// 题意：给定一个整数数组 card_points ，每次可以从左或者从右拿走 1 个数，
//       直至拿走 k 个数，求拿走数之和的最大值？


// 数据限制：
//   1 <= card_points.length <= 10 ^ 5
//   1 <= card_points[i] <= 10 ^ 4
//   1 <= k <= card_points.length


// 输入： card_points = [1,2,3,4,5,6,1], k = 3
// 输出： 12
// 解释： 拿走最右边的 3 个数，和为 1 + 6 + 5 = 12

// 输入： card_points = [2,2,2], k = 2
// 输出： 4
// 解释： 无论怎么拿，最终拿走的两个数的和都是 4

// 输入： card_points = [9,7,7,9,7,7,9], k = 7
// 输出： 55
// 解释： 拿走全部的 7 个数，和为 55


// 思路： 滑动窗口
//
//      如果一道题目需要在所有满足某种状态的连续子串/连续子数组中，
//      找到满足题意的一个，那么可以考虑滑动窗口。
//
//      本题类似 LeetCode 1658 ，可以采用类似的方法处理，
//      基本修改一下思路中的文字和代码就可以完美适配。
//
//      如果我们将 card_points 拼接一次，得到一个长度为 2 * len(card_points) 的新数组，
//      那么本题就转化为：在新数组中找到一个长度为 k 连续子数组的最大和。
//
//      可以使用滑动窗口的方法解决本题，这样的空间复杂度为 O(n) 。
//      不过可以特殊处理右边界 r 的情况，能让空间复杂度优化为 O(1) ，
//      但不便于理解且容易出错。
//
//      此时可以考虑原问题的镜像问题：求长度为 len(card_points) - k 的子数组的最小和 ans ，
//      那么 sum(card_points) - ans 就是原问题的答案。
//
//      设 target = sum(card_points) - x ，
//      我们使用滑动窗口 [l, r] 表示一个数字和小于等于 target 的连续子数组，
//      初始化为左边界 l = 0 ，右边界 r = target - 1 ，滑动窗口数字和 total = sum(card_points[:target]) 。
//
//      我们不断右移右边界 r ，将其纳入到滑动窗口中考虑， total += card_points[r] ，
//      同时右移左边界 l 一次，保持滑动窗口长度为 target ，total -= card_points[l] 。
//
//      获取所有滑动窗口数字和的最小值 ans ，最后返回 sum(card_points) - ans 即可。
//
//
//      时间复杂度：O(n)
//          1. 需要遍历 card_points 中全部 O(n) 个数字
//      空间复杂度：O(1)
//          1. 只需用维护常数个额外变量


impl Solution {
    pub fn max_score(card_points: Vec<i32>, k: i32) -> i32 {
        let total = card_points.iter().sum();
        let n = card_points.len();
        // 原问题可以转化为求长度为 target 的子数组的最小和
        let target = n - k as usize;
        // 如果 target 等于 0 ，则原问题的解为 total - 0 = total
        if target == 0 {
            return total
        }

        // 初始化滑动窗口为 [0, target - 1] ，那么滑动窗口内数字和为 sum(card_points[:target])
        let mut amount: i32 = card_points.iter().take(target).sum();
        // 初始连续子数组的最小和为 amount
        let mut ans = amount;
        // 不断右移右边界 r
        for r in target..n {
            // 将 card_points[r] 纳入到滑动窗口中考虑
            amount += card_points[r];
            // 同时 card_points[r - target] 要移出滑动窗口，使滑动窗口长度保持为 target
            amount -= card_points[r - target];

            // 更新长度为 k 的连续子数组的最小和
            ans = ans.min(amount);
        }

        total - ans
    }
}
