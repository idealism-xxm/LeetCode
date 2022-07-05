// 链接：https://leetcode.com/problems/candy/
// 题意：有 n 个孩子站成一排，第 i 个孩子的评分为 ratings[i] ，
//		现在要给这些孩子发糖果，要满足以下限制：
//			1. 每个孩子至少要有一个糖果
//			2. 若 ratings[i] > ratings[i - 1] ，则第 i 个孩子的糖果数要多余第 i - 1 个孩子的糖果数；
//			   若 ratings[i] > ratings[i + 1] ，则第 i 个孩子的糖果数要多余第 i + 1 个孩子的糖果数；
//		求最少要发多少个糖果？


// 数据限制：
//  n == ratings.length
//  1 <= n <= 2 * 10 ^ 4
//  0 <= ratings[i] <= 2 * 10 ^ 4


// 输入： [1,0,2]
// 输出： 5
// 解释： 分别发 2, 1, 2 个糖果

// 输入： [1,2,2]
// 输出： 4
// 解释： 分别发 1, 2, 1 （第三个孩子发一个糖果，因为满足上述两个限制）


// 思路1：DP + 排序
//
//		设 dp[i] 表示第 i 个孩子最少能发的糖果数，
//      那么 dp[i] 可以通过两边的 dp 值进行状态转移获得。
//
//      但如果按照通常的下标顺序遍历的话，那么总有一边的 dp 值无法确定，
//      状态转移也就无法进行。
//
//      不过本题也加了其他限制，所以我们可以自定顺序，
//      即可以通过 rating 从小到大的顺序处理，来进行状态转移，
//      这种处理方法与 LeetCode 329 一致。
//
//      如果我们按照 rating 从小到大的顺序进行状态转移，
//      所有 ratings[j] 小于 ratings[i] 的 dp[j] 都已确定，
//      所以 dp[i] 就能通过两边的的 dp 值转移得到。
//
//      ratings 不会改变，所以我们最开始就将 ratings 转成单元格数组 cells ，
//      其中 cells[i] = (matrix[i], i) ，然后按照 ratings[i] 升序排序。
//
//      初始化令所有 dp[i] = 1 ，即至少要分一个糖果。
//
//      然后遍历 cells 进行状态转移，此时能保证在处理 i 之前，
//      所有 rating 小于 ratings[i] 的 dp 值都已处理完成。
//
//
//      时间复杂度：O(nlogn)
//          1. 需要遍历收集 ratings 全部 O(n) 个数
//          2. 需要对 cells 进行排序，时间复杂度是 O(nlogn)
//          3. 需要遍历 cells 全部 O(n) 个元素进行状态转移
//      空间复杂度：O(n)
//          1. 需要维护 cells 全部 O(n) 个元素


type RatingNode struct {
	Rating, Index int
}


func candy(ratings []int) int {
	n := len(ratings)
	// 将 ratings 收集成 (rating, index) 的列表，
	// 然后按照 rating 升序排序，方便后续状态转移
	cells := make([]*RatingNode, n)
	// dp[i] 表示第 i 个孩子分到的糖果数，初始化最少每人一个
	dp := make([]int, n)
	for i, rating := range ratings {
		cells[i] = &RatingNode{ rating, i }
		dp[i] = 1
	}
	sort.Slice(cells, func(i, j int) bool { return cells[i].Rating < cells[j].Rating })

	// ans 维护所有孩子分到的糖果数之和
	ans := 0
	// 按 rating 升序进行状态转移，这样就能保证在更新 dp[i] 时，
	// 其左右两侧的 dp 值均已确定
	for _, cell := range cells {
		rating, i := cell.Rating, cell.Index
		// 如果其评分大于左侧的评分，则 dp[i] 至少为 dp[i - 1] + 1 ，
		// 此时不用更新最大值，因为 dp[i] = 1 < 1 + 1 <= dp[i - 1] + 1 
		if i > 0 && rating > ratings[i - 1] {
			dp[i] = dp[i - 1] + 1
		}
		// 如果其评分大于左侧的评分，则 dp[i] 至少为 dp[i + 1] + 1 ，
		// 注意此时要更新最大值
		if i + 1 < n && rating > ratings[i + 1] {
			dp[i] = max(dp[i], dp[i + 1] + 1)
		}
		// 计入第 i 个孩子分到的糖果
		ans += dp[i]
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}


// 思路2： 贪心
//
//		我们先给第一个孩子发一个糖果，然后依次给后边的第 i 个孩子发糖果，使得 i 及其左边的孩子都满足限制
//		1. 我们记录现在评分处于上升阶段、下降阶段还是平缓阶段，以及这一阶段的开始点
//		2. 当处于平缓阶段时：
//			(1) ratings[i] == ratings[i - 1] 时，第 i 个孩子的糖果 = 1
//			(2) ratings[i] <  ratings[i - 1] 时，第 i 个孩子的糖果 = 第 i - 1 个孩子的糖果 + 1 ，并改为上升阶段
//			(3) ratings[i] >  ratings[i - 1] 时，第 i 个孩子的糖果 = 第 i - 1 个孩子的糖果 - 1 ，并改为下降阶段
//		3. 当处于上升阶段时：
//			(1) ratings[i] == ratings[i - 1] 时，第 i 个孩子的糖果 = 1 ，并改为平缓阶段
//			(2) ratings[i] <  ratings[i - 1] 时，第 i 个孩子的糖果 = 第 i - 1 个孩子的糖果 + 1
//			(3) ratings[i] >  ratings[i - 1] 时，第 i 个孩子的糖果 = 第 i - 1 个孩子的糖果 - 1 ，并改为下降阶段
//		4. 当处于平缓阶段时：
//			(0) 若此时下降阶段结束（即： ratings[i] <=  ratings[i - 1] ），需要先处理下降阶段中所发的糖果
//					若第 i - 1 个孩子的糖果数 num < 1 ，则下降阶段的所有孩子均补发 1 - num 颗糖果
//					若第 i - 1 个孩子的糖果数 num > 1 ，则下降阶段的除开始的孩子均收回 num - 1 颗糖果
//			(1) ratings[i] == ratings[i - 1] 时，第 i 个孩子的糖果 = 1 ，并改为平缓阶段
//			(2) ratings[i] <  ratings[i - 1] 时，第 i 个孩子的糖果 = 第 i - 1 个孩子的糖果 + 1
//			(3) ratings[i] >  ratings[i - 1] 时，第 i 个孩子的糖果 = 第 i - 1 个孩子的糖果 - 1 ，并改为下降阶段
//
//		时间复杂度： O(n)
//		空间复杂度： O(1)

const (
	FLAT = 1
	UP   = 2
	DOWN = 3
)

func candy(ratings []int) int {
	// 最后添加一个，变为平缓，方便处理最后一个阶段
	ratings = append(ratings, ratings[len(ratings)-1])
	stage, stageStartIndex := FLAT, 0
	// 第一个小孩发 1 个糖果
	lastCandyNum, sum := 1, 1
	for i, n := 1, len(ratings); i < n; i++ {
		switch stage {
		case FLAT:
			if ratings[i-1] < ratings[i] {
				stage, stageStartIndex = UP, i-1
				// 上升阶段，每次比前一个孩子多一颗
				lastCandyNum++
			} else if ratings[i] == ratings[i-1] {
				// 平缓阶段，直接一颗
				lastCandyNum = 1
			} else {
				stage, stageStartIndex = DOWN, i-1
				// 下降阶段，每次比前一个孩子少一颗
				lastCandyNum--
			}
		case UP:
			if ratings[i-1] < ratings[i] {
				// 上升阶段，每次比前一个孩子多一颗
				lastCandyNum++
			} else if ratings[i] == ratings[i-1] {
				stage, stageStartIndex = FLAT, i-1
				// 平缓阶段，直接一颗
				lastCandyNum = 1
			} else {
				stage, stageStartIndex = DOWN, i-1
				// 下降阶段，每次比前一个孩子少一颗
				lastCandyNum--
			}
		case DOWN:
			// 下降阶段结束，需要特殊处理这一阶段所发的糖果数
			if ratings[i-1] <= ratings[i] {
				// 下降阶段最后一次的糖果数量小于 1 ，则对下降阶段的每一个孩子均补发 1 - lastCandyNum 个糖果
				if lastCandyNum < 1 {
					sum += (i - stageStartIndex) * (1 - lastCandyNum)
					// 最后一个孩子更新为 1 颗糖果
					lastCandyNum = 1
				} else if lastCandyNum > 1 {
					// 下降阶段最后一次的糖果数量大于 1 ，则对下降阶段的非起始孩子均拿回 lastCandyNum - 1 个糖果
					sum -= (i - stageStartIndex - 1) * (lastCandyNum - 1)
					// 最后一个孩子更新为 1 颗糖果
					lastCandyNum = 1
				}
			}

			if ratings[i-1] < ratings[i] {
				stage, stageStartIndex = UP, i-1
				// 上升阶段，每次比前一个孩子多一颗
				lastCandyNum++
			} else if ratings[i] == ratings[i-1] {
				stage, stageStartIndex = FLAT, i-1
				// 平缓阶段，直接一颗
				lastCandyNum = 1
			} else {
				// 下降阶段，每次比前一个孩子少一颗
				lastCandyNum--
			}
		}
		sum += lastCandyNum
	}
	// 因为最后添加了一个平缓阶段，所以要减去最后一个孩子的糖果
	return sum - 1
}
