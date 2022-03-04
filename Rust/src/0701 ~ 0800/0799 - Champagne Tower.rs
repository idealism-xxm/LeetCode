// 链接：https://leetcode.com/problems/champagne-tower/
// 题意：有一个 100 层的香槟塔，第一层有 1 个玻璃杯，
//      第二层有 2 个玻璃杯， ... ，第 100 层有 100 个玻璃杯。
//
//      每一个玻璃杯都能装下一杯香槟，如果一个玻璃杯装满了，
//      多余的香槟会立刻平均流入到它下面左右两边的玻璃杯中。
//
//      现在给第一层的玻璃杯倒入 poured  杯香槟，
//      求第 query_row 层第 query_glass 杯的香槟量。
//      （ query_row 和 query_glass 都从 0 开始）


// 数据限制：
//  0 <= poured <= 10 ^ 9
//  0 <= query_glass <= query_row < 100


// 输入： poured = 1, query_row = 1, query_glass = 1
// 输出： 0.00000
// 解释： 在香槟塔顶层（下标是 (0, 0) ）倒了一杯香槟后，
//       没有溢出，因此所有在顶层以下的玻璃杯都是空的。

// 输入： poured = 2, query_row = 1, query_glass = 1
// 输出： 0.50000
// 解释： 在香槟塔顶层（下标是 (0, 0) ）倒了两杯香槟后，
//       有一杯量的香槟将从顶层溢出，
//       位于 (1, 0) 的玻璃杯和 (1, 1) 的玻璃杯平分了这一杯香槟，
//       所以每个玻璃杯有一半的香槟。

// 输入： poured = 100000009, query_row = 33, query_glass = 17
// 输出： 1.00000


// 思路： DP
//
//      定义一个二维数组 dp ，
//      其中 dp[i][j] 表示第 i 层第 j 杯会接收到的香槟量。
//
//      初始只有第 0 层第 0 杯会接收到 poured 杯香槟，
//      其他玻璃杯需要后续在循环中计算会接收到的香槟量。
//
//      初始化： dp[i][j] = 0, dp[0][0] = poured
//      状态转移：如果 dp[i][j] > 1 ，
//          则说明第 i 层第 j 杯会有多余香槟流出，
//          流出量为 dp[i][j] - 1 。
//
//          并且会平均流入到它下面左右两边的玻璃杯中，
//          即第 i + 1 层的第 j 杯和第 j + 1 杯
//          均会流入 (dp[i][j] - 1) / 2 杯香槟。
//
//
//		时间复杂度： O(query_row ^ 2)
//          1. 需要遍历处理二维数组 dp 的全部 O(query_row ^ 2) 个元素
//		空间复杂度： O(query_row ^ 2)
//          1. 需要维护一个大小为 O(query_row ^ 2) 二维数组 dp
//          2. （可以使用滚动数组将空间复杂度优化为 O(query_row) ）


impl Solution {
    pub fn champagne_tower(poured: i32, query_row: i32, query_glass: i32) -> f64 {
        // 先转成 usize 类型，方便后续处理
        let (query_row, query_glass) = (query_row as usize, query_glass as usize);
        // 定义一个二维数组 dp ，
        // 初始化均为 0 ，表示没有香槟
        let mut dp = vec![vec![0.0; query_row + 2]; query_row + 2];
        // 初始只有第 0 层第 0 杯会接收到 poured 杯香槟
        dp[0][0] = poured as f64;
        for i in 0..=query_row {
            for j in 0..=i {
                // 如果当前玻璃杯接收的香槟量大于 1 ，
                // 则多余的香槟会平均地流入到它下面左右两边的玻璃杯中。
                if dp[i][j] > 1.0 {
                    // 将多余的香槟均分
                    let extra = (dp[i][j] - 1.0) / 2.0;
                    // 它下面左右两边的玻璃杯均会接收到 extra 杯香槟
                    dp[i + 1][j] += extra;
                    dp[i + 1][j + 1] += extra;
                    // 当前玻璃杯多余的香槟流出后，只剩 1 杯香槟
                    dp[i][j] = 1.0;
                }
            }
        }

        dp[query_row][query_glass]
    }
}
