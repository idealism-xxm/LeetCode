// 链接：https://leetcode.com/problems/remove-stones-to-minimize-the-total/
// 题意：给定一个整型数组 nums ，每次可以将其中一个数 nums[i] 变为 ceil(nums[i] / 2) ，
//      求 k 次这样的操作后，所有数的和最小是多少？


// 数据限制：
//   1 <= piles.length <= 10 ^ 5
//   1 <= piles[i] <= 10 ^ 4
//   1 <= k <= 10 ^ 5


// 输入： piles = [5,4,9], k = 2
// 输出： 12
// 解释： [5,4,9] -> [5,4,5] -> [3,4,5]

// 输入： piles = [4,3,6,7], k = 3
// 输出： 12
// 解释： [4,3,6,7] -> [4,3,3,7] -> [4,3,3,4] -> [2,3,3,4]



// 思路： 贪心 + 堆
//
//      为了使最终的数最小，那么每次都尽量将当前最大的数执行这个操作.
//
//      我们维护一个最大堆，每次从堆中取当前最大的数 cur ，
//      然后将 ceil(cur / 2) 放回堆，这样操作 k 次后，所有数的和就是答案。
//
//
//      时间复杂度： O(klogn)
//          1. 需要初始化堆中全部 O(n) 个数，初始化堆时间复杂度为 O(n)
//          2. 需要对堆执行 O(k) 次 pop/push 操作，每次操作的时间复杂度为 O(logn)
//          3. 需要求堆中全部 O(n) 个数的和
//      空间复杂度： O(n)
//          1. 需要维护堆中全部 O(n) 个数


use std::collections::BinaryHeap;
use std::iter::FromIterator;


impl Solution {
    pub fn min_stone_sum(piles: Vec<i32>, k: i32) -> i32 {
        // 建立最大堆，时间复杂度为 O(m)
        let mut heap = BinaryHeap::from_iter(piles.into_iter());
        // 执行 k 次操作
        for _ in 0..k {
            // 从堆中取当前最大的数 cur
            let cur = heap.pop().unwrap();
            // 将 ceil(cur / 2) 放回堆
            heap.push((cur + 1) >> 1);
        }

        // 堆中所有数的和就是答案
        heap.into_iter().sum()
    }
}
