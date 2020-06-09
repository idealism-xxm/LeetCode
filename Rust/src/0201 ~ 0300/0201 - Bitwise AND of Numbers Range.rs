// 链接：https://leetcode.com/problems/bitwise-and-of-numbers-range/
// 题意：给定非负整数的范围 [m, n] ，求其中所有数按位与的结果？

// 输入：[5,7]
// 输出：4

// 输入：[0,1]
// 输出：0

// 思路1：位运算
//
//		数据范围大，所以不可能将所有的数都与起来，
//		可以发现与操作的特点，如果这一位为 0 ，则以后都不需要考虑这一位，
//		所以我们可以从 m 开始，每次对其加上最后一个为 1 的二进制位对应的整数，
//		直至 m 最后一个为 1 的二进制位对应的整数大于 n - m
//
//		看了题解发现，其实这题就是求所有数二进制对最长公共前缀，
//		那么只用考虑 m 和 n 即可
//		我这里是对 m 不停加最后一个为 1 的二进制位对应的整数，
//		同样也可以对 n 不停抹去最后一个为 1 的二进制位，直至 n <= m
//
//		时间复杂度： O(1)
//		空间复杂度： O(1)

impl Solution {
	pub fn range_bitwise_and(m: i32, n: i32) -> i32 {
		// 如果 m 是 0 ，则直接返回 0 即可
		if m == 0 {
			return 0
		}
		// 让 m 可以改变
		let mut m = m;
		// 先初始化结果
		let mut result = m;
		// 当 m 最后一个为 1 的二进制位对应的整数小于等于 n - m 时，
		// 可以继续处理
		while m & (-m) <= n - m {
			// 对 m 加上其最后一个为 1 的二进制位对应的整数
			m += m & (-m);
			// 将结果进行与操作
			result &= m;
		}
		// 返回结果
		result
	}
}

// 思路2：位移
//
//		看了题解发现，其实这题就是求所有数二进制对最长公共前缀，
//		那么只用考虑 m 和 n 即可
//		所以我们可以同时对 m 和 n 都往右移动一位，直至 m 和 n 相等，
//		最后再将 m 往回移动相同对位数即可
//
//		时间复杂度： O(1)
//		空间复杂度： O(1)

impl Solution {
	pub fn range_bitwise_and(m: i32, n: i32) -> i32 {
		// 让 m 和 n 可以改变
		let mut m = m;
		let mut n = n;
		// 记录移动对位数
		let mut shift = 0;
		// m 和 n 不停同时右移，直至 m 等于 n
		while m != n {
			m >>= 1;
			n >>= 1;
			shift += 1;
		}
		// 将 m 往回移动 shift 位
		m << shift
	}
}