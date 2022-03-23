// 链接：https://leetcode.com/problems/broken-calculator/
// 题意：给定两个整数 startValue 和 target ，
//      现在我们可以不断对 startValue 做以下两种操作：
//          1. 将 startValue 乘 2
//          2. 将 startValue 减 1
//
//      求能得到 target 的最小操作数？

// 数据限制：
//  1 <= startValue, target <= 10 ^ 9


// 输入： startValue = 2, target = 3
// 输出： 2
// 解释： 先乘 2 ，再减 1 ： 2 -> 4 -> 3

// 输入： startValue = 5, target = 8
// 输出： 2
// 解释： 先减 1 ，再乘 2 ： 5 -> 4 -> 8

// 输入： startValue = 3, target = 10
// 输出： 3
// 解释： 先乘 2 ，再减 1 ，再乘 2 ： 3 -> 6 -> 5 -> 10

// 思路： 贪心
//
//      我们从 target 开始处理，想着执行最少的操作得到 startValue ，
//      那么两个操作将变为相反的操作：
//          1. 将 target 除以 2
//          2. 将 target 加上 1
//
//      1. 如果 target <= startValue ，
//          那么必定是不断对 startValue 执行加 1 操作，
//          所需操作数为 startValue - target 。
//      2. 如果 target > startValue ，
//          那么必定是需要不断减小 target 的值，转换为第 1 种情况。
//          有以下两种情况：
//
//          (1) 若 target 是奇数，那么只能执行加 1 操作，
//              即 target += 1
//          (2) 若 target 是偶数，那么我们考虑一下减小 target 的操作，
//              先对 target 执行加 1 ，那么变为奇数，只能继续执行加 1 操作，
//              然后才能执行除以 2 的操作，即 (target + 1 + 1) / 2 ，
//              需要 3 次操作。
//
//              而 (target + 1 + 1) / 2 = target / 2 + 1 ，
//              那么先进行除以 2 的操作，就只需要 2 次操作。
//
//              综合这两种情况，我们此时贪心地对 target 执行除以 2 的操作。
//
//
//      时间复杂度：O(log(target / startValue))
//          1. 需要不断对 target 执行除以 2 的操作，
//              直至 target <= startValue
//      空间复杂度：O(1)
//          1. 只需要维护常数个额外变量

func brokenCalc(startValue int, target int) int {
	// ans 统计所需的操作次数
	ans := 0
	// 当 target 比 startValue 大时，要继续循环处理
	for startValue < target {
		if target&1 == 0 {
			// 如果 target 是偶数，贪心地执行除以 2 操作
			target >>= 1
		} else {
			// 如果 target 是奇数，只能执行加 1 操作
			target += 1
		}
		// 操作次数加 1
		ans += 1
	}

	// 现在 startValue >= target ，
	// 还需要对 target 执行 startValue - target 次加 1 操作，
	// 才能得到 startValue
	return ans + startValue - target
}