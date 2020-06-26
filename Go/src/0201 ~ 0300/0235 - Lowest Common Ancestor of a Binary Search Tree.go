// 链接：https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/
// 题意：给定一个二叉搜索树，找到两个节点的最近公共祖先？

// 输入： root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
// 输出： 6
// 解释： 节点 2 和 8 的 LCA 是 6
//       6
//    /      \
//   2        8
//  / \      / \
// 0   4    7   9
//    / \
//   3   5

// 输入： root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
// 输出： 2
// 解释： 节点 2 和 4 的 LCA 是 2
//       6
//    /      \
//   2        8
//  / \      / \
// 0   4    7   9
//    / \
//   3   5

// 思路： 迭代
//
//      因为是二叉搜索树，所以要寻找的两个节点第一次分开的节点就是它们的 LCA ，
//      首先为了方便，令 p.val < q.val ，然后迭代判断即可，
//		1. p.Val <= cur.Val && p.Val >= cur.Val ，则两者有不同的走向，
//			cur 就是两者的 LCA ，直接返回
//		2. p.Val <= q.Val < cur.Val ，则两者的 LCA 在左子树
//		3. cur.Val < p.Val <= q.Val ，则两者的 LCA 在右子树
//
//      时间复杂度： O(h)
//      空间复杂度： O(1)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 让 p.Val <= q.Val ，方便后续处理
	if p.Val > q.Val {
		p, q = q, p
	}
	for cur := root; cur != nil; {
		// 此时两个节点将有不同的走向，所以 cur 是 LCA
		if p.Val <= cur.Val && q.Val >= cur.Val {
			return cur
		}

		if q.Val < cur.Val {
			// 如果 p.Val <= q.Val < cur.Val ，则两者的 LCA 在左子树
			cur = cur.Left
		} else {
			// 此时 cur.Val < p.Val <= q.Val ，则两者的 LCA 在右子树
			cur = cur.Right
		}
	}

	// 题目保证不会走到这
	return nil
}
