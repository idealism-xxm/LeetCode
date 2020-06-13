// 链接：https://leetcode.com/problems/combination-sum-iii/
// 题意：在 1～9 中选择 k 个数，使其和为 n ，求所有这样的组合？

// 输入： k = 3, n = 7
// 输出： [[1,2,4]]

// 输入： k = 3, n = 9
// 输出： [[1,2,6], [1,3,5], [2,3,4]]

// 思路： 递归
//
//		直接递归模拟即可，定义一个 dfs 函数，
//      在 digit~9 中选择 k 个数，使其和为 n ，当前收集数在 list 中，
//      最后返回收集的所有结果即可

impl Solution {
    pub fn combination_sum3(k: i32, n: i32) -> Vec<Vec<i32>> {
        let mut list = vec![0; k as usize];
        // 递归收集结果
        Solution::dfs(k, n, 1, &mut list)
    }

    // 在 digit~9 中选择 k 个数，使其和为 n
    pub fn dfs(k: i32, n: i32, digit: i32, list: &mut Vec<i32>) -> Vec<Vec<i32>> {
        // 不再需要新的数了，准备返回结果
        if k == 0 {
            return if n == 0 {
                // 如果已满足要求，则返回当前收集的列表
                vec![list.clone(); 1]
            } else {
                // 还有未满足的，返回空列表
                Vec::new()
            }
        }

        // 剪枝，如果还剩的数，不够，或者最小也比 n 大，或者最大也比 n 小，则直接返回空列表
        let min_sum = digit * k + k * (k - 1) / 2;
        let max_sum = 9 * k - k * (k - 1) / 2;
        if k > 9 - digit + 1 || min_sum > n || max_sum < n {
            return Vec::new()
        }

        // 遍历收集结果
        let mut result = Vec::new();
        // 总共要收集 list.len() 个数，还需要收集 k 个，则当前收集的数下标为 list.len() - k
        let i = list.len() - k as usize;
        for num in digit..=9 {
            // 确定当前位置的数
            list[i] = num;
            // 递归收集结果
            result.append(&mut Solution::dfs(k - 1, n - num, num + 1, list))
        }

        // 返回当前收集的结果
        result
    }
}
