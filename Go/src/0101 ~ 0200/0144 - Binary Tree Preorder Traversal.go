// 链接：https://leetcode.com/problems/binary-tree-preorder-traversal/
// 题意：给定一个二叉树，返回前序遍历的序列（通过迭代完成）？

// 输入： [1,null,2,3]
// 输出： [1,2,3]
// 解释：
//	 1
//    \
//     2
//    /
//   3

// 思路： Morris
//
//		当然可以使用栈模拟递归，但现在已经知道了额外空间复杂度为 O(1) 的 Morris
//		所以就再次熟悉一下
//		Morris 的思路就是
//		第一次遍历到某个结点时，将其左子树中的最后一个结点的右子结点指向其本身，
//		第二次遍历到该结点时，将其左子树中的最后一个结点的右子结点复原为 nil
//		根据具体的先中后序遍历，决定收集结果的顺序
//		本题需要进行先序遍历，所以第一次遍历到某个结点时，就直接收集结果，第二次仅复原
//
//		时间复杂度： O(n)
//		空间复杂度： O(n) （额外空间复杂度为 O(1) ）

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	var result []int
	for cur := root; cur != nil; {
		// 左子树不存在，则直接收集结果，然后指向右子树
		if cur.Left == nil {
			result = append(result, cur.Val)
			cur = cur.Right
			continue
		}
		// 左子树存在，则找到左子树最后一个结点
		leftLast := cur.Left
		// 不停寻找右子结点，
		// 直到某一结点的右子结点为 nil （第一次遍历到 cur）
		// 或者 为 cur （第二次遍历到 cur）
		for ; leftLast.Right != nil && leftLast.Right != cur; {
			leftLast = leftLast.Right
		}

		if leftLast.Right == nil {
			// 第一次遍历到 cur ，先收集 cur 的值
			result = append(result, cur.Val)
			// 然后让左子树的最后一个结点的右子结点指向 cur ，
			// 方便后续直接转移到 cur
			leftLast.Right = cur
			// 遍历左子树
			cur = cur.Left
		} else {
			// 第二次遍历到 cur ，复原即可
			leftLast.Right = nil
			// 遍历右子树
			cur = cur.Right
		}
	}
	return result
}
