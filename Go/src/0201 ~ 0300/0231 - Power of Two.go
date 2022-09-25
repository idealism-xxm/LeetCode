// 链接：https://leetcode.com/problems/power-of-two/
// 题意：判断一个整数是不是 2 的幂次方？
//
//      进阶：不使用循环或递归。


// 数据限制：
//  -(2 ^ 31) <= n <= 2 ^ 31 - 1


// 输入： n = 1
// 输出： true
// 解释： 2 ^ 0 = 1

// 输入： n = 16
// 输出： true
// 解释： 2 ^ 4 = 16

// 输入： n = 218
// 输出： false
// 解释： 2 ^ 7 < n < 2 ^ 8


// 思路： 位运算
//
//      如果 n 是 2 的幂次方，则 n 一定只有一个 1 （假设是从右边起第 i 位），
//      那么 n - 1 的右边 i - 1 位都是 1 ，
//      因此 n & (n - 1) == 0
//
//      如果 n 不是 2 的幂次方，则 n 一定至少有两个 1 （假设从右往左的第一个 1 在第 i 位），
//      那么 n - 1 与 n 在大于第 i 位处，必定存在相同的位为 1 ，那么 n & (n - 1) != 0
//
//
//      时间复杂度： O(1)
//          1. 只需要使用常数次位运算和布尔运算
//      空间复杂度： O(1)
//          2. 只需要使用常数个额外变量即可


func isPowerOfTwo(n int) bool {
    // 非正数一定不是 2 的幂次方
    if n <= 0 {
        return false
    }

    return (n & (n - 1)) == 0
}
