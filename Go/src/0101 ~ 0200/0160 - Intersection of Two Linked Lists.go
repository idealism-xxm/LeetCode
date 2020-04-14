// 链接：https://leetcode.com/problems/intersection-of-two-linked-lists/
// 题意：给定两个链表，找到这两个链表第一个相交的结点 ？

// 输入：
//	      a1 -> a2 ↘
//	                 c1 -> c2 -> c3
//	b1 -> b2 -> b3 ↗
// 输出： c1 -> c2 -> c3

// 输入：
//	      a1 -> a2
//	b1 -> b2 -> b3
// 输出： nil

// 思路： 双指针
//
//		先用双指针同时遍历两个链表，直到其中一个遍历完毕
//		若还有链表未遍历完，则其比较长，对应的 head 需要提前走，
//		保证 headA 和 headB 剩余的结点数一样，
//		然后再同时走，遇见同一个结点，则是第一个相交的结点，否则不相交
//
//		时间复杂度： O(m + n)
//		空间复杂度： O(1)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 用双指针遍历完其中一个
	a, b := headA, headB
	for ; a != nil && b != nil; {
		a = a.Next
		b = b.Next
	}
	// 如果 a 还未遍历完，则 headA 要提前走几步，保证剩余结点数一样
	for ; a != nil; {
		a = a.Next
		headA = headA.Next
	}
	// 如果 b 还未遍历完，则 headB 要提前走几步，保证剩余结点数一样
	for ; b != nil; {
		b = b.Next
		headB = headB.Next
	}
	// 同步走，如果遇到相同的结点，则它们在此处相交
	for ; headA != nil && headB != nil; {
		if headA == headB {
			return headA
		}
		headA = headA.Next
		headB = headB.Next
	}
	// 没有遇到相同的结点，则没有相交
	return nil
}
