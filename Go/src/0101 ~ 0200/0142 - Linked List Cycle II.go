// 链接：https://leetcode.com/problems/linked-list-cycle-ii/
// 题意：给定一个链表，判断是否存在环？

// 输入： head = [3,2,0,-4], pos = 1
// 输出： 1
// 解释： 尾结点连上第二个结点

// 输入： head = [1,2], pos = 0
// 输出： 0
// 解释： 尾结点连上第一个结点

// 输入： head = [1], pos = -1
// 输出： -1
// 解释： 不存在环

// 思路1： 快慢指针 + 二分
//
//		还是使用 0141 的快慢指针方式，
//		若初始状态下没有环，则直接返回 nil
//		否则找到出现环的第一个结点设为 tail，
//		则：形成环的结点必定在 [head, tail] 之间，
//		对其进行二分，每次判断以 mid 开始的链表中， mid 是否在环上
//			若不存在，则开始环的结点必定在 (mid, tail] 之间，令 head = mid.Next
//			若存在，则开始环的结点必定在 [head, mid] 之间，令 tail = mid
//		当 head == tail 时， head 就是进入环的第一个结点
//
//		时间复杂度： O(nlogn)
//		空间复杂度： O(1)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	tail := getTail(head)
	// 如果无环，直接返回 nil
	if tail == nil {
		return tail
	}
	// 如果有环，则对 [head, tail] 进行二分
	// 统计每次 [head, tail] 的链表长度，方便后续二分找到 mid
	length := 1
	for cur := head; cur != tail; cur = cur.Next {}
	for ; head != tail; {
		// 找到 mid
		mid := head
		for i := (length - 1) >> 1; i > 0; i-- {
			mid = mid.Next
		}

		if isHeadInCircle(mid) {
			// 如果 mid 在环中，则开始环的结点必定在 [head, mid] 之间，令 tail = mid
			tail = mid
			length = ((length - 1) >> 1) + 1
		} else {
			// 则开始环的结点必定在 (mid, tail] 之间，令 head = mid.Next
			head = mid.Next
			length = length - ((length - 1) >> 1) - 1
		}
	}
	return head
}

// 判断 head 是否在环中
func isHeadInCircle(head *ListNode) bool {
	fast, slow := head.Next, head
	for ; fast != slow; {
		// 如果快指针为 head 或者 快指针的下一个结点为 head ，则 head 在环中
		if fast == head || fast.Next == head {
			return true
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	// 此时若 fast == head ，则 head 在环中
	return fast == head
}

// 有环时，返回 fast 和 slow 遇见时的结点；无环时，返回 nil
func getTail(head *ListNode) *ListNode {
	// 如果头结点为 nil 或者 头结点的下一个结点为 nil ，则不存在环
	if head == nil || head.Next == nil {
		return nil
	}
	fast, slow := head.Next, head
	for ; fast != slow; {
		// 如果快指针为 nil 或者 快指针的下一个结点为 nil ，则不存在环
		if fast == nil || fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	// 此时必有 fast = slow ，即存在环
	return fast
}


// 思路2： 快慢指针 + Floyd
//
//		还是题解总能找到各种隐藏的关系，让人大开眼界
//		找到一个数学证明的题解
//		https://leetcode-cn.com/problems/linked-list-cycle-ii/solution/shuang-zhi-zhen-qing-xi-ti-jie-zhen-zheng-cong-shu/
//
//		假设非环部分的长度是 x ，从环起点到相遇点的长度是 y ，环的长度是 c
//		则相遇时，慢指针走过的长度为 x + n1 * c + y
//		快指针的速度是慢指针的速度的两倍，所以快指针走过的长度为 2 * (x + n1 * c + y)
//
//		由于相遇时，快指针比慢指针多走的长度必定是环的长度的整数倍，
//		则有： (2 * (x + n1 * c + y)) - (x + n1 * c + y) = n2 * c
//		化简可得： x + n1 * c + y = n2 * c
//		即： x + y = (n2 - n1) * c
//		即： 非环部分的长度 + 环起点到相遇点的长度 = 环长度的整数倍
//		所以我们只要再走 x 就能从相遇点到达环起点，
//		此时让快指针从 head 开始，并且每次只走一步，那么它们就会在环起点处相遇
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
func detectCycle(head *ListNode) *ListNode {
	// 如果头结点为 nil 或者 头结点的下一个结点为 nil ，则不存在环
	if head == nil || head.Next == nil {
		return nil
	}
	// 注意快慢指针必须从同一个结点开始走
	fast, slow := head.Next.Next, head.Next
	for ; fast != slow; {
		// 如果快指针为 nil 或者 快指针的下一个结点为 nil ，则不存在环
		if fast == nil || fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		slow = slow.Next
	}

	// 快指针从头开始，并且每次走一步
	fast = head
	for ; fast != slow; {
		fast = fast.Next
		slow = slow.Next
	}
	// 此时相遇在环起点
	return fast
}
