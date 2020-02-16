// 链接：https://leetcode.com/problems/reverse-linked-list-ii/
// 题意：给定一个单链表，只遍历一次将 [m, n] (1 <= m <= n <= length) 内的结点翻转？

// 输入：1->2->3->4->5->NULL, m = 2, n = 4
// 输出：1->4->3->2->5->NULL

// 思路：模拟
//		先找到第 m 个结点的前一个结点，
//		将连同 m 在内的后续 n - m 个结点翻转，
//		然后将第三部分挂在第二部分后面即可
// 		时间复杂度： O(n)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	headPre := &ListNode{
		Val:  0,
		Next: head,
	}
	// 先找到第 m 个结点的前一个结点，即第一部分第尾部
	firstPartTail := headPre
	for i := m - 1; i > 0; i-- {
		firstPartTail = firstPartTail.Next
	}
	// 将连同 m 在内的后续 n - m + 1 个结点翻转
	secondPartTail := firstPartTail.Next
	cur := firstPartTail.Next
	for i := n - m; i >= 0; i-- {
		next := cur.Next
		cur.Next = firstPartTail.Next
		firstPartTail.Next = cur

		cur = next
	}
	// cur 指向第三部分表头
	secondPartTail.Next = cur
	return headPre.Next
}
