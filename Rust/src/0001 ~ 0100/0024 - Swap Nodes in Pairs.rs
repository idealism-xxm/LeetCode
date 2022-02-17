// 链接：https://leetcode.com/problems/swap-nodes-in-pairs/
// 题意：给定一个链表，对链表中每两个结点进行交换，返回结果链表的头结点。


// 数据限制：
//  链表的结点数在 [0, 100] 中
//  0 <= Node.val <= 100


// 输入：head = [1,2,3,4]
// 输出：[2,1,4,3]

// 输入：head = []
// 输出：[]

// 输入：head = [1]
// 输出：[1]


// 思路：模拟
//
//		如果直接按照 2 个一组进行处理，那么非常简单，
//      只需要记录一个结点：待交换的两个结点的前一个结点 pre ，
//      然后每次通过 pre 找到 cur 和 next
//
//      交换时注意赋值顺序即可，然后往后移动 pre 两个结点，循环处理
//
//
//      但本题可以更通用地解答，即将 2 个一组变成 K 个一组进行处理，
//      本题中 K 一直是 2 ，针对其他题，不同的 K 可以动态变化。
//
//      针对结果链表，我们需要定义哨兵结点 head_pre 和 尾结点 tail 。
//      然后在外层循环遍历 head ，如果 head 非空，则还有结点需要处理。
//
//      在外层循环体中，我们维护分组链表的哨兵结点 group_head_pre ，
//      然后使用头插法，将 head 后续的 K 个结点放入到分组链表中。
//      （如果发现空结点，则说明到原链表尾部，直接跳出内层循环）
//
//      内层循环结束后，现在 group_head_pre 的下一个结点，
//      指向 K 个反转的分组链表的头结点。
//
//      然后我们用尾插法将分组链表插入到结果链表尾部，
//      最后将尾结点移动至结果链表的最新尾结点即可。
//
//
//      【进阶】 K 个一组反转链表？
//      这就是 LeetCode 25 这题，本解法可以直接使用。
//
//		时间复杂度： O(n)
//		空间复杂度： O(1)

const K: usize = 2;

// Definition for singly-linked list.
// #[derive(PartialEq, Eq, Clone, Debug)]
// pub struct ListNode {
//   pub val: i32,
//   pub next: Option<Box<ListNode>>
// }

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
    pub fn swap_pairs(mut head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        // 结果链表的哨兵结点，方便后续处理
        let mut head_pre = Some(Box::new(ListNode::new(0)));
        // 结果链表的尾结点的可变引用，方便将每一个分组链表用尾插法插入结果链表
        let mut tail = head_pre.as_mut().unwrap();
        // 当原链表还有结点时，循环处理
        while head.is_some() {
            // 分组链表的哨兵结点，方便后续处理
            let mut group_head_pre = Box::new(ListNode::new(0));
            // 需要对接下来的 K 个结点用头插法插入分组链表中
            for _ in 0..K {
                // 如果原链表还有结点时，将其用头插法插入分组链表中
                if let Some(mut cur) = head {
                    // 原链表头结点移动到下一个结点
                    head = cur.next;

                    // 头插法插入分组链表中
                    cur.next = group_head_pre.next.take();
                    group_head_pre.next = Some(cur);
                } else {
                    // 如果原链表已没有结点，则直接跳出循环
                    break;
                }
            }
            // 将分组链表用尾插法插入结果链表中
            tail.next = group_head_pre.next.take();
            // 将 tail 变为现在尾结点的可变引用
            while tail.next.is_some() {
                // 如果当前 tail 还有下一个结点，则变为下一个结点的可变引用
                tail = tail.next.as_mut().unwrap();
            }
        }

        // 哨兵结点的下一个结点就是结果链表的头结点
        head_pre.unwrap().next
    }
}
