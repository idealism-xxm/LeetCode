// 链接：https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/
// 题意：给定一棵二叉树，找到两个结点的最近公共祖先？


// 数据限制：
//  树的结点数在 [2, 10 ^ 5] 之间
//	-(10 ^ 9) <= Node.val <= 10 ^ 9
//	所有的 Node.val 都各不相同
//	p != q
//	p 和 q 必定在树中


// 输入： root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
// 输出： 3
// 解释： 结点 5 和 1 的 LCA 是 3
//       3
//     /   \
//   5       1
//  / \     / \
// 6   2   0   8
//    / \
//   7   4

// 输入： root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
// 输出： 5
// 解释： 结点 5 和 4 的 LCA 是 5
//       3
//     /   \
//   5       1
//  / \     / \
// 6   2   0   8
//    / \
//   7   4


// 思路： 递归
//
//      我们使用递归处理即可，返回值有两个：
//          cnt 表示当前子树中含有的 p 和 q 的结点数
//          node 表示 cnt 为 2 时当前子树中 p 和 q 的 LCA ，其他时候无意义
//
//      然后按照以下步骤继续处理：
//          1. root 为空：直接返回 nil, 0
//          2. 递归计算左子树结果 leftNode, leftCnt 。
//              若 leftCnt == 2 ，则 leftNode 就是 LCA ，直接返回即可
//          3. 递归计算右子树结果 rightNode, rightCnt 。
//              若 rightCnt == 2 ，则 rightNode 就是 LCA ，直接返回即可
//          4. 计算当前子树含有的 p 和 q 的结点数 curCnt = leftCnt + rightCnt 。
//              若 root == p || root == q ，则还需要对 curCnt + 1
//          5. 返回 root, curCnt （若 curCnt == 2 ，则 root 就是 LCA ）
//
//      如果本题是一棵树有多次查询，那么可以使用 tarjan 或者倍增进行处理。
//
//
//      时间复杂度： O(n)
//            1. 需要遍历全部 O(n) 个结点一次
//      空间复杂度： O(h)
//            1. 栈递归深度就是树高，最差情况下，全部 O(n) 个结点在一条链上


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	lca, _ := dfs(root, p, q)
	return lca
}

func dfs(root, p, q *TreeNode) (*TreeNode, int) {
	// root 为空，则直接返回 nil, 0
	if root == nil {
		return nil, 0
	}

	// 递归计算左子树结果 leftNode, leftCnt
	leftNode, leftCnt := dfs(root.Left, p, q)
	// 若 leftCnt == 2 ，则 leftNode 就是 LCA ，直接返回即可
	if leftCnt == 2 {
		return leftNode, 2
	}

	// 递归计算右子树结果 rightNode, rightCnt
	rightNode, rightCnt := dfs(root.Right, p, q);
	// 若 rightCnt == 2 ，则 rightNode 就是 LCA
	if rightCnt == 2 {
		return rightNode, 2
	}

	// 计算当前子树含有的 p 和 q 的结点数
	curCnt := leftCnt + rightCnt
	// 同时要统计 root 可能为 p 或 q 的情况
	if root == p || root == q {
		curCnt += 1;
	}

	// 若 curCnt == 2 ，则 root 就是 LCA
	return root, curCnt
}
