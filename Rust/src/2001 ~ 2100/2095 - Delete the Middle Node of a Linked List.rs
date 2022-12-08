// 链接：https://leetcode.com/problems/delete-the-middle-node-of-a-linked-list/
// 题意：给定一个单链表，删除中间结点。
//      如果存在两个中间结点，则删除第二个。


// 数据限制：
//  链表的结点数范围为 [1, 10 ^ 5]
//  1 <= Node.val <= 10 ^ 5


// 输入： head = [1,3,4,7,1,2,6]
// 输出： [1,3,4,1,2,6]
// 解释： 7 是链表的中间结点。

// 输入： head = [1,2,3,4]
// 输出： [1,2,4]
// 解释： 链表有两个中间结点 2 和 3 ，
//       删除第二个中间结点 3 。

// 输入： head = [2,1]
// 输出： [2]
// 解释： 链表有两个中间结点 2 和 1 ，
//       删除第二个中间结点 1 。


// 思路： 快慢指针/双指针
//
//      本题是 LeetCode 876 和 LeetCode 203 的加强版，数据范围加大，并需要删除中间结点。
//
//      对于链表的题目，一般都可以使用一个哨兵结点。
//      本题使用哨兵结点方，便处理删除头结点这种边界情况。
//
//      因为本题需要删除中间结点，所以要找到待删除结点的前一个结点。
//
//      初始化快指针为头结点，慢指针为哨兵结点
//
//      如果快指针还能走两步，则继续循环处理：快指针走两步，慢指针走一步。
//
//      结束循环时，慢指针就指向中间结点的前一个结点。
//      （如果链表结点数是偶数，那么慢指针必定是第一个中间结点）
//
//      删除中间结点后返回。
//
//
//      时间复杂度：O(n)
//          1. 需要遍历全部 O(n) 个结点
//          2. Rust 中需要 clone 原链表 O(n) 个结点
//      空间复杂度：O(n)
//          1. 只需要维护常数个额外变量即可
//          2. Rust 中需要 clone 原链表 O(n) 个结点    


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
    pub fn delete_middle(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        // 初始化哨兵结点，方便处理删除头结点这种边界情况。
        let mut dummy = Some(Box::new(ListNode{ val: 0, next: head}));
        // 初始化快指针为头结点，慢指针为哨兵结点
        // （由于 Rust 中 mutable borrow 和 borrow 不能同时存在，
        //  所以我们只能克隆原链表用于快指针遍历）
        let mut fast = &(dummy.as_ref().unwrap().next.clone());
        let mut slow = &mut dummy;
        // 如果快指针还能走两步，则继续循环
        while fast.is_some() && fast.as_ref().unwrap().next.is_some() {
            // 快指针每次走两步
            fast = &fast.as_ref().unwrap().next.as_ref().unwrap().next;
            // 慢指针每次走一步
            slow = &mut slow.as_mut().unwrap().next;
        }
        // 慢指针就指向中间结点的前一个结点，删除中间结点
        //（如果链表结点数是偶数，那么慢指针必定是第一个中间结点）
        slow.as_mut().unwrap().next = slow.as_mut().unwrap().next.as_mut().unwrap().next.take();

        dummy.as_mut().unwrap().next.take()
    }
}
