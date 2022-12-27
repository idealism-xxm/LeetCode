// 链接：https://leetcode.com/problems/maximum-bags-with-full-capacity-of-rocks/
// 题意：给定两个长度为 n 的数组 capacity 和 rocks ，和一个正整数 additionalRocks 。
//
//      其中， capacity[i] 表示第 i 个袋子的容量， rocks[i] 表示第 i 个袋子初始装的石头数，
//      additionalRocks 表示现在额外还有的石头数。
//
//      求最多能装满多少个袋子？


// 数据限制：
//  n == capacity.length == rocks.length
//  1 <= n <= 5 * 10 ^ 4
//  1 <= capacity[i] <= 10 ^ 9
//  0 <= rocks[i] <= capacity[i]
//  1 <= additionalRocks <= 10 ^ 9


// 输入： capacity = [2,3,4,5], rocks = [1,2,4,4], additionalRocks = 2
// 输出： 3
// 解释： 往袋子 0 和袋子 1 中各放一块石头，所有袋子的石头数变为 [2,3,4,4] 。
//       袋子 0, 1 和 2 已装满，共 3 个装满的袋子。
//
//       【注意】还有其他方案也能装满 3 个袋子

// 输入： capacity = [10,2,2], rocks = [2,2,0], additionalRocks = 100
// 输出： 3
// 解释： 往袋子 0 中放 8 块石头，往袋子 2 中放 2 块石头，
//       所有袋子的石头数变为 [10,2,2] 。
//       袋子 0, 1 和 2 已装满，共 3 个装满的袋子。
//
//       【注意】没必要用完全部额外的石头


// 思路： 贪心 + 排序
//
//      要让最终装满的袋子数最多，我们可以贪心地优先装满需要更少石头的袋子。
//
//
//      所以我们先计算出每个袋子还需要的石头数到 required 中，
//      然后对 required 按照升序排序。
//
//      再按顺序装满袋子，直至所需的石头数不足，或者装满了全部袋子。
//
//
//      时间复杂度：O(nlogn)
//          1. 需要遍历计算 required 中全部 O(n) 个元素
//          2. 需要对 required 中全部 O(n) 个元素排序，时间复杂度为 O(nlogn)
//      空间复杂度：O(n)
//          1. 需要维护 required 中全部 O(n) 个元素


impl Solution {
    pub fn maximum_bags(capacity: Vec<i32>, rocks: Vec<i32>, mut additional_rocks: i32) -> i32 {
        let n = capacity.len();
        // required[i] 表示第 i 个袋子还需要的石头数
        let mut required = Vec::with_capacity(n);
        for (c, rock) in capacity.into_iter().zip(rocks.into_iter()) {
            required.push(c - rock);
        }
        // 按升序排序，贪心地优先装满需要更少石头的袋子
        required.sort();

        // ans 维护最多能装满的袋子数
        let mut ans = 0;
        for r in required {
            // 如果所需的石头数不足，则无法继续装满袋子，跳出循环
            if r > additional_rocks {
                break;
            }
            
            // 拿出 r 块石头装满当前袋子
            additional_rocks -= r;
            ans += 1;
        }
        
        ans
    }
}
