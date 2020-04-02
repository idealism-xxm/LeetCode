// 链接：https://leetcode.com/problems/binary-tree-postorder-traversal/
// 题意：给定一个二叉树，返回后序遍历的序列（通过迭代完成）？

// 输入： [1,null,2,3]
// 输出： [3,2,1]
// 解释：
//	 1
//    \
//     2
//    /
//   3

// 思路1： 前序 Morris
//
//		Morris 完成后序遍历不太好想，
//		但是后序遍历（左右中）可以看作前序遍历（中右左）这种形式的逆序
//		由于本题不需要直接输出，且需要返回结果列表，所以可以按照 0144 进行前序遍历（中右左）
//		最后逆序后返回即可
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
func postorderTraversal(root *TreeNode) []int {
	var result []int
	for cur := root; cur != nil; {
		// 右子树不存在，则直接收集结果，然后指向左子树
		if cur.Right == nil {
			result = append(result, cur.Val)
			cur = cur.Left
			continue
		}
		// 右子树存在，则找到右子树最后一个结点
		// （由于是 中右左 遍历，所以最后一个结点是右子树最左边的结点）
		rightLast := cur.Right
		// 不停寻找左子结点，
		// 直到某一结点的左子结点为 nil （第一次遍历到 cur）
		// 或者 为 cur （第二次遍历到 cur）
		for ; rightLast.Left != nil && rightLast.Left != cur; {
			rightLast = rightLast.Left
		}

		if rightLast.Left == nil {
			// 第一次遍历到 cur ，先收集 cur 的值
			result = append(result, cur.Val)
			// 然后让右子树的最后一个结点的左子结点指向 cur ，
			// 方便后续直接转移到 cur
			rightLast.Left = cur
			// 遍历右子树
			cur = cur.Right
		} else {
			// 第二次遍历到 cur ，复原即可
			rightLast.Left = nil
			// 遍历左子树
			cur = cur.Left
		}
	}

	// 逆序
	for i, j := 0, len(result) - 1; i < j; i, j = i + 1, j - 1 {
		result[i], result[j] = result[j], result[i]
	}
	return result
}

// 思路2： 后序 Morris
//
//		使用了取巧的方式转换成前序 Morris 完成了本题，
//		但是如果严格限制额外空间复杂度为 O(1) ，
//		且不是返回结果，而是输出结果，那么就必须使用 后序 Morris
//
//		Morris 完成后序遍历难在右子树完成后回到当前结点，
//		因为右子树最后一个结点恰好是当前结点的右子结点，
//		如果右子结点有两个子结点，那么就无法运用前序和中序使用的方法
//
//		这里太复杂了，需要用到特殊的技巧，可以发现最后一定是当前结点至最右结点的逆序
//		所以我们可以根据这一特性并使用 Morris 找前驱的方法完成后序遍历
//
//		首先建立一个 rootParent 结点，使 root 结点为它的左子结点
//		然后从 rootParent 开始循环进行以下流程直至 cur 为 nil ：
//			1. 若 cur 无左子结点，则转移到右子结点 (cur = cur.Right)，重复流程
//			2. 若 cur 有左子结点，则找到左子树中前序遍历的最后一个结点 leftLast
//				（即左子树中最右边的一个结点）
//				(1) 若 leftLast 的右子结点为 nil ，那么是第一次遍历到 cur ，
//					让 leftLast 的右子结点为 cur (leftLast.Right = cur) ，
//					并转移到左子结点 (cur = cur.Left) ，重复流程
//				(2) 若 leftLast 的右子结点为 cur ，那么是第二次遍历到 cur ，
//					将 leftLast 的右子结点置为 nil (leftLast.Right = nil) ，
//					倒序输出从 cur.Left 到 leftLast 这条路径上的所有结点，
//					并转移到右子结点 (cur = cur.Right) ，重复流程
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
func postorderTraversal(root *TreeNode) []int {
	var result []int
	rootParent := &TreeNode {
		Val: 0,
		Left: root,
		Right: nil,
	}
	for cur := rootParent; cur != nil; {
		// 左子树不存在，则指向右子树
		if cur.Left == nil {
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
			// 第一次遍历到 cur
			// 让左子树的最后一个结点的右子结点指向 cur ，
			// 方便后续直接转移到 cur
			leftLast.Right = cur
			// 遍历左子树
			cur = cur.Left
		} else {
			// 第二次遍历到 cur ，复原即可
			leftLast.Right = nil
			// 倒序收集从 cur.Left 到 leftLast 这条路径上的所有结点
			result = append(result, collectResult(cur.Left, leftLast)...)
			// 遍历左子树
			cur = cur.Right
		}
	}

	return result
}

func collectResult(from, to *TreeNode) []int {
	// 翻转路径上的结点（保证输出情况下严格满足额外空间复杂度为 O(1) ）
	reverse(from, to)

	// 收集结果
	var result []int
	for cur := to; ; cur = cur.Right {
		result = append(result, cur.Val)
		// 已收集完
		if cur == from {
			break
		}
	}

	// 复原路径上的结点
	reverse(to, from)
	return result
}

func reverse(from, to *TreeNode) {
	// 如果是同一个结点，直接返回
	if from == to {
		return
	}

	// 0143 使用了 头插法 进行翻转
	// 此处再使用 双指针 进行翻转
	pre, tail := to.Right, to.Right
	for cur := from; cur != tail; {
		// 翻转前后关系
		next := cur.Right
		cur.Right = pre

		// 移动指向，处理下一个相邻的结点
		pre = cur
		cur = next
	}
}
