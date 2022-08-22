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
    pub fn is_palindrome(mut head: Option<Box<ListNode>>) -> bool {
        // 1. 先用快慢指针找到后半部分
        let mut fast = &head;
        let mut slow = &head;
        while fast.is_some() {
            // 快指针每次走两步
            fast = &fast.as_ref().unwrap().next;
            if let Some(fast_node) = fast.as_ref() {
                fast = &fast_node.next;
            }
            // 慢指针每次走一步
            slow = &slow.as_ref().unwrap().next;
        }
        // 此时慢指针就指向后半部分的头结点
        //（如果链表结点数是奇数，那么此时必定是正中间结点的后一个）
        let slow = slow.clone();

        // 2. 翻转后半部分
        let slow = Self::reverse_list(slow);

        // 3. 对比前后是否相等即可
        let mut l = &head;
        let mut r = &slow;
        while l.is_some() && r.is_some() {
            let l_node = l.as_ref().unwrap();
            let r_node = r.as_ref().unwrap();
            // 如果值不相等，则必定不是回文链表，直接返回 false
            if l_node.val != r_node.val {
                return false;
            }
            // 如果值相等，则都移动至下一个结点继续对比
            l = &l_node.next;
            r = &r_node.next;
        }

        // 所有值都相等，则是回文链表
        true
    }

    pub fn reverse_list(mut head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        // 使用头插法翻转链表
        let mut head_pre = Some(Box::new(ListNode::new(0)));
        // 若 head 不是空结点，则继续处理
        while head.is_some() {
            // 先获取下一个结点
            let next = head.as_mut().unwrap().next.take();
            // 再将 head 用头插法放入结果链表中
            head.as_mut().unwrap().next = head_pre.as_mut().unwrap().next.take();
            head_pre.as_mut().unwrap().next = head;
            // 接下来处理下一个结点
            head = next;
        }

        // 返回翻转后链表的头结点
        head_pre.unwrap().next.take()
    }
}
