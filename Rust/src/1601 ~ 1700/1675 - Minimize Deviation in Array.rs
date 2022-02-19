// 链接：https://leetcode.com/problems/minimize-deviation-in-array/
// 题意：给定一个长度为 n 的正整数数组 nums ，
//      现在可以对其中的每个数 num 执行任意次以下操作：
//          1. 如果 num 是偶数，则可以对其 除以 2
//          2. 如果 num 是奇数，则可以对其 乘上 2
//  	求最后能形成的数组中 max(nums) - min(nums) 的最小值？


// 数据限制：
//  n == nums.length
//  2 <= n <= 10 ^ 5
//  1 <= nums[i] <= 10 ^ 9


// 输入：nums = [1,2,3,4]
// 输出：1
// 解释：将数组转换成 [ 1 ,2,3,(2)] ，
//      然后再转换成 [(2),2,3, 2 ] ，
//      偏移量是 3 - 2 = 1

// 输入：nums = [4,1,5,20,3]
// 输出：3
// 解释：将数组转换成 [4, 2 ,5,(5),3] ，
//      然后再转换成 [4,(2),5, 5 ,3] ，
//      偏移量是 5 - 2 = 3


// 思路：贪心 + 优先队列（堆）
//
//		每个数可以执行任意次操作，所以我们可以先把所有奇数乘上 2 。
//
//      那么对奇数乘 2 的操作就没了，只剩下对每个偶数进行任意次除以 2 的操作。
//
//      而除法每次只能减半一个数的值，所以肯定贪心选择 nums 中最大的偶数执行这个操作，
//      这样才有可能让新数组中 max(nums) - min(num) 的值变小。
//
//      我们可以维护一个最大堆 q 和数组中的最小值 min_num ，
//      并维护结果 ans ，初始化为 i32::MAX ，
//      那么每次都从 q 取出最大值 max_num ，
//      更新答案 ans = min(ans, max_num - min_num) 。
//
//      然后根据 max_num 的奇偶继续处理：
//          1. max_num 是偶数：则 max_num / 2 可能使得答案更优，可以继续循环。
//              更新 min_num = min(min_num, max_num / 2) ，
//              并将 max_num / 2 放入到最大堆 q 中。
//          2. max_num 是奇数：则已无法优化可能的答案，跳出循环
//
//
//      设 m 为 max(nums) ，最大可以取到 10 ^ 9
//
//      时间复杂度：O(n * logn * logm) 
//          1. 入堆/出堆的时间复杂度是 O(logn)
//          2. 最大数字可能最后需要入堆/出堆 O(logm) 次
//          3. 所有数字可能都是最大数字，那么这 O(n) 个数字都要执行前面的操作
//      空间复杂度：O(n)
//          1. 最大堆最多只会存 O(n) 个数

use std::collections::BinaryHeap;

impl Solution {
    pub fn minimum_deviation(nums: Vec<i32>) -> i32 {
        // 维护一个最大堆，最大长度为 nums.len()
        let mut q = BinaryHeap::with_capacity(nums.len());
        // 维护最大堆中的最小数字 min_num ，初始化为 i32::MAX
        let mut min_num = i32::MAX;
        // 遍历所有 nums ，将所有数处理后放入最大堆中
        for mut num in nums {
            // 如果 num 是奇数，就让其加倍。
            // 这样就消除了乘法操作，后续流程只用做除法即可
            if num & 1 == 1 {
                num <<= 1;
            }
            // 将 num 放入最大堆 q 中
            q.push(num);
            // 获取 nums 中所有数的最小值
            min_num = min_num.min(num);
        }

        // 维护答案 max(nums) - min(nums) ，初始化为 i32::MAX
        let mut ans = i32::MAX;
        loop {
            // 将 nums 中最大数出堆
            let max_num = q.pop().unwrap();
            // 更新此时 nums 中的 max(nums) - min(nums) 的最小值
            ans = ans.min(max_num - min_num);

            // 如果 max_num 是奇数，则已无法优化可能的答案，跳出循环
            if max_num & 1 == 1 {
                break;
            }
            // 更新 nums 中的最小值
            min_num = min_num.min(max_num >> 1);
            // 如果 max_num 是偶数，则 max_num >> 1 入堆，
            // 后续可能继续更新答案
            q.push(max_num >> 1);
        }

        ans
    }
}