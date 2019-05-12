// 链接：https://leetcode.com/problems/longest-substring-without-repeating-characters/
// 题意：看题目就知道要求的是最长没有重复字符的子串的长度
// 思路：双指针法，每次移动 r 后，调整 l，使得 [l, r] 内不存在重复字符，统计每一次移动后 r - l + 1 的最大值即可

func lengthOfLongestSubstring(s string) int {
	length := len(s)
	if length == 0 {
		return 0
	}

	exists := [127]bool{} // 标记字符是否出现过
	l, r := 0, -1 // 从第一个字符开始计算
	res := 0 // 此时长度为 1
	// 遍历每个字符
	for r += 1; r < length; r += 1{
		// 如果字符 s[r] 已经出现
		if exists[s[r]] {
			// 不停移动 l，直到 s[l - 1] == s[r]
			for ; ; {
				exists[s[l]] = false // 标记 s[l] 未出现
				l += 1
				if s[l - 1] == s[r] {
					break
				}
			}
		}

		// 标记 s[r] 已出现
		exists[s[r]] = true
		// 更新结果
		res = max(res, r - l + 1)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}