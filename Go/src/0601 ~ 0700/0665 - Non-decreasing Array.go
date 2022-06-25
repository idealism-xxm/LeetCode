// 链接：https://leetcode.com/problems/non-decreasing-array/
// 题意：给定一个整数数组 nums ，最多可以将一个数字修改为任意值，
//      求能否将其变为非递减数组？


// 数据限制：
//   n == nums.length
//   1 <= n <= 10 ^ 4
//   -(10 ^ 5) <= nums[i] <= 10 ^ 5


// 输入： nums = [4,2,3]
// 输出： true
// 解释： 将第一个 4 改为 1 ，数组变为 [1,2,3] ，是非递减数组

// 输入： nums = [4,2,1]
// 输出： false
// 解释： 最多修改一个数字无法将其变为非递减数组


// 思路： 贪心
//
//       维护 modified 表示是否已经修改。
//
//		枚举当前数 nums[i] ，并和前一个数 nums[i - 1] 做比较：
//           1. nums[i] < nums[i - 1]: 需要贪心地修改一个数，如果此时 modified 为 true ，
//               则无法修改，直接返回 false 。
//
//               修改时贪心地将 nums[i - 1] 减小为 nums[i] ，
//               但要保证修改后 nums[i - 1] 仍大于等于 nums[i - 2] ；
//               否则只能将 nums[i] 增大为 nums[i - 1] ， nums[i - 1] 不变。
//           2. nums[i] >= nums[i - 1]: 无需修改任何数。
//
//       最后结束循环时，则说明最多修改一个数字能将其变为非递减数组，返回 true
//
//
//		时间复杂度： O(n)
//           1. 需要遍历全部 O(n) 个数字
//		空间复杂度： O(1)
//           1. 只需要使用常数个额外变量


func checkPossibility(nums []int) bool {
    // modified 表示是否已经修改过
	modified := false
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i - 1] {
			// 如果当前数小于 nums[i - 1] ，则需要修改一个数
			if modified {
				// 如果已经修改过，则无法再修改，直接返回 false
				return false
			}

			modified = true
			if i < 2 || nums[i] >= nums[i - 2] {
				// 此时贪心地将 nums[i - 1] 减小为 nums[i] ，
				// 能保证修改后 nums[i - 1] >= nums[i - 2] 。
				// 这样能在不破坏原有单调性的情况下，尽可能使 nums[i] 最小
				nums[i - 1] = nums[i]
			} else {
				// 此时只能将 nums[i] 增大为 nums[i - 1]
				nums[i] = nums[i - 1]
			}
		}
	}

	return true
}
