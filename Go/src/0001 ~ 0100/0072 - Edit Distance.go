// 链接：https://leetcode.com/problems/edit-distance/
// 题意：给定两个字符串，每次可以对第一个字符串做三个操作：
//		1. 插入一个字符
//		2. 删除一个字符
//      3. 替换一个字符
//		求能将第一个字符串变成第二个字符串的最小操作数？

// 输入：word1 = "horse", word2 = "ros"
// 输出：3

// 输入：word1 = "intention", word2 = "execution"
// 输出：5

// 思路：DP
//		dp[i][j] 表示将 word1[0:i] 转换成 word2[0:j] 所需的最小操作数
//		初始化状态：dp[0][j] = j, dp[i][0] = i
//		状态转移方程：
//		1. 删除一个字符：dp[i - 1][j] + 1
//		2. 插入一个字符：dp[i][j - 1] + 1
//		3. 替换一个字符：dp[i - 1][j - 1] + extra （最后一个字符相同时为 0， 否则为 1）
//		时间复杂度： O(len(word1) * len(word2))
//		空间复杂度： O(len(word1) * len(word2)) 【可以用滚动数组优化为 O(len(word2))】

func minDistance(word1 string, word2 string) int {
	len1, len2 := len(word1), len(word2)
	dp := make([][]int, len1 + 1)
	for i := 0; i <= len1; i++ {
		dp[i] = make([]int, len2 + 1)
		dp[i][0] = i
	}
	for j := 0; j <= len2; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			// 当前最后一步从 删除字符和插入字符 中选取一个最小的操作数
			dp[i][j] = min(dp[i - 1][j], dp[i][j - 1]) + 1
			// 当前最后一步再从 替换字符中 选取一个最小的操作数
			extra := 0  // 默认不需要替换
			if word1[i - 1] != word2[j - 1] {  // 最后一个字符不同，则需要替换
				extra = 1
			}
			dp[i][j] = min(dp[i][j], dp[i - 1][j - 1] + extra)
		}
	}
	return dp[len1][len2]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}