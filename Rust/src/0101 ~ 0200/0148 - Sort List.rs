// 链接：https://leetcode.com/problems/sort-list/
// 题意：对链表进行排序。
//       要求时间复杂度为 O(nlogn) ，空间复杂度为 O(1) 。


// 数据限制：
//   链表中的结点数范围是 [0, 5 * 10 ^ 4]
//   -(10 ^ 5) <= Node.val <= 10 ^ 5


// 输入： head = [4,2,1,3]
// 输出： [1,2,3,4]
// 解释： 4->2->1->3
//            ↓
//       1->2->3->4

// 输入： head = [-1,5,3,4,0]
// 输出： [-1,0,3,4,5]
// 解释： (-1)->5->3->4->0
//              ↓
//       (-1)->0->3->4->5

// 输入： head = []
// 输出： []


// 思路： 归并排序 + 倍增法
//
//		平均时间复杂度为 O(nlogn) 的只有希尔排序、堆排序、快速排序和归并排序，
//		而最快时间复杂度为 O(nlogn) 的只有堆排序、快速排序和归并排序
//
//      我们再具体分析一下：
//
//		1. 堆排序利用了数组可以 O(1) 查找元素性质，所以链表中无法使用。
//		2. 快速排序最差时间复杂度是 O(n ^ 2) ，
//          且每趟重排后两边的长度不确定，
//          所以很难转成空间复杂度为 O(1) 的迭代
//      3. 希尔排序可以进行空间复杂度为 O(1) 迭代处理，
//          但最坏时间复杂度为 O(n * (logn) ^ 2)
//		4. 归并排序每次归并后长度翻倍，能准确定位到每一段并进行处理，
//			所以可以使用倍增法转成空间复杂度为 O(1) 的迭代
//
//      归并排序 + 迭代法就是先按长度为 1 的区间划分，每两个区间进行一次合并，
//      这样就能保证所有长度为 2 的区间有序。
//
//      然后按长度为 2 的区间划分，每两个区间进行一次合并成长度为 4 的区间。
//
//      以此类推，直至区间长度 大于等于 链表长度，此时整个链表就是有序的。
//		
//
//		时间复杂度： O(nlogn)
//          1. 总共有 O(logn) 次整体的区间合并，长度分别为 1, 2, 4, ...
//          2. 每次整体的区间合并，都会遍历链表的全部 O(n) 个结点
//		空间复杂度： O(1)
//          1. 没有使用栈，只用了常数个变量


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
    pub fn sort_list(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        // length 表示链表的长度
        let mut length = 0;
        // 遍历链表，计算链表的长度
        let mut cur = &head;
        // 如果当前还有结点，则继续处理
        while let Some(node) = cur {
            // 链表长度 +1
            length += 1;
            // 移动到下一个结点
            cur = &node.next;
        }

        // 使用哨兵结点，方便后续操作
        let mut head_pre = Some(Box::new(ListNode{val: 0, next: head}));
        // 倍增法进行归并排序，初始区间长度为 1
        let mut interval = 1;
        // 至少还有两个区间时，每两个区间归并一次
        while interval < length {
            // tail 表示已经归并完成的区间的尾结点，初始化为 head_pre
            let mut tail = head_pre.as_mut().unwrap();
            // 如果还有剩余的结点，则需要继续处理。
            // 每次找到长度为 interval 的两个区间 first 和 second ，
            // 然后进行归并
            while tail.next.is_some() {
                // 上一个区间尾结点 tail 的下一个结点，
                // 就是本次第一个区间的头结点 first_head
                let mut first_head = tail.next.take();
                // 同时维护第一个区间的尾结点 first_tail ，
                // 方便后续处理
                let mut first_tail = first_head.as_mut().unwrap();
                // 找到第一个区间的最后一个结点
                for _ in 1..interval {
                    // 如果下一个结点不存在，则说明已经遍历完，跳出循环
                    if first_tail.next.is_none() {
                        break;
                    }
                    // 此时还有下一个结点，则移动到下一个结点
                    first_tail = first_tail.next.as_mut().unwrap();
                }
                // 如果第一个区间就包含剩余的全部结点，不需要归并处理，
                // 将其放回到链表中，然后跳出内层循环即可
                if first_tail.next.is_none() {
                    // 将第一个区间的头结点放回到链表中
                    tail.next = first_head; 
                    break
                }

                // 第一个区间尾结点 first_tail 的下一个结点，
                // 就是第二个区间的头结点 second_head
                let mut second_head = first_tail.next.take();
                // 同时维护第二个区间的尾结点 second_tail ，
                // 方便后续处理
                let mut second_tail = second_head.as_mut().unwrap();
                // 找到第二个区间的最后一个结点
                for _ in 1..interval {
                    // 如果下一个结点不存在，则说明已经遍历完，跳出循环
                    if second_tail.next.is_none() {
                        break;
                    }
                    // 此时还有下一个结点，则移动到下一个结点
                    second_tail = second_tail.next.as_mut().unwrap();
                }

                // 第二个区间尾部和未处理区间断开，
                // 先记录未处理区间的头结点，方便复原
                let next = second_tail.next.take();

                // 合并 first 和 second ，获得合并链表的头结点，
                // 并将其插入已合并区间的尾部
                tail.next = Self::merge_two_sorted_lists(first_head, second_head);
                // 如果 tail 还有下一个结点
                while tail.next.is_some() {
                    // 移动 tail 到下一个结点
                    tail = tail.next.as_mut().unwrap();
                }
                
                // 将已合并区间 和 未合并区间连起来，防止链表断裂
                tail.next = next;
            }

            // 区间扩大一倍
            interval <<= 1;
        }

        // 返回排序后的头结点
        head_pre.unwrap().next
    }

    // 合并两个有序链表，返回合并后的链表的头结点
    pub fn merge_two_sorted_lists(mut first: Option<Box<ListNode>>, mut second: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        // 使用一个哨兵结点 head_pre ，方便后续处理
        let mut head_pre = Some(Box::new(ListNode::new(0)));
        // 使用尾插法，所以需要尾部结点的引用
        let mut tail = head_pre.as_mut().unwrap();
        // 如果两个链表都还有结点，则继续处理
        while first.is_some() && second.is_some() {
            // 这里取等保证排序是稳定的
            if first.as_ref().unwrap().val <= second.as_ref().unwrap().val {
                // 如果 first 的值更小，则将 first 的头结点取出，放入新的链表中
                let next = first.as_mut().unwrap().next.take();
                tail.next = first;
                // 移动 first 到下一个结点
                first = next;
            } else {
                // 如果 second 的值更小，则将 second 的头结点取出，放入新的链表中
                let next = second.as_mut().unwrap().next.take();
                tail.next = second;
                // 移动 second 到下一个结点
                second = next;
            }
            // 移动结果链表的 tail 到下一个结点
            tail = tail.next.as_mut().unwrap();
        }

        if first.is_some() {
            // 如果 first 还有结点，表明 second 已遍历完成，
            // 则将 first 直接放在 tail 后面
            tail.next = first;
        } else {
            // 如果 first 没有结点，表明 second 已遍历完成，
            // 则将 second 直接放在 tail 后面
            tail.next = second;
        }

        // 返回合并后的链表的头结点
        head_pre.unwrap().next
    }
}
