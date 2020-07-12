// 链接：https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/
// 题意：给定一个二叉树，找到两个节点的最近公共祖先？

// 输入： root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
// 输出： 3
// 解释： 节点 5 和 1 的 LCA 是 3
//       3
//    /      \
//   5        1
//  / \      / \
// 6   2    0   8
//    / \
//   7   4

// 输入： root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
// 输出： 5
// 解释： 节点 5 和 4 的 LCA 是 5
//       3
//    /      \
//   5        1
//  / \      / \
// 6   2    0   8
//    / \
//   7   4

// 思路： 递归
//
//		我们使用递归处理即可，返回值有两个：
//     		cnt 表示当前子树中含有的 p 和 q 的节点数
//			node 表示 cnt 为 2 时当前子树中 p 和 q 的 LCA ，其他时候无意义
//		1. root == nil: 直接返回 nil, 0
//		2. 递归计算左子树结果 leftNode, leftCnt ，
//			若 leftCnt == 2 ，则直接返回
//		3. 递归计算右子树结果 rightNode, rightCnt ，
//			若 rightCnt == 2 ，则直接返回
//		4. 计算当前子树含有的 p 和 q 的节点数 curCnt = leftCnt + rightCnt
//			若 root == p || root == q ，则还需要对 curCnt + 1
//		5. 返回 root, curCnt
//
//		如果本题是一棵树有多次查询，那么可以使用 tarjan 或者倍增进行处理
//
//      时间复杂度： O(n)
//      空间复杂度： O(h)

/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	lca, _ := dfs(root, p, q)
	return lca
}

func dfs(root, p, q *TreeNode) (*TreeNode, int) {
	if root == nil {
		return nil, 0
	}

	leftNode, leftCnt := dfs(root.Left, p, q)
	if leftCnt == 2 {
		return leftNode, 2
	}
	rightNode, rightCnt := dfs(root.Right, p, q)
	if rightCnt == 2 {
		return rightNode, 2
	}
	curCnt := leftCnt + rightCnt
	if root == p || root == q {
		curCnt++
	}
	return root, curCnt
}