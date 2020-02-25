// 链接：https://leetcode.com/problems/powx-n/
// 题意：给定浮点数 x 和整数 n ，求 x ^ n ？

// 输入：2.00000, 10
// 输出：1024.00000

// 输入：2.10000, 3
// 输出：9.26100

// 输入：2.00000, -2
// 输出：0.25000

// 思路：快速幂
//		时间复杂度： O(logn)

func myPow(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}

	result := 1.0
	for ; n > 0; {
		if n & 1 == 1 {
			result *= x
		}
		n >>= 1
		x *= x
	}
	return result
}
