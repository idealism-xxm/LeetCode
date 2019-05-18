// 链接：https://leetcode.com/problems/remove-nth-node-from-end-of-list/
// 题意：给定一个链表，删除第 n 个元素

// 输入：list: 1->2->3->4->5, n = 2
// 输出：1->2->3->5

// 思路1：先循环遍历一遍求出长度，再循环删除第 n 个元素即可	

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    length := 0
    cur := head
    // 先求出链表长度
    for ; cur != nil; cur = cur.Next {
        length++
    }

    // 算出要删除的元素是第几个
    index := length - n + 1
    if index == 1 {
    	return head.Next // 如果是删除头结点，则直接返回头结点的下一个结点即可
    }

    // 移动到待删除元素的前一个元素
    cur = head
    for i := 1; i < index - 1; i++ {
        cur = cur.Next
    }
    cur.Next = cur.Next.Next // 删除第 index 个元素
    return head
}

// 思路2：快慢指针
//		初始化一个新头结点，防止移除头结点时出现问题，快慢指针初始化为新头结点
//		先将快指针移动 n 次，然后快慢指针每次同时往后移动一个结点，直到快指针到达尾结点
//		此时，慢指针指向待删除结点的前一个结点
//		（快慢指针也可以用来快速找到中间的元素，快指针每次移动两个结点，慢指针每次移动一个结点）

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 防止移除的是头结点时出现问题
    newHead := &ListNode{Next: head}

    fast, slow := newHead, newHead // 初始化快慢指针为新头结点
    // 先移动快指针 n 次
    for i := 0; i < n; i++ {
        fast = fast.Next
    }

    // 同时移动快慢指针，直到快指针到达尾结点
    for ; fast.Next != nil; fast = fast.Next {
        slow = slow.Next
    }

    if slow == newHead { // 如果移除的是头结点，则直接返回头结点的下一个结点
    	return head.Next
    }

    // 移除指定结点
    slow.Next = slow.Next.Next
    return head
}