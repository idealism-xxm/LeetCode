// 链接：https://leetcode.com/problems/find-the-difference-of-two-arrays/
// 题意：给定两个整数数组 nums1 和 nums2 ，返回一个长度为 2 的数组。
//      其中 answer[0] 是在 nums1 中且不在 nums2 中的所有不同数的列表，
//          answer[1] 是在 nums2 中且不在 nums1 中的所有不同数的列表。
//      注意：列表中的数字可以是任意顺序的。


// 数据限制：
//  1 <= nums1.length, nums2.length <= 1000
//  -1000 <= nums1[i], nums2[i] <= 1000


// 输入： nums1 = [1,2,3], nums2 = [2,4,6]
// 输出： [[1,3],[4,6]]
// 解释： nums1[0] 和 nums1[2] 不在 nums2 中，所以 answer[0] = [1,3] 。
//       nums2[1] 和 nums2[2] 不在 nums1 中，所以 answer[1] = [4,6] 。

// 输入： nums1 = [1,2,3,3], nums2 = [1,1,2,2]
// 输出： [[3],[]]
// 解释： nums1[2] 和 nums1[3] 不在 nums2 中，而 nums1[2] == nums2[3] ，
//           所以 answer[0] = [3] 。
//       nums2 中没有数 不在 nums1 中，所以 answer[1] = [] 。


// 思路： 模拟
//
//      先获取 nums1 和 nums2 的数字集合，保证每个数字只出现一次。
//
//      然后分别枚举两个数字集合中的数字，将不在另一个数字集合中的数字收集成列表，
//      最后返回即可。
//
//
//      时间复杂度：O(n + m)
//          1. 需要遍历 nums1 中全部 O(n) 个数字
//          2. 需要遍历 nums2 中全部 O(m) 个数字
//      空间复杂度：O(n + m)
//          1. 需要维护 nums1 中全部不同的数字，最差情况下有 O(n) 个
//          2. 需要维护 nums2 中全部不同的数字，最差情况下有 O(m) 个


use std::collections::HashSet;
use std::iter::FromIterator;


impl Solution {
    pub fn find_difference(nums1: Vec<i32>, nums2: Vec<i32>) -> Vec<Vec<i32>> {
        // 获取 nums1 和 nums2 的数字集合
        let nums1: HashSet<i32> = HashSet::from_iter(nums1);
        let nums2: HashSet<i32> = nums2.into_iter().collect();
        // 枚举数字集合保证一个数字只出现一次，同时去除在另一个集合中的数字
        return vec![
            nums1.difference(&nums2).copied().collect(),
            nums2.difference(&nums1).cloned().collect(),
        ]
    }
}
