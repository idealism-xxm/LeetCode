// 链接：https://leetcode.com/problems/top-k-frequent-elements/
// 题意：给定一个整数数组 nums 和一个正整数 k ，
//      返回 nums 中出现次数最多的 k 个数字。


// 数据限制：
//  1 <= nums.length <= 10 ^ 5
//  k 在 [1, nums 中不同数字的个数] 范围内
//  确保答案是唯一的


// 输入： nums = [1,1,1,2,2,3], k = 2
// 输出： [1,2]

// 输入： nums = [1], k = 1
// 输出： [1]


// 思路： 优先队列（堆）
//
//      先统计每个数字出现的次数到 num_to_cnt 中，
//      时间复杂度为 O(n) ，空间复杂度为 O(n) 。
//
//      然后再将 num_to_cnt 中的数字及其出现次数，
//      初始化一个大小为 n 的最大堆，
//      时间复杂度为 (n) ，空间复杂度为 O(n) 。
//
//      最后将最大堆中的数字收集成一个数组并返回，
//      时间复杂度为 O(klogk) ，空间复杂度为 O(k) 。
//
//
//      时间复杂度：O(n + klogk)
//          1. 需要遍历 nums 中全部 O(n) 个数字
//          2. 需要根据 num_to_cnt 中全部 O(n) 个数字初始化最大堆 heap
//          3. 需要移除最大堆 heap 中前 k 大的数字，时间复杂度为 O(klogk)
//      空间复杂度：O(n + k)
//          1. 需要维护 num_to_cnt 统计全部 O(n) 个数字的出现次数，
//              最差情况下所有数字都不同
//          2. 需要维护一个包含 O(n) 个元素的最大堆
//          3. 需要维护一个包含 O(k) 个数字的数组


func topKFrequent(nums []int, k int) []int {
    // 统计 nums 中每个数字出现的次数，
    // 时间复杂度为 O(n) ，空间复杂度为 O(n)
    numToCnt := make(map[int]int)
    for _, num := range nums {
        // num 如果不在 num_to_cnt 中，则初始化为 0 ，
        // 然后对 num 的出现次数加 1
        numToCnt[num] += 1;
    }

    // 将 num_to_cnt 中的数字及其出现次数收集到切片中，
    // 时间复杂度为 O(n) ，空间复杂度为 O(n)
    numCountHeap := make(NumCountHeap, 0, len(numToCnt))
    for num, cnt := range numToCnt {
        numCountHeap = append(numCountHeap, &NumCount{num, cnt})
    }
    // 初始化最大堆，时间复杂度为 O(n)
    heap.Init(&numCountHeap)

    // 初始化结果数组，空间复杂度为 O(k)
    ans := make([]int, 0, k)
    // 将最大堆中前 k 大的数字收集到 ans 中，
    // 时间复杂度为 O(klogk)
    for ; k > 0; k-- {
        ans = append(ans, heap.Pop(&numCountHeap).(*NumCount).Num);
    }

    return ans
}


type NumCount struct {
    Num int
	Count int
}

type NumCountHeap []*NumCount

func (h NumCountHeap) Len() int {
	return len(h)
}

func (h NumCountHeap) Less(i, j int) bool {
    // heap 默认是最小堆，而我们需要最大堆，
    // 所以这里将最大数字放在堆顶
	return h[i].Count > h[j].Count
}

func (h NumCountHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// 由于 Push 会修改切片，所以需要传指针
func (h *NumCountHeap) Push(x interface{}) {
	*h = append(*h, x.(*NumCount))
}

// 由于 Pop 会修改切片，所以需要传指针
func (h *NumCountHeap) Pop() interface{} {
	n := h.Len()
	// 堆顶已被移动到切片最后，方便删除
	x := (*h)[n-1]
	*h = (*h)[0 : n-1]
	return x
}
