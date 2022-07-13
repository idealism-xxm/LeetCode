// 链接：https://leetcode.com/problems/matchsticks-to-square/
// 题意：给定一个数组 matchsticks ，其中 matchsticks[i] 表示第 i 根火柴棍的长度。
//      求恰好用每根火柴棍一次，能否拼出一个正方形？


// 数据限制：
//  1 <= matchsticks.length <= 15
//  1 <= matchsticks[i] <= 10 ^ 8


// 输入：matchsticks = [1,1,2,2,2]
// 输出：true
// 解释：可以拼出一个正方形，每边长度为 2 。

// 输入：matchsticks = [3,3,3,3,4]
// 输出：false
// 解释：无法使用全部火柴棍拼出一个正方形。


// 思路：递归/回溯/DFS
//
//      如果这些火柴棍能拼出一个正方形，必定先满足下面的两个条件：
//          1. 火柴棍的个数至少为 4 根
//          2. 所有火柴棍的长度之和 total 能被 4 整除
//
//      然后我们使用 dfs(matchsticks, i, remain) 回溯遍历所有可能的组合：
//          1. matchsticks: 所有火柴棍，直接透传
//          2. i: 当前需要放置的火柴棍的下标，初始化为 len(matchsticks) - 1
//          3. remain: 长度为 4 的整型数组， remain[j] 表示第 j 条边还需的长度，
//                  初始化均为 total / 4
//
//      在 dfs 中，我们按照如下逻辑处理：
//          1. i < 0: 已放置完全部火柴棍，则必定有 remain[0~4] = 0 ，
//                  即此时可以拼出一个正方形，直接返回 true
//          2. i >= 0: 枚举第 i 根火柴棍放置的边 j ，如果 remain[j] >= 0 ，
//                  则将其放置在第 j 条边，然后递归处理下根火柴棍。
//
//                  如果递归返回结果为 true ，则说明此时能拼出一个正方形，直接返回 true ；
//                  否则恢复状态，继续处理下一条边。
//
//                  如果放置在所有边时，都不能拼出正方形，则返回 false 。
//
//      我们在调用 dfs 前可以进一步优化：先对 matchsticks 按照升序排序，
//      然后在回溯时先枚举更长的火柴棍，这样能减小搜索空间。
//
//
//		时间复杂度： O(4 ^ n)
//          1. 需要对 matchsticks 全部 O(n) 个元素排序，时间复杂度为 O(nlogn)
//          2. 需要对 matchsticks 全部 O(n) 个元素回溯，
//              每次都有 4 种选择，时间复杂度为 O(4 ^ n)
//		空间复杂度： O(n)
//          1. 栈递归深度最大为 O(n)


impl Solution {
    pub fn makesquare(mut matchsticks: Vec<i32>) -> bool {
        // 如果不足 4 个火柴棍，则不能拼出正方形
        if matchsticks.len() < 4 {
            return false;
        }
        // 计算所有火柴棍的长度之和，如果不能被 4 整除，则不能拼出正方形
        let total = matchsticks.iter().sum::<i32>();
        if total % 4 != 0 {
            return false;
        }
        // 初始化 4 条边都还需要 total / 4 的长度
        let mut remain = vec![total / 4; 4];
        // 将火柴棍长度升序排序，回溯时先枚举更长的火柴棍，能减小搜索空间
        matchsticks.sort();
        
        // 回溯搜索
        let last_index = matchsticks.len() as i32 - 1;
        Self::dfs(&mut matchsticks, last_index, &mut remain)
    }

    fn dfs(matchsticks: &Vec<i32>, i: i32, remain: &mut Vec<i32>) -> bool {
        // 如果放置完全部火柴棍，则必定有 remain[0~4] = 0 ，
        // 即此时可以拼出一个正方形，直接返回 true
        if i < 0 {
            return true;
        }

        let matchstick = matchsticks[i as usize];
        // 枚举第 i 根火柴棍放置的边
        for j in 0..4 {
            // 如果第 j 条边的所需的长度小于当前火柴棍的长度，则直接处理下一条边
            if remain[j] < matchstick {
                continue;
            }

            // 当前边所需的长度减去当前火柴棍的长度
            remain[j] -= matchstick;
            // 递归处理下根火柴棍，如果能成功拼出正方形，则直接返回
            if Self::dfs(matchsticks, i - 1, remain) {
                return true;
            }
            // 此时不能拼出正方形，需要恢复状态，将当前边所需的长度加回去
            remain[j] += matchstick;
        }

        // 当前情况下，第 i 根火柴棍放在哪条边都不能拼出正方形，返回 false
        false
    }
}
