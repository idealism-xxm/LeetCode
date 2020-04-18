// 链接：https://leetcode.com/problems/longest-consecutive-sequence/
// 题意：给定一个未排序的整数数组，找到最长连续序列（可重排）的长度（在 O(n) 内完成）？

// 输入： [100, 4, 200, 1, 3, 2]
// 输出： 4
// 解释： 最长连续序列为： [1, 2, 3, 4] ，长度为： 4

// 思路1： 并查集
//		1. 第一反应就是排序后直接判断，但时间复杂度不满足，
//		2. 又想到我每次添加一个数字后，需要把连续的区间合并起来，这个和并查集很像，
//		可以直接使用并查集求解，但是并查集的复杂度是 O(log* n) ，所以最后的时间复杂度是 O(n ✖ log* n)
//		（虽然 O(log* n) 只比 O(1) 慢一点点，但是还不是严格满足要求）
//		3. 最后想着并查集的处理方式，发现其实我们只需要直到区间的开始和结束即可，
//		这样在新数字进来的时候也方便进行合并，
//		所以可以先用一个 map 对数字进行去重，再用另一个 map 执行类似并查集的操作，
//		由于数字不会重复，那么我们只需要记录区间的起止数字即可，
//		新数字 num 进来的时候不会和在已知的区间内，所以只需要判断 num - 1 和 num + 1 是否在 map 中，
//			(1) 如果在，则合并，并移除无用的端点
//			(2) 如果不在，则当前数字为一个区间
//
//		本解法先使用并查集求解
//
//		时间复杂度： O(n ✖ log* n)
//		空间复杂度： O(n)

func longestConsecutive(nums []int) int {
	// 1. 初始化并查集
	parent, count := make(map[int]int), make(map[int]int)
	for _, num := range nums {
		parent[num] = num
		count[num] = 1
	}

	// 2. 合并区间，并记录最大区间长度
	result := 0
	for _, num := range nums {
		// 如果 num - 1 在集合中，则可进行合并
		if _, exists := parent[num-1]; exists {
			union(parent, count, num-1, num)
		}
		// 如果 num + 1 在集合中，则可进行合并
		if _, exists := parent[num+1]; exists {
			union(parent, count, num, num+1)
		}
		// 更新最大区间的长度
		result = max(result, count[find(parent, num)])
	}
	return result
}

func union(parent, count map[int]int, a, b int) {
	// 如果 a, b 不在同一个区间，则进行合并，并更新区间的长度
	aParent, bParent := find(parent, a), find(parent, b)
	if aParent != bParent {
		parent[bParent] = aParent
		count[aParent] += count[bParent]
	}
}

func find(parent map[int]int, a int) int {
	// 如果 a 不是区间的根，则递归找根并更新 a 的父亲为根
	if parent[a] != a {
		parent[a] = find(parent, parent[a])
	}
	return parent[a]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 思路2： map
//		使用 思路1 中提到的第三个方式处理
//		3. 最后想着并查集的处理方式，发现其实我们只需要直到区间的开始和结束即可，
//		这样在新数字进来的时候也方便进行合并，
//		所以可以先用一个 map 对数字进行去重，再用另一个 map 执行类似并查集的操作，
//		由于数字不会重复，那么我们只需要记录区间的起止数字即可，
//		新数字 num 进来的时候不会和在已知的区间内，所以只需要判断 num - 1 和 num + 1 是否在 map 中，
//			(1) 如果在，则合并，并移除无用的端点
//			(2) 如果不在，则当前数字为一个区间
//
//		当然题解中先放入 map 中，然后遍历 nums ，
//		若 num - 1 不在 map 中，则依次对 num + 1, num + 2, ... 判断，直至 num + k 不在 map 中，
//		则区间最大长度 = max(k)
//
//		时间复杂度： O(n)
//		空间复杂度： O(n)

func longestConsecutive(nums []int) int {
	// 1. 初始化 used (标记这个数字是否以及处理过) interval (标记当前区间的另一个端点)
	used, interval := make(map[int]bool), make(map[int]int)
	for _, num := range nums {
		// 如果未使用过，则开始处理
		if !used[num] {
			used[num] = true
			// 由于不会处理重复的数字，且只有相邻的数字才会合并区间，
			// 所以新出现的数字 1. 要么自己独立为一个区间， 2. 要么和其他区间相邻
			// 自己独立为一个区间
			interval[num] = num
			// 区间相邻则进行合并
			if _, exists := interval[num-1]; exists {
				merge(interval, num-1, num)
			}
			if _, exists := interval[num+1]; exists {
				merge(interval, num, num+1)
			}
		}
	}

	// 2. 记录最大区间长度
	result := 0
	for num := range interval {
		// 更新最大区间的长度（假设 num 为区间左端点）
		result = max(result, interval[num]-num+1)
	}
	return result
}

// 合并两个相邻的区间，其中 leftMax 是左边区间的较大一点， rightMin 是右边区间的较小一点
func merge(interval map[int]int, leftMax, rightMin int) {
	// 合并区间，并移除合并后区间中间的点（注意区间只有一个数字时不能进行移除）
	// 先记录合并后区间的两端
	leftMin, rightMax := interval[leftMax], interval[rightMin]
	// 再删除区间中间的点
	if leftMin != leftMax {
		delete(interval, leftMax)
	}
	if rightMin != rightMax {
		delete(interval, rightMin)
	}
	// 最后再设置新的区间
	interval[leftMin], interval[rightMax] = rightMax, leftMin
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
