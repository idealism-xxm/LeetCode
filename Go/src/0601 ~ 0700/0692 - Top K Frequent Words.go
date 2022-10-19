// 链接：https://leetcode.com/problems/top-k-frequent-words/
// 题意：给定一个字符串数组 words 和一个正整数 k ，返回 k 个出现次数最多的字符串。
//      结果按照出现次数降序排序，出现次数相同时按字典序升序排序。
//
//      进阶：使用时间复杂度为 O(nlogk) ，空间复杂度为 O(n) 的算法。


// 数据限制：
//  1 <= words.length <= 500
//  1 <= words[i].length <= 10
//  words[i] 仅由英文小写字母组成
//  1 <= k <= words 中不同字符串的数量


// 输入： words = ["i","love","leetcode","i","love","coding"], k = 2
// 输出： ["i","love"]
// 解释： "i" 和 "love" 是出现次数最多的两个，其中 "i" 字典序更小

// 输入： words = ["the","day","is","sunny","the","the","the","sunny","is","is"], k = 4
// 输出： ["the","is","sunny","day"]
// 解释： "the", "is", "sunny", "day" 是出现次数最多的四个，出现次数分别为 4, 3, 2, 1


// 思路： Map + 优先队列/堆
//
//      先用 wordToCnt 统计 wrods 不同字符串的出现次数。
//
//      然后用一个优先队列/堆 top 维护这些字符串及出现次数，
//      并保持堆最多有 k 个元素。
//
//      堆底元素是第 1 多的，堆顶顶元素是第 k 多的，
//      这样当堆内元素超过 k 个时，可以快速移除不满足题意的。
//
//
//      时间复杂度：O(nlogk)
//          1. 需要遍历 word 中的全部 O(n) 个字符串
//          2. 需要遍历 wordToCnt 中全部不同的字符串，
//              最差情况下有 O(n) 个。
//              每次都需要进行至多两次堆操作，时间复杂度为 O(logk)
//      空间复杂度：O(n + k)
//          1. 需要维护 wordToCnt 中全部不同的字符串，
//              最差情况下有 O(n) 个。
//          2. 需要维护 top 中全部 O(k) 个元素


func topKFrequent(words []string, k int) []string {
    // wordToCnt 统计每个字符串的出现次数
    wordToCnt := make(map[string]int)
    for _, word := range words {
        wordToCnt[word] += 1
    }

    // top 维护前 k 个出现次数最多的字符串及其出现次数，
    // 堆顶元素是出现次数第 k 多的，方便移除不满足题意的元素
    var top StringInfoHeap
    for word, cnt := range wordToCnt {
        // 将当前字符串及其出现次数放入 top 中
        heap.Push(&top, &StringInfo{ cnt, word })
        // 保持 top 只有前 k 个出现次数最多的
        if top.Len() > k {
            heap.Pop(&top)
        }
    }

    // 从 top 中取出所有字符串并收集成一个列表
    ans := make([]string, 0, top.Len())
    for top.Len() > 0 {
        stringInfo := heap.Pop(&top).(*StringInfo)
        ans = append(ans, stringInfo.s)
    }
    // ans 的顺序是次数少的在前，所以要反向
    for i, j := 0, len(ans) - 1; i < j; i, j = i + 1, j - 1 {
        ans[i], ans[j] = ans[j], ans[i]
    }
    return ans
}

type StringInfo struct {
    cnt int
    s string
}

type StringInfoHeap []*StringInfo

func (h StringInfoHeap) Len() int {
    return len(h)
}

func (h StringInfoHeap) Less(i, j int) bool {
    // heap 默认是最小堆，所以让出现次数少的在堆顶
    if h[i].cnt != h[j].cnt {
        return h[i].cnt < h[j].cnt
    }
    // 出现次数相同时，让字典序大的在堆顶
    return h[i].s > h[j].s
}

func (h StringInfoHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

// 由于 Push 会修改切片，所以需要传指针
func (h *StringInfoHeap) Push(x interface{}) {
    *h = append(*h, x.(*StringInfo))
}

// 由于 Pop 会修改切片，所以需要传指针
func (h *StringInfoHeap) Pop() interface{} {
    n := h.Len()
    // 堆顶已被移动到切片最后，方便删除
    x := (*h)[n-1]
    *h = (*h)[0 : n-1]
    return x
}
