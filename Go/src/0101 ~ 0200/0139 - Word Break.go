// 链接：https://leetcode.com/problems/word-break/
// 题意：给一个非空字符串 s 和一个单词列表 wordDict ，
//		判断 s 能否断成由 wordDict 中的单词构成的序列（每个单词可用任意次）？
//
// 输入： s = "leetcode", wordDict = ["leet", "code"]
// 输出： true
// 解释： "leet code"

// 输入： s = "applepenapple", wordDict = ["apple", "pen"]
// 输出： true
// 解释： "apple pen apple"

// 输入： s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
// 输出： false

// 思路： DP
//
//		感觉和 0131 一样，只不过把需要分割的方式从回文串判定变成给定的
//		直接递归，然后完美 TLE
//		然后恍然大悟，其实和 0132 可以直接用 DP
//		设 dp[i] 表示 s[:i] 能否分割成 wordDict 中的单词
//		初始化 dp[0] = true
//		若 dp[i] = true ，则 s[i:j] 在 wordDict 中，那么 dp[j] = true
//
//		时间复杂度： O(n ^ 2)
//		空间复杂度： O(n)

func wordBreak(s string, wordDict []string) bool {
	// 先转换成 map
	wordMap := make(map[string]bool)
	for _, word := range wordDict {
		wordMap[word] = true
	}

	// 初始化 dp
	length := len(s)
	dp := make([]bool, length + 1)
	dp[0] = true

	for i := 0; i < length; i++ {
		if dp[i] {
			for j := i + 1; j <= length; j++ {
				if wordMap[s[i:j]] {
					dp[j] = true
				}
			}
		}
	}

	return dp[length]
}
