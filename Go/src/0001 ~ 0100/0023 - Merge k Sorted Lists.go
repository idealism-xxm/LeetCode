// 链接：https://leetcode.com/problems/merge-k-sorted-lists/
// 题意：给定 k 个升序链表，把它们合成一个升序链表。

// 数据限制：
//  k == lists.length
//  0 <= k <= 10 ^ 4
//  0 <= lists[i].length <= 500
//  - (10 ^ 4) <= lists[i][j] <= 10 ^ 4
//  lists[i] 是升序排序的
//  lists[i].length 的和不超过 10 ^ 4

// 输入：lists = [[1,4,5],[1,3,4],[2,6]]
// 输出：[1,1,2,3,4,4,5,6]
// 解释：原始链表为：
//          [
//            1->4->5,
//            1->3->4,
//            2->6
//          ]
//      合成后的链表为：
//          1->1->2->3->4->4->5->6

// 输入：lists = []
// 输出：[]

// 输入：lists = [[]]
// 输出：[]

// 思路1：优先队列（堆）
//
//      我们回想合并两个链表时的场景，每次对比两个链表头结点大小，
//      将较小的结点加入结果链表中，然后后移一个，直至所有结点都在结果链表中。
//
//      那么合并多个链表时也是如此，不过由于存在多个可比较的结点，
//      可以使用最小堆将取最小结点的时间复杂度从 O(k) 降低为 O(logk) 。
//
//      把所有链表的头结点放入堆，每次取最小的结点出来插入结果链表，
//      再将其后的结点放入堆中，直至所有结点都在结果链表中。
//
//
//      设 n 表示所有链表总长度， k 表示升序链表个数
//
//      时间复杂度：O(nlogk)
//      空间复杂度：O(n + k) 。每个结果链表中的结点都是新建的，共有 n 个；同时最小堆最多有 k 个结点

import "container/heap"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    listNodeHeap := ListNodeHeap([]*ListNode{})
    for _, head := range lists  { // 将所有链表头结点推入堆
        if head != nil {
            heap.Push(&listNodeHeap, head)
        }
    }

    resultHeadPre := &ListNode{} // 结果链表头结点的前一个结点，方便统一处理
    pre := resultHeadPre
    for ; listNodeHeap.Len() != 0;  { // 当堆中还有结点时
        minNode := heap.Pop(&listNodeHeap).(*ListNode) // 取出堆中最小的元素
        pre.Next = &ListNode{Val: minNode.Val} // 最小的元素插入结果链表
        pre = pre.Next // 结果链表后移

        if minNode.Next != nil { // 当取出的结点的链表还有下一个元素时，入堆
            heap.Push(&listNodeHeap, minNode.Next)
        }
    }

    return resultHeadPre.Next
}

type ListNodeHeap []*ListNode

func (h ListNodeHeap) Len() int {
    return len(h)
}

func (h ListNodeHeap) Less(i, j int) bool {
    return h[i].Val < h[j].Val
}

func (h ListNodeHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

// 由于 Push 会修改切片，所以需要传指针
func (h *ListNodeHeap) Push(x interface{}) {
    *h = append(*h, x.(*ListNode))
}

// 由于 Pop 会修改切片，所以需要传指针
func (h *ListNodeHeap) Pop() interface{} {
    n := h.Len()
    // 堆顶已被移动到切片最后，方便删除
    x := (*h)[n-1]
    *h = (*h)[0 : n-1]
    return x
}

// 思路2：分治
//      
//      看到题解还可以用分治，思路和归并排序一样，每次只和并两个列表
//
//      设 n 表示所有链表总长度， k 表示升序链表个数
//
//      时间复杂度：O(nlogk)
//      空间复杂度：O(logk) 。结果结点不新建，使用原有结点，不过这种方法修改入参了，有点不优雅；同时栈深度为 O(logk)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    length := len(lists)
    if length == 0 { // 列表为空时递归结束
        return nil
    }
    if length == 1 { // 只有一个列表时直接返回第一个列表即可
        return lists[0]
    }

    // 其他情况都可以合并入下面操作中
    middle := len(lists) >> 1
    return merge2Lists(mergeKLists(lists[0: middle]), mergeKLists(lists[middle:]))
}

func merge2Lists(l1, l2 *ListNode) *ListNode {
    headPre := &ListNode{} // 头结点的前一个结点
    pre := headPre

    for ; l1 != nil && l2 != nil;  {
        if l1.Val < l2.Val {
            pre.Next = l1
            l1 = l1.Next // l1 链表后移
        } else {
            pre.Next = l2
            l2 = l2.Next // l2 链表后移
        }
        pre = pre.Next // 结果链表后移
    }

    if l1 == nil { // 方便统一处理
        l1 = l2
    }
    pre.Next = l1

    return headPre.Next
}