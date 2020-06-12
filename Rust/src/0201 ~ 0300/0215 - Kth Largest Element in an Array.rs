// 链接：https://leetcode.com/problems/kth-largest-element-in-an-array/
// 题意：给定一个整数数组，返回第 k 大的数？

// 输入： [3,2,1,5,6,4] and k = 2
// 输出： 5

// 输入： [3,2,3,1,2,4,5,5,6] and k = 4
// 输出： 4

// 思路1： 优先队列
//
//		维护一个优先队列，遍历 nums 数组，
//      每次将 num[i] 放入优先队列中，
//      若队列中的数多余 k 个，则将最小的一个数移除
//
//      时间复杂度： O(nlogk)
//      空间复杂度： O(k)

use std::collections::BinaryHeap;
use std::cmp::Reverse;

impl Solution {
    pub fn find_kth_largest(nums: Vec<i32>, k: i32) -> i32 {
        let k = k as usize;
        // 最大堆
        let mut heap = BinaryHeap::new();
        nums.iter().for_each(|num| {
            // 由于是最大堆，所以想删除最小的元素就得变成最小堆
            heap.push(Reverse(num));
            // 如果元素多余 k 个，则把最小的移除
            if heap.len() > k {
                heap.pop();
            }
        });
        // 现在堆中最小的元素，就是第 k 大的元素
        *heap.pop().unwrap().0
    }
}

// 思路2： 快速排序
//
//		由于我们只需要找到第 k 大的数，不必知道其他数的顺序，
//      我们就可以利用快排的快速选择方法，
//      每次选择第一个数作为基准，将小于它的数移到最左边，将大于它的数移到右边，
//      然后看它的排名 rank ，
//          1. rank == k: 直接返回当前数
//          2. rank < k: 在右边查找第 k - rank 大的数
//          3. rank > k: 在左边查找第 k 大的数
//
//      时间复杂度： 平均 O(n) / 最差 O(n ^ 2)
//      空间复杂度： O(1)

impl Solution {
    pub fn find_kth_largest(nums: Vec<i32>, k: i32) -> i32 {
        // 让入参可以修改
        let mut nums = nums;
        let mut k = k as usize;
        // 每次要快速选择的范围
        let mut l: usize = 0;
        let mut r: usize = nums.len() - 1;
        while l < r {
            // 选第一个数作为基准，空出一个位置
            let pivot = nums[l];
            let mut ll = l;
            let mut rr = r;
            // 当范围内还有数需要确定的时候继续处理
            while ll < rr {
                // 从右边找到第一个 大于等于 pivot 的数
                while ll < rr && nums[rr] < pivot {
                    rr -= 1;
                }
                // 将这个数移动至左边的空位置 ll ，空出 rr 这个位置
                nums[ll] = nums[rr];
                // 从左边找到第一个 小于 pivot 的数
                while ll < rr && nums[ll] >= pivot {
                    ll += 1;
                }
                // 将这个数移动至左边的空位置 rr ，空出 ll 这个位置
                nums[rr] = nums[ll];
            }
            // 最后还有一个空位置 ll ，就是 pivot 的位置
            nums[ll] = pivot;
            // 算出 pivot 的排名
            let rank = ll - l + 1;
            // 如果是刚好要找的数，则直接返回
            if rank == k {
                return pivot;
            }
            // 如果不足 k 个数，则需要在右边找第 k - rank 个数
            if rank < k {
                // 在右边区间继续处理
                l = ll + 1;
                // 减去左边区间的数的个数
                k -= rank;
            } else {
                // 如果超过 k 个数，则需要在左边找第 k 个数
                r = ll - 1;
            }

        }
        // 最后只剩一个数，就是满足要求的
        nums[l]
    }
}
