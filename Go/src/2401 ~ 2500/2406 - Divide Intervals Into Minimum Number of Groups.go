// 链接：https://leetcode.com/problems/divide-intervals-into-minimum-number-of-groups/
// 题意：给定一个二维整型数组 intervals ，其中 intervals[i] = [left_i, right_i] ，
//      表示一个左开右闭区间 [left_i, right_i] 。
//
//      求将其分成最少多少组，使得每组中的区间都不相交。
//
//      两个区间相交当且仅当至少存在一个数都在这两个区间中。
//      例如：区间 [1,5] 和区间 [5,8] 是相交的。


// 数据限制：
//  1 <= intervals.length <= 10 ^ 5
//  intervals[i].length == 2
//  1 <= left_i <= right_i <= 10 ^ 6


// 输入： intervals = [[5,10],[6,8],[1,5],[2,3],[1,10]]
// 输出： 3
// 解释： 可以划分成以下三组：
//       - 第 1 组： [1, 5], [6, 8]
//       - 第 2 组： [2, 3], [5, 10]
//       - 第 3 组： [1, 10]

// 输入： intervals = [[1,3],[5,6],[8,10],[11,13]]
// 输出： 1
// 解释： 全部区间都不相交。


// 思路： 优先队列（堆） + 排序
//
//      如果我们按照区间左边界升序的顺序遍历所有区间，
//      那么在处理区间 intervals[i] 时，所有左边界小于 left_i 的区间都已处理过。
//
//      按照这种顺序进行分组（组内按左边界升序排序）时，
//      我们就只用关心该分组的最后一个区间右边界。
//
//      因为这些分组的最后一个区间左边界必定小于等于 left_i ，
//      所以相交只有一种可能：分组的最后一个区间右边界的大于等于 left_i 。
//
//      所以我们可以贪心地找到所有分组最后一个区间右边界的最小值 smallest ：
//          1. smallest < left_i: 则当前区间与该分组的所有区间都不相交，可以放入当前分组，
//              更新该分组的最后一个区间右边界为 right_i 。
//          2. smallest >= right_i: 则现存的所有分组中，必定都会与当前区间相交。
//              （因为现存分组中的最后一个区间的左边界都小于等于 left_i ，
//              右边界都大于等于 right_i ）
//              
//              此时需要加入一个新的分组，其最后一个区间右边界为 right_i 。
//
//      不过直接寻找 smallest 需要遍历全部分组，时间复杂度为 O(n) 。
//
//      我们可以定义一个最小堆 heap ，维护当前已产生的所有分组的最后一个区间右边界，
//      这样就能在 O(logn) 内找到最后一个区间右边界的最小值 smallest 。
//
//
//      时间复杂度：O(n)
//          1. 需要遍历 intervals 中全部 O(n) 个区间
//      空间复杂度：O(n)
//          1. 需要用 heap 维护全部组的区间右边界，最差情况下有 O(n) 个


import "container/heap"
import "sort"


func minGroups(intervals [][]int) int {
    // 按区间左边界升序排序，保证后续遍历时，
    // 所有已遍历过的区间的左边界都不大于当前区间的左边界
    sort.SliceStable(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    // 最小堆 q ，维护当前已产生的所有分组中的区间右边界
    var q IntHeap
    // 初始化一个右边界为 0 的分组，方便后续处理
    heap.Push(&q, 0)
    // 遍历每个区间 [left, right]
    for _, interval := range intervals {
        left, right := interval[0], interval[1]
        // 获取所有分组中，区间右边界的最小值
        smallest := heap.Pop(&q).(int)
        if smallest < left {
            // 如果不相交，则当前区间可以加入该分组，
            // 那么该分组的区间右边界变为 right
            heap.Push(&q, right)
        } else {
            // 如果相交，则现存的所有分组中，必定都会与当前区间相交。
            // （因为现存分组中的最后一个区间的左边界都小于等于 left ，
            // 右边界都大于等于 right ）
            //
            // 此时需要将原分组放回，并加入一个新的分组。
            heap.Push(&q, smallest)
            heap.Push(&q, right)
        }
    }

    return q.Len()
}

type IntHeap []int

func (h IntHeap) Len() int {
    return len(h)
}

func (h IntHeap) Less(i, j int) bool {
    // heap 默认是最小堆，将最小数字放在堆顶
    return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

// 由于 Push 会修改切片，所以需要传指针
func (h *IntHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}

// 由于 Pop 会修改切片，所以需要传指针
func (h *IntHeap) Pop() interface{} {
    n := h.Len()
    // 堆顶已被移动到切片最后，方便删除
    x := (*h)[n-1]
    *h = (*h)[0 : n-1]
    return x
}
