// 链接：https://leetcode.com/problems/interleaving-string/
// 题意：给定三个字符串 s1, s2, s3 ，判断 s3 是否能由 s1, s2 交错组成？

// 输入：s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
// 输出：true

// 输入：s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
// 输出：false

// 思路：DP
//		很容以就能想到可以用 DP 做
//
//		设 dp[i][j] 表示 s1[:i] 和 s2[:j] 能否交错组成 s3[:i + j]
//		初始化：dp[0][0] = true （空串必定符合题意）
//		状态转移：dp[i][j] 可由两种状态转移而来，具体看 s[i + j - 1] 对应哪个字符串的字符
//			1. s3[i + j - 1] 对应 s1[i - 1] ：
//				dp[i - 1][j] && s1[i - 1] == s3[i + j - 1]
//			2. s3[i + j - 1] 对应 s2[j - 1] ：
//				dp[i][j - 1] && s2[j - 1] == s3[i + j - 1]
//
//		时间复杂度： O(len(s1) * len(s2))
//		空间复杂度： O(len(s1) * len(s2)) 【当然可以优化为一维，因为每次更新一行，只会用到它左边和上边的数据，所以直接用一维，其他不需要改】

func isInterleave(s1 string, s2 string, s3 string) bool {
	s1Len, s2Len, s3Len := len(s1), len(s2), len(s3)
	if s1Len + s2Len != s3Len {
		return false
	}

	dp := make([][]bool, s1Len + 1)
	for i := 0; i <= s1Len; i++ {
		dp[i] = make([]bool, s2Len + 1)
	}

	dp[0][0] = true  // 空串必定符合题意
	for i := 0; i <= s1Len; i++ {
		for j := 0; j <= s2Len; j++ {
			// 初始状态不处理
			if i == 0 && j == 0 {
				continue
			}
			// 下一个字符由 s1 对应
			if i != 0 && dp[i - 1][j] && s1[i - 1] == s3[i + j - 1] {
				dp[i][j] = true
			}
			// 下一个字符由 s2 对应
			if j != 0 && dp[i][j - 1] && s2[j - 1] == s3[i + j - 1] {
				dp[i][j] = true
			}
		}
	}
	return dp[s1Len][s2Len]
}
