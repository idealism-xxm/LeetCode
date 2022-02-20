// 链接：https://leetcode.com/problems/merge-k-sorted-lists/
// 题意：给定 k 个升序链表，把它们合成一个升序链表。

// 数据限制：
//  k == lists.length
//  0 <= k <= 10 ^ 4
//  0 <= lists[i].length <= 500
//  - (10 ^ 4) <= lists[i][j] <= 10 ^ 4
//  lists[i] 是升序排序的
//  lists[i].length 的和不超过 10 ^ 4

// 输入：lists = [[1,4,5],[1,3,4],[2,6]]
// 输出：[1,1,2,3,4,4,5,6]
// 解释：原始链表为：
//          [
//            1->4->5,
//            1->3->4,
//            2->6
//          ]
//      合成后的链表为：
//          1->1->2->3->4->4->5->6

// 输入：lists = []
// 输出：[]

// 输入：lists = [[]]
// 输出：[]

// 思路1：优先队列（堆）
//
//      我们回想合并两个链表时的场景，每次对比两个链表头结点大小，
//      将较小的结点加入结果链表中，然后后移一个，直至所有结点都在结果链表中。
//
//      那么合并多个链表时也是如此，不过由于存在多个可比较的结点，
//      可以使用最小堆将取最小结点的时间复杂度从 O(k) 降低为 O(logk) 。
//
//      把所有链表的头结点放入堆，每次取最小的结点出来插入结果链表，
//      再将其后的结点放入堆中，直至所有结点都在结果链表中。
//
//
//      设 n 表示所有链表总长度， k 表示升序链表个数
//
//      时间复杂度：O(nlogk)
//      空间复杂度：O(n + k) 。每个结果链表中的结点都是新建的，共有 n 个；同时最小堆最多有 k 个结点


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

use std::{collections::BinaryHeap, cmp::Reverse};

impl Ord for ListNode {
    fn cmp(&self, other: &Self) -> std::cmp::Ordering {
        self.val.cmp(&other.val)
    }
}

impl PartialOrd for ListNode {
    fn partial_cmp(&self, other: &Self) -> Option<std::cmp::Ordering> {
        self.val.partial_cmp(&other.val)
    }
}

impl Solution {
    pub fn merge_k_lists(lists: Vec<Option<Box<ListNode>>>) -> Option<Box<ListNode>> {
        // 初始化一个最小堆，把非空链表的头结点放入
        let mut heap = BinaryHeap::new();
        for list in lists {
            if let Some(list) = list {
                // 由于 BinaryHeap 是最大堆，所以对所有结点使用 Reverse 包一层
                heap.push(Reverse(list));
            }
        }

        // 定义一个哨兵结点，方便处理空链表的情况
        let mut head_pre: Box<ListNode> = Box::new(ListNode::new(0));
        // 定义结果链表的尾结点，方便后续加入新结点
        let mut tail = &mut head_pre;
        // 不停从最小堆中获取结点，直至最小堆为空
        while let Some(Reverse(cur)) = heap.pop() {
            // 以当前结点 cur 的值创建新结点，加入到结果链表尾部
            tail.next = Some(Box::new(ListNode::new(cur.val)));
            // 将尾结点往后移动一个
            tail = tail.next.as_mut().unwrap();
            // 如果当前结点 cur 还有后续结点，则放入到最小堆中
            if let Some(next) = cur.next {
                heap.push(Reverse(next));
            }
        }

        // 哨兵结点的下一个结点就是结果链表的头结点
        head_pre.next
    }
}
