// 链接：https://leetcode.com/problems/remove-linked-list-elements/
// 题意：给定一个链表 head 和一个整数 val ，删除所有值等于 val 的结点。


// 数据限制：
//  链表的结点数范围为 [0, 10 ^ 4]
//  1 <= Node.val <= 50
//  0 <= val <= 50


// 输入： head = [1,2,6,3,4,5,6], val = 6
// 输出： [1,2,3,4,5]

// 输入： head = [], val = 1
// 输出： []

// 输入： head = [7,7,7,7], val = 7
// 输出： []


// 思路：模拟
//
//      对于链表的题目，一般都可以使用一个哨兵结点。
//      本题使用哨兵结点，方便处理删除头结点这种边界情况。
//
//      要删除链表的一个结点 a ，那么需要找到其前驱 b ，
//      令 b.next = a.next 即可。
//
//      我们可以初始化当前结点 cur 为哨兵结点 dummy ，
//      当 cur.next 存在时，不断进行以下处理：
//          1. cur.next.val == val: 那么需要删除 cur.next ，
//              即 cur.next = cur.next.next
//          2. cur.next.val != val: 那么只需要向后移动 cur ，
//              即 cur = cur.next
//
//
//      时间复杂度： O(n)
//          1. 需要遍历全部 O(n) 个结点
//      空间复杂度： O(1)
//          1. 只需要维护常数个额外变量即可


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
    pub fn remove_elements(head: Option<Box<ListNode>>, val: i32) -> Option<Box<ListNode>> {
        // 定义一个哨兵结点，方便处理删除头结点这种边界情况
        let mut dummy = Some(Box::new(ListNode{ val: 0, next: head }));
        // 初始化当前结点为哨兵结点
        let mut cur = &mut dummy;
        // 如果下一个结点存在，则可以继续处理
        while let Some(next_node) = &mut cur.as_mut().unwrap().next {
            if next_node.val == val {
                // 如果下一个结点是需要删除的结点，则 cur 的 next 指向下一个结点的 next
                cur.as_mut().unwrap().next = next_node.next.take();
                // 此时，不需要再移动结点，删除相当于往后移动了
            } else {
                // 如果下一个结点不是需要删除的结点，则 cur 移向下一个结点
                cur = &mut cur.as_mut().unwrap().next;
            }
        }

        dummy.unwrap().next
    }
}
