// 链接：https://leetcode-cn.com/problems/power-of-two/
// 题意：判断一个整数是不是 2 的幂次方？

// 输入： 1
// 输出： true
// 解释： 2 ^ 0 = 1

// 输入： 16
// 输出： true
// 解释： 2 ^ 4 = 16

// 输入： 218
// 输出： false

// 思路： 位运算
//
//      如果 n 是 2 的幂次方，则 n 一定只有一个 1 （假设是从右边起第 i 位），
//      那么 n - 1 的右边 i - 1 位都是 1 ，
//      因此 n & (n - 1) 是 0
//
//      如果 n 不是 2 的幂次方，则 n 一定不只有一个 1 ，那么 n - 1 与 n 必定有相同的位位 1
//
//      时间复杂度： O(1)
//      空间复杂度： O(1)

impl Solution {
    pub fn is_power_of_two(n: i32) -> bool {
        // 非正数一定不是
        if n <= 0 {
            return false
        }

        (n & (n - 1)) == 0
    }
}
