// 链接：https://leetcode.com/problems/longest-common-prefix/
// 题意：求多个字符串的最长公共前缀

// 输入：["flower","flow","flight"]
// 输出："fl"

// 思路：模拟即可

// 扩展：题解最后给出了问题的扩展如下
// 		给定一个字符串集合 S=[s1, s2, ..., sn]，找到串 q 和 S 的最长公共前缀
//		第一反应就是 Trie树，其实仔细一想还是和本题一样的做法，先求出 S 的最长公共前缀 LCP
//		每次查询的时候，再求 q 和 LCP 的最长公共前缀即可，时间复杂度一样，空间复杂度为常数
func longestCommonPrefix(strs []string) string {
	// 初始化数组和字符串长度
	length := len(strs)
	if length == 0 {
		return ""
	}

	lens := make([]int, length)
	for i := 0; i < length; i++ {
		lens[i] = len(strs[i])
	}

	// 枚举第一个串的长度
	for end := 0; end < lens[0]; end++ {
		for i := 0; i < length; i++ { // 为了减少多余代码判断，此处还要对第一个串进行判断
			// 如果有 一个串长度不够 或者 当前字符不同，则返回最长公共前缀
			if end >= lens[i] || strs[0][end] != strs[i][end] {
				return strs[0][0:end]
			}
		}
	}
	// 第一个串就是最长公共前缀
	return strs[0]
}