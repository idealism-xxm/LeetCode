// 链接：https://leetcode.com/problems/reorder-list/
// 题意：给定一个链表 0->1->...->n - 1->n ，将其重新排序为 0->n->1->n - 1->... ？

// 输入： 1->2->3->4
// 输出： 1->4->2->3

// 输入： 1->2->3->4->5
// 输出： 1->5->2->4->3

// 思路： 快慢指针
//
//		先通过快慢指针找到前一半（最后一个结点到 Next 设置为 nil），
//		然后将后一半反转，再间隔插入到前一半到列表中
//
//		时间复杂度： O(n)
//		空间复杂度： O(1)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode)  {
	// 如果只有 0 个 或 1 个 结点，则直接返回
	if head == nil || head.Next == nil {
		return
	}
	fast, slow := head, head
	for ; ; {
		// 如果 fast 已经走到尾部，则当前 slow 则在前半尾部
		if fast == nil || fast.Next == nil {
			break
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	// 此时 slow 指向链表前半尾部，拿到后半部分反转后的头部
	back := reverseList(slow.Next)
	// 前半部分断开
	slow.Next = nil
	// 开始间隔插入
	for front := head; back != nil; front = front.Next.Next {
		// 先保存下一个结点
		backNext := back.Next

		// 后半部分的当前结点插入
		back.Next = front.Next
		front.Next = back

		// 处理下一个结点
		back = backNext
	}
}

// 头插法反转链表
func reverseList(head *ListNode) *ListNode {
	headPre := &ListNode {
		Val: 0,
		Next: nil,
	}
	for ; head != nil; {
		// 先保存下一个结点
		headNext := head.Next

		// 头结点插入
		head.Next = headPre.Next
		headPre.Next = head

		// 处理下一个结点
		head = headNext
	}
	return headPre.Next
}
