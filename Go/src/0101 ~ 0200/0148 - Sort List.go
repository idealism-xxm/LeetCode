// 链接：https://leetcode.com/problems/sort-list/
// 题意：对链表进行排序，要求时间复杂度为 O(nlogn) ，空间复杂度为 O(1) ？

// 输入： 4->2->1->3
// 输出： 1->2->3->4

// 输入： -1->5->3->4->0
// 输出： -1->0->3->4->5

// 思路： 归并排序 + 倍增法
//
//		平均时间复杂度为 O(nlogn) 的只有希尔排序、堆排序、
//		而最快时间复杂度为 O(nlogn) 的只有快速排序和归并排序
//
//		堆排序利用了数组可以 O(1) 查找元素性质，所以链表中无法使用
//		快速排序每趟重排后两边的长度不确定，所以很难转成空间复杂度为 O(1) 的迭代
//		归并排序每次归并后长度翻倍，能准确定位到每一段并进行处理，
//			所以可以使用倍增法转成空间复杂度为 O(1) 的迭代
//			（习惯了平时递归分治的形式，这次用倍增法到又开了眼界）
//		希尔排序也可以进行空间复杂度为 O(1) 迭代处理，但最坏时间复杂度为 O(n * (logn)^2)
//
//		时间复杂度： O(nlogn)
//		空间复杂度： O(1)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	// 找到链表长度
	length := 0
	for cur := head; cur != nil; cur = cur.Next {
		length++
	}

	// 在头结点前添加一个，方便后续操作
	headPre := new(ListNode)
	headPre.Next = head
	// 初始每个区间但长度为 1 ，每相邻两个区间进行归并，后续倍增区间长度继续处理即可
	for interval := 1; interval < length; interval <<= 1 {
		// 从 headPre 开始，每次找到长度为 interval 的两个区间，取出来进行合并
		for last := headPre; last.Next != nil; {
			// 上一段末尾的下一个结点，就是本次第一个区间的头结点
			firstHead, firstTail := last.Next, last.Next
			// 找到第一个区间的最后一个结点
			for i := interval - 1; i > 0 && firstTail.Next != nil; i-- {
				firstTail = firstTail.Next
			}
			// 如果第一个区间就包含了剩下所有的结点，则跳出内层循环，处理下一个长度的区间
			if firstTail.Next == nil {
				break
			}

			// 第一个区间末尾的下一个结点，就是本次第二个区间的头结点
			secondHead, secondTail := firstTail.Next, firstTail.Next
			// 找到第二个区间的最后一个结点
			for i := interval - 1; i > 0 && secondTail.Next != nil; i-- {
				secondTail = secondTail.Next
			}

			// 第一个区间前后断开
			last.Next, firstTail.Next = nil, nil
			// 第二个区间前后断开断开，并记录第二个区间的下一个结点，方便复原
			next := secondTail.Next
			secondTail.Next =  nil

			// 合并链表
			mergedHead, mergedTail := mergeList(firstHead, secondHead)
			// 复原链表
			last.Next, mergedTail.Next = mergedHead, next
			// 更新 last ，方便下一次处理
			last = mergedTail
		}
	}
	return headPre.Next
}

// 将两个有序链表进行合并，返回一个有序链表的头结点和尾结点
func mergeList(first, second *ListNode) (*ListNode, *ListNode) {
	// 在头结点前添加一个，方便后续操作
	headPre := new(ListNode)
	cur := headPre
	for ; first != nil && second != nil; {
		// 取等保证稳定
		if first.Val <= second.Val {
			cur.Next = first
			first = first.Next
		} else {
			cur.Next = second
			second = second.Next
		}
		cur = cur.Next
	}
	// 如果是第一个链表用完了，则让其指向第二个链表剩余部分
	if first == nil {
		first = second
	}
	// 剩余的结点直接挂在最后即可
	cur.Next = first
	// 找到尾结点
	for ; first.Next != nil; first = first.Next {}

	return headPre.Next, first
}
