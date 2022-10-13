// 链接：https://leetcode.com/problems/delete-node-in-a-linked-list/
// 题意：有一个单链表，只给定待删除的结点（非尾结点），
//      将链表变成删除该结点后的链表？


// 数据限制：
//  链表中的结点数量为 [2, 1000]
//  -1000 <= Node.val <= 1000
//  每个结点的值各不相同
//  待删除的结点不是尾结点


// 输入： head = [4,5,1,9], node = 5
// 输出： [4,1,9]
// 解释：  4 -> 5 -> 1 -> 9
//               ↓
//        4   ->    1 -> 9

// 输入： head = [4,5,1,9], node = 1
// 输出： [4,5,9]
// 解释：  4 -> 5 -> 1 -> 9
//               ↓
//        4 -> 5    ->   9


// 思路： 模拟
//
//      由于给定的是待删除结点，所以无法直接删除该结点。
//      
//      但限制了删除的不是尾结点，所以我们可以将下一个结点的值赋给当前结点，
//      再删除下一个结点即可。
//
//
//      时间复杂度： O(1)
//          1. 只需要常数次操作即可
//      空间复杂度： O(1)
//          1. 只需要维护常数个额外变量即可


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {
	// 将下一个结点的值赋给当前结点
	node.Val = node.Next.Val
	// 删除下一个结点
	node.Next = node.Next.Next   
}
