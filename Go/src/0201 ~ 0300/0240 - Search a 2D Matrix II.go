// 链接：https://leetcode-cn.com/problems/search-a-2d-matrix-ii/
// 题意：给定一个排序后的 m * n 的矩阵（每行递增，每列递增），
//      判断指定的数字 target 是否存在？

// [
//   [1,   4,  7, 11, 15],
//   [2,   5,  8, 12, 19],
//   [3,   6,  9, 16, 22],
//   [10, 13, 14, 17, 24],
//   [18, 21, 23, 26, 30]
// ]
//
// target = 5 -> 返回 true
// target = 20 -> 返回 false

// 思路1： 分治
//
//      我们以矩阵的中心为基准分成四个小矩阵，递归分治处理
//		并且在每一层递归前进行剪枝，由于矩阵是有序的，
//		所以可以知道该矩阵的最小值（左上角）和最大值（右下角），
//		如果 target 不在这个范围内则可以直接返回 false

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	return dfs(matrix, target, 0, 0, len(matrix[0]) - 1, len(matrix) - 1)
}

func dfs(matrix [][]int, target int, up, left, bottom, right int) bool {
	// 矩阵无元素
	if bottom < up || right < left {
		return false
	}
	// 若 target < 矩阵最小值（左上角） 或者 target > 矩阵最大值（右下角）
	// 则矩阵不可能包含 target ，直接返回 false
	if target < matrix[up][left] || target > matrix[bottom][right] {
		return false
	}
	// 此时有 matrix[up][left] <= target <= matrix[bottom][right]
	// 因为只有一个元素，所以必定是 target ，直接返回 true
	if up == bottom && left == right {
		return true
	}

	// 分成四个子矩阵，递归处理
	row, col := (up + bottom) / 2, (left + right) / 2
	return dfs(matrix, target, up, left, row, col) ||
		dfs(matrix, target, up, col + 1, row, right) ||
		dfs(matrix, target, row + 1, left, bottom, col) ||
		dfs(matrix, target, row + 1, col + 1, bottom, right)
}

// 思路2： 减治
//
//      我们从左下角 (r, c) 开始不断进行判断
//		1. 若 target < matrix[r][c]: 则 r 这一行都不满足题意，
//			可以转换为求解子矩阵 [:r-1][c:]
//		2. 若 target > matrix[r][c]: 则 c 这一列都不满足题意，
//			可以转换为求解子矩阵 [:r][c+1:]
//		3. 此时必有 target == matrix[r][c]: 直接返回 true
//
//		最后矩阵为空仍未找到答案，则返回 false
//
//		时间复杂度： O(m + n)
//		空间复杂度： O(1)

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	m, n := len(matrix), len(matrix[0])
	r, c := m - 1, 0
	// 若矩阵中还有元素，则继续减治处理
	for r >= 0 && c < n {
		if target < matrix[r][c] {
			// r 这一行都不满足题意
			r--
		} else if target > matrix[r][c] {
			// c 这一列都不满足题意
			c++
		} else {
			// target == matrix[r][c] ，直接返回 true
			return true
		}
	}
	// 矩阵中不存在 target
	return false
}
