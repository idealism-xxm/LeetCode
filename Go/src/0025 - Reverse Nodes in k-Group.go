// 链接：https://leetcode.com/problems/reverse-nodes-in-k-group/
// 题意：给定一个链表，对链表中每 k 个结点进行逆序处理

// 输入：list: 1->2->3->4->5, k = 2
// 输出：2->1->4->3->5

// 输入：list: 2->1->4->3->5，k = 3
// 输出：3->2->1->4->5

// 思路：遍历一次即可，用头插法 + 计数，每 k 个结点一次循环
//      如果最后一次不足 k 个结点，则反转最后一次结果
//
//      记录一个结点即可：待交换的两个结点的前一个结点 pre，每次通过 pre 找到 cur 和 next
//      交换时注意赋值顺序即可，然后往后移动 pre 两个结点，循环处理
//      时间复杂度：O(n + k)，空间复杂度：O(1)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
    headPre := &ListNode{} // 便于统一处理
    cur := head // cur 用于遍历链表，表示当前需要插入的结点
    currentHeadPre := headPre // currentHeadPre 表示当前循环进行插入部分的链表头结点的前一个结点
    count := 0 // 链表总长度
    for ; cur != nil;  { // 当还有结点时，进行循环处理
        currentTail := cur // 第一个结点将变成当前循环内反转后的链表尾结点
        for ; count < k && cur != nil; { // 当前反转 最多只反转 k 个 或者 反转剩余的全部
            next := cur.Next // 保存下一个遍历的结点
            // 头插法插入结点
            cur.Next = currentHeadPre.Next
            currentHeadPre.Next = cur

            cur = next // 待遍历结点后移
            count++
        }
        if count == k { // 如果 count 是 k，则最后一次循环反转的结点是 k 个，需要移动 currentHeadPre
            count = 0
            currentHeadPre = currentTail // 当前循环内的链表尾结点 是 下一次循环的链表头结点的前一个结点
        }
    }
    if count != 0 { // 如果 count 没有重置尾 0，则最后一次循环反转的结点 小于 k 个，需要再反转
        currentHeadPre.Next = reverseList(currentHeadPre.Next) // 反转以 currentHeadPre.Next 开头的链表
    }

    return headPre.Next
}

// 反转链表（用 reverseList(nil, head) 调用即可，最后返回反转后的头结点）
func reverseList(head *ListNode) *ListNode {
    var pre *ListNode
    cur := head
    for ; cur != nil; {
        next := cur.Next
        cur.Next = pre
        pre, cur = cur, next
    }

    return pre
}