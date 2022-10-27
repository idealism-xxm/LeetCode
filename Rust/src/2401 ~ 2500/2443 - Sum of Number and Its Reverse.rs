// 链接：https://leetcode.com/problems/sum-of-number-and-its-reverse/
// 题意：给定一个非负整数 num ，判断是否存在一个数 i 满足 i + reverse(i) 为 num ？
//      其中， reverse(x) 表示对 x 按十进制位翻转。


// 数据限制：
//  0 <= num <= 10 ^ 5


// 输入： num = 443
// 输出： true
// 解释： 172 + 271 = 443

// 输入： num = 63
// 输出： 不存在任何 i 使得 i + reverse(i) 为 63

// 输入： num = 181
// 输出： 140 + 041 = 181


// 思路： 模拟
//
//      枚举 [0, num] 内的每个数，如果 i + reverse(i) == num 成立，则直接返回 true 。
//
//      枚举完还未返回时，就表明无满足题意的数，直接返回 false 。
//
//
//      时间复杂度：O(nlogn)
//          1. 需要遍历 [0, n] 内全部 O(n) 个数，每次都要遍历全部 O(logn) 个十进制位
//      空间复杂度：O(1)
//          1. 只需要维护常数个额外变量


impl Solution {
    pub fn sum_of_number_and_reverse(num: i32) -> bool {
        // 枚举 [0, num] 内的每个数
        for i in 0..=num {
            // 如果当前 i 满足题意，则直接返回 true
            if i + Self::reverse(i) == num {
                return true;
            }
        }

        // 此时表明无满足题意的数，直接返回 false
        false
    }

    fn reverse(mut num: i32) -> i32 {
        // 对 num 按十进制位翻转
        let mut res = 0;
        while num > 0 {
            res = res * 10 + num % 10;
            num /= 10;
        }
        
        res
    }
}
