// 链接：https://leetcode.com/problems/power-of-three/
// 题意：给定一个数 n ，判断其是否是 3 的幂次方？
//
//      进阶：不使用循环或者递归进行判断。


// 数据限制：
//  -(2 ^ 31) <= n <= 2 ^ 31 - 1


// 输入： n = 27
// 输出： true
// 解释： 27 = 3 ^ 3

// 输入： n = 0
// 输出： 0 < 3 ^ 0

// 输入： n = 9
// 输出： 9 = 3 ^ 2


// 思路： 数学
//
//      由于 n 的上限是 MAX = 2 ^ 31 - 1 ，
//      那么 3 的幂次方的最大值为 3 ^ floor(log3(MAX)) 
//                           ≈ 3 ^ floor(19.56)
//                           = 3 ^ 19
//                           = 1162261467
//
//      所以只要 n 是 3 ^ 0, 3 ^ 1, ..., 3 ^ 19 之一，就必定是 3 的幂次方，
//      其他情况都不是。
//
//      又因 3 是质数，所以 3 ^ x 的所有因数都是 3 的幂次方，
//      即只有 3 ^ 0, 3 ^ 1, ..., 3 ^ 19 能整除 3 ^ 19 ，
//      所以只要 n 能整除 3 ^ 19 ，那么 n 就必定是 3 的幂次方。
//
//
//      这个方法可以推广到判断 n 是否为一个质数 prime 的幂次方。
//
//
//      时间复杂度：O(1)
//          1. 只需要使用常数次布尔运算和模运算即可
//      空间复杂度：O(1)
//          1. 只需要使用常数个额外变量即可


impl Solution {
    pub fn is_power_of_three(n: i32) -> bool {
        // 只要 n 是 3 ^ 0, 3 ^ 1, ..., 3 ^ 19 之一，就必定是 3 的幂次方。
        //
        // 而 3 是质数，所以 3 ^ x 的所有因数都是 3 的幂次方，
        // 即只有 3 ^ 0, 3 ^ 1, ..., 3 ^ 19 能整除 3 ^ 19 ，
        // 所以只要 n 能整除 3 ^ 19 ，那么 n 就必定是 3 的幂次方。
        n > 0 && 1162261467 % n == 0
    }
}
