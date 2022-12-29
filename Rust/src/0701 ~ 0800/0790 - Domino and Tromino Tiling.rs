// 链接：https://leetcode.com/problems/domino-and-tromino-tiling/
// 题意：有两种类型的瓷砖，求铺满 2 * n 的地板共有多少种方案？
//          1. 第一种类型的瓷砖大小为 2 ，形状为 ⠰
//          2. 第二种类型的瓷砖大小为 3 ，形状为 ⠚


// 数据限制：
//  1 <= n <= 1000


// 输入： n = 3
// 输出： 5

// 输入： n = 1
// 输出： 1
// 解释： 只能铺 1 个大小为 2 的瓷砖，且只有一种朝向。


// 思路1： DP
//
//      如果能看出是 DP ，但一时想不到简单的状态和转移方程时，
//      可以细化状态，提升状态维度，从而降低复杂度。
//
//
//      设 dp[i][j] 表示前 i 列地板中，前 i - 1 列已铺满的情况下，
//      最后一列的状态为 j 时的方案数
//          1. j == 0: 表示第 j 列铺满两个瓷砖
//          2. j == 1: 表示第 j 列上面没铺瓷砖，下面铺了瓷砖
//          3. j == 2: 表示第 j 列下面没铺瓷砖，上面铺了瓷砖
//
//      初始化：
//          1. 前 0 列铺满，共有一种方案数
//          2. 前 1 列铺满，共有一种方案数
//
//      状态转移：
//          1. dp[i][0] = (sum(dp[i - 1]) + dp[i - 2][0]) % MOD
//             前 i 列全部铺满的方案数，可由以下状态转移而来：
//              (1) dp[i - 1][0]: 前 i - 1 列全部铺满，第 i 列竖着放一块瓷砖 ⠰
//              (2) dp[i - 1][1]: 前 i - 1 列最后一列上面没铺瓷砖，第 i 列放一块瓷砖 ⠙
//              (3) dp[i - 1][2]: 前 i - 1 列最后一列下面没铺瓷砖，第 i 列放一块瓷砖 ⠚
//              (4) dp[i - 2][0]: 前 i - 2 列全部铺满，后两列横着放两块瓷砖 ⠭
//          2. dp[i][1] = (dp[i - 2][0] + dp[i - 1][2]) % MOD
//             前 i 列最后一列上面没铺瓷砖的方案数，可由以下状态转移而来：
//              (1) dp[i - 2][0]: 前 i - 2 列全部铺满，后两列放一块瓷砖 ⠓
//              (2) dp[i - 1][2]: 前 i - 1 列最后一列下面没铺瓷砖，第 i 列下面横着放一块瓷砖 ⠤
//          3. dp[i][2] = (dp[i - 2][0] + dp[i - 1][1]) % MOD   
//             前 i 列最后一列下面没铺瓷砖的方案数，可由以下状态转移而来：
//              (1) dp[i - 2][0]: 前 i - 2 列全部铺满，后两列放一块瓷砖 ⠋
//              (2) dp[i - 1][1]: 前 i - 1 列最后一列上面没铺瓷砖，第 i 列上面横着放一块瓷砖 ⠉
//
//      最后 dp[n][0] 就是前 n 列全部铺满的方案数。
//
//
//      DP 常见的三种优化方式见 LeetCode 583 这题的思路，
//      本题只能采用滚动数组的方式进行优化，将空间复杂度从 O(n) 优化为 O(1) 。
//      本实现为了便于理解，不做优化处理。
//
//
//      进阶优化：使用矩阵快速幂处理，可以将时间复杂度优化为 (logn * C ^ 3) ，其中 C 等于 3 。
//
//
//		时间复杂度： O(n)
//          1. 需要遍历 dp 中全部 O(n) 个状态
//		空间复杂度： O(n)
//          1. 需要维护一个大小为 O(n) 的二维数组 dp


const MOD: i32 = 1_000_000_007;


impl Solution {
    pub fn num_tilings(n: i32) -> i32 {
        let n = n as usize;
        // 设 dp[i][j] 表示前 i 列地板中，前 i - 1 列已铺满的情况下，
        // 最后一列的状态为 j 时的方案数
        //  1. j == 0: 表示第 j 列铺满两个瓷砖
        //  2. j == 1: 表示第 j 列上面没铺瓷砖，下面铺了瓷砖
        //  3. j == 2: 表示第 j 列下面没铺瓷砖，上面铺了瓷砖
        let mut dp = vec![vec![0, 0, 0]; n + 1];
        // 初始化：
        //  1. 前 0 列铺满，共有一种方案数
        //  2. 前 1 列铺满，共有一种方案数
        dp[0][0] = 1;
        dp[1][0] = 1;
        for i in 2..=n {
            // 前 i 列全部铺满的方案数，可由以下状态转移而来：
            //  1. dp[i - 1][0]: 前 i - 1 列全部铺满，第 i 列竖着放一块瓷砖 ⠰
            //  2. dp[i - 1][1]: 前 i - 1 列最后一列上面没铺瓷砖，第 i 列放一块瓷砖 ⠙
            //  3. dp[i - 1][2]: 前 i - 1 列最后一列下面没铺瓷砖，第 i 列放一块瓷砖 ⠚
            //  4. dp[i - 2][0]: 前 i - 2 列全部铺满，后两列横着放两块瓷砖 ⠭
            dp[i][0] = (Self::sum_with_mod(&dp[i - 1]) + dp[i - 2][0]) % MOD;
            // 前 i 列最后一列上面没铺瓷砖的方案数，可由以下状态转移而来：
            //  1. dp[i - 2][0]: 前 i - 2 列全部铺满，后两列放一块瓷砖 ⠓
            //  2. dp[i - 1][2]: 前 i - 1 列最后一列下面没铺瓷砖，第 i 列下面横着放一块瓷砖 ⠤
            dp[i][1] = (dp[i - 2][0] + dp[i - 1][2]) % MOD;
            // 前 i 列最后一列下面没铺瓷砖的方案数，可由以下状态转移而来：
            //  1. dp[i - 2][0]: 前 i - 2 列全部铺满，后两列放一块瓷砖 ⠋
            //  2. dp[i - 1][1]: 前 i - 1 列最后一列上面没铺瓷砖，第 i 列上面横着放一块瓷砖 ⠉
            dp[i][2] = (dp[i - 2][0] + dp[i - 1][1]) % MOD;
        }
        
        // 最后返回前 n 列全部铺满的方案数
        dp[n][0]
    }

    fn sum_with_mod(nums: &Vec<i32>) -> i32 {
        let mut ans = 0;
        for &num in nums {
            ans = (ans + num) % MOD;
        }

        ans
    }
}
