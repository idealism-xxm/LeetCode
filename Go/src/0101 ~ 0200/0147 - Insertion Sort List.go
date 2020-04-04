// 链接：https://leetcode.com/problems/insertion-sort-list/
// 题意：使用插入排序对链表进行排序 ？

// 输入： 4->2->1->3
// 输出： 1->2->3->4

// 输入： -1->5->3->4->0
// 输出： -1->0->3->4->5

// 思路： 插入排序
//
//		我们新建一个 headPre 结点，
//		然后不停从链表中拿头结点 head ，直至链表为空
//		然后找到到第一个值 大于等于 head 的结点的前一个结点 cur ，
//		将 head 插入到 cur 之后
//
//		时间复杂度： O(n ^ 2)
//		空间复杂度： O(1)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
	headPre := new(ListNode)
	for ; head != nil; {
		// 提前保存下一个结点
		next := head.Next

		// 找到第一个值 大于等于 head 的结点的前一个结点
		cur := headPre
		for ; cur.Next != nil && cur.Next.Val < head.Val; {
			cur = cur.Next
		}
		// 将 head 插入到 cur 之后
		cur.Next, head.Next = head, cur.Next

		// 处理下一个结点
		head = next
	}
	return headPre.Next
}
