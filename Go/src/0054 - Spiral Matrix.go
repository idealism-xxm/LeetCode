// 链接：https://leetcode.com/problems/spiral-matrix/
// 题意：给定 m * n 的矩阵，返回一个长度为 m * n 的一维数组，
//		包含所有元素，顺序按照顺时针从外向内螺旋。

// 输入：
// [
//   [ 1, 2, 3 ],
//   [ 4, 5, 6 ],
//   [ 7, 8, 9 ]
// ]
// 输出：[1,2,3,6,9,8,7,4,5]

// 输入：
// [
//   [1, 2, 3, 4],
//   [5, 6, 7, 8],
//   [9,10,11,12]
// ]
// 输出：[1,2,3,4,8,12,11,10,9,5,6,7]

// 思路：模拟即可
//		顺时针螺旋加入所有元素，注意换方向的边界即可
//		时间复杂度： O(m * n) ，空间复杂度： O(m * n)

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	m, n := len(matrix), len(matrix[0])  // 矩阵大小
	length := m * n  // 总元素个数
	result := make([]int, length)  // 结果数组
	r, c := 0, 0  // 当前选择的元素下标
	dr := []int{0, 1, 0, -1}  // 各方向 row 变化值
	dc := []int{1, 0, -1, 0}  // 各方向 col 变化值
	direction := 0  // 移动方向
	rowMin, rowMax := 0, m - 1  // 行下标的范围
	colMin, colMax := 0, n - 1  // 列下标的范围
	// 以下两个数组与 dc 和 dr 对应（写出来便于理解）
	rowRangeDelta := []int{1, 0, -1, 0}  // 当前方向移动到边界后， rowMin (正数时) 和 rowMax (负数时) 变化值
	colRangeDelta := []int{0, -1, 0, 1}  // 当前方向移动到边界后， colMin (正数时) 和 rowMax (负数时) 变化值

	for i := 0; i < length; i++ {
		result[i] = matrix[r][c]  // 放入元素
		// 假设不换方向移动到下一个元素
		nextR, nextC := r + dr[direction], c + dc[direction]
		// 如果行列都超出边界，则需要换方向
		if nextR < rowMin || nextR > rowMax || nextC < colMin || nextC > colMax {
			// 更新边界值
			if rowRangeDelta[direction] < 0 {
				rowMax += rowRangeDelta[direction]
			} else {
				rowMin += rowRangeDelta[direction]
			}
			if colRangeDelta[direction] < 0 {
				colMax += colRangeDelta[direction]
			} else {
				colMin += colRangeDelta[direction]
			}
			// 切换方向
			direction = (direction + 1) % 4
		}
		// 移动到下一个元素
		r, c = r + dr[direction], c + dc[direction]
	}
	return result
}
