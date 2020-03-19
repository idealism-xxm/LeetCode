// 链接：https://leetcode.com/problems/palindrome-partitioning/
// 题意：给定一个字符串，将其分成一些回文子串，求所有这样的分割结果 ？

// 输入： "aab"
// 输出：
// [
//   ["aa","b"],
//   ["a","a","b"]
// ]

// 思路： 遍历 + 递归
//
//		先在 O(n ^ 2) 内找出所有的回文串
//		再递归枚举每一次使用的回文串，组装所有可能的结果
//

func partition(s string) [][]string {
	length := len(s)
	// palindromes[i] 收集所有以 i 为开始字符的回文串的 结束字符 + 1 的下标
	// 即： palindromes[i][j] 表示 s[i:palindromes[i][j]] 是一个回文串
	palindromes := make([][]int, length)
	palindromes[0] = []int{1}
	// 在 O(n ^ 2) 内找出所有回文串
	for i := 1; i < length; i++ {
		// 首先自身是一个回文串
		palindromes[i] = []int{i + 1}
		// 找以 s[i] 为中心的回文串（长度为奇数）
		for l, r := i - 1, i + 1; l >= 0 && r < length && s[l] == s[r]; l, r = l - 1, r + 1 {
			// 将当前回文串的 结束下标 + 1 放入 palindromes[i] 中
			palindromes[l] = append(palindromes[l], r + 1)
		}
		// 找以 s[i - 1:i + 1] 为中心的回文串（长度为偶数）
		for l, r := i - 1, i; l >= 0 && r < length && s[l] == s[r]; l, r = l - 1, r + 1 {
			// 将当前回文串的 结束下标 + 1 放入 palindromes[i] 中
			palindromes[l] = append(palindromes[l], r + 1)
		}
	}

	list := make([]string, length)
	return dfs(palindromes, s, 0, list, 0)
}

// 划分 s[start:] 成回文串，并且每一种结果列表的前缀都是 list[:count]
func dfs(palindromes [][]int, s string, start int, list []string, count int) [][]string {
	// 如果所有字符以用完，则直接返回当前这种结果
	if start == len(s) {
		return [][]string{append(list[:0:0], list[:count]...)}
	}

	var result [][]string
	for i, length := 0, len(palindromes[start]); i < length; i++ {
		end := palindromes[start][i]
		// 枚举当前使用的回文串为 s[start:end] ，递归收集后续的结果
		list[count] = s[start:end]
		result = append(result, dfs(palindromes, s, end, list, count + 1)...)
	}
	return result
}
