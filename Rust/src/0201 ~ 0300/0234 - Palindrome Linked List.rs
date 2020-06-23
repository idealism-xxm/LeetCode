// 链接：https://leetcode.com/problems/palindrome-linked-list/
// 题意：给定一个单链表，判断是不是回文的？

// 输入： 1->2
// 输出： false

// 输入： 1->2->2->1
// 输出： true

// 思路： 快慢指针
//
//      先用快慢指针找到后半部分，然后将后半部分翻转，再对比前后是否相等即可
//
//      时间复杂度： O(n)
//      空间复杂度： O(1)

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
    pub fn is_palindrome(head: Option<Box<ListNode>>) -> bool {
        let mut head = head;
        // 1. 先用快慢指针找到后半部分
        let mut fast = &head;
        let mut slow = &head;
        while fast.is_some() {
            fast = &fast.as_ref().unwrap().next;
            if let Some(fast_node) = fast.as_ref() {
                fast = &fast_node.next;
            }
            slow = &slow.as_ref().unwrap().next;
        }
        let slow = slow.clone();

        // 2. 翻转后半部分
        let slow = Solution::reverse_list(slow);

        // 3. 对比前后是否相等即可
        let mut l = &head;
        let mut r = &slow;
        while l.is_some() && r.is_some() {
            let l_node = l.as_ref().unwrap();
            let r_node = r.as_ref().unwrap();
            // 如果值相等，则继续比较下一个
            if l_node.val == r_node.val {
                l = &l_node.next;
                r = &r_node.next;
                continue;
            }
            return false;
        }

        // 所有值都相等，则是回文链表
        true
    }

    pub fn reverse_list(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        // 让 pre 和 cur 可以修改
        let mut pre = None;
        let mut cur = head;
        // 若 cur 不是空节点，则还可以翻转 pre 和 cur
        while cur.is_some() {
            // 获取下一个节点
            let next = cur.as_mut().unwrap().next.take();
            // 翻转 pre 和 cur 的前后关系
            cur.as_mut().unwrap().next = pre;
            // 接下来翻转 cur 和 cur.next 的前后关系
            pre = cur;
            cur = next;
        }
        pre
    }
}
