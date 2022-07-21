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
//          1. 先找到第 left 个结点的前一个结点 first_tail ，
//              并记录第二部分翻转后的尾部结点 second_tail = first_tail.next
//          2. 再将接下来 right - left + 1 个结点用头插法插入到 first_tail 后面
//          3. 最后将剩余部分挂在第二部分翻转后的尾部结点 second_tail 后面即可
//
//      时间复杂度： O(n)
//          1. 只需要遍历链表全部 O(n) 个结点一次
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
    pub fn reverse_between(head: Option<Box<ListNode>>, left: i32, right: i32) -> Option<Box<ListNode>> {
        // 定义一个哨兵结点，方便后续处理
        let mut head_pre = Some(Box::new(ListNode{ val: 0, next: head }));
        // 先找到第 m 个结点的前一个结点，即第一部分的尾部结点
        let mut first_tail = &mut head_pre;
        for _ in 1..left {
            first_tail = &mut first_tail.as_mut().unwrap().next;
        }
        // 将接下来 right - left + 1 个结点用头插法插入到 first_tail 后面
        let mut cur = first_tail.as_mut().unwrap().next.take();
        for _ in 0..right - left + 1 {
            // 先保存下一个结点
            let next = cur.as_mut().unwrap().next.take();
            // 将当前结点插入到 first_tail 后面
            cur.as_mut().unwrap().next = first_tail.as_mut().unwrap().next.take();
            first_tail.as_mut().unwrap().next = cur;

            cur = next
        }
        // 接着找到第二部分的尾部结点
        //  （由于 Rust 中只能同时存在一个 mutable borrow ，
        //  所以我们无法在前面提前保存第二部分的尾部结点）
        let mut second_tail = first_tail;
        while second_tail.as_mut().unwrap().next.is_some() {
            second_tail = &mut second_tail.as_mut().unwrap().next;
        }
        // 将剩余部分挂在第二部分翻转后的尾部结点 second_tail 后面
        second_tail.as_mut().unwrap().next = cur;

        head_pre.unwrap().next
    }
}
