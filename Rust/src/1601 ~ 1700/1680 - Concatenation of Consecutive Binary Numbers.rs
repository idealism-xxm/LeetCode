// 链接：https://leetcode.com/problems/concatenation-of-consecutive-binary-numbers/
// 题意：给定一个整数 n ，将 [1, n] 内的所有数按照二进制形式拼接在一起，
//      求对应的数的十进制值？
//      结果模 10 ^ 9 + 7 。


// 数据限制：
//  1 <= n <= 10 ^ 5


// 输入： n = 1
// 输出： 1
// 解释： 将 1 的二进制拼接在一起得 "1" ，对应的十进制值为 1 。

// 输入： n = 3
// 输出： 27
// 解释： 将 1, 2, 3 的二进制拼接在一起得 "11011" ，对应的十进制值为 27 。

// 输入： n = 12
// 输出： 505379714
// 解释： 将 [1, 12] 的二进制拼接在一起得 "1101110010111011110001001101010111100" ，
//       对应的十进制值为 118505380540 ，模 10 ^ 9 + 7 得 505379714 。


// 思路： 位运算
//
//      设 dp[i] 为 [1, i] 的二进制拼接在一起后对应的十进制值，
//      则初始有 dp[1] = 1 。
//
//      那么我们只要找到 dp[i] 与 dp[i - 1] 的状态转移关系，就能递推出 dp[n] 。
//
//      设 binary_len(num) 表示 num 的二进制位的数量。
//
//      可以发现 dp[i] 的二进制位就是在 dp[i - 1] 的二进制位后面拼上了 i 的二进制位，
//      那么二进制表示下就是 dp[i - 1] 左移 binary_len(i) 后，再按位或上 i ，
//      即 dp[i] = (dp[i - 1] << binary_len(i)) | i 。
//
//      【注意】 binary_len(i) 最大可取到 17 ，
//      那么 32 位整型直接左移 binary_len(i) 位可能溢出，
//      所以递推时需要采用 64 位整型处理。
//
//
//      根据目前这个状态转移方程，我们就可以在 O(nlogn) 内求出 dp[n] 。
//      （需要处理全部 O(n) 个数，每次处理时都需要遍历全部 O(logn) 个二进制位）
//
//      不过还有优化的空间，可以注意到 binary_len(i) 在大多数情况下都等于 binary_len(i - 1) 。
//
//      只有在 i 是 2 的幂次方时（此时 (i - 1) + 1 的最高位会产生进位），
//      才会等于 binary_len(i - 1) 。
//
//      所以我们只要能在 O(1) 内判断 i 是否为 2 的幂次方，就能在 O(1) 内递推出 binary_len(i) 。
//
//      有两种常见的方法判断一个数是否为 2 的幂次方：
//          1. 利用位运算：只要 i & (i - 1) == 0 ，那么 i 就是 2 的幂次方。
//
//              这个方法常用来判断一个数是否为 2 的幂次方，
//              原理见 LeetCode 231 和 LeetCode 342 这两题。
//          2. 利用模运算：本题中 2 的幂次方的最大值为 2 ^ 17 = 131072 >= 10 ^ 5 ，
//              所以只要 131072 % i == 0 ，那么 i 就是 2 的幂次方。
//
//              这个方法常用来判断一个数是否为素数的幂次方，
//              原理见 LeetCode 326 这题。
//
//
//      时间复杂度：O(n)
//          1. 需要处理全部 O(n) 个数
//      空间复杂度：O(1)
//          1. 只需要维护常数个额外变量即可


const MOD: i64 = 1_000_000_007;


impl Solution {
    pub fn concatenated_binary(n: i32) -> i32 {
        // ans 维护最终结果
        let mut ans = 0i64;
        // binary_len 表示当前二进制位的数量
        let mut binary_len = 0;
        // 遍历 [1, n] 内的数字
        for i in 1..=n as i64 {
            // 当 i 是 2 的幂次方时，最高位恰好发生了进位，二进制位数量 +1
            if i & (i - 1) == 0 {
                binary_len += 1;
            }
            // 将 ans 左移 binary_len 位后，再按位或上 i
            ans = ((ans << binary_len) | i) % MOD;
        }

        ans as i32
    }
}
