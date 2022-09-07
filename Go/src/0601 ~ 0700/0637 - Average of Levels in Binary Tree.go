// 链接：https://leetcode.com/problems/average-of-levels-in-binary-tree/
// 题意：给定一棵二叉树 root ，返回其每一层所有结点值的平均值。


// 数据限制：
//  树中结点的数量范围是 [1, 10 ^ 4]
//  -(2 ^ 31) <= Node.val <= 2 ^ 31 - 1


// 输入：root = [3,9,20,null,null,15,7]
// 输出：[3.00000,14.50000,11.00000]
// 解释： 3      ->  第 0 层所有结点的平均值为 3
//      / \
//     9  20    ->  第 1 层所有结点的平均值为 14.5
//       /  \
//      15   7  ->  第 2 层所有结点的平均值为 11.0

// 输入：root = [3,9,20,15,7]
// 输出：[3.00000,14.50000,11.00000]
// 解释： 3      ->  第 0 层所有结点的平均值为 3
//      / \
//     9  20    ->  第 1 层所有结点的平均值为 14.5
//    / \
//   15  7      ->  第 2 层所有结点的平均值为 11.0


// 思路：BFS
//
//      本题其实就是需要层序遍历，是 LeetCode 102 的变形，
//      可以直接化用对应的思路和代码。
//
//      普通的层序遍历可以直接使用 BFS 处理，但本题需要收集同一层的结点平均值，
//      所以需要特殊处理一下。
//
//      只要我们知道当前层的结点数 num ，从队列 queue 中只取前 num 个结点处理，
//      这样就能将 BFS 转换成满足题意的层序遍历。
//
//      那么如何知道当前层的结点数 num 呢？
//
//      可以发现，只要我们是一层一层处理的，那么在处理当前层的结点前，
//      queue 中的所有结点都是当前层的结点，即当前层的结点数 num = queue.len() 。
//
//      所以我们在 BFS 循环中再加一层循环进行遍历，外部循环仅控制层序遍历是否结束，
//      内部循环才进行真正的处理逻辑，内部循环中只取 queue 中前 num 个结点，
//      收集这些结点的值，并将其非空左右子结点加入 queue 中。
//
//
//      时间复杂度： O(n)
//          1. 需要遍历全部 O(n) 个结点
//      空间复杂度： O(n)
//          1. 需要维护结果数组 ans ，保存全部 O(h) 层的结点平均值，
//              最差情况下，有 O(n) 层
//          2. 需要维护队列 queue ，保存全部 O(n) 个结点


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
	var ans []float64
	// 如果根结点为空，则直接返回空的结果
	if root == nil {
		return ans
	}

	// 把根结点放入队列中，准备进行 BFS
	queue := []*TreeNode{ root }
	// 队列中还有结点，则可以继续进行 BFS
	for len(queue) > 0 {
		// levelSum 收集当前层结点值的和， levelCnt 表示当前层的结点数
		levelSum := 0.0
		levelCnt := len(queue)
		// 此时 queue 中的结点数量为当前层的结点数量，
		// 只从队列中这些数量的结点，这样就变成了层序遍历
		for _ = range queue {
			node := queue[0]
			queue = queue[1:]
			// 收集这些结点的值到 level 中
			levelSum += float64(node.Val)
			// 将其非空左右子结点放入队列中
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		ans = append(ans, levelSum / float64(levelCnt))
	}

	return ans
}
