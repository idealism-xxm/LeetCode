// 链接：https://leetcode.com/problems/the-k-weakest-rows-in-a-matrix/
// 题意：给定一个 m * n 的二维 01 数组 mat ，每一行中所有的 1 都在 0 之前，
//      现在返回最小的前 k 行，按从小到大排序。
//
//      第 i 行比第 j 行小时，满足以下两个条件之一：
//          1. 第 i 行中 1 的数量小于第 j 行中 1 的数量
//          2. 第 i 行中 1 的数量等于第 j 行中 1 的数量，且 i < j


// 数据限制：
//  m == mat.length
//  n == mat[i].length
//  2 <= n, m <= 100
//  1 <= k <= m
//  matrix[i][j] 是 1 或 0


// 输入： mat = 
//          [[1,1,0,0,0],
//          [1,1,1,1,0],
//          [1,0,0,0,0],
//          [1,1,0,0,0],
//          [1,1,1,1,1]], 
//       k = 3
// 输出： [2,0,3]
// 解释： 每一行的 1 的数量如下：
//          第 0 行： 2
//          第 1 行： 4
//          第 2 行： 1
//          第 3 行： 2
//          第 4 行： 5
//      行号按从小到大排序后为： [2,0,3,1,4]

// 输入： mat = 
//          [[1,0,0,0],
//          [1,1,1,1],
//          [1,0,0,0],
//          [1,0,0,0]], 
//      k = 2
// 输出： [0,2]
// 解释： 每一行的 1 的数量如下：
//          第 0 行： 1
//          第 1 行： 4
//          第 2 行： 1
//          第 3 行： 1
//      行号按从小到大排序后为： [0,2,3,1]


// 思路： 二分 + 优先队列（堆）
//
//      最简单的想法就是统计每一行的 1 的数量，时间复杂度为 O(mn) ，
//      然后按照 1 的数量维护一个最多有 k 个元素的最大堆，时间复杂度为 O(mlogk) ，
//      最后收集最大堆中的 k 个行号，时间复杂度为 O(klogk)。
//
//      这样总时间复杂度为 O(mn + mlogk + klogk) ，但还可以继续优化：
//
//      1. 可以发现每一行中，所有的 1 都在 0 之前，
//          所以可以用二分找到第一个 0 的下标，那这个下标就是 1 的个数，
//          时间复杂度优化为 (mlogn) 。
//
//      2. 我们可以先收集所有 O(m) 个元素，然后直接建一个最小堆，
//          而通过数组直接建堆的时间复杂度为 O(m) 。
//
//          那么后续收集最小堆中的 k 个元素的时间复杂度为 O(klogm) ，
//          则后两步时间复杂度优化为 O(m + klogm) ，
//          但空间复杂度上升为 O(m) 。
//
//
//      时间复杂度：O(mlogn + klogm)
//          1. 用二分统计全部 O(m) 行中 1 的个数，时间复杂度为 O(mlogn)
//          2. 收集全部 O(m) 行中的元素，直接建立最小堆，时间复杂度为 O(m)
//          3. 收集最小堆中的前 O(k) 个元素，时间复杂度为 O(klogm)
//      空间复杂度：O(m + k)
//          1. 需要收集全部 O(m) 行中的元素
//          2. 需要维护 O(m) 个元素的最小堆
//          3. 需要收集前 O(k) 个最小的元素

use std::{collections::BinaryHeap, iter::FromIterator, cmp::Reverse};

impl Solution {
    pub fn k_weakest_rows(mat: Vec<Vec<i32>>, k: i32) -> Vec<i32> {
        // 统计每一行的 1 的数量
        let arr_iter = mat
            .iter()
            .enumerate()
            .map(|(i, row)| {
                // 使用二分找到第一个 0 的下标，
                // 那么这个就是第 i 行中 1 的个数

                // 二分区间左边界 l
                let mut l = 0;
                // 二分区间右边界 r   
                let mut r = row.len() as i32 - 1;
                // 当区间不为空时，继续二分
                // （注意这里取等号是因为我们的区间是左闭右闭区间）
                while l <= r {
                    // 计算区间中点下标 mid
                    let mid = (l + r) >> 1;
                    if row[mid as usize] == 0 {
                        // 如果 mid 为 0，
                        // 则说明第一个 0 在左边区间中
                        r = mid - 1;
                    } else {
                        // 如果 mid 不为 0，
                        // 则说明第一个 0 在右边区间中
                        l = mid + 1;
                    }
                }

                // 最后返回当前行的 1 的数量，以及行下标
                //  （需要用 Reverse 包一层，
                //  因为 BinaryHeap 默认是最大堆）
                Reverse((l, i))
            });

        // 建立最小堆，时间复杂度为 O(m)
        let mut heap = BinaryHeap::from_iter(arr_iter);
        // 收集最小堆中的前 k 个元素，时间复杂度为 O(klogm)
        (0..k)
            // 取出对应的行下标
            .map(|_| heap.pop().unwrap().0.1 as i32)
            // 收集成结果数组
            .collect()
    }
}