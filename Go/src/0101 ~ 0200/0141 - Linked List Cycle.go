// 链接：https://leetcode.com/problems/linked-list-cycle/
// 题意：给定一个链表，判断是否存在环？

// 输入： head = [3,2,0,-4], pos = 1
// 输出： true

// 输入： head = [1,2], pos = 0
// 输出： true

// 输入： head = [1], pos = -1
// 输出： false

// 思路： 快慢指针
//
//		维护两个指针 fast 和 slow ，每次循环时 fast 往前走两步， slow 往前走一步
//		若 fast 遇到 nil 则链表无环
//		若 fast 遇到 slow 则链表有环
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
func hasCycle(head *ListNode) bool {
	// 如果头结点为 nil 或者 头结点的下一个结点为 nil ，则不存在环
	if head == nil || head.Next == nil {
		return false
	}
	fast, slow := head.Next, head
	for ; fast != slow; {
		// 如果快指针为 nil 或者 快指针的下一个结点为 nil ，则不存在环
		if fast == nil || fast.Next == nil {
			return false
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	// 此时必有 fast = slow ，即存在环
	return true
}
