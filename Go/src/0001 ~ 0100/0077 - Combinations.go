// 链接：https://leetcode.com/problems/combinations/
// 题意：给定两个数字 n 和 k ，求在 1~n 中选取 k 个数的所有情况？

// 输入：n = 4, k = 2
// 输出：
// [
//   [2,4],
//   [3,4],
//   [2,3],
//   [1,2],
//   [1,3],
//   [1,4],
// ]

// 思路：递归
//		每一层选取比上一层数大的数，保证升序，减少重复情况

func combine(n int, k int) [][]int {
	list := make([]int, k + 1)
	used := make([]bool, n + 1)  // 标记每个数是否已使用
	list[0] = 0  // 第一个放入 0 ，方便直接进入递归
	return dfs(n, k, used, 1, list)
}

func dfs(n, k int, used []bool, count int, list []int) [][]int {
	if count > k {  // 由于多放入一个 0 ，所以当 count > k 时，刚好放入 k 个数
		return [][]int {append(list[:0:0], list[1:count]...)}
	}

	var result [][]int
	// 从前一个数的下一个数开始，保证 list 中升序
	for i := list[count - 1] + 1; i <= n; i++ {
		if !used[i] {  // 若未使用，则放入列表，进行递归
			used[i] = true
			list[count] = i  // 第 count 个数为 i
			result = append(result, dfs(n, k, used, count + 1, list)...)
			used[i] = false
		}
	}
	return result
}
