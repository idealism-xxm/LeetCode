// 链接：https://leetcode.com/problems/count-primes/
// 题意：给定一个非负整数 n ，求小于 n 的所有质数的个数 ？

// 输入： 10
// 输出： 4
// 解释： 有 4 个小于 10 的质数： 2, 3, 5, 7

// 思路：埃式筛
//
//		用一个数组表示下标对应的数字是否为质数
//		然后从 2 开始遍历
//			若当前数字 i 不是质数，则直接处理下一个
//			若当前数字 i 是质数，则质数数量 + 1 ，
//				并将 i 的所有倍数都标记为非质数即可
//
//		时间复杂度： O(n * loglogn)
//		空间复杂度： O(n)

impl Solution {
	pub fn count_primes(n: i32) -> i32 {
		// 标记每个数字是否为质数
		let mut prime: Vec<bool> = vec![true; n as usize];
		// 当前质数的数量
		let mut count = 0;
		// 遍历 2 ~ n 的所有数
		for i in 2..n as usize {
			// 如果当前不是质数，则直接处理下一个即可
			if !prime[i] {
				continue;
			}
			// 质数数量 + 1
			count += 1;
			// 如果当前是质数，则将其所有的倍数都标记为非质数
			let mut num = i << 1;
			while num < n as usize {
				prime[num] = false;
				num += i;
			}
		}
		return count;
	}
}
