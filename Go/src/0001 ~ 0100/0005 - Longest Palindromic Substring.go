// 链接：https://leetcode.com/problems/longest-palindromic-substring/
// 题意：求最长回文子串

// 思路1：很容易就能想到 O(n^2) 的暴力解法：枚举回文串中点，不停向两端扩张，直到字符不同，统计所有这些长度的最大值即可
func longestPalindrome(s string) string {
	length := len(s)
	res , start, end := 0, 0, 0
	for i := 0; i < length; i += 1 {
		// 奇数回文串
		l, r := i, i
		// 不停向两端扩张
		for ; l >= 0 && r < length && s[l] == s[r]; l, r = l - 1, r + 1 {}
		// 更新最大值
		cur := r - l - 1
		if res < cur {
			res = cur
			start = l + 1
			end = r // 左闭右开
		}
		
		// 偶数回文串
		l, r = i, i + 1
		// 不停向两端扩张
		for ; l >= 0 && r < length && s[l] == s[r]; l, r = l - 1, r + 1 {}
		// 更新最大值
		cur = r - l - 1
		if res < cur {
			res = cur
			start = l + 1
			end = r // 左闭右开
		}
	}
	return s[start: end]
}


// 思路2：依旧记得可以用 O(n) 的 Manacher 解，但太久没有复习，已经忘了插入字符后的原理了，需要继续复习
// 没法画图，不好直观讲解，大家自己搜索理解吧
// 其实只要看了图，就知道这个算法的精妙之处了，充分利用了回文串的特点
func longestPalindrome(s string) string {
	// 预处理字符串
	t := "$#"
	for _, ch := range s {
		t += string(ch) + "#"
	}

	length := len(t)
	radius := make([]int, length)
	center, right := 0, 0
	resCenter, resRadius := 0, 0
	for i := 1; i < length; i += 1 {
		if i < right { // 如果以前找到的回文子串包含了 i，则可以进行计算
			// 此处需要取与 right - i 的较小值，因为即便镜像位置的串有更长的半径
			// 但是 right 没有包含进来，所以无法判断，需要后续扩张
			radius[i] = min(radius[(center << 1) - i], right - i)
		} else { // 如果以前找到的回文子串还没包含 i，那么以 i 为中心的回文子串半径设为 1
			radius[i] = 1
		}
		// 扩张回文子串
		for ; i - radius[i] >= 0 && i + radius[i] < length && t[i - radius[i]] == t[i + radius[i]]; radius[i] += 1 {}
		// 若当前回文子串的右端更靠右，则更新最右端下标及其中心下标
		if right < i + radius[i] {
			center = i
			right = i + radius[i] - 1
		}

		// 如果当前回文子串的半径更大，则更新答案中心下标和半径
		if resRadius < radius[i] {
			resCenter = i
			resRadius = radius[i]
		}
	}

	start := (resCenter - resRadius) >> 1
	end := start + (resRadius - 1)
	return s[start: end]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}