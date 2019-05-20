// 链接：https://leetcode.com/problems/merge-k-sorted-lists/
// 题意：给定多个有序链表，把它们合成一个有序链表
// 额外收获：Go 中使用堆

// 输入：[
//          1->4->5,
//          1->3->4,
//          2->6
//       ]
// 输出：1->1->2->3->4->4->5->6

// 思路1：想法和合并两个有序链表时类似，不过由于存在多个可比较的结点，需要用到堆
//      把所有链表的第一个结点放入堆，每次取最小的结点出来插入结果链表，在将其后的节点放入堆即可
//      时间复杂度：O(nlogk) （n表示所有链表总长度，k表示有序链表个数）
//      空间复杂度：O(n) （创建新链表需要 O(n)，堆需要 O(k)）

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
    x := (*h)[n-1] // 最后一个元素是最小的
    *h = (*h)[0 : n-1]
    return x
}

// 思路2：看到题解还可以用分治，思路和归并排序一样，每次只和并两个列表
//      时间复杂度：O(nlogk) （n表示所有链表总长度，k表示有序链表个数）
//      空间复杂度：O(1) （结果结点不新建，使用原有结点，不过这种方法修改入参了，有点不优雅）

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
    headPre := &ListNode{} // 头节点的前一个结点
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