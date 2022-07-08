// 链接：https://leetcode.com/problems/interleaving-string/
// 题意：给定三个字符串 s1, s2, s3 ，求 s3 是否能由 s1 和 s2 交错组成？
//
//      进阶：使用空间复杂度为 O(|s2|) 的算法


// 数据限制：
//  0 <= s1.length, s2.length <= 100
//  0 <= s3.length <= 200
//  s1, s2, and s3 仅由英文小写字母组成


// 输入： s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
// 输出： true
// 解释： s1: aa____bc_c
//       s2: __dbbc__a_
//       s3: aadbbcbcac

// 输入： s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
// 输出： false

// 输入： s1 = "", s2 = "", s3 = ""
// 输出： true


// 思路： DP
//
//      很容易就能想到可以用 DP 做。
//
//      设 dp[i][j] 表示 s1[:i] 和 s2[:j] 能否交错组成 s3[:i + j]
//      初始化：dp[0][0] = true （空串必定符合题意）
//      状态转移：dp[i][j] 可由两种状态转移而来，具体看 s3[i + j - 1] 对应哪个字符串的字符
//          1. s3[i + j - 1] 对应 s1[i - 1] ：则状态转移自
//              dp[i - 1][j] && s1[i - 1] == s3[i + j - 1]
//          2. s3[i + j - 1] 对应 s2[j - 1] ：则状态转移自
//              dp[i][j - 1] && s2[j - 1] == s3[i + j - 1]
//
//      DP 常见的三种优化方式见 LeetCode 583 这题的思路，
//      本题化用一维数组 + 倒序转移这种方法进行优化，
//      采用一维数组 + 顺序转移的方式处理，即可将空间复杂度从 O(mn) 优化为 O(n) 。 
//
//      因为状态 dp[i][j] 仅由 dp[i - 1][j] 和 dp[i][j - 1] 两个状态转移而来，
//      所以必须使用顺序转移，从前往后计算，以便使用当前行的状态进行转移。
//
//
//      时间复杂度：O(mn)
//          1. 需要遍历计算全部 O(mn) 个状态
//      空间复杂度：O(n)
//          1. 需要用 dp 维护全部 O(n) 个状态


impl Solution {
    pub fn is_interleave(s1: String, s2: String, s3: String) -> bool {
        let (s1, s2, s3) = (s1.as_bytes(), s2.as_bytes(), s3.as_bytes());
        // 如果长度不相等，则 s3 必定不可能由 s1 和 s2 交错组成
        let (m, n, s3_len) = (s1.len(), s2.len(), s3.len());
        if m + n != s3_len {
            return false
        }

        // dp[j] 表示 s1[..i] 与 s2[..j] 是否能交错组成 s3[..i + j]
        let mut dp = vec![false; n + 1];
        for i in 0..=m {
            for j in 0..=n {
                // 初始化只有 s1[..0] 和 s2[..0] 能交错组成 s3[..0] ，即空串必定符合题意
                if i == 0 && j == 0 {
                    dp[0] = true;
                    continue;
                }

                // 标记 s3 的最后一个字符是否能来源于 s1/s2
                let (mut is_last_ch_from_s1, mut is_last_ch_from_s2) = (false, false);
                // 如果 dp[i - 1][j] 为 true ，且 s1 的最后一个字符与 s3 最后一个字符相等，
                // 则状态 dp[i][j] 可由 dp[i - 1][j] 转移而来
                if i > 0 && dp[j] && s1[i - 1] == s3[i + j - 1] {
                    is_last_ch_from_s1 = true;
                }
                // 如果 dp[i][j - 1] 为 true ，且 s2 的最后一个字符与 s3 最后一个字符相等，
                // 则状态 dp[i][j] 可由 dp[i][j - 1] 转移而来
                if j > 0 && dp[j - 1] && s2[j - 1] == s3[i + j - 1] {
                    is_last_ch_from_s2 = true;
                }
                // 只要能从 dp[i - 1][j] 和 dp[i][j - 1] 任意一个状态转移而来，
                // dp[i][j] 就为 true
                dp[j] = is_last_ch_from_s1 || is_last_ch_from_s2;
            }
        }

        dp[n]
    }
}
