// 链接：https://leetcode.com/problems/remove-linked-list-elements/
// 题意：给定一个链表，移除所有值等于给定值的节点 ？

// 输入： 1->2->6->3->4->5->6, val = 6
// 输出： 1->2->3->4->5

// 思路：模拟
//
//		先定义一个 head_pre 节点， next 指向 head
//		然后循环直到 pre 的 next 是 None
//			若 pre.next.val == val ，则 pre.next 指向 pre.next.next
//			否则， pre 指向 pre.next
//
//		时间复杂度： O(n)
//		空间复杂度： O(1)

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
		// 定义一个 head_pre 节点，方便后续操作
		let mut head_pre = Some(Box::new(
			ListNode {
				val: 0,
				next: head,
			}
		));
		let mut pre = &mut head_pre;
		// 当前节点不是空节点时，则可以继续处理
		while let Some(cur_node) = &mut pre.as_mut().unwrap().next {
			if cur_node.val == val {
				// 如果当前节点是需要移除的节点，则 pre 的 next 指向 cur_node 的 next
				pre.as_mut().unwrap().next = cur_node.next.take();
				// 不需要再移动节点，移除相当于往后移动了
			} else {
				// 如果当前节点不是需要移除的节点，则 pre 移向下一个节点
				pre = &mut pre.as_mut().unwrap().next;
			}
		}

		head_pre.unwrap().next
	}
}
