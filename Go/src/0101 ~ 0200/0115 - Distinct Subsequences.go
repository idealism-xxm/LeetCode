// 链接：https://leetcode.com/problems/distinct-subsequences/
// 题意：给定两个字符串 s 和 t ，求在 s 中删除字符后能形成 t 的方案数？

// 输入： S = "rabbbit", T = "rabbit"
// 输出： 3
// 解释：
// rabbbit
// ^^^^ ^^
// rabbbit
// ^^ ^^^^
// rabbbit
// ^^^ ^^^

// 输入： S = "babgbag", T = "bag"
// 输出： 5
// 解释：
// babgbag
// ^^ ^
// babgbag
// ^^    ^
// babgbag
// ^    ^^
// babgbag
//   ^  ^^
// babgbag
//     ^^^

// 思路1：DP
//
//		一看这种题就知道要用 DP 了，应该是题感
//
//		设 dp[i][j] 表示 s[:i] 删除非 s[i - 1] 字符后能形成 t[:j] 的方案数
//		初始化：dp[0][0] = 1 （不删除字符就相等）
//		状态转移：
//			1. 若 s[i - 1] != t[j - 1] ，则必定无法通过操作满足题意
//				即：dp[i][j] = 0
//			2. 若 s[i - 1] == t[j - 1] ，则等于 s[k] 删除非 s[k - 1] 字符后形成 t[:j - 1] 的方案数和
//				即：dp[i][j] = sum(dp[k][j - 1]) (k = j - 1, j, ..., i - 1)
//
//		时间复杂度： O(|s|^2 * |t|)
//		空间复杂度： O(|s| * |t|)

func numDistinct(s string, t string) int {
	// 最后添加一个相同的字母，可以直接计算出最后结果，不需要计算后再循环一遍
	s, t = s + "a", t + "a"
	sLen, tLen := len(s), len(t)
	dp := make([][]int, sLen + 1)
	for i := range dp {
		dp[i] = make([]int, tLen + 1)
	}
	// 初始化
	dp[0][0] = 1
	for i := 1; i <= sLen; i++ {
		for j := 1; j <= tLen && j <= i; j++ {
			// 最后一个字符相同，则可以转移
			if s[i - 1] == t[j - 1] {
				for k := j - 1; k < i; k++ {
					dp[i][j] += dp[k][j - 1]
				}
			}
		}
	}
	return dp[sLen][tLen]
}

// 思路2：DP
//
//		这种想法还是复杂了，添加了一个限制，导致时空复杂度都无法减小
//		看了题解区还有 O(|s|^2 * |t|) 的 DP ，才恍然大悟
//
//		设 dp[i][j] 表示 s[:i] 删除字符后能形成 t[:j] 的方案数
//		初始化：dp[0~|s|][0] = 1 （任何串都可以删除全部字符形成空串）
//		状态转移：
//			1. 若 s[i - 1] != t[j - 1] ，则：当前字符无存在感，无需关注，
//				即：dp[i][j] = dp[i - 1][j]
//			2. 若 s[i - 1] == t[j - 1] ，则：不要当前字符时和要当前字符时两种情况的方案数和
//				则：dp[i][j] = dp[i - 1][j] + dp[i - 1][j - 1]
//
//		时间复杂度： O(|s|^2 * |t|)
//		空间复杂度： O(|s| * |t|) 【当然可以优化为一维，因为每次更新一行，只会用到它左边和左上的数据，所以直接用一维，其他不需要改】

func numDistinct(s string, t string) int {
	sLen, tLen := len(s), len(t)
	dp := make([][]int, sLen + 1)
	// 初始化
	for i := range dp {
		dp[i] = make([]int, tLen + 1)
		dp[i][0] = 1
	}

	for i := 1; i <= sLen; i++ {
		for j := 1; j <= tLen && j <= i; j++ {
			// 最后一个字符相同，有两种情况：不要当前字符和要当前字符
			if s[i - 1] == t[j - 1] {
				dp[i][j] = dp[i - 1][j] + dp[i - 1][j - 1]
			} else {
				// 最后一个字符不同，则最后一个字符只能不要
				dp[i][j] = dp[i - 1][j]
			}
		}
	}
	return dp[sLen][tLen]
}
