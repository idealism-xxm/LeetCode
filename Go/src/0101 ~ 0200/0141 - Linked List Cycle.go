// 链接：https://leetcode.com/problems/linked-list-cycle/
// 题意：给定一个链表，判断是否存在环？


// 数据限制：
//  链表的结点数在 [0, 10 ^ 4] 内
//	-(10 ^ 5) <= Node.val <= 10 ^ 5
//	pos 是 -1 或者链表中合法的结点下标


// 输入： head = [3,2,0,-4], pos = 1
// 输出： true
// 解释： 链表有环，第 3 个结点会连接到第 1 个结点 
//       3 → 2 → 0 → (-4)
//           ↑         ↓
//           └────←────┘       

// 输入： head = [1,2], pos = 0
// 输出： true
// 解释： 链表有环，第 1 个结点会连接到第 0 个结点 
//       1 → 2
//       ↑   ↓
//       └─←─┘    

// 输入： head = [1], pos = -1
// 输出： false
// 解释： 链表只有一个结点，无环


// 思路： 双指针
//
//		维护两个指针 fast 和 slow ，
//      每次循环时 fast 往前走两步， slow 往前走一步：
//		  	1. 若 fast 遇到 nil ，则链表无环
//		    2. 若 fast 遇到 slow ，则链表有环
//
//
//		时间复杂度： O(n)
//          1. 需要遍历链表中的全部 O(n) 个结点
//		空间复杂度： O(1)
//          1. 只需要维护常数个额外变量


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	// 如果头结点为空 或者 头结点的下一个结点为空，
	// 则不存在环，直接返回 false
	if head == nil || head.Next == nil {
		return false
	}
	// 定义快指针 fast ，让其先走一步，方便后续判断
	fast := head.Next
	// 定义慢指针 slow ，初始化为头结点
	slow := head
	// 如果快指针 fast 不等于 慢指针 slow ，
	// 则需要继续循环处理
	for ; fast != slow; {
		// 如果快指针为空 或者 快指针的下一个结点为空，
		// 则不存在环，直接返回 false
		if fast == nil || fast.Next == nil {
			return false
		}

		// 快指针 fast 往前走两步
		fast = fast.Next.Next
		// 慢指针 slow 往前走一步
		slow = slow.Next
	}

	// 此时必有 fast == slow ，即存在环
	return true
}
