// 链接：https://leetcode.com/problems/partition-list/
// 题意：给定一个单链表和一个数字 x ，将单链表中小于 x 的结点都移动至大于等于 x 的结点左边，
//		且保持链表原有顺序，返回结果单链表头？

// 输入：head = 1->4->3->2->5->2, x = 3
// 输出：1->2->2->4->3->5

// 思路：模拟即可
//		维护两个链表表尾，一个是小于 x 的链表，一个是大于等于 x 的链表，遍历插入即可
//		时间复杂度： O(n)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	smallHeadPre := &ListNode{
		Val:  0,
		Next: nil,
	}
	otherHeadPre := &ListNode{
		Val:  0,
		Next: nil,
	}

	smallTail, otherTail := smallHeadPre, otherHeadPre
	for ; head != nil; head = head.Next {
		if head.Val < x { // 若当前结点 小于 x ，则放入 small 链表
			smallTail.Next = head
			smallTail = head
		} else { // 否则，放入 other 链表
			otherTail.Next = head
			otherTail = head
		}
	}
	otherTail.Next = nil               // 为 other 链表最终指向 nil
	smallTail.Next = otherHeadPre.Next // 将 other 链表挂在 small 链表后
	return smallHeadPre.Next
}
