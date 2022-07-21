// 链接：https://leetcode.com/problems/reverse-linked-list-ii/
// 题意：给定一个单链表，只遍历一次将位置在 [left, right] 内的结点翻转？


// 数据限制：
//  链表中的结点数为 n
//  1 <= n <= 500
//  -500 <= Node.val <= 500
//  1 <= left <= right <= n


// 输入：head = [1,2,3,4,5], left = 2, right = 4
// 输出：[1,4,3,2,5]
// 解释：1 -> (2 -> 3 -> 4) -> 5 -> NULL
//                 ↓
//      1 -> (4 -> 3 -> 2) -> 5 -> NULL

// 输入：head = [5], left = 1, right = 1
// 输出：[5]


// 思路：模拟
//
//      为了方便处理，我们在 head 前面添加一个哨兵结点，
//      然后直接按照题意模拟即可，对不同的三段结点分别处理：
//          1. 先找到第 left 个结点的前一个结点 firstTail ，
//              并记录第二部分翻转后的尾部结点 secondTail = firstTail.next
//          2. 再将接下来 right - left + 1 个结点用头插法插入到 firstTail 后面
//          3. 最后将剩余部分挂在第二部分翻转后的尾部结点 secondTail 后面即可
//
//      时间复杂度： O(n)
//          1. 只需要遍历链表全部 O(n) 个结点一次
//      空间复杂度： O(1)
//          1. 只需要使用常数个额外变量即可


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	// 定义一个哨兵结点，方便后续处理
	headPre := &ListNode{ Val: 0, Next: head }
	// 先找到第 m 个结点的前一个结点，即第一部分的尾部结点
	firstTail := headPre
	for i := 1; i < left; i++ {
		firstTail = firstTail.Next
	}
	// 记录第二部分翻转后的尾部结点
	secondTail := firstTail.Next
	// 将接下来 right - left + 1 个结点用头插法插入到 firstTail 后面
    cur := firstTail.Next
	for i := right - left; i >= 0; i-- {
		// 先保存下一个结点
		next := cur.Next
		// 将当前结点插入到 firstTail 后面
		cur.Next = firstTail.Next
		firstTail.Next = cur

		cur = next
	}
	// 将剩余部分挂在第二部分翻转后的尾部结点 secondTail 后面
	secondTail.Next = cur

	return headPre.Next
}
