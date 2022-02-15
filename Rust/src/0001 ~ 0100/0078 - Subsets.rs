// 链接：https://leetcode.com/problems/subsets/
// 题意：给定一个数字集合，求所有子集？


// 数据限制：
//  1 <= nums.length <= 10
//  -10 <= nums[i] <= 10
//  nums 中的数字各不相同


// 输入：nums = [1,2,3]
// 输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

// 输入：nums = [0]
// 输出：[[],[0]]


// 思路：回溯
//
//		按照顺序递归处理每一个数，每一层可做两个操作：
//		    1. 放入结果待选列表 list
//          2. 不放入结果待选列表 list
//
//      递归终止条件是：所有数字都已枚举完毕，
//      这时候结果待选列表 list 就是一种合法的子集
//
//      时间复杂度：O(n * 2 ^ n)
//          1. 每个数字都要进行两种选择，总共有 n 个数字要这样处理，时间复杂度时 O(2 ^ n)
//          2. 最后收集每个子集时都要克隆 list 中的数字，最长为 O(n)
//      空间复杂度：O(n * 2 ^ n)
//          1. 总共有 O(2 ^ n) 个子集，每个子集最长为 O(n)
//          2. 实际计算所有子集的元素个数的公式为： sum(i * C(n, i)) = n * 2 ^ (n - 1)

impl Solution {
    pub fn subsets(nums: Vec<i32>) -> Vec<Vec<i32>> {
        // list 用于收集当前子集内的数字，最大长度为 nums.len()
        let mut list = Vec::with_capacity(nums.len());
        // ans 用于收集所有可能的子集
        let mut ans = vec![];
        // 回溯遍历所有可能情况
        Self::dfs(&nums[..], &mut list, &mut ans);

        ans
    }

    fn dfs(nums: &[i32], mut list: &mut Vec<i32>, mut ans: &mut Vec<Vec<i32>>) {
        // 如果已经遍历完所有的数字，则当前 list 就是一种合法的子集
        if nums.len() == 0 {
            ans.push(list.clone());
            return
        }

        // 不选第一个数
        Self::dfs(&nums[1..], list, ans);
        // 选第一个数，则先把 nums[0] 放入到 list 中
        list.push(nums[0]);
        // 递归处理剩余的数
        Self::dfs(&nums[1..], list, ans);
        // 移除最后一个数
        list.pop();
    }
}
