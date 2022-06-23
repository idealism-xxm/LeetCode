// 链接：https://leetcode.com/problems/merge-sorted-array/
// 题意：给定两个升序排序的数组 nums1 和 nums2 ，以及两个整数 m 和 n 。
//      nums1 的长度为 m + n ， nums2 的长度为 n 。
//      nums1 中只有前 m 个数是有效的，后 n 个数都是无效的，均为 0 。
//
//      现在需要将 nums2 合并到 nums1 中，使得 nums1 成为一个升序数组。
//
//      进阶：使用时间复杂度为 O(m + n) ，空间复杂度为 O(1) 的算法。


// 数据限制：
//  nums1.length == m + n
//  nums2.length == n
//  0 <= m, n <= 200
//  1 <= m + n <= 200
//  -(10 ^ 9) <= nums1[i], nums2[j] <= 10 ^ 9


// 输入： nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
// 输出： [1,2,2,3,5,6]
// 解释： 待合并的数组是 [1,2,3] 和 [2,5,6] ，
//       合并的结果为 [(1),(2),2,(3),5,6] ，其中带小括号的来自 nums1 。

// 输入： nums1 = [1], m = 1, nums2 = [], n = 0
// 输出： [1]
// 解释： 待合并的数组是 [1] 和 [] ，
//       合并的结果为 [(1)] ，其中带小括号的来自 nums1 。

// 输入： nums1 = [0], m = 0, nums2 = [1], n = 1
// 输出： [1]
// 解释： 待合并的数组是 [] 和 [1] ，
//       合并的结果为 [1] 。


// 思路： 双指针
//
//      因为两个数组都是升序的，所以可以用双指针来操作。
//
//      我们可以从后往前向 nums1 中填充当前最大的数字，
//      因为 nums1 中后面 n 个位置是空的，所以从后往前填充不会覆盖未使用的数，
//      也能避免使用过多额外空间。
//      
//      然后不断循环填充，直到 nums2 中的数字全部填充完毕，
//      每次将两个数组中还未使用的最大数字的较大值放到 nums1[k] 即可。
//
//      结束循环后， i 和 k 必定相等：
//          1. i == k == -1: 所有数字全部填充完成
//          2. i == k >= 0: nums1 中前 i + 1 个数字还未填充，但本身已有序，
//              继续填充位置也不会改变，所以无需处理
//
//
//      时间复杂度：O(m + n)
//          1. 需要遍历 nums1 中全部 O(m) 个数字
//          2. 需要遍历 nums2 中全部 O(n) 个数字
//      空间复杂度：O(1)
//          1. 只需要使用常数个额外变量


impl Solution {
    pub fn merge(nums1: &mut Vec<i32>, m: i32, nums2: &mut Vec<i32>, n: i32) {
        // i/j 分别表示 nums1/nums2 中还未使用的最大数的下标
        let (mut i, mut j) = (m - 1, n - 1);
        // k 表示 nums1 中下一个该填充的位置。
        let mut k = m + n - 1;
        // 如果 nums2 中还有数字，则继续向 nums1[k] 填充当前最大的数
        while j >= 0 {
            if i >= 0 && nums1[i as usize] > nums2[j as usize] {
                // nums1 和 num2 中都还有数字，且 nums1[i] > nums2[j] ，
                // 则将 nums1[i] 放到 k 处
                nums1[k as usize] = nums1[i as usize];
                i -= 1;
            } else {
                // 此时有两种情况，都需要将 nums2[j] 放到 k 处：
                //    1. 只有 nums2 中还有数字
                //    2. nums1 和 num2 中都还有数字，且 nums1[i] <= nums2[j]
                nums1[k as usize] = nums2[j as usize];
                j -= 1;
            }
            k -= 1;
        }
        // 此时 i 和 k 必定相等，后续即使 nums1 中还有数字，位置也不会再改变
    }
}