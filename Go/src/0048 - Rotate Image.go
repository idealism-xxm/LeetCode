// 链接：https://leetcode.com/problems/rotate-image/
// 题意：给定一个 n * n 的矩阵，求其顺时针旋转 90° 后的矩阵？

// 输入：
// [
//   [1,2,3],
//   [4,5,6],
//   [7,8,9]
// ]
// 输出：
// [
//   [7,4,1],
//   [8,5,2],
//   [9,6,3]
// ]

// 输入：
// [
//   [ 5, 1, 9,11],
//   [ 2, 4, 8,10],
//   [13, 3, 6, 7],
//   [15,14,12,16]
// ]
// 输出：
// [
//   [15,13, 2, 5],
//   [14, 3, 4, 1],
//   [12, 6, 8, 9],
//   [16, 7,10,11]
// ]

// 思路：模拟即可
//		从最外圈开始枚举，再枚举每一圈上边的元素
//		若当前是圈上边的元素 (rUp, cUp) ，则很容易可得对应三个需要旋转的元素的坐标
//		找到坐标后按照题意旋转即可

func rotate(matrix [][]int)  {
	n := len(matrix)
	for rUp := 0; rUp < (n >> 1); rUp++ {  // 从最外圈开始枚举左上角行下标
		for cUp := rUp; cUp < (n - 1 - rUp); cUp++ {  // 该圈最右边元素的下标是有 n - 1 - rUp （该元素算成右边一列的方便操作）
			// 找到 (rUp, cUp) 对应的其余的三个坐标
			rRight, cRight := cUp, n - 1 - rUp
			rDown, cDown := n - 1 - rUp, n - 1 - cUp
			rLeft, cLeft := n - 1 - cUp, rUp

			// 旋转
			valueUp, valueRight := matrix[rUp][cUp], matrix[rRight][cRight]
			valueDown, valueLeft := matrix[rDown][cDown], matrix[rLeft][cLeft]
			matrix[rUp][cUp], matrix[rRight][cRight] = valueLeft, valueUp
			matrix[rDown][cDown], matrix[rLeft][cLeft] = valueRight, valueDown
		}
	}
}