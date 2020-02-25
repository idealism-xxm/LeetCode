// 链接：https://leetcode.com/problems/remove-duplicates-from-sorted-list/
// 题意：给定一个升序的整型单链表，若一个数字重复出现，则删除至只出现一次，
//		返回剩余结点单单链表头？

// 输入：1->1->2
// 输出：1->2

// 输入：1->1->2->3->3
// 输出：1->2->3

// 思路：模拟即可
//		维护单链表表尾即可，每次待插入值和表尾的值一样，则跳过
//		时间复杂度： O(n)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	headPre := &ListNode {
		Val: head.Val - 1,  // 确保表头元素能够插入
		Next: head,
	}
	tail := headPre
	for cur := head; cur != nil; cur = cur.Next {
		if cur.Val != tail.Val {  // 若与表尾值不同，则插入
			tail.Next = cur
			tail = cur  // 更新表尾
		}
	}
	tail.Next = nil  // 表尾无后继结点
	return headPre.Next
}
