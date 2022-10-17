// 链接：https://leetcode.com/problems/bitwise-xor-of-all-pairings/
// 题意：给定两个非负整数数组 nums1 和 nums2 ，
//      nums3 含有 len(nums1) * len(nums2) 个数字，
//      分别为 nums1 和 nums2 所有数对的异或和。
//      求 nums3 所有数字的异或和？


// 数据限制：
//  1 <= nums1.length, nums2.length <= 10 ^ 5
//  0 <= nums1[i], nums2[j] <= 10 ^ 9


// 输入： nums1 = [2,1,3], nums2 = [10,2,5,0]
// 输出： 13
// 解释： 一个合法的 nums3 是 [8,0,7,2,11,3,4,1,9,1,6,3] ，
//       所有数的异或和为 13

// 输入： nums1 = [1,2], nums2 = [3,4]
// 输出： 0
// 解释： 一个合法的 nums3 是 [2,5,1,6] ，
//       所有数的异或和为 0


// 思路： 位运算
//
//      关于异或的题目，都要牢记 a ^ a = 0 这一点。
//
//      本题也是基于此进行优化，只需要遍历 nums1 和 nums2 各一次即可求得结果。
//
//      设 a_i = nums1[i], b_i = nums2[i] ， nums1 的长度为 m, nums2 的长度为 n 。
//
//      nums3 里面必定有 n 个数与 a_i 有关，
//      即： a_i ^ b_0, ..., a_i ^ b_(n - 1) ，
//      那么当 n 为偶数时，所有 a_i 的异或和为 0 ，否则为 a_i 。
//
//      同理可得，当 m 为偶数时，所有 b_i 的异或和为 0 ，否则为 b_i 。
//
//      且最终结果中，同一个数组里的数的异或和要么同时为 0 ，要么同时为自身。
//
//      设 xor(nums) 表示数组 nums 中所有数的异或和。
//
//      那么当 n 为奇数时，结果要异或上 xor(nums1) ；
//      当 m 为奇数时，结果要异或上 xor(nums2) 。
//
//
//      时间复杂度：O(m + n)
//          1. 需要遍历 nums1 中全部 O(m) 个数字一次
//          2. 需要遍历 nums2 中全部 O(n) 个数字一次
//      空间复杂度：O(1)
//          1. 只需要维护常数个额外变量


func xorAllNums(nums1 []int, nums2 []int) int {
    // ans 维护 nums3 所有数的异或和
    ans := 0
    // 如果 nums2 含有奇数个数，则 nums1 中每个数对 ans 都有一次贡献
    if len(nums2) & 1 == 1 {
        for _, num := range nums1 {
            ans ^= num
        }
    }
    // 如果 nums1 含有奇数个数，则 nums2 中每个数对 ans 都有一次贡献
    if len(nums1) & 1 == 1 {
        for _, num := range nums2 {
            ans ^= num
        }
    }

    return ans
}
