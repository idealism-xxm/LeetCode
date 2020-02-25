// 链接：https://leetcode.com/problems/edit-distance/
// 题意：给定一个 m * n 的矩阵，如果 matrix[i][j] == 0
//		则：将 matrix[i][0 ~ n-1] 和 matrix[0 ~ m-1][j] 全置为 0
//		在输入数组中操作。

// 输入：
// [
//   [1,1,1],
//   [1,0,1],
//   [1,1,1]
// ]
// 输出：
// [
//   [1,0,1],
//   [0,0,0],
//   [1,0,1]
// ]

// 输入：
// [
//   [0,1,2,0],
//   [3,4,5,2],
//   [1,3,1,5]
// ]
//输出：
// [
//   [0,0,0,0],
//   [0,4,5,0],
//   [0,3,1,0]
// ]

// 思路：模拟
//		很容易就能想到用一个矩阵中不存在的值表示被置为 0 的元素，
//		第二次扫描时再将其替换为 0 ，只是用 O(1) 的空间
//		但是这样做肯定不太优雅，还是没能想到其他的方法
//		最后看了题解才焕然大悟：
//		我们使用第一行和第一列作为标记列，
//		若 matrix[i][j] == 0 ，则将 matrix[i][0] 和 matrix[0][j] 置为 0
//		然后第二次扫描时，只要该行或该列的第一个元素为 0 ，则当前元素置为 0
//		时间复杂度：O(m * n) ，空间复杂度： O(1)

func setZeroes(matrix [][]int)  {
	m, n := len(matrix), len(matrix[0])
	shouldFirstColZero := false
	for i := 0; i < m; i++ {
		// 由于共用 matrix[0][0] ，所以需要额外变量标识第一列是否需要置 0
		if matrix[i][0] == 0 {
			shouldFirstColZero = true
		}
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				// 使用第一行和第一列进行标记
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	// 从右下角操作，最后处理标记的行
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j > 0; j-- {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	// 再处理标记的列
	if shouldFirstColZero {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}
