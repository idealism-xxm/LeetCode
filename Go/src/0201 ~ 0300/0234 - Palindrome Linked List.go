// 链接：https://leetcode.com/problems/palindrome-linked-list/
// 题意：给定一个单链表，判断是不是回文的？
//
//      进阶：使用时间复杂度为 O(n) 空间复杂度为 O(1) 的算法求解。


// 数据限制：
//  链表的结点数范围为 [1, 10 ^ 5]
//  0 <= Node.val <= 9


// 输入： head = [1,2,2,1]
// 输出： true

// 输入： head = [1,2]
// 输出： false


// 思路： 快慢指针
//
//      先用快慢指针找到后半部分，然后将后半部分翻转，再对比前后是否相等即可
//
//      时间复杂度： O(n)
//          1. 只需要遍历全部 O(n) 个结点
//      空间复杂度： O(1)
//          1. 只需要使用常数个额外变量即可


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    // 1. 先用快慢指针找到后半部分
    fast := head
    slow := head
    for fast != nil {
        // 快指针每次走两步
        fast = fast.Next
        if fast != nil {
            fast = fast.Next
        }
        // 慢指针每次走一步
        slow = slow.Next
    }
    // 此时慢指针就指向后半部分的头结点
    //（如果链表结点数是奇数，那么此时必定是正中间结点的后一个）

    // 2. 翻转后半部分
    slow = reverseList(slow)

    // 3. 对比前后是否相等即可
    l := head
    r := slow
    for l != nil && r != nil {
        // 如果值不相等，则必定不是回文链表，直接返回 false
        if l.Val != r.Val {
            return false
        }
        // 如果值相等，则都移动至下一个结点继续对比
        l = l.Next
        r = r.Next
    }

    // 所有值都相等，则是回文链表
    return true
}

func reverseList(head *ListNode) *ListNode {
    // 使用头插法翻转链表
    headPre := &ListNode{}
    // 若 head 不是空结点，则继续处理
    for head != nil {
        // 先获取下一个结点
        next := head.Next
        // 再将 head 用头插法放入结果链表中
        head.Next = headPre.Next
        headPre.Next = head
        // 接下来处理下一个结点
        head = next
    }

    // 返回翻转后链表的头结点
    return headPre.Next
}
