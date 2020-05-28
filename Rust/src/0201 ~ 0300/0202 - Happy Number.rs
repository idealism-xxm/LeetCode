// 链接：https://leetcode.com/problems/happy-number/
// 题意：给定一个正整数，不断求每一位的平方和，判断最终是否能变为 1 ？

// 输入：19
// 输出：true
// 解释：
// 1^2 + 9^2 = 82
// 8^2 + 2^2 = 68
// 6^2 + 8^2 = 100
// 1^2 + 0^2 + 0^2 = 1

// 思路：模拟
//
//		用一个 bool 数组表示每个数字是否出现过，
//			若出现过，则直接返回 false
//			若没出现，则继续计算平方和，直至结果为 1
//
//		估算第一次计算后结果的上限：
//			32 位正整数直接按照 10 个 9 计算，
//			则最终能产生的数为 81 * 10 = 810 ，
//			直接使用长度为 1000 的 bool 数组标记就不会越界

impl Solution {
	pub fn is_happy(n: i32) -> bool {
		// 赋值给可变遍历，方便后续操作
		let mut num = n;

		// 创建一个长度为 1000 的 bool 数组
		let mut used = [false; 1000];
		// 无限循环
		loop {
			// 计算下一个数
			let mut next_num = 0;
			// 当前数不为 0 时，则还可以计算平方和
			while num > 0 {
				// 获取前一个数字的最后一位
				let digit: i32 = &num % 10;
				// 计算平方和加入下一个数字
				next_num += &digit * &digit;
				num /= 10;
			}
			// 如果下一个数字是 1 ，则直接返回 true
			if next_num == 1 {
				return true
			}
			// 如果下一个数字出现过，则直接返回 false
			if used[next_num as usize] {
				return false
			}
			// 准备计算下一个数
			num = next_num;
			// 标记这个数字已出现过
			used[num as usize] = true;
		}
	}
}
