// 链接：https://leetcode.com/problems/convert-sorted-list-to-binary-search-tree/
// 题意：给定一个排好序的单链表，将其转换成一颗二叉搜索树，并保证每个结点的左右子树高度差不大于 1 ？

// 输入：[-10,-3,0,5,9],
// 输出：
//      0
//     / \
//   -3   9
//   /   /
// -10  5

// 思路1：递归 + 快慢指针
//
//		思路 0109 一样，每次让中间的数作为当前子树的根结点，
//		该数左边的数递归形成左子树，该数右边的数递归形成右子树
//		用快慢指针每一层总共可以在 O(n) 内找到所有的中间结点，总共有 O(logn) 层
//
//		时间复杂度： O(nlogn)
//		空间复杂度： O(n)  实际只需要 O(logn) 的额外空间

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
	return toBst(head, nil)
}

// 将 头结点 和 最后一个结点的下一个 传入，构造中间元素形成的二叉搜索树
// 加入 tailNext 可以不用修改原有链表，且方便双指针运行
func toBst(head *ListNode, tailNext *ListNode) *TreeNode {
	// 如果没有可用结点，直接返回 nil
	if head == tailNext {
		return nil
	}

	// 通过快慢指针找到中间结点
	fast, slow := head.Next, head
	for ; fast != tailNext; {
		slow = slow.Next
		fast = fast.Next
		if fast != tailNext {
			fast = fast.Next
		}
	}

	// 中间数为根结点，左边的数递归处理成左子树，右边的数递归处理成右子树
	return &TreeNode {
		Val: slow.Val,
		Left: toBst(head, slow),
		Right: toBst(slow.Next, tailNext),
	}
}

// 思路2：递归 + 中序模拟
//
//		看了题解后感觉思路很巧妙，其实这种方法以前写线段树经常用
//		我们可以发现对于一颗二叉搜索树进行中序遍历，结果就是一个有序的链表
//		所以我们可以反向操作，在中序遍历的时候，不是输出当前结点，而是读入当前结点
//		这样就能在 O(n) 内建立其一颗二叉树
//
//		时间复杂度： O(n)
//		空间复杂度： O(n)  实际只需要 O(logn) 的额外空间

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
	length := 0
	for cur := head; cur != nil; cur = cur.Next {
		length ++
	}

	root, _ := toBst(0, length - 1, head)
	return root
}

func toBst(l, r int, cur *ListNode) (*TreeNode, *ListNode) {
	// 如果已是空结点，则直接返回
	if l > r {
		return nil, cur
	}

	// 如果不是空结点
	mid := (l + r) >> 1
	left, root := toBst(l, mid - 1, cur)  // 先处理左子树，获得根结点的值，
	right, next := toBst(mid + 1, r, root.Next)  // 再处理右子树，获取需要返回的结点值
	// 最后直接构造根结点并返回即可
	return &TreeNode {
		Val: root.Val,
		Left: left,
		Right: right,
	}, next
}
