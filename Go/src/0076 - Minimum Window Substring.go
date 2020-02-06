// 链接：https://leetcode.com/problems/minimum-window-substring/
// 题意：给定两个字符串 s 和 t ，
//		求 s 中最短子串，使得其中每个字符的数量都大于等于 t 中对应字符的数量？

// 输入：S = "ADOBECODEBANC", T = "ABC"
// 输出："BANC"

// 思路：双指针
//		设定左右指针 l, r ，在 [l, r] 范围内
//		count[ch] 表示窗口中每个字符还需出现的次数
//		remain 表示还剩余多少个字符不符合要求
//		每次向右移动 r ，纳入 s[r] 到窗口内，直到 r == length
//			s[r] 还需出现的字符数 - 1 ： count[s[r]]--
//			若 count[s[r]] == 0 ，则 s[r] 在窗口内已满足要求，剩余不符合要求的字符数 - 1
//			再不断右移动 l ，从窗口排除 s[l] ，直到 remain != 0
//				更新答案
//				s[l] 还需出现的字符数 + 1 ： count[s[l]]++
//				若 count[s[l]] == 1 ，则 s[l] 排除后， s[l] 抢号不满足要求，剩余不符合要求的字符数 + 1
//		时间复杂度：O(len(s) + len(t)) ，空间复杂度： O(len(s) + len(t))

func minWindow(s string, t string) string {
	count := make(map[byte]int)  // 表示窗口中每个字符还需出现的次数
	for i := len(t) - 1; i >= 0; i-- {
		count[t[i]]++
	}
	remain := len(count)  // t 中不同的字符数
	length := len(s)
	resultL, resultR, resultLen := 0, 0, length + 1
	l, r := 0, 0  // 前后指针 [l, r] 内的字符串为答案待选字符串
	for ; r < length; r++ {  // 移动右指针，扩大窗口范围
		count[s[r]]--
		if count[s[r]] == 0 {  // 如果等于 0 ，则表示 s[r] 在窗口中出现的次数满意要求，剩余数字 - 1
			remain--
		}
		for ; remain == 0; l++ {  // 移动做指针，缩小窗口范围
			if r - l + 1 < resultLen {
				resultL, resultR, resultLen = l, r, r - l + 1
			}
			count[s[l]]++
			if count[s[l]] == 1 {  // 如果等于 1 ，则表示 s[l] 去除后， s[l] 在窗口中出现的次数抢号不满足要求，剩余数字 + 1
				remain++
			}
		}
	}
	// 如果从未更新过，则无答案
	if resultLen == length + 1 {
		return ""
	}
	return s[resultL:resultR + 1]
}
