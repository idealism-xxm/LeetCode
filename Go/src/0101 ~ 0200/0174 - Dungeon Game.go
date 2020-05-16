// 链接：https://leetcode.com/problems/dungeon-game/
// 题意：有一个 M * N 的地牢，骑士初始在左上角，每一个格子有一个值，
//		正数代表会给骑士加相应的生命值，负数则会给骑士减相应的生命值，
//		求骑士初始生命值最小为多少时，
//		存在一条路径使得他能走到右下角并且任意时刻的生命值都为正？

// 输入：
//	[
//		[-2, -3,  3],
//		[-5, -10, 1],
//		[10, 30,  -5]
//	]
//	输出：7
//  解释：最优路径： 右 -> 右 -> 下 -> 下

// 思路1： DP + 二分
//
//		按照常规思路很容易就能想到
//		dp[i][j][k] 表示从左上角走到 (i, j) 时生命值为 k 时路径上的最小生命值，
//		其中 k 这一维用 map 存储，这样可以勉强进行状态转移，但是时间复杂度很大，
//		一直卡在这不知如何处理；
//		又想 dp[i][j] 表示从左上角走到 (i, j) 时的最大生命值和路径上的最小生生命值，
//		但这样无法进行状态转移，因为不满足最优子结构，
//		可能是中间的某个值转移能达到更优的下一个状态
//
//		经过朋友的点拨之后，发现可以对答案（初始生命值）进行二分，
//		然后使用 dp 求出是否存在一条路径满足最小生命值不低于 1 的路径即可，
//		若存在这样的路径，则答案可能是更小的值；若不存在这样的路径，则答案只能是更大的值
//
//		dp[i][j] 表示从左上角走到 (i, j) 时生命的最大值（生命值为 0 是无法进行状态转移）
//		状态转移方程大致如下（未考虑无法转移的情况）：
//			dp[i][j] = max(dp[i - 1][j], dp[i][j - 1])
//		只要 dp[m - 1][n - 1] 为正数，即可成功营救公主
//
//		时间复杂度： O(m * n * log(math.MaxInt32))
//		空间复杂度： O(m * n)

import "math"

func calculateMinimumHP(dungeon [][]int) int {
	// 由于骑士必须任何时刻都有正数生命值，所以初始生命值最小为 1
	l, r := 1, math.MaxInt32
	for l < r {
		mid := (l + r) >> 1
		if canRescue(dungeon, mid) {
			// 如果能成功营救公主，则 [l, mid] 范围内必定可以营救
			r = mid
		} else {
			// 如果能成功营救公主，则 [mid + 1, r] 范围内才可以营救
			l = mid + 1
		}
	}
	// 由于二分时，右边界值是满足拯救公主的，所以需要取右边界值
	return r
}

func canRescue(dungeon [][]int, initHP int) bool {
	m, n := len(dungeon), len(dungeon[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	// 初始生命值 + 左上角的值 = 骑士在左上角后的生命值
	dp[0][0] = dungeon[0][0] + initHP
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 当前 dp[i][j] 的值已经计算出来了，
			// 所以可以继续向下和向右的状态转移

			if dp[i][j] <= 0 {
				// 如果当前生命值不为正数，则不能进行状态转移
				continue
			}
			// 向右状态转移
			if j + 1 < n {
				dp[i][j + 1] = max(dp[i][j + 1], dp[i][j] + dungeon[i][j + 1])
			}
			// 向下状态转移
			if i + 1 < m {
				dp[i + 1][j] = max(dp[i + 1][j], dp[i][j] + dungeon[i + 1][j])
			}
		}
	}
	return dp[m - 1][n - 1] > 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 思路2： DP
//		看了题解后才恍然大悟，正向操作无法保证最优子结构，但是反向操作却可以，
//		dp[i][j] 表示从 (i, j) 开始到右下角保证存在一条生命值恒为正数的路径的所需最小生命值
//		因为 dp[i][j] 已经计算了到右下角所需最小生命值，
//		而无论从 (i - 1, j) 还是从 (i, j - 1) 到达 (i, j)
//		都不会影响从 (i, j) 到右下角所需的最小生命值，具备了最优子结构，
//		所以可从 dp[i][j] 转移到 dp[i - 1][j] 和 dp[i][j - 1]
//		初始化：
//			dp[m - 1][n - 1] = max(1, 1 - dungeon[m - 1][n - 1])
//			1. 当 dungeon[m - 1][n - 1] >= 0 时，不需要额外的生命值，
//			只需要 1 点生命值保证刚进入能活下来即可
//			2. 当 dungeon[m - 1][n - 1] < 0 时，需要补齐缺少的值，
//			并需要额外 1 点生命值保证最低时也为正数即可
//		状态转移时也是这个思路，具体见代码注释
//
//		时间复杂度： O(m * n)
//		空间复杂度： O(m * n)

import "math"

func calculateMinimumHP(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			// 初始化为 math.MaxInt32 ，方便后续直接更新
			dp[i][j] = math.MaxInt32
		}
	}
	// 从右下角直接开始所需的最小生命值
	//	1: dungeon[m - 1][n - 1] >= 0 ，
	//		即最后一个格子不扣生命值，最小也需要 1
	//	1 - dungeon[m - 1][n - 1]: dungeon[m - 1][n - 1] < 0 ，
	//		即最后一个格子扣生命值，补齐后还需要额外 1 点生命值保证一直为正数
	dp[m - 1][n - 1] = max(1, 1 - dungeon[m - 1][n - 1])
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			// 当前 dp[i][j] 的值已经计算出来了，
			// 所以可以继续向上和向左的状态转移

			// 向左状态转移
			if j - 1 >= 0 {
				// max(1, dp[i][j] - dungeon[i][j - 1])
				// 表示从 (i, j - 1) 经过 (i, j) 时所需的最小生命值
				// 1: dungeon[i][j - 1] >= dp[i][j]
				//		即 dungeon[i][j - 1] 提供的生命值已够保证后续生命值一直为正数，
				//		但 (i, j - 1) 本身还需要 1 点生命值
				// dp[i][j] - dungeon[i][j - 1]: dungeon[i][j - 1] < dp[i][j]
				//		即 dungeon[i][j - 1] 提供的生命值无法保证后续生命值一直为正数，
				//		还需要 dp[i][j] - dungeon[i][j - 1] 补齐生命值，保证后续生命值一直为正数
				//		由于最开始即状态转移的过程中已给每个格子都有额外的 1 点生命值，
				//		所以本次不需要额外的 1 点生命值（相当于将其静默转移到当前格子）
				dp[i][j - 1] = min(dp[i][j - 1], max(1, dp[i][j] - dungeon[i][j - 1]))
			}
			// 向上状态转移
			if i - 1 >= 0 {
				// max(1, dp[i][j] - dungeon[i - 1][j])
				// 表示从 (i - 1, j) 经过 (i, j) 时所需的最小生命值
				// 1: dungeon[i - 1][j] >= dp[i][j]
				//		即 dungeon[i - 1][j] 提供的生命值已够保证后续生命值一直为正数，
				//		但 (i - 1, j) 本身还需要 1 点生命值
				// dp[i][j] - dungeon[i - 1][j]: dungeon[i - 1][j] < dp[i][j]
				//		即 dungeon[i - 1][j] 提供的生命值无法保证后续生命值一直为正数，
				//		还需要 dp[i][j] - dungeon[i - 1][j] 补齐生命值，保证后续生命值一直为正数
				//		由于最开始即状态转移的过程中已给每个格子都有额外的 1 点生命值，
				//		所以本次不需要额外的 1 点生命值（相当于将其静默转移到当前格子）
				dp[i - 1][j] = min(dp[i - 1][j], max(1, dp[i][j] - dungeon[i - 1][j]))
			}
		}
	}
	return dp[0][0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
