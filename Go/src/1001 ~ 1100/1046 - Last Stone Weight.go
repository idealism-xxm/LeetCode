// 链接：https://leetcode.com/problems/last-stone-weight/
// 题意：给定一个非负整数数组 stones ，
//      每次从中移除最大的两个数 x 和 y (x <= y) ，
//      如果 x != y ，则再将 y - x 放回 stones 中。
//      不断重复这个过程，直至 stones 中的数字不足 2 个。
//      如果此时 stones 还有数字，则返回该数字，否则返回 0 。


// 数据限制：
//  1 <= stones.length <= 30
//  1 <= stones[i] <= 1000


// 输入： stones = [2,7,4,1,8,1]
// 输出： 1
// 解释： 移除 7 和 8 ， 8 - 7 = 1 ，所以 stones 转变为 [2,4,1,1,1] 。
//       移除 2 和 4 ， 4 - 2 = 2 ，所以 stones 转变为 [2,1,1,1] 。
//       移除 2 和 1 ， 2 - 1 = 1 ，所以 stones 转变为 [1,1,1] 。
//       移除 1 和 1 ， 1 - 1 = 0 ，所以 stones 转变为 [1] 。
//       stones 还剩一个数字，所以返回 1 。

// 输入： stones = [1]
// 输出： 1
// 解释： stones 只有一个数字，直接返回 1 。


// 思路： 优先队列（堆）
//
//      定义一个最大堆 heap ，维护数组中剩余的数字，
//      初始为 stones 中的数字。
//
//      当堆中数字个数大于 1 时，不断循环处理，
//      每次移除堆顶两个数字 x 和 y (x <= y) ，
//      如果 x != y ，则再将 y - x 放回堆中。
//
//      最后，如果堆中没有数字，返回 0 ；
//      如果堆中剩余一个数字，返回该数字。
//      
//
//		时间复杂度： O(nlogn)
//          1. 通过数组直接建立堆，时间复杂度为 O(n)
//          2. 每次循环至少会移除一个数字，总共会有 O(n) 次循环
//          3. 每次循环时，需要执行出堆和入堆操作，时间复杂度为 O(logn)
//		空间复杂度： O(n)
//          1. 需要维护一个包含 O(n) 个数字的堆


func lastStoneWeight(stones []int) int {
    // 通过数组直接建立堆
    q := IntHeap(stones)
    heap.Init(&q)
    // 至少还有两个数字时，继续循环处理
    for len(q) > 1 {
        // 获取最大数字
        y := heap.Pop(&q).(int)
        // 获取次大数字
        x := heap.Pop(&q).(int)
        // 如果 x != y ，则再将 y - x 放回堆中
        if x != y {
            heap.Push(&q, y - x);
        }
    }

    // 如果此时还有数字，则返回该数字，否则返回 0
    if len(q) != 0 {
        return q[0]
    }
    return 0
}

type IntHeap []int

func (h IntHeap) Len() int {
    return len(h)
}

func (h IntHeap) Less(i, j int) bool {
    // heap 默认是最小堆，而我们需要最大堆，
    // 所以这里将最大数字放在堆顶
    return h[i] > h[j]
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
    // 最后一个数字是堆顶
    x := (*h)[n-1]
    *h = (*h)[0 : n-1]
    return x
}
