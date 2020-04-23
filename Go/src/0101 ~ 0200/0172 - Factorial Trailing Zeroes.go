// 链接：https://leetcode.com/problems/factorial-trailing-zeroes/
// 题意：求 n! 的末尾有几个零？

// 输入： 3
// 输出： 0
// 解释： 3! = 6 ，末尾没有零

// 输入： 5
// 输出： 1
// 解释： 5! = 120 ，末尾有 1 个零

// 思路： 模拟
//
//		只有一个 2 和 一个 5 才会产生一个零，
//		而每 2 个数就有 2 ，但每 5 个数才有 5 ，
//		所以零的数量取决于 5 的数量，只需要统计 5 的个数即可，
//		需要注意的是有些数含有多个 5 的因子，例如： 25 含有两个 5 ，
//		可以连续除以 5 ，将所有的结果加起来即可
//
//		时间复杂度： O(logn / log5)
//		空间复杂度： O(1)

func trailingZeroes(n int) int {
	result := 0
	for n > 0 {
		n /= 5
		result += n
	}
	return result
}
