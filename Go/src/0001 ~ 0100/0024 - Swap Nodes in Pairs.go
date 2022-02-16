// 链接：https://leetcode.com/problems/swap-nodes-in-pairs/
// 题意：给定一个链表，对链表中每两个结点进行交换

// 输入：1->2->3->4
// 输出：2->1->4->3

// 思路：模拟
//
//      记录一个结点即可：待交换的两个结点的前一个结点 pre，每次通过 pre 找到 cur 和 next
//      交换时注意赋值顺序即可，然后往后移动 pre 两个结点，循环处理
//
//      时间复杂度：O(n)
//      空间复杂度：O(1)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    if head == nil { // 空链表直接返回
        return head
    }
    
    headPre := &ListNode{Next: head} // 方便统一处理
    pre := headPre
    for ; pre.Next != nil && pre.Next.Next != nil;  { // 当待交换的两个结点都存在时
        cur, next := pre.Next, pre.Next.Next
        
        pre.Next = next
        cur.Next = next.Next
        next.Next = cur
        
        pre = pre.Next.Next // pre 后移两个结点
    }
    
    return headPre.Next
}