// 链接：https://leetcode.com/problems/scramble-string/
// 题意：给定两个字符串 s1 和 s2，我们可以把 s1 递归地分割成两个非空子字符串，从而将其表示为二叉树。
//		我们可以挑选任何一个非叶节点，然后交换它的两个子节点。
//		求 s1 经过若干次操作是否能变成 s2 ？

// 输入：s1 = "great", s2 = "rgeat"
// 输出：true

// 输入：s1 = "abcde", s2 = "caebd"
// 输出：false

// 思路1：递归 + 记忆化
//		被题目整懵了，都忘记可以递归搜索了，看了可以用 递归 就能立刻想到递归搜索
//		每一层先进行剪枝，判断当前对应都字符串字符是否对应
//		然后枚举左半部分都字符数，分 不交换和交换 s1 左右两部分 递归进行

import (
	"sort"
	"strings"
)

var ok = make(map[string]bool) // ok[s1 + s2] 表示 s1 和 s2 是否满足题意

func isScramble(s1 string, s2 string) bool {
	return dfs(s1, s2)
}

func dfs(s1 string, s2 string) bool {
	if s1 == s2 {  // 如果当前字符串一样，则两者符合题意
		return true
	}
	s1s2 := s1 + s2
	if isOk, exists := ok[s1s2]; exists {  // 如果以前遇到过，则直接返回结果
		return isOk
	}
	// 剪枝，如果每个字符出现次数不一样，则必定不可能相同
	if sortString(s1) != sortString(s2) {
		ok[s1s2] = false
		return false
	}

	length := len(s1)
	// 枚举 s1 左半部分的字符数
	for i := 1; i < length; i++ {
		// 1. s1 当前结点不交换子结点
		// 若左右两部分都可以符合题意，则直接返回 true
		if dfs(s1[:i], s2[:i]) && dfs(s1[i:], s2[i:]) {
			ok[s1s2] = true
			return true
		}

		// 2. s1 当前结点交换子结点
		// 若 s1 右半 和 s2 左半可以符合题意，且 s1 左半 和 s2 右半可以符合题意，
		// 则直接返回 true （注意字符数）
		if dfs(s1[i:], s2[:length - i]) && dfs(s1[:i], s2[length - i:]) {
			ok[s1s2] = true
			return true
		}
	}
	// 都不符合题意，返回 false
	ok[s1s2] = false
	return false
}

func sortString(str string) string {
	strs := strings.Split(str, "")
	sort.Strings(strs)
	return strings.Join(strs, "")
}

// 思路2：DP
//		既然已经可以递归 + 记忆化了，那么可以想着用 DP 实现
//		考虑我们记忆化其实记录 s1 和 s2 长度相同的子串是否满意题意
//		（因为每次只能将一个串分成两部分，两部分内部可以继续处理，所以必定是子串和子串相比）
//		设 dp[i][j][k] 表示 s1[i:i+k] 和 s2[j:j+k] 是否满足题意
//		初始化： dp[i][j][1] = s1[i] == s2[j]
//		状态转移方程：
//		1. 不交换 s1 当前子串左右两部分
//			只要存在一个 l (0 < l < k) 使得：dp[i][j][l] 和 dp[i+l][j+l][k-l] 都为 true
//			那么 dp[i][j][k] = true ，否则 dp[i][j][k] = false
//		2. 交换 s1 当前子串左右两部分
//			只要存在一个 l (0 < l < k) 使得：dp[i+l][j][k-l] 和 dp[i][j+k-l][l] 都为 true
//			那么 dp[i][j][k] = true ，否则 dp[i][j][k] = false
//
//		还是感觉 DP 想不到就很难，但是只要有一点思路，努力往 DP 想一想，还是很简单都
//		时间复杂度： O(len(s1)^4) ，空间复杂度： O(len(s1)^3)

func isScramble(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	length := len(s1)
	dp := make([][][]bool, length)
	for i := 0; i < length; i++ {
		dp[i] = make([][]bool, length)
		for j := 0; j < length; j++ {
			dp[i][j] = make([]bool, length + 1)
			dp[i][j][1] = s1[i] == s2[j]
		}
	}

	for k := 2; k <= length; k++ {  // 长度放在最外层，因为大长度依赖小长度
		for i := 0; i + k <= length; i++ {
			for j := 0; j + k <= length; j++ {
				for l := 1; l < k; l++ {
					// 不交换左右两部分
					if dp[i][j][l] && dp[i+l][j+l][k-l] {
						// 当前情形满足题意，直接跳出本层循环
						dp[i][j][k] = true
						break
					}
					// 交换左右两部分
					if dp[i+l][j][k-l] && dp[i][j+k-l][l] {
						// 当前情形满足题意，直接跳出本层循环
						dp[i][j][k] = true
						break
					}
				}
			}
		}
	}
	return dp[0][0][length]
}
