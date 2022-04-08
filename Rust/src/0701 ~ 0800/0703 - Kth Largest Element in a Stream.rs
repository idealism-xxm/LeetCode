// 链接：https://leetcode.com/problems/kth-largest-element-in-a-stream/
// 题意：设计一个数据结构，支持添加一串数字流，
//      每次添加时返回当前第 k 大的数字。


// 数据限制：
//  1 <= k <= 10 ^ 4
//  0 <= nums.length <= 10 ^ 4
//  -(10 ^ 4) <= nums[i] <= 10 ^ 4
//  -(10 ^ 4) <= val <= 10 ^ 4
//  最多有 10 ^ 4 次 add 函数的调用
//  确保查找时，第 k 大的数字必定存在


// 输入： ["KthLargest", "add", "add", "add", "add", "add"]
//       [[3, [4, 5, 8, 2]], [3], [5], [10], [9], [4]]
// 输出： [null, 4, 5, 5, 8, 8]
// 解释： KthLargest kthLargest = new KthLargest(3, [4, 5, 8, 2]);
//       kthLargest.add(3);   // return 4
//       kthLargest.add(5);   // return 5
//       kthLargest.add(10);  // return 5
//       kthLargest.add(9);   // return 8
//       kthLargest.add(4);   // return 8


// 思路： 优先队列（堆）
//
//      定义一个最小堆，维护数字流中最大的 k 个数字。
//
//      在 new 中初始化时直接使用 nums 数组建立最小堆，
//      然后不断从堆顶中移除数字，
//      直至堆中的数字个数小于等于 k 。
//
//      在 add 中，将新数字 val 插入堆中，
//      如果此时堆中数字个数大于 k ，再移除堆顶的数字。
//
//      此时堆顶数字就是第 k 大的数字，直接返回即可。
//
//
//      设初始化时的数组长度为 n ， add 总共调用 m 次。
//
//		时间复杂度： O(nlogn + mlogk)
//          1. new 中通过数组直接建立堆，时间复杂度为 O(n)
//          2. new 中需要从堆中移除数字，直至剩余 k 个数字，
//              时间复杂度为 O(nlogn)
//          3. add 每次需要往堆中添加/移除一个数字，
//              时间复杂度为 O(logk)
//          4. add 总共会调用 O(m) 次
//		空间复杂度： O(n)
//          1. new 中初始化时需要维护一个包含 O(n) 个数字的堆

use std::{collections::BinaryHeap, cmp::Reverse, iter::FromIterator};

struct KthLargest {
    // 最小堆中最少的数字个数
    k: usize,
    // 需要使用最小堆维护最大的 k 个数字。
    // 由于 BinaryHeap 默认是最大堆，
    // 所以需要使用 Reverse 结构体转换成最小堆。
    heap: BinaryHeap<Reverse<i32>>,
}


/** 
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl KthLargest {

    fn new(k: i32, nums: Vec<i32>) -> Self {
        let k = k as usize;
        // 根据 nums 初始化最小堆，时间复杂度为 O(n)
        let mut heap = BinaryHeap::from_iter(nums.iter().map(|&num| Reverse(num)));
        // 不断移除堆顶数字，直到堆中数字个数小于等于 k
        while heap.len() > k {
            heap.pop();
        }

        // 返回 KthLargest 实例
        KthLargest { k, heap }
    }
    
    fn add(&mut self, val: i32) -> i32 {
        // 将当前数字放入堆中
        self.heap.push(Reverse(val));
        // 如果堆中数字个数大于 k ，则移除堆顶的数字
        if self.heap.len() > self.k {
            self.heap.pop();
        }

        // 返回堆顶数字
        self.heap.peek().unwrap().0
    }
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * let obj = KthLargest::new(k, nums);
 * let ret_1: i32 = obj.add(val);
 */
