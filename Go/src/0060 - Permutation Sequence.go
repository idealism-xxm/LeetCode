// 链接：https://leetcode.com/problems/permutation-sequence/
// 题意：给定正整数 n 和 k ，返回 1 ~ n 的第 k 个排列对应的字符串？

// 输入：n = 3, k = 3
// 输出："213"

// 输入：n = 4, k = 9
// 输出："2314"

// 思路：模拟
//		从最高位开始计算，
//		假设还剩余 i + 1 个数字需要填充，需要找到第 k 小（从 0 开始）的排列
//		则：第一个数字为第 k / i! 小（从 0 开始）的数字
//		k 需要转换为剩余 i 个数中的排列数， k -= (k / i!) * i!
//		时间复杂度： O(n ^ 2) ，空间复杂度： O(n)

import "strconv"

func getPermutation(n int, k int) string {
	k --  // 将 k 变成以 0 为开始的，方便后续计算
	// 计算阶乘
	factorial := make([]int, n + 1)
	factorial[0] = 1
	for i := 1; i <= n; i++ {
		factorial[i] = factorial[i - 1] * i
	}

	// 计算每一位的数字
	result := ""
	used := make([]bool, n + 1)
	for i := n - 1; i >= 0; i-- {
		order := k / factorial[i]  // 计算当前位是第几小的（从 0 开始）
		k -= order * factorial[i]  // 去除前几位的结果
		num := 0
		for j := 1; j <= n && order >= 0; j++ {
			if !used[j] {
				order--
				num = j
			}
		}
		used[num] = true  // 标记已使用
		result += strconv.Itoa(num)  // 添加当前位数字
	}
	return result
}
