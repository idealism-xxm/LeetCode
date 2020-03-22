// 链接：https://leetcode.com/problems/gas-station/
// 题意：在一条环形路上有 n 个加油站，给定每个加油站 i 可以加油的数量 gas[i] 和
//		到下一个加油站 i + 1 所需消耗的油的数量 cost[i] ，
//		求油箱无限容量的车从哪个加油站开始可以绕环形路一圈，不存在则返回 -1 ？

// 输入：
// gas  = [1,2,3,4,5]
// cost = [3,4,5,1,2]
// 输出： 3
// 解释： 从加油站 3 开始出发，油箱里油的数量 = 4
//		运行到加油站 4 ，油箱里油的数量 = 4 - 1 + 5 = 8
//		运行到加油站 0 ，油箱里油的数量 = 8 - 2 + 1 = 7
//		运行到加油站 1 ，油箱里油的数量 = 7 - 3 + 2 = 6
//		运行到加油站 2 ，油箱里油的数量 = 6 - 4 + 3 = 5
//		运行到加油站 3 需要花费 5 ，而油箱里刚好的油还有 5 ，可以返回出发点

// 输入：
// gas  = [2,3,4]
// cost = [3,4,3]
// 输出： -1

// 思路1： 线段树
//
//		第一反应就是这两个数组可以合二为一， gas - cost 后变为到下一个加油站的油箱中油的变化量
//		然后对这个数组求和（代表从加油站 0 开始，到每个加油站所剩的油量），
//			1. 若不存在负数，则表示可以完整绕弯一圈，直接返回当前
//			2. 若存在一个为负数，则表示无法完整绕弯一圈，需要和数组转换成从下一个加油站开始的形式
//				除了 i 对应的和外的的每个和 = 原本和 - 当前加油站的油变化量 (gas[i] - cost[i])
//				i 对应的和 = i - 1 最新对应的和 + 当前加油站的油变化量 (gas[i] - cost[i])
//				然后重复判断和数组的元素即可
//
//		想到这的时候发现既需要区间修改，又需要求区间最小值，很容易就能想到线段树，O(nlogn) 内能求出答案
//
//		时间复杂度： O(nlogn)
//		空间复杂度： O(n)

import "math"

func canCompleteCircuit(gas []int, cost []int) int {
	// 1. 求出和数组，并构建线段树
	length := len(gas)
	sum := make([]int, length)
	sum[0] = gas[0] - cost[0]
	for i := 1; i < length; i++ {
		sum[i] = sum[i - 1] + gas[i] - cost[i]
	}
	tree := BuildSegmentTree(sum)

	// 2. 找到能使 minSum 不为负的加油站
	for i := 0; i < length; i++ {
		// 如果都不为负，则加油站 i 可以作为起始点
		if tree.QueryMin(1, 1, length) >= 0 {
			return i
		}
		// 当前加油站不满足，转换到下一个加油站，先对全部和减去 gas[i] - cost[i]
		tree.Add(1, 1, length, -(gas[i] - cost[i]))
		// 获取加油站 i - 1 在新起点下的值
		lastIndex := (i - 1 + length) % length + 1
		lastSum := tree.QueryMin(1, lastIndex, lastIndex)
		// 获取加油站 i 当前的值
		curSum := tree.QueryMin(1, i + 1, i + 1)
		// 算出加油站 i 转换成以 i + 1 起点时对应的增量，并对其加上这个增量
		tree.Add(1, i + 1, i + 1, lastSum + gas[i] - cost[i] - curSum)
	}
	// 所有的加油站都不满足，则返回 -1
	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type Node struct {
	// 区间左右下标
	l, r int
	// 当前区间的最小值
	minVal int
	// lazy 标记，不为 0 时表示需要将 lazy 传递下去
	lazy int
}

type SegmentTree []*Node

func BuildSegmentTree(sum []int) SegmentTree {
	nodes := make([]*Node, len(sum) * 4)
	tree := SegmentTree(nodes)
	tree.Build(sum, 1, 1, len(sum))
	return tree
}

// 根据 sum 数组构建线段树，并返回整个数组的最小值
func (t SegmentTree) Build(sum []int, i, l, r int) int {
	// 叶子结点的最小值就是本身
	if l == r {
		t[i] = &Node{
			l: l,
			r: r,
			// 线段数区间是 [1, n] ， sum 区间是 [0, n - 1]
			minVal: sum[l - 1],
			lazy:   0,
		}
		return t[i].minVal
	}

	// t[i] 的最小值就是两个子树最小值的较小值
	mid := (l + r) >> 1
	t[i] = &Node{
		l: l,
		r: r,
		minVal: min(t.Build(sum, i << 1, l, mid), t.Build(sum, (i << 1) + 1, mid + 1, r)),
		lazy:   0,
	}
	return t[i].minVal
}

// 对区间 [l, r] 内所有的值加上 val
func (t SegmentTree) Add(i, l, r, val int) {
	// 当前区间在所选区间范围内，使用 lazy 处理
	if l <= t[i].l && t[i].r <= r {
		t[i].minVal += val
		t[i].lazy += val
		return
	}
	
	// 处理子树前先传递 lazy
	t.updateLazy(i)
	// 如果左子树部分在区间 [l, r] 内，则递归处理
	if r <= t[i << 1].r {
		t.Add(i << 1, l, r, val)
	}
	// 如果右子树部分在区间 [l, r] 内，则递归处理
	if t[i << 1].r < l {
		t.Add((i << 1) + 1, l, r, val)
	}
	t[i].minVal = min(t[i << 1].minVal, t[(i << 1) + 1].minVal)
}

// 查询区间 [l, r] 内的最小值
func (t SegmentTree) QueryMin(i, l, r int) int {
	// 当前区间在所选区间范围内，直接返回 minVal
	if l <= t[i].l && t[i].r <= r {
		return t[i].minVal
	}

	// 处理子树前先传递 lazy
	t.updateLazy(i)
	minVal := math.MaxInt64
	// 如果左子树部分在区间 [l, mid] 内，则递归处理
	if r <= t[i << 1].r {
		minVal = min(minVal, t.QueryMin(i << 1, l, r))
	}
	// 如果右子树部分在区间 [l, r] 内，则递归处理
	if t[i << 1].r < l {
		minVal = min(minVal, t.QueryMin((i << 1) + 1, l, r))
	}
	return minVal
}

func (t SegmentTree) updateLazy(i int) {
	// 不需要传递
	if t[i].lazy == 0 {
		return
	}
	lazy := t[i].lazy
	// 去除自身的标记
	t[i].lazy = 0
	// 更新左子树
	t[i << 1].minVal += lazy
	t[i << 1].lazy += lazy
	// 更新右子树
	t[(i << 1) + 1].minVal += lazy
	t[(i << 1) + 1].lazy += lazy
}

// 思路2： 一次遍历
//
//		看完官方题解 (https://leetcode-cn.com/problems/gas-station/solution/jia-you-zhan-by-leetcode/)，
//		发现自己还是没有对题目的各个条件进行足够对分析
//
//		若存在答案，则有： sum(gas - cost) >= 0 ，且起点下标可由以下方式计算
//		设从 i 能到达 j ，
//		若 i 不能到达 j + 1 ，则必定是 j 不能到达 j + 1 ，此时我们只能让 j + 2 作为起始点，继续计算
//
//		时间复杂度： O(n)
//		空间复杂度： O(1)

func canCompleteCircuit(gas []int, cost []int) int {
	length := len(gas)
	// totalSum 为 sum(gas - cost) ， curSum 为从 start 开始到当前加油站的 gas - cost 综合
	totalSum, curSum := 0, 0
	start := 0
	for i := 0; i < length; i++ {
		totalSum += gas[i] - cost[i]
		curSum += gas[i] - cost[i]
		if curSum < 0 {
			curSum = 0
			start = i + 1
		}
	}
	if totalSum < 0 {
		return -1
	}
	return start
}
