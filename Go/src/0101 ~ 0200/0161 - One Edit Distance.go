// 链接：https://leetcode.com/problems/one-edit-distance/
// 题意：给定两个字符串 s 和 t ，判断他们的编辑距离是否为 1 ？
//
//		满足编辑距离等于 1 有三种可能的情形：
//
//		1. 往 s 中插入一个字符得到 t
//		2. 从 s 中删除一个字符得到 t
//		3. 在 s 中替换一个字符得到 t

// 输入： s = "ab", t = "acb"
// 输出： true
// 解释： 可以将 'c' 插入字符串 s 来得到 t 。

// 输入： s = "cab", t = "ad"
// 输出： false
// 解释： 无法通过 1 步操作使 s 变为 t 。

// 输入： s = "1203", t = "1213"
// 输出： true
// 解释： 可以将字符串 s 中的 '0' 替换为 '1' 来得到 t 。

// 思路： 双指针
//
//		分别统计两个字符串的长度，如果长度差大于 1 ，
//		则编辑距离必定不可能为 1
//
//		如果 s 更短，则只能是 s 插入一个字符
//		如果长度相同，则只能是 s 替换一个字符
//		如果 s 更长，则只能是 s 删除一个字符
//
//		同时遍历 s 和 t ，遇到第一个不匹配的字符，
//		然后通过长度差判断编辑操作，然后更新对应的指针即可。
//		最后编辑过一次才返回 true ，否则都返回 false
//
//		时间复杂度： O(n)
//		空间复杂度： O(1)

func isOneEditDistance(s string, t string) bool {
	lenS, lenT := len(s), len(t)
	lenDiff := lenS - lenT
	// 长度差不在 1 以内，直接返回 false
	if lenDiff < -1 || lenDiff > 1 {
		return false
	}

	// 标记是否编辑过
	edited := false
	l, r := 0, 0
	for ; l < lenS && r < lenT; {
		// 如果当前字符相同，则直接处理下一个
		if s[l] == t[r] {
			l++
			r++
			continue
		}
		// 如果已经编辑过，则不能继续编辑，直接返回 false
		if edited {
			return false
		}

		if lenDiff == -1 {
			// 如果 s 更短，则需要插入一个字符
			// 假装插入，只需对 r + 1 即可
			r++
		} else if lenDiff == 0 {
			// 如果长度相同，则需要替换字符
			// 假装替换， l 和 r 都需要 + 1
			l++
			r++
		} else {
			// 如果 t 更短，则需要删除一个字符
			// 假装删除， 只需对 l + 1 即可
			l++
		}
		edited = true
	}
	// 如果两个串，同时遍历完成，则必须要编辑过才行
	if l == lenS && r == lenT {
		return edited
	}
	// 如果两个串，没有同时遍历完成，则必须要没有编辑过才行
	return !edited
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
