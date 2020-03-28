// 链接：https://leetcode.com/problems/word-break-ii/
// 题意：给一个非空字符串 s 和一个单词列表 wordDict ，
//		找到将 s 断成由 wordDict 中的单词构成的序列的所有情况（每个单词可用任意次）？

// 输入：
// s = "catsanddog"
// wordDict = ["cat", "cats", "and", "sand", "dog"]
// 输出：
// [
//   "cats and dog",
//   "cat sand dog"
// ]

// 输入：
// s = "pineapplepenapple"
// wordDict = ["apple", "pen", "applepen", "pine", "pineapple"]
// 输出：
// [
//   "pine apple pen apple",
//   "pineapple pen apple",
//   "pine applepen apple"
// ]

// 输入：
// s = "catsandog"
// wordDict = ["cats", "dog", "sand", "and", "cat"]
// 输出： []

// 思路： 递归
//
//		可以说和 0131 一样，需要先处理出所有每个位置所有可能的单词，然后一样 dfs 即可
//		又完美 TLE ，看来还是思考不够深
//		这题其实和 0131 还是有点区别的，这题可能无法合法断开，
//		所以就有一个不断重复且无法合法断开的样例
//		看到样例后就能发现，递归时会不停对同一种情况处理，造成复杂度指数级上升，
//		所以需要进行剪枝标记， splittable[i] 表示 s[i:] 能否断成由 wordDict 中的单词构成的序列
//		当某一次递归结束后，没有收集到任何结果，则需要 splittable[i] 标记为 false
//
//		当然这种其实就是只记忆化了非法的结果，其实也可以记忆化全部的结果，
//		这样后续遇到合法的结果可以提前结束递归，保证合法情况下也能只遍历一次

func wordBreak(s string, wordDict []string) []string {
	// 先转换成 map
	wordMap := make(map[string]bool)
	for _, word := range wordDict {
		wordMap[word] = true
	}

	// availableWords[i] 收集所有以 i 为开始字符的在 wordDict 中的单词的 结束字符 + 1 的下标
	// 即： availableWords[i][j] 表示 s[i:availableWords[i][j]] 在 wordDict 中
	length := len(s)
	availableWords := make([][]int, length)
	// splittable[i] 表示 s[i:] 能否断成由 wordDict 中的单词构成的序列，用于剪枝
	splittable := make([]bool, length)
	for i := 0; i < length; i++ {
		// 初始都认为可以合法断开
		splittable[i] = true
		availableWords[i] = []int{}
		for j := i + 1; j <= length; j++ {
			if wordMap[s[i:j]] {
				availableWords[i] = append(availableWords[i], j)
			}
		}
	}

	return dfs(availableWords, splittable, s, 0, "")
}

// 划分 s[start:] 成 wordDict 中出现的单词，并且每一种结果列表的前缀都是 list[:count]
func dfs(availableWords [][]int, splittable []bool, s string, start int, sentence string) []string {
	// 如果所有字符以用完，则直接返回当前这种结果
	if start == len(s) {
		return []string{sentence[1:]}
	}
	// 剪枝，当 s[start] 已经处理过，且发现不能合法断开时，直接返回
	if !splittable[start] {
		return nil
	}

	var result []string
	for i, length := 0, len(availableWords[start]); i < length; i++ {
		end := availableWords[start][i]
		// 枚举当前使用的单词为 s[start:end] ，递归收集后续的结果
		nextSentence := sentence + " " + s[start:end]
		result = append(result, dfs(availableWords, splittable, s, end, nextSentence)...)
	}
	// 标记无法合法断开
	if len(result) == 0 {
		splittable[start] = false
	}
	return result
}
