// 链接：https://leetcode.com/problems/delete-node-in-a-linked-list/
// 题意：有一个单链表，只给定待删除待节点（非尾节点），
//		将链表变成删除该节点后的链表？

// 输入： head = [4,5,1,9], node = 5
// 输出： [4,1,9]
// 解释： 在调用我们的函数后，单链表将变为 4 -> 1 -> 9

// 输入： head = [4,5,1,9], node = 1
// 输出： [4,5,9]
// 解释： 在调用我们的函数后，单链表将变为 4 -> 5 -> 9

// 思路： 模拟
//
//		按照题意操作即可，由于删除的不是尾节点，
//     	所以我们可以将下一个节点的值赋给当前节点，并删除下一个节点即可
//
//      时间复杂度： O(1)
//      空间复杂度： O(1)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
