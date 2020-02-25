// 链接：https://leetcode.com/problems/rotate-list/
// 题意：给定一个单链表，所有结点向右移动 k 个，使得倒数第 k 个变为第一个结点？

// 输入：1->2->3->4->5->NULL, k = 2
// 输出：4->5->1->2->3->NULL

// 输入：0->1->2->NULL, k = 4
// 输出：2->0->1->NULL

// 思路：双指针
//		先遍历一遍得到链表长度，再对其求余得到哪一结点会变成第一个结点
//		最后用双指针遍历，前指针到达尾结点后，后指针刚好到达目标结点的前一个结点
//		时间复杂度： O(n) ，空间复杂度： O(1)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	// 设置一个挂在头结点前的结点，方便操作
	headPre := &ListNode{
		Val:  0,
		Next: head,
	}

	// 获取链表长度
	n, tail := 0, headPre
	for ; tail.Next != nil; tail = tail.Next {
		n++
	}
	if n <= 1 {
		return head
	}

	// 获取截断位置
	k = k % n
	if k == 0 {
		return head
	}

	// 双指针，前指针先往前走 k 个
	front, back := headPre, headPre
	for ; k > 0; k-- {
		front = front.Next
	}
	// 同时往前走，直到前指针到达尾结点
	for ; front.Next != nil; front = front.Next {
		back = back.Next
	}
	// 翻转前后关系
	result := back.Next
	back.Next = nil
	tail.Next = head
	return result
}
