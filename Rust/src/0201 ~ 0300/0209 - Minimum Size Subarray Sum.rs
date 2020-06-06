// 链接：https://leetcode.com/problems/minimum-size-subarray-sum/
// 题意：给定一个正整数数组和一个整数 s ，求数组中和大于等于 s 的长度最小的子数组的长度 ？

// 输入：s = 7, nums = [2,3,1,2,4,3]
// 输出：2
// 解释：子数组 [4, 3] 是满足题意的长度最小子数组

// 思路1：双指针
//
//		维护前后指针，前后指针范围内的数为满足要求的子数组，不断更新子数组的长度最小值。
//
//		每次我们移动前指针，纳入右边一个新的数，然后调整后指针，
//		使得子数组内的和刚好大于等于 s 即可，然后取这些长度的最小值即可
//
//		时间复杂度： O(n)
//		空间复杂度： O(1)

use std::cmp;

impl Solution {
	pub fn min_sub_array_len(s: i32, nums: Vec<i32>) -> i32 {
		// 初始结果比数组长度大 1
		let mut result = nums.len() + 1;
		// 初始化后指针为 0 ，最开始就必须包含第一个数
		let mut back: usize = 0;
		// 初始化指针范围内的和为 0 ，最开始没有包含任何数字
		let mut sum = 0;
		// 遍历前指针
		for (front, num) in nums.iter().enumerate() {
			// 纳入当前前指针指向的数
			sum += num;
			// 不断移动后指针，直至子数组内的和刚好大于等于 s
			while sum - nums[back] >= s {
				sum -= nums[back];
				back += 1;
			}
			// 更新子数组长度的最小值
			if sum >= s {
				result = cmp::min(result, front - back + 1);
			}
		}

		// 如果不存在这样的子数组，则直接返回 0
		if result == nums.len() + 1 {
			0
		} else {
			result as i32
		}
	}
}

// 思路2：二分
//
//		二分最终结果子数组的长度，初始化 l, r = 1, nums.len()
//		然后看所有长度为 mid 的子数组是否存在和大于等于 s 的，
//		存在则在 [l, mid - 1] 继续二分，不存在则在 [mid + 1, r] 继续二分，
//		直至 l > r
//		最后若 l == nums.len() + 1 ，则返回 0 ，否则返回 l 即可
//
//		时间复杂度： O(nlogn)
//		空间复杂度： O(n)

impl Solution {
	pub fn min_sub_array_len(s: i32, nums: Vec<i32>) -> i32 {
		// 初始化前缀和
		let mut sums = vec![0; nums.len()];
		for (i, num) in nums.iter().enumerate() {
			if i == 0 {
				sums[0] = *num;
			} else {
				sums[i] = sums[i - 1] + *num;
			}
		}
		// 初始化子数组长度范围为 [1, nums.lens()]
		let mut l: usize = 1;
		let mut r = nums.len();
		// 开始二分
		while l <= r {
			let mid = (l + r) >> 1;
			// 默认当前长度不可以
			let mut is_valid = false;
			for i in mid - 1..nums.len() {
				let j = i - mid + 1;
				// 如果存在一个长度为 mid 的子数组和大于等于 s ，则进行标记并退出循环
				if sums[i] - sums[j] + nums[j] >= s {
					is_valid = true;
					break;
				}
			}
			if is_valid {
				// 如果可以，则下次在 [l, mid - 1] 内二分
				r = mid - 1;
			} else {
				// 如果不可以，则下次在 [mid + 1, r] 内二分
				l = mid + 1;
			}
		}

		// 如果不存在这样的子数组，则直接返回 0
		if l == nums.len() + 1 {
			0
		} else {
			l as i32
		}
	}
}
