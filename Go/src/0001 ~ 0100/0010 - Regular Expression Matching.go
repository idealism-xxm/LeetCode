// 链接：https://leetcode.com/problems/regular-expression-matching/
// 题意：给定一个字符串（空串或只含有小写字母）和一个正则（空串或只含有 小写字母 '.' '*'），判断字符串是否满足正则匹配

// 输入：aab c*a*b
// 输出：true

// 输入：mississippi mis*is*p*.
// 输出：false

// 思路1：第二反映就是有限状态自动机（DFA），在编译原理（龙书）里学过，忘了怎么合并状态和化简了，还是暴力和DP最简单
// 思路2：知道肯定可以用暴力递归，但是觉得复杂度太高，不太优雅，还是想用效率比较高的方法（这样也错过了优化成记忆化搜索的机会）
func isMatch(s string, p string) bool {
	// 若 p 为空，则 s 为空时匹配，s 不为空时不匹配
	sLen, pLen := len(s), len(p)
	if pLen == 0 {
		return sLen == 0
	}

	// 若 p 的下一个字符是 *，则可以跳过当前字符继续进行匹配
	if pLen > 1 && p[1] == '*' && isMatch(s, p[2:]) {
		return true
	}

	// 此时 p 不为空，且 p 的下一个字符不是 *，但 s 未空，则不匹配
	if sLen == 0 {
		return false
	}

	// 若第一个字符 匹配，则继续递归
	if s[0] == p[0] || p[0] == '.' {
		// 如果下一个字符是 * ，则 p 可以不移动，也可以不匹配
		if pLen > 1 && p[1] == '*' {
			return isMatch(s[1:], p) || isMatch(s[1:], p[2:])
		}
		// 如果下一个字符不是 * ，则 p 必须移动
		return isMatch(s[1:], p[1:])
	}

	// 若第一个字符 不匹配，且 p 下一个字符也不是 *，则不匹配
	return false
}

// 思路3：用记忆化搜索优化暴力递归，由于递归时很多子状态由于没有记录会重复递归，影响效率（类似递归求斐波那契数列）
// 		所以需要记录所有的计算过的状态
// 		时间复杂度：O(mn)
func isMatch(s string, p string) bool {
	// 初始化 dp
	sLen, pLen := len(s), len(p)
	dp := make([][]int, sLen + 1) // dp[i][j] 表示 s[i:] 和 p[j:] 的匹配结果，-1 表示还未计算，0 表示不匹配，1 表示匹配
	for i := 0; i <= sLen; i += 1 {
		dp[i] = make([]int, pLen + 1)
		for j := 0; j <= pLen; j += 1 {
			dp[i][j] = -1 // -1 表示没计算过
		}
	}

	// 记忆化搜索
	// 判断 s[i:] 与 p[j:] 是否匹配
	var isSubstringMatch func (i, j int) bool
	isSubstringMatch = func (i, j int) bool{
		// 已经计算过，则直接返回
		if dp[i][j] != -1 {
			return dp[i][j] == 1
		}
	
		// 若 p 已经遍历完，则 s 遍历完时匹配，s 未遍历完时不匹配
		sLen, pLen := len(s), len(p)
		if j == pLen {
			if i == sLen {
				dp[i][j] = 1
			} else {
				dp[i][j] = 0
			}
			return dp[i][j] == 1
		}
	
		// 若 p 的下一个字符是 *，则可以跳过当前字符继续进行匹配
		if j + 1 < pLen && p[j + 1] == '*' && isSubstringMatch(i, j + 2) {
			dp[i][j] = 1
			return true
		}
	
		// 此时 p 未遍历完，且 p 的下一个字符不是 *，但 s 已遍历完，则不匹配
		if i == sLen {
			dp[i][j] = 0
			return false
		}
	
		// 若当前字符 匹配，则继续递归
		if s[i] == p[j] || p[j] == '.' {
			// 如果 p 的下一个字符是 * ，则 p 可以不移动，也可以不匹配
			if j + 1 < pLen && p[j + 1] == '*' {
				if isSubstringMatch(i + 1, j) || isSubstringMatch(i + 1, j + 2) {
					dp[i][j] = 1
				} else {
					dp[i][j] = 0
				}
				return dp[i][j] == 1
			}
			// 如果 p 的下一个字符不是 * ，则 p 必须移动
			if isSubstringMatch(i + 1, j + 1) {
				dp[i][j] = 1
			} else {
				dp[i][j] = 0
			}
			return dp[i][j] == 1
		}
	
		// 若当前字符 不匹配，且 p 下一个字符也不是 *，则不匹配
		dp[i][j] = 0
		return false
	}
	
	return isSubstringMatch(0, 0)
}

// 思路4：dp，dp[i][j] 表示 s[:i] 与 p[:j] 是否匹配（最终发现还是dp最简单）
//		若 dp[i][j] == true，则根据 s[i] 和 p[j] p[j + 1] 的字符进行状态转移
func isMatch(s string, p string) bool {
	// 初始化 dp
	sLen, pLen := len(s), len(p)
	dp := make([][]bool, sLen + 1)
	for i := 0; i <= sLen; i += 1 {
		dp[i] = make([]bool, pLen + 1)
	}
	dp[0][0] = true // 两个都是空串，肯定匹配

	for i := 0; i <= sLen; i += 1 { // i 要等于 sLen，sLen 空串没有继续进行状态转移
		for j := 0; j < pLen; j += 1  { // j 不需要等于 pLen，只有 dp[i + 1][j] 才需要状态转移，然而需要 p[j + 1] 存在
			// 当 s[:i] 与 p[:j] 不匹配时，则不进行状态转移
			if !dp[i][j] {
				continue
			}

			// 当 p 下一个字符是 * 时，s[:i] 与 p[: j + 2] 匹配
			if j + 1 < pLen && p[j + 1] == '*' {
				dp[i][j + 2] = true
			}

			if i < sLen && (s[i] == p[j] || p[j] == '.') {
				// 当前字符匹配时，s[: i + 1] 与 p[: j + 1] 也匹配
				dp[i + 1][j + 1] = true
				// 当 p 下一个字符是 * 时， s[: i + 1]  与 （p[:j] 和 p[: j + 2]） 也匹配
				if j + 1 < pLen && p[j + 1] == '*' {
					dp[i + 1][j] = true
					dp[i + 1][j + 2] = true
				}
			}
		}
	}

	return dp[sLen][pLen]
}