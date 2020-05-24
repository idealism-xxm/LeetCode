// 链接：https://leetcode.com/problems/longest-substring-with-at-most-two-distinct-characters/
// 题意：给定一个字符串 s ，找到最多含有两个不同字符的最长子串的长度？

// 输入： "eceba"
// 输出： 3
// 解释： "ece" 长度为 3

// 输入： "ccaabbb"
// 输出： 5
// 解释： "aabbb" 长度为 5

// 思路： 双指针
//
//		维护前后指针 l 和 r ，
//		每次移动前指针，并记录新字符出现次数，
//			如果当前不同字符出超过两个，则移动后指针，
//			直至不同字符数变为 2
//		更新最长子串的长度
//
//		时间复杂度： O(n)
//		空间复杂度： O(1)
package main

func lengthOfLongestSubstringTwoDistinct(s string) int {
	length := len(s)
	// 当前前后在指针中每个字符出现的次数，以及不同的字符数
	charCount, distinctCharNum := make([]int, 128), 0
	result := 0
	for l, r := 0, 0; r < length; r++ {
		// 新字符纳入窗口考虑
		newChar := s[r]
		// newChar 字符出现次数 + 1
		charCount[newChar]++
		// newChar 字符第一次出现，则 distinctCharNum + 1
		if charCount[newChar] == 1 {
			distinctCharNum++
		}
		// 不同的字符数大于 2 时，不断移动后指针，
		// 直至不同的字符数等于 2
		for ; distinctCharNum > 2; l++ {
			// 旧字符剔除考虑
			oldChar := s[l]
			// oldChar 字符出现次数 - 1
			charCount[oldChar]--
			// oldChar 字符全部没了，则 distinctCharNum - 1
			if charCount[oldChar] == 0 {
				distinctCharNum--
			}
		}
		// 更新最长子串的长度
		result = max(result, r-l+1)
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
