// 链接：https://leetcode.com/problems/merge-intervals/
// 题意：给定一些线段的起止点，合并重叠的线段，返回不重叠的线段。

// 输入：[[1,3],[2,6],[8,10],[15,18]]
// 输出：[[1,6],[8,10],[15,18]]

// 输入：[[1,4],[4,5]]
// 输出：[[1,5]]

// 思路：模拟
//		按起点和终点升序排序，然后遍历合并即可
//      时间复杂度： O(nlogn) ，空间复杂度： O(n)

import "sort"

type IntervalSlice [][]int

func merge(intervals [][]int) [][]int {
	length := len(intervals)
	if length == 0 {
		return [][]int{}
	}

	sort.Sort(IntervalSlice(intervals))
	result := [][]int{intervals[0]}
	lastIndex := 0
	for i := 1; i < length; i++ {
		if intervals[i][0] <= result[lastIndex][1] {
			result[lastIndex][1] = max(result[lastIndex][1], intervals[i][1])
		} else {
			result = append(result, intervals[i])
			lastIndex++
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (h IntervalSlice) Len() int {
	return len(h)
}

func (h IntervalSlice) Less(i, j int) bool {
	if h[i][0] != h[j][0] {
		return h[i][0] < h[j][0]
	}
	return h[i][1] < h[j][1]
}

func (h IntervalSlice) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
