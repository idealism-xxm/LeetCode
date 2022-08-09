// 链接：https://leetcode.com/problems/binary-trees-with-factors/
// 题意：给定不含重复数字的整数数组 arr ，其中每个数字都严格大于 1 。
//
//      现在需要用这些数字组成一个二叉树，每个数字都用任意次，
//      要求该二叉树的所有非叶子结点值等于其子结点值的乘积。
//
//      求共能组成多少种不同的二叉树？


// 数据限制：
//  1 <= arr.length <= 1000
//  2 <= arr[i] <= 10 ^ 9
//  arr 中所有的值都各不相同


// 输入： arr = [2,4]
// 输出： 3
// 解释： 所有可能的二叉树如下：
//      ①     ②     ③
//      2      4      4
//                   / \
//                  2   2

// 输入： arr = [2,4,5,10]
// 输出： 7
// 解释： 所有可能的二叉树如下：
//      ①     ②     ③      ④      ⑤      ⑥      ⑦
//      2      4      5      10      4      10      10
//                                  / \     / \     / \
//                                 2   2   2   5   5   2


// 思路： DP + 排序
//
//      对于满足题意的二叉树的来说，如果其非叶子结点值为 x ，
//      则其左右子结点的值 left 和 right 必定小于 x 。
//
//      这种隐含的约束表明我们可以使用 DP 来进行状态转移，
//      只要我们按照从小到大的顺序处理 arr 中的数，
//      那么所有需要的子状态在转移时都已确定。
//
//
//      我们先对 arr 按升序排序，保证后续状态转移时，子状态都已确定。
//
//      设 dp[i] 表示以 i 为根的二叉树的个数，
//      由于 i 是离散的，而且非常大，我们可以使用 Map 来维护这个状态。
//
//      枚举 arr 中的每个数作为二叉树的根 root ，并初始化 dp[root] = 1 ，
//      因为只含有根 root 一个结点的二叉树必定满足题意。
//
//      然后，我们统计以 root 为根 且 至少有三个结点的二叉树。
//
//      枚举 arr 中的每个数作为二叉树的左子结点 left ，必有 left < root 。
//          1. root % left != 0: 则无法组成满足题意的二叉树，直接处理下一个
//          2. root % left == 0: 则先计算对应的右子结点的值 right 。
//              如果 right 在 arr 中（即 right 必定在 dp 中），则可以组成满足题意的二叉树。
//
//              此时，需要统计以 left 和 right 作为 root 左右子结点的二叉树的个数，
//              即进行状态转移： dp[root] += dp[left] * dp[right] 。
//
//      最后，对 dp 中所有的状态值求和即可得到答案。 
//
//
//      时间复杂度：O(n ^ 2)
//          1. 需要对 arr 全部 O(n) 个数字排序，时间复杂度为 O(nlogn)
//          2. 需要枚举 arr 全部 O(n) 个数作为二叉树根结点 root ，
//              枚举每个 root 时，都需要枚举 arr 全部 O(n) 个数作为其左子结点 left 。
//              综上：时间复杂度为 O(n ^ 2)
//      空间复杂度：O(n)
//          1. 需要维护 dp 的全部 O(n) 个状态  


use std::collections::HashMap;


const MOD: i64 = 1_000_000_007;


impl Solution {
    pub fn num_factored_binary_trees(mut arr: Vec<i32>) -> i32 {
        // 对数组按升序排序
        arr.sort();
        // dp[i] 表示以 i 为根的二叉树的个数
        let mut dp: HashMap<i32, i64> = HashMap::with_capacity(arr.len());
        let mut ans = 0;
        for (i, &root) in arr.iter().enumerate() {
            // cnt 表示以 root 为根的二叉树的个数，
            // 初始时，只含有根 root 一个结点的二叉树必定满足题意
            let mut cnt = 1;
            // 枚举 root 左子树的根结点的值 left ，其必定比 root 小
            for &left in arr.iter().take(i) {
                // 如果 root 不能整除 left ，则不满足题意，直接处理下一个
                if root % left != 0 {
                    continue;
                }
                // 计算 root 的右子树的根结点的值
                let right = root / left;
                // 如果 left 和 right 可以作为 root 的左右子结点，
                // 则 right 必定在 arr 中，即 right 必定在 dp 中
                if let Some(right_count) = dp.get(&right) {
                    let left_count = dp.get(&left).unwrap();
                    // 加上以 left 和 right 作为 root 左右子结点的二叉树的个数
                    cnt += left_count * right_count;
                    cnt %= MOD;
                }
            }

            // 将 cnt 加入 dp 中，方便后续进行状态转移
            dp.insert(root, cnt);
            // 统计以所有 root 为根的二叉树的个数
            ans = (ans + cnt) % MOD;
        }

        ans as i32
    }
}
