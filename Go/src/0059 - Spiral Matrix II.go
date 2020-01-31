// 链接：https://leetcode.com/problems/spiral-matrix-ii/
// 题意：给定正整数 n ，返回一个包含 1 ~ n^2 的矩阵，顺序为从外向内螺旋。

// 输入：3
// 输出：
// [
//   [ 1, 2, 3 ],
//   [ 8, 9, 4 ],
//   [ 7, 6, 5 ]
// ]

// 思路：模拟
//		将 (https://leetcode.com/problems/spiral-matrix/) 的代码修改即可
//		顺时针螺旋加入所有元素，注意换方向的边界即可
//		时间复杂度： O(n ^ 2)
package main

func generateMatrix(n int) [][]int {
	result := make([][]int, n) // 结果矩阵
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}

	length := n * n          // 总元素个数
	r, c := 0, 0             // 当前选择的元素下标
	dr := []int{0, 1, 0, -1} // 各方向 row 变化值
	dc := []int{1, 0, -1, 0} // 各方向 col 变化值
	direction := 0           // 移动方向
	rowMin, rowMax := 0, n-1 // 行下标的范围
	colMin, colMax := 0, n-1 // 列下标的范围
	// 以下两个数组与 dc 和 dr 对应（写出来便于理解）
	rowRangeDelta := []int{1, 0, -1, 0} // 当前方向移动到边界后， rowMin (正数时) 和 rowMax (负数时) 变化值
	colRangeDelta := []int{0, -1, 0, 1} // 当前方向移动到边界后， colMin (正数时) 和 rowMax (负数时) 变化值

	for i := 0; i < length; i++ {
		result[r][c] = i + 1 // 放入元素
		// 假设不换方向移动到下一个元素
		nextR, nextC := r+dr[direction], c+dc[direction]
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
		r, c = r+dr[direction], c+dc[direction]
	}
	return result
}
