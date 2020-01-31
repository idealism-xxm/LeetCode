// 链接：https://leetcode.com/problems/insert-interval/
// 题意：给定一些有序的不重叠的线段的起止点，返回插入一个新的线段的列表？

// 输入：intervals = [[1,3],[6,9]], newInterval = [2,5]
// 输出：[[1,5],[6,9]]

// 输入：intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
// 输出：[[1,2],[3,10],[12,16]]

// 思路：模拟
//		找到插入的位置，然后合并重叠线段即可
//      时间复杂度： O(n) ，空间复杂度： O(n)

func insert(intervals [][]int, newInterval []int) [][]int {
	length := len(intervals)
	if length == 0 {
		return [][]int {newInterval}
	}
	// 1. 找到插入的位置
	insertIndex := 0
	for ; insertIndex < length; insertIndex++ {
		if newInterval[0] <= intervals[insertIndex][0] {
			break
		}
	}

	// 2. 合并重叠线段
	minStart := min(intervals[0][0], newInterval[0]) - 1
	result := [][]int {{minStart, minStart}}  // 先放入一个不重合的点，合并初始情况
	lastIndex := 0
	for i := 0; i <= length; i++ {
		// 如果是插入位置
		if insertIndex == i {
			// 如果插入的线段和前一个重叠，则合并
			if newInterval[0] <= result[lastIndex][1] {
				result[lastIndex][1] = max(result[lastIndex][1], newInterval[1])
			} else {
				// 否则直接放入新线段
				result = append(result, newInterval)
				lastIndex++
			}
		}

		if i == length {
			continue
		}
		// 如果当前的线段和前一个重叠，则合并
		if intervals[i][0] <= result[lastIndex][1] {
			result[lastIndex][1] = max(result[lastIndex][1], intervals[i][1])
		} else {
			// 否则直接放入当前线段
			result = append(result, intervals[i])
			lastIndex++
		}
	}
	return result[1:]
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
