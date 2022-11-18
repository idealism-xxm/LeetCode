// 链接：https://leetcode.com/problems/ugly-number/
// 题意：给定一个数 n ，判断其质因数是否只有 2, 3 和 5 ？


// 数据限制：
//  -(2 ^ 31) <= n <= 2 ^ 31 - 1


// 输入： n = 6
// 输出： true
// 解释： 6 = 2 * 3

// 输入： n = 1
// 输出： false
// 解释： 1 不含任何质因数。

// 输入： n = 1
// 输出： false
// 解释： 14 = 2 * 7 ，含有质因数 7 。


// 思路： 模拟
//
//      如果 n 是非正数，则必定不满足题意，直接返回 false 。
//
//      如果 n 是正数，则枚举 [2, 3, 5] 三个全部质因数 factor ，
//      不断执行 n /= factor ，直至 n 不含质因数 factor 。
//
//      最后 n 为 1 时，则说明不含其他质因数，返回 true ；
//      否则， n 还含有其他质因数，返回 false 。
//
//
//      时间复杂度： O(logn)
//          1. 需要执行最多 O(logn) 次除法
//      空间复杂度： O(1)
//          1. 只需要使用常数个额外变量即可


impl Solution {
    pub fn is_ugly(mut n: i32) -> bool {
        // 非正数不满足题意，直接返回 false
        if n <= 0 {
            return false;
        }
        
        // 遍历所有质因数
        for factor in [2, 3, 5] {
            // 如果 n 还含有质因数 factor ，则从中移除
            while n % factor == 0 {
                n /= factor;
            }
        }
        
        // 最后当且仅当 n 为 1 时，说明不含其他质因数，满足题意
        n == 1
    }
}
