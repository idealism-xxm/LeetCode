// 链接：https://leetcode.com/problems/merge-two-sorted-lists/
// 题意：给定两个升序排序的链表 list1 和 list2 ，
//      将它们合并成一个升序链表并返回。


// 数据限制：
//  两个链表中的结点数均在 [0, 50] 内
//  -100 <= Node.val <= 100
//  list1 和 list2 均为非降序链表


// 输入： list1 = [1,2,4], list2 = [1,3,4]
// 输出： [1,1,2,3,4,4]
// 解释： 1 -> 2 -> 4 | merge | 1 -> 3 -> 4
//                       ↓
//           1 -> 1 -> 2 -> 3 -> 4 -> 4

// 输入： list1 = [], list2 = []
// 输出： []
// 解释： (empty) | merge | (empty)
//                   ↓
//                (empty)

// 输入： list1 = [], list2 = [0]
// 输出： [0]
// 解释： (empty) | merge | 0
//                   ↓
//                   0


// 思路： 模拟
//
//      由于两个链表已经是升序的，所以可以按题意直接模拟处理。
//      当 list1 和 list2 均还有结点时，
//      取它们中较小的头结点放入结果链表中，然后不断循环。
//
//      最后当其中一个链表为空时，
//      将另一个链表剩余的部分全部插入结果链表尾部即可。
//
//      进阶：
//          1. LeetCode 23: 合并 k 个有序链表
//          2. LeetCode 148: 对无序链表排序
//
//		时间复杂度： O(|list1| + |list2|)
//          1. 需要遍历 list1 中的全部 O(|list1|) 个结点
//          2. 需要遍历 list2 中的全部 O(|list2|) 个结点
//		空间复杂度： O(1)
//          1. 只需要维护常数个额外变量


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
    pub fn merge_two_lists(mut list1: Option<Box<ListNode>>, mut list2: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        // 使用一个哨兵结点 head_pre ，方便后续处理
        let mut head_pre = ListNode::new(0);
        // 使用尾插法，所以需要尾部结点的引用
        let mut tail = &mut head_pre;
        // 如果两个链表都还有结点，则继续处理
        while list1.is_some() && list2.is_some() {
            if list1.as_ref().unwrap().val <= list2.as_ref().unwrap().val {
                // 如果 list1 的值更小，则将 list1 放入结果链表中
                tail.next = list1;
                // 移动 list1 到下一个结点，从 tail.next 中取出 list1 剩余的结点
                list1 = tail.next.as_mut().unwrap().next.take();
            } else {
                // 如果 list2 的值更小，则将 list2 放入结果链表中
                tail.next = list2;
                // 移动 list2 到下一个结点，从 tail.next 中取出 list2 剩余的结点
                list2 = tail.next.as_mut().unwrap().next.take();
            }
            // 移动结果链表的 tail 到下一个结点
            tail = tail.next.as_mut().unwrap();
        }

        if list1.is_some() {
            // 如果 list1 还有结点，表明 list2 已遍历完成，
            // 则将 list1 直接放在 tail 后面
            tail.next = list1;
        } else {
            // 如果 list1 没有结点，表明 list1 已遍历完成，
            // 则将 list2 直接放在 tail 后面
            tail.next = list2;
        }

        // 返回合并后的链表的头结点
        head_pre.next
    }
}
