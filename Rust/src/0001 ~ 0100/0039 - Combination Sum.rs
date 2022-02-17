// 链接：https://leetcode.com/problems/combination-sum/
// 题意：给定一些互不相同的正整数列表 candidates 和一个目标整数 target ，
//      candidates 中的每个数可以无限选取，从中选取若干个数形成一个列表，
//      求所有和为 target 的列表？


// 数据限制：
//  1 <= candidates.length <= 30
//  1 <= candidates[i] <= 200
//  candidates 中所有的数字均不相同
//  1 <= target <= 500


// 输入：candidates = [2,3,6,7], target = 7
// 输出：[[2,2,3],[7]]
// 解释：只有两种组合
//      2 + 2 + 3 = 7 ，注意 2 可以被使用多次
//      7 = 7


// 输入：candidates = [2,3,5], target = 8
// 输出：[[2,2,2,2],[2,3,3],[3,5]]

// 输入：candidates = [2], target = 1
// 输出：[]


// 思路：回溯
//
//		按照顺序递归处理 candidates 中的第 cur 个数，每一层可做两个操作：
//		    1. candidates[cur] 不放入结果待选列表 list 中，
///             下一层递归处理 candidates[cur + 1]
//          2. candidates[cur] 放入结果待选列表 list 中，
//              下一层递归处理 candidates[cur] ，因为一个数可以重复选择
//
//      递归终止条件是：
//          1. target == 0: 这时候结果待选列表 list 就是一种合法的子集
//          2. candidates.len() == cur || target < candidates[cur]: 
//              已经遍历完所有的数字，且当前还未返回，
//              则当前 list 就是一种不合法的列表，直接返回
//
//
//      【进阶】需要保证数字不能重复选择。
//
//      这就是 [LeetCode 40](https://leetcode.com/problems/combination-sum-ii/) 这题，只需要修改递归 cur 的改变逻辑即可，
//      无论是否选择了第 cur 个数，下一层都要从 cur + 1 开始。
//
//      设结果列表大小为 S
//
//		时间复杂度： O(S * target)
//          1. 要遍历完 O(S) 个结果
//          2. 最后收集每个结果时都要克隆 list 中的数字，最长为 O(target)
//          3. 实际难以计算比较上界
//		空间复杂度： O(S * target)
//          1. 总共有 O(S) 个结果，每个结果最长为 O(target)
//          2. 实际难以计算比较上界


impl Solution {
    pub fn combination_sum(mut candidates: Vec<i32>, target: i32) -> Vec<Vec<i32>> {
        // 先对 candidates 按升序排序，方便后续优化和处理
        candidates.sort();

        // list 用于收集当前列表内的数字，最大长度为 target / candidates[0] + 1
        let mut list = Vec::with_capacity((target / candidates[0] + 1) as usize);
        // ans 用于收集所有可能的列表
        let mut ans = Vec::new();
        Self::dfs(&candidates, target, 0, &mut list, &mut ans);

        ans
    }

    fn dfs(candidates: &Vec<i32>, target: i32, cur: usize, list: &mut Vec<i32>, ans: &mut Vec<Vec<i32>>) {
        // 如果列表中的和已满足题意，则当前 list 就是一种合法的列表
        if target == 0 {
            ans.push(list.clone());
            return;
        }
        // 如果已经遍历完所有的数字 或者 能选的最小数字都比 target 大，
        // 则当前 list 就是一种不合法的列表，直接返回
        if cur == candidates.len() || target < candidates[cur] {
            return;
        }

        // 不选第 cur 个数，递归处理目标数为 target 的情况，从 cur + 1 开始
        Self::dfs(candidates, target, cur + 1, list, ans);
        // 选择第 cur 个数 ，将其加入 list 中
        list.push(candidates[cur]);
        // 递归处理目标数为 target - candidates[cur] 的情况，
        // 从 cur 开始，保证可重复选
        Self::dfs(candidates, target - candidates[cur], cur, list, ans);
        // 移除最后一个数
        list.pop();
    }
}
