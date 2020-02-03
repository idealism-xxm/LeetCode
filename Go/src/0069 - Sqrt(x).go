// 链接：https://leetcode.com/problems/sqrtx/
// 题意：模拟 int sqrt(int x) ，结果截断。

// 输入：4
// 输出：2

// 输入：8
// 输出：2

// 思路：二分
//		对范围 [1, x] 进行二分计算
//		时间复杂度： O(logx) ，空间复杂度： O(1)

func mySqrt(x int) int {
	l, r := 1, x
	for ; l <= r; {
		mid := (l + r) >> 1
		if mid * mid <= x {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return r
}
