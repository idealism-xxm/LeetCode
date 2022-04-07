// 链接：https://leetcode.com/problems/last-stone-weight/
// 题意：给定一个非负整数数组 stones ，
//      每次从中移除最大的两个数 x 和 y (x <= y) ，
//      如果 x != y ，则再将 y - x 放回 stones 中。
//      不断重复这个过程，直至 stones 中的数字不足 2 个。
//      如果此时 stones 还有数字，则返回该数字，否则返回 0 。


// 数据限制：
//  1 <= stones.length <= 30
//  1 <= stones[i] <= 1000


// 输入： stones = [2,7,4,1,8,1]
// 输出： 1
// 解释： 移除 7 和 8 ， 8 - 7 = 1 ，所以 stones 转变为 [2,4,1,1,1] 。
//       移除 2 和 4 ， 4 - 2 = 2 ，所以 stones 转变为 [2,1,1,1] 。
//       移除 2 和 1 ， 2 - 1 = 1 ，所以 stones 转变为 [1,1,1] 。
//       移除 1 和 1 ， 1 - 1 = 0 ，所以 stones 转变为 [1] 。
//       stones 还剩一个数字，所以返回 1 。

// 输入： stones = [1]
// 输出： 1
// 解释： stones 只有一个数字，直接返回 1 。


// 思路： 优先队列（堆）
//
//      定义一个最大堆 heap ，维护数组中剩余的数字，
//      初始为 stones 中的数字。
//
//      当堆中数字个数大于 1 时，不断循环处理，
//      每次移除堆顶两个数字 x 和 y (x <= y) ，
//      如果 x != y ，则再将 y - x 放回堆中。
//
//      最后，如果堆中没有数字，返回 0 ；
//      如果堆中剩余一个数字，返回该数字。
//      
//
//		时间复杂度： O(nlogn)
//          1. 通过数组直接建立堆，时间复杂度为 O(n)
//          2. 每次循环至少会移除一个数字，总共会有 O(n) 次循环
//          3. 每次循环时，需要执行出堆和入堆操作，时间复杂度为 O(logn)
//		空间复杂度： O(n)
//          1. 需要维护一个包含 O(n) 个数字的堆

use std::collections::BinaryHeap;

impl Solution {
    pub fn last_stone_weight(stones: Vec<i32>) -> i32 {
        // 通过数组直接建立堆
        let mut heap = BinaryHeap::from(stones);
        // 至少还有两个数字时，继续循环处理
        while heap.len() > 1 {
            // 获取最大数字
            let y = heap.pop().unwrap();
            // 获取次大数字
            let x = heap.pop().unwrap();
            // 如果 x != y ，则再将 y - x 放回堆中
            if x != y {
                heap.push(y - x);
            }
        }

        // 如果此时还有数字，则返回该数字，否则返回 0
        heap.pop().unwrap_or(0)
    }
}
