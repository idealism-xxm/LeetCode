// 链接：https://leetcode.com/problems/search-a-2d-matrix/
// 题意：给定一个 m * n 的整数矩阵，每一行从左到右升序，
//		每一行的第一个数字比上一行的最后一个数字大，
//		判断一个数是否在这个矩阵中？

// 输入：
// matrix = [
//   [1,   3,  5,  7],
//   [10, 11, 16, 20],
//   [23, 30, 34, 50]
// ]
// target = 3
// 输出：true

// 输入：
// matrix = [
//   [1,   3,  5,  7],
//   [10, 11, 16, 20],
//   [23, 30, 34, 50]
// ]
// target = 13
//输出：false

// 思路：二分
//		矩阵是有序的，所以二分即可
//		先对第一列二分找到数字可能所在的行，
//		再对当前行进行二分判断数字是否存在
//		时间复杂度：O(logm + logn) ，空间复杂度： O(1)

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	m, n := len(matrix), len(matrix[0])
	// 对第一列二分找到数字可能所在的行
	l, r := 0, m-1
	for l <= r {
		mid := (l + r) >> 1
		if target < matrix[mid][0] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	// 目标数比任何数字都小，则直接返回
	if r < 0 {
		return false
	}

	// 当前行进行二分判断数字是否存在
	row := r
	l, r = 0, n-1
	for l <= r {
		mid := (l + r) >> 1
		if target < matrix[row][mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	// 必定有 0 <= r && r < n
	return matrix[row][r] == target
}
