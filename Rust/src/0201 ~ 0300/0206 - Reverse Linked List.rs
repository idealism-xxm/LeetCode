// 链接：https://leetcode.com/problems/reverse-linked-list/
// 题意：翻转单链表 ？

// 输入： 1->2->3->4->5->NULL
// 输出： 5->4->3->2->1->NULL

// 思路1：递归
//
//		创建一个 reverse(pre, cur) 函数，用于递归处理
//		1. cur 是空节点， pre 就是新的头节点，直接返回 pre
//		2. cur 不是头节点，获取 next = cur.next ，
//			将 cur.next 置为 pre ，然后返回 reverse(cur, next)
//
//		时间复杂度： O(n)
//		空间复杂度： O(n)

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
	pub fn reverse_list(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
		Solution::reverse(None, head)
	}

	// 翻转 pre 和 cur 的前后关系，并返回新的链表头
	fn reverse(pre: Option<Box<ListNode>>, cur: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
		// 让 cur 可以修改
		let mut cur = cur;
		// 如果 cur 不是空节点，则先翻转 pre 和 cur 的前后关系
		// 然后递归调用翻转 cur 和 cur.next 的前后关系
		if cur.is_some() {
			// 获取下一个节点
			let next = cur.as_mut().unwrap().next.take();
			// 翻转 pre 和 cur 的前后关系
			cur.as_mut().unwrap().next = pre;
			// 递归调用翻转 cur 和 cur.next 的前后关系，并返回其链表头
			Solution::reverse(cur, next)
		} else {
			// cur 已是空节点，则 pre 就是链表头
			pre
		}
	}
}

// 思路2：循环
//
//		思路和递归一样，由于递归是尾递归，所以复制过来修改为循环即可
//
//		不停循环翻转 pre 和 cur 前后关系，直至 cur 是空节点
//		获取 next = cur.next ，将 cur.next 置为 pre ，
//		然后令 pre = cur, cur = next 即可
//
//		最后 pre 是新的头节点，返回 pre 节点即可
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
