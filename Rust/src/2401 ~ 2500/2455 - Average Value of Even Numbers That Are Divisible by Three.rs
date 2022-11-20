// 链接：https://leetcode.com/problems/average-value-of-even-numbers-that-are-divisible-by-three/
// 题意：给定一个正整数数组 nums ，返回能被 3 整除的偶数的平均数（向下取整）。


// 数据限制：
//  1 <= nums.length <= 1000
//  1 <= nums[i] <= 1000


// 输入： nums = [1,3,6,10,12,15]
// 输出： 9
// 解释： 6 和 12 是能被 3 整除的偶数，
//       平均数 = (6 + 12) / 2 = 9

// 输入： nums = [1,2,4,7,10]
// 输出： 0
// 解释： 没有能被 3 整除的偶数，平均数为 0


// 思路： 模拟
//
//      能被 3 整除的偶数就是能被 6 整除的数，
//      所以统计所有能被 6 整除的数的和 sum ，以及这些数的个数 cnt 。
//
//      最后如果 cnt 为 0 ，则没有满足题意的数，直接返回 0 ；
//      否则，返回 sum / cnt 即可。 
//
//
//      时间复杂度：O(n)
//          1. 需要遍历 nums 中全部 O(n) 个数
//      空间复杂度：O(1)
//          1. 只需要维护常数个额外变量即可


impl Solution {
    pub fn average_value(nums: Vec<i32>) -> i32 {
        // 能被 6 整除的数的和
        let mut sum = 0;
        // 能被 6 整除的数的个数
        let mut cnt = 0;
        for num in nums {
            // 如果 num 是能被 6 整除的数，则进行统计
            if num % 6 == 0 {
                sum += num;
                cnt += 1;
            }
        }

        if cnt == 0 {
            // 没有满足题意的数，直接返回 0
            0
        } else {
            // 否则，返回这些数的平均数
            sum / cnt
        }
    }
}
