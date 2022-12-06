// 链接：https://leetcode.com/problems/middle-of-the-linked-list/
// 题意：给定一个单链表，返回中间结点。
//      如果存在两个中间结点，则返回第二个。


// 数据限制：
//  链表的结点数范围为 [1, 100]
//  1 <= Node.val <= 100


// 输入： head = [1,2,3,4,5]
// 输出： [3,4,5]
// 解释： 3 是链表的中间结点。

// 输入： head = [1,2,3,4,5,6]
// 输出： [4,5,6]
// 解释： 链表有两个中间结点 3 和 4 ，
//       返回第二个中间结点 4 。


// 思路： 快慢指针/双指针
//
//      本题是 LeetCode 19 的加强版，需要找到中间结点（与链表长度相关）。
//
//
//      初始化快慢指针均为头结点。
//
//      如果快指针还能走两步，则继续循环处理：快指针走两步，慢指针走一步。
//
//      结束循环时，慢指针就指向中间结点。
//      （如果链表结点数是偶数，那么慢指针必定是第二个中间结点）
//
//
//      时间复杂度：O(n)
//          1. 需要遍历全部 O(n) 个结点
//      空间复杂度：O(1)
//          1. 只需要维护常数个额外变量即可      


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    // 初始化快慢指针均为头结点
    fast := head
    slow := head
    // 如果快指针还能走两步，则继续循环
    for fast != nil && fast.Next != nil {
        // 快指针每次走两步
        fast = fast.Next.Next
        // 慢指针每次走一步
        slow = slow.Next
    }
    // 此时慢指针就指向中间结点
    //（如果链表结点数是偶数，那么慢指针必定是第二个中间结点）
    return slow
}
