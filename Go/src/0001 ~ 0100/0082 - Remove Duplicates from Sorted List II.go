// 链接：https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/
// 题意：给定一个升序的整型单链表，若一个数字重复出现，则删除所有该数字，
//		返回剩余结点单单链表头？

// 输入：1->2->3->3->4->4->5
// 输出：1->2->5

// 输入：1->1->1->2->3
// 输出：2->3

// 思路：模拟即可
//		维护当前指针和其前一指针，并记录前一指针指向值出现的次数
//		当两个指针指向的值不同时，若出现次数为 1 ，则可放入结果链表
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
	headPre := &ListNode{
		Val:  0,
		Next: nil,
	}
	tail := headPre
	pre, cur := head, head.Next
	preCount := 1
	for ; cur != nil; pre, cur = cur, cur.Next {
		if pre.Val == cur.Val {  // 如果当前值和前一个值一样，则次数 + 1
			preCount++
		} else {
			if preCount == 1 {  // 如果前一个值只出现一次，则放入结果链表
				tail.Next = pre
				pre.Next = nil  // 最新加入的结点是尾结点
				tail = pre
			}
			preCount = 1  // 重置前一个值出现次数为 1
		}
	}
	if preCount == 1 {  // 如果最后一个值也值出现一次，则放入结果链表
		tail.Next = pre
		pre.Next = nil  // 最新加入的结点是尾结点
	}
	return headPre.Next
}