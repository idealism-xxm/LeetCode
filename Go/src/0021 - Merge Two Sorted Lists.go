// 链接：https://leetcode.com/problems/merge-two-sorted-lists/
// 题意：合并两个有序链表

// 输入：1->2->4, 1->3->4
// 输出：1->1->2->3->4->4

// 输入：(]
// 输出：false

// 思路：模拟即可，每次取当前结点的较小值插入链表
//      当某一个链表为空时，将剩下的链表全部插入结果链表尾部

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    headPre := &ListNode{} // 初始化头结点的前一个结点，方便统一处理
    pre := headPre

    for ; l1 != nil && l2 != nil;  {
        if l1.Val < l2.Val {
            cur := &ListNode{Val: l1.Val}
            pre.Next = cur
            l1 = l1.Next // 链表后移
        } else {
            cur := &ListNode{Val: l2.Val}
            pre.Next = cur
            l2 = l2.Next // 链表后移
        }
        pre = pre.Next // 结果链表后移
    }

    if l1 == nil { // 便于统一处理
        l1 = l2
    }

    for ; l1 != nil; l1 = l1.Next { // 剩余的值全部插入结果链表尾部
        cur := &ListNode{Val: l1.Val}
        pre.Next, pre = cur, cur
    }

    return headPre.Next
}