// 链接：https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/
// 题意：给定一个升序的整型单链表，若一个数字重复出现，
//      则删除所有该数字，最后返回剩余结点的单链表。


// 数据限制：
//  链表中的结点数在 [0, 300] 内
//  -100 <= Node.val <= 100
//  链表是升序的


// 输入： head = [1,2,3,3,4,4,5]
// 输出： [1,2,5]
// 解释： 链表中只有 1, 2, 5 这个三个数字只出现过一次，
//       所以最后只保留这三个数对应的结点
//       1 -> 2 -> 3 -> 3 -> 4 -> 4 -> 5
//                      ↓
//                 1 -> 2 -> 5

// 输入： head = [1,1,1,2,3]
// 输出： [2,3]
// 解释： 链表中只有 2, 3 这个两个数字只出现过一次，
//       所以最后只保留这两个数对应的结点
//       1 -> 1 -> 1 -> 2 -> 3
//                 ↓
//               2 -> 3


// 思路： 模拟
//
//		使用 pre 记录当前结点 cur 的前一个结点，
//      并用 is_pre_duplicated 记录前一个结点的值 pre.val 是否重复。
//
//      当原始链表还有结点时，不断循环处理：
//          1. cur.val == pre.val: 前一个结点的值必定重复，
//              标记 is_pre_duplicated = true
//          2. cur.val != pre.val: 如果此时 is_pre_duplicated 是 false ，
//              就说明 pre 的值不重复，加入到结果链表中。
//
//              然后重新记录前一个结点的信息： pre = cur; is_pre_deplicated = false
//
//      结束循环后，最后一个结点可能还没有处理，要再判断 is_pre_duplicated 的值。
//      如果 is_pre_duplicated 是 false ，
//      则最后一个结点的值不重复，加入到结果链表中
//
//		时间复杂度： O(n)
//          1. 只需要遍历链表中全部 O(n) 个结点一次
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
    pub fn delete_duplicates(mut head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        // 如果是空链表，则直接返回 None
        if head.is_none() {
            return None
        }

        // 定义一个哨兵结点，方便后续处理
        let mut head_pre = ListNode::new(0);
        // 定义结果链表的尾结点，方便实用尾插法插入结点
        let mut tail = &mut head_pre;

        // 定义前一个结点，用于统计是否重复
        let mut pre = head;
        // 原始链表头结点移动到下一个结点
        head = pre.as_mut().unwrap().next.take();
        // 记录前一个结点的值是否已经重复
        let mut is_pre_duplicated = false;

        // 当原始链表还有结点的时候，循环处理
        while head.is_some() {
            // 取下当前链表的头结点
            let mut cur = head;
            // 将 head 向后移动一个结点
            head = cur.as_mut().unwrap().next.take();

            if cur.as_mut().unwrap().val == pre.as_mut().unwrap().val {
                // 如果当前结点的值 cur.val 和前一个结点的值 pre.val 相同，
                // 标记其已重复
                is_pre_duplicated = true;
            } else {
                // 此时，当前结点的值 cur.val 和前一个结点的值 pre.val 不同

                // 如果前一个结点的值未重复
                if !is_pre_duplicated {
                    // 将 pre 插入结果链表尾部
                    tail.next = pre;
                    // 结果链表尾结点向后移动一个结点
                    tail = tail.next.as_mut().unwrap();
                }

                // 修改前一个结点为 pre
                pre = cur;
                // 标记其未重复
                is_pre_duplicated = false;
            }
        }

        // 如果最后一个结点的值未重复
        if !is_pre_duplicated {
            // 将 pre 插入结果链表尾部
            tail.next = pre;
            // 由于已是最后一次插入，所以无需再移动尾结点
        }
        
        // 返回结果链表的头结点
        head_pre.next
    }
}
