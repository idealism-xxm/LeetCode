// 链接：https://leetcode.com/problems/wildcard-matching/
// 题意：给定一个字符串 s 和一个模式串 p ，
//      s 只含有小写字母 a-z ， p 含有小写字母 a-z 以及 ？和 * ，
//      其中： ? 匹配任意一个字符， * 匹配任意长度的字符串（包括空串）
//      判断 p 是否完全匹配 s ？

// 输入：s = "aa", p = "a"
// 输出：false

// 输入：s = "aa", p = "*"
// 输出：true

// 输入：s = "cb", p = "?a"
// 输出：false

// 输入：s = "adceb", p = "*a*b"
// 输出：true

// 输入：s = "acdcb", p = "a*c?b"
// 输出：false

// 思路：DP
//      dp[i][j] 表示 s[0:i] 是否和 p[0:j] 完全匹配
//      初始化:
//      1. dp[0][0] = true
//      2. 若 p 开头有连续 k 个 * ，则 dp[0][1~k] = true （* 可以匹配空串）
//
//      若：
//      1. p[j] == '?' : dp[i][j] = dp[i - 1][j - 1]
//          (1) s[0:i - 1] 和 p[0:j - 1] 匹配 （? 匹配一个字符）
//      2. p[j] == '*' : dp[i][j] = dp[i - 1][j - 1] || dp[i - 1][j] || dp[i][j - 1]
//          (1) dp[i - 1][j - 1] ： s[0:i - 1] 和 p[0:j - 1] 匹配 （* 只匹配一个字符）
//          (2) dp[i - 1][j] ： s[0:i - 1] 和 p[0:j] 匹配（* 匹配多个字符）
//          (3) dp[i][j - 1] ： s[0:i] 和 p[0:j - 1] 匹配（* 匹配空串）
//      3. 当前模式串的字符为字母，则 dp[i][j] = dp[i - 1][j - 1] && s[i - 1] == p[j - 1]
//          (1) s[0:i - 1] 和 p[0:j - 1] 匹配，且当前字符匹配 （sChar == pChar）
//      时间复杂度： O(len(s) * len(p)) ，空间复杂度： O(len(s) * len(p))

func isMatch(s string, p string) bool {
    // 初始化 dp
    sLen, pLen := len(s), len(p)
    dp := make([][]bool, sLen + 1)
    for i := 0; i <= sLen; i++ {
        dp[i] = make([]bool, pLen + 1)
    }
    dp[0][0] = true  // 空串和空模式串匹配
    for j := 1; j <= pLen && p[j - 1] == '*'; j++ {
        dp[0][j] = true  // 空串和开头连续的通配符匹配
    }
    for i := 1; i <= sLen; i++ {
        for j := 1; j <= pLen; j++ {
            pChar := p[j - 1]
            switch pChar {
            case '?':  // ? 匹配一个字符
                dp[i][j] = dp[i - 1][j - 1]  // 只要 s[0:i - 1] 和 p[0:j - 1] 匹配，则当前必定匹配
            case '*':  // * 匹配任意长度的字符串
                // 1. s[0:i - 1] 和 p[0:j - 1] 匹配 （* 只匹配一个字符）
                // 2. s[0:i - 1] 和 p[0:j] 匹配（* 匹配多个字符）
                // 3. s[0:i] 和 p[0:j - 1] 匹配（* 匹配空串）
                // 则当前必定匹配
                dp[i][j] = dp[i - 1][j - 1] || dp[i - 1][j] || dp[i][j - 1]
            default:  // 只匹配相同的字符
                sChar := s[i - 1]  // 因为 dp[i][j] 表示 s[0:i] 是否和 p[0:j] 完全匹配，所以要比较 s[i - 1] 和 p[j - 1]
                dp[i][j] = dp[i - 1][j - 1] && sChar == pChar  // s[0:i - 1] 和 p[0:j - 1] 匹配 且 字符相同，则当前必定匹配
            }
        }
    }
    return dp[sLen][pLen]
}
