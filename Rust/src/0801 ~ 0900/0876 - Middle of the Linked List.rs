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
//          2. Rust 中需要 clone 链表后半部分的 O(n) 个结点
//      空间复杂度：O(n)
//          1. 只需要维护常数个额外变量即可
//          2. Rust 中需要 clone 链表后半部分的 O(n) 个结点


// Definition for singly-linked list.
// #[derive(PartialEq, Eq, Clone, Debug)]
// pub struct ListNode {
//   pub val: i32,
//   pub next: Option<Box<ListNode>>
// }
// 
// impl ListNode {
//   #[inline]
//   fn new(val: i32) -> Self {
//     ListNode {
//       next: None,
//       val
//     }
//   }
// }
impl Solution {
    pub fn middle_node(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        // 初始化快慢指针均为头结点
        let mut fast = &head;
        let mut slow = &head;
        // 如果快指针还能走两步，则继续循环
        while fast.is_some() && fast.as_ref().unwrap().next.is_some() {
            // 快指针每次走两步
            fast = &fast.as_ref().unwrap().next.as_ref().unwrap().next;
            // 慢指针每次走一步
            slow = &slow.as_ref().unwrap().next;
        }
        // 此时慢指针就指向中间结点
        //（如果链表结点数是偶数，那么慢指针必定是第二个中间结点）
        slow.clone()
    }
}
