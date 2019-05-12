// 链接：https://leetcode.com/problems/add-two-numbers/
// 题意：给定两个非空链表表示的非负整数，从个位开始每个结点存储一位数，求两数之和，用相同方式表示。
// 思路：其实就是简易版的大数加法，直接模拟加法运算即可

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 初始化头结点（不存储数据）和前一个结点
	head := new(ListNode)
	pre := head
	// 进位初始化为 0
	carry := 0
	// 当两个链表都还有结点时，进行加法操作
	for ; l1 != nil && l2 != nil;  {
		// 计算当前位的值
		carry, pre.Next = addListNode(l1.Val, l2.Val, carry)
		// 移动结点
		pre = pre.Next

		// 移动链表结点
		l1 = l1.Next
		l2 = l2.Next
	}
	// 如果 l2 还有结点，则赋给 l1 继续操作
	if l2 != nil {
		l1 = l2
	}
	// 当 l1 还有结点时，进行加法操作（其实并不用执行，除了第一次会有进位影响，后面的都不进位）
	for ; l1 != nil;  {
		// 计算当前位的值
		carry, pre.Next = addListNode(l1.Val, 0, carry)
		// 移动结点
		pre = pre.Next

		// 移动链表结点
		l1 = l1.Next
	}
	// 还有进位时，再执行一次
	if carry == 1 {
		_, pre.Next = addListNode(0, 0, carry)
	}

	return head.Next
}

// 一位数加法，返回进位和当前位结果
func addListNode(a int, b int, carry int) (resultCarry int, result *ListNode) {
	// 执行加法
	val := a + b + carry
	// 重置进位为 0
	carry = 0
	// 如果需要进位，则进行相应处理
	if val >= 10 {
		val -= 10
		carry = 1
	}

	node := new(ListNode)
	node.Val = val
	// 返回结果
	return carry, node
}