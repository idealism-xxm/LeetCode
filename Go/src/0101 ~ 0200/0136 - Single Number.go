// 链接：https://leetcode.com/problems/single-number/
// 题意：给一个整数数组，有一个数出现一次，其他数均出现两次，找出出现一次的数？
//
// 输入： [2,2,1]
// 输出： 1

// 输入： [4,1,2,1,2]
// 输出： 4

// 思路： 异或
//
//		a ^ a = 0 ，所以出现两次的数异或后会相互抵消，最后剩余的就是只出现一次的数
//
//		时间复杂度： O(n)
//		空间复杂度： O(1)

func singleNumber(nums []int) int {
	result := 0
	for _, num := range nums {
		result ^= num
	}
	return result
}
