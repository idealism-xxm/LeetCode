// 链接：https://leetcode.com/problems/validate-binary-search-tree/
// 题意：给定一个二叉树，判断是不是二叉搜索树？

// 输入：[2,1,3]
// 输出：true

// 输入：[5,1,4,null,null,3,6]
// 输出：false

// 思路1：递归
//
//		根据题意可得：
//			左子树所有的值 < root.Val < 右子树所有的值
//		按照上述条件递归模拟判断即可
//
//		时间复杂度： O(n)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	return isValidBst(root, nil, nil)
}

func isValidBst(root *TreeNode, minVal, maxVal *int) bool {
	// 空子树满足题意
	if root == nil {
		return true
	}
	// 如果存在最小值，则根结点的值不能 <= 最小值
	if minVal != nil && root.Val <= *minVal {
		return false
	}
	//如果存在最大值，则 根结点的值不能 >= 最大值
	if maxVal != nil && root.Val >= *maxVal {
		return false
	}
	// 此处：
	// 若存在最小值，则右子树的最小值不能 >= max(*minVal, root.Val) = root.Val ，可以和不存在最小值并在一起处理
	// 若存在最大值，则左子树的最大值不能 <= min(*maxVal, root.Val) = root.Val ，可以和不存在最大值并在一起处理
	return isValidBst(root.Left, minVal, &root.Val) && isValidBst(root.Right, &root.Val, maxVal)
}


// 思路2：循环
//
//		根据题意可得：
//			左子树所有的值 < root.Val < 右子树所有的值
//		而中序遍历顺序为：左子树的结点 -> 根结点 -> 右子树的结点
//		所以中序遍历所得的结点值必定是严格单调递增的
//		我们可以采用 0094 这题的方法采用循环进行遍历，然后每次比较并保存最后遍历的结点的值即可
//
//		时间复杂度： O(n) 空间复杂度： O(n)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	var stack []*TreeNode
	var lastVal *int

	for cur := root; cur != nil || len(stack) != 0; {
		// 若当前结点存在，则把当前结点放入，处理左子结点
		for ; cur != nil; {
			stack = append(stack, cur)
			cur = cur.Left
		}

		// 栈顶元素出栈
		length := len(stack)
		cur = stack[length - 1]
		stack = stack[:length - 1]
		if lastVal != nil {
			// 中序遍历的值必定严格递增
			if cur.Val <= *lastVal {
				return false
			}
		}

		// 左子树已处理完，更新 lastVal
		lastVal = &cur.Val
		// 处理右子树
		cur = cur.Right
	}
	// 全部满足严格递增
	return true
}
