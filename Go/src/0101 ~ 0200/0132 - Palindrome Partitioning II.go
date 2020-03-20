// 链接：https://leetcode.com/problems/palindrome-partitioning-ii/
// 题意：给定一个字符串，求将其分成一些回文子串的最小分割次数 ？

// 输入： "aab"
// 输出： 1
// 解释： 分割为 ["aa", "b"] 下，需要分割一次

// 思路： DP
//
//		先在 O(n ^ 2) 内找出所有的回文串
//		再进行 DP
//		设 dp[i] 表示 s[:i] 分割成回文子串的最小分割次数
//		初始化： dp[i] = i - 1 （ 0 <= i <= n ，有 n 个字符，最多就需要分割 n - 1 次）
//				（ dp[0] 初始化为 -1 ，因为从左边没有字符串，所以使用它转移时会多一次，提前减去）
//		状态转移：
//			对于已经确定的 dp[i] ，找出所有以 i 开始的回文串的结束下标 + 1 （设为 j）
//			则有： dp[j] = min(dp[j], dp[i] + 1)
//
//		为了降低常数部分的复杂度和空间复杂度，我们可以在找回文字串的时候就进行 DP ，
//		这空间复杂度降低为 O(n) ，时间复杂度也少了部分常数
//
//		时间复杂度： O(n ^ 2)
//		空间复杂度： O(n)

func minCut(s string) int {
	length := len(s)
	// 初始化 dp
	dp := make([]int, length + 1)
	for i := 0; i <= length; i++ {
		dp[i] = i - 1
	}
	// 状态转移 （dp[0 ~ i] 都已计算完毕）
	for i := 0; i < length; i++ {
		// 找以 s[i] 为中心的回文串（长度为奇数）
		for l, r := i, i; l >= 0 && r < length && s[l] == s[r]; l, r = l - 1, r + 1 {
			// 可以更新将当前回文串的 结束下标 + 1 的 dp 值
			dp[r + 1] = min(dp[r + 1], dp[l] + 1)
		}
		// 找以 s[i - 1:i + 1] 为中心的回文串（长度为偶数）
		for l, r := i - 1, i; l >= 0 && r < length && s[l] == s[r]; l, r = l - 1, r + 1 {
			// 可以更新将当前回文串的 结束下标 + 1 的 dp 值
			dp[r + 1] = min(dp[r + 1], dp[l] + 1)
		}
	}

	return dp[length]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
