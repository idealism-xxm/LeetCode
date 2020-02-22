// 链接：https://leetcode.com/problems/recover-binary-search-tree/
// 题意：有一个二叉搜索树，交换了其中两个结点的值，将其恢复成二叉搜索树？

// 输入：[1,3,null,null,2]
// 输出：[3,1,null,null,2]

// 输入：[3,1,4,null,null,2]
// 输出：[2,1,4,null,null,3]

// 思路1：递归
//
//		按照 0098 思路2 的思想：一个二叉搜索树的中序遍历值必定严格递增
//		所以我们只要进行中序遍历，然后比较相邻两个结点的值
//		1. 若当前结点的值小于前一个结点的值，则前一个结点（设为 large ）必定是其中一个结点
//		2. 在找到 1 中结点的前提下，找到第一个比 large 值大的结点，
//			则其前一个结点（设为 small ）必定是选择的另一个结点
//			（若找不到则为最后一个结点）
//		3. 递归完成后，交换 large 和 small 结点的值即可
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
func recoverTree(root *TreeNode)  {
	// 1. 找到交换的两个结点
	large, small, last := find(root, nil, nil)
	// 2. 进行交换
	if small == nil {
		// 如果没找到第一个比 large 值大的结点，则 最后一个结点 是另一个结点
		small = last
	}
	large.Val, small.Val = small.Val, large.Val
}

func find(root, large, last *TreeNode) (*TreeNode, *TreeNode, *TreeNode) {
	// 对 root 子树进行中序遍历
	// large, last 分别表示前面遍历所得的 large 结点和上一个结点

	// 空子树不改变对应对值
	if root == nil {
		return large, nil, last
	}

	// 先处理左子树
	lLarge, lSmall, lLast := find(root.Left, large, last)
	// 如果左边子树已找到交换的两个结点，则直接结束递归
	if lSmall != nil {
		return lLarge, lSmall, lLast
	}

	// 处理根结点
	if lLarge == nil {
		// 若前面还没有找到 large 结点，则需要先找到它
		if lLast != nil && root.Val < lLast.Val {
			lLarge = lLast
		}
	} else {
		// 若前面已找到 large 结点，则需要找到 small 结点
		// 找到则直接结束递归
		if root.Val > lLarge.Val {
			return lLarge, lLast, nil
		}
	}

	// 最后处理右子树
	return find(root.Right, lLarge, root)
}

// 思路2：Morris
//		基本思路还是同上
//
//		按照 0098 思路2 的思想：一个二叉搜索树的中序遍历值必定严格递增
//		所以我们只要进行中序遍历，然后比较相邻两个结点的值
//		1. 若当前结点的值小于前一个结点的值，则前一个结点（设为 large ）必定是其中一个结点
//		2. 在找到 1 中结点的前提下，找到第一个比 large 值大的结点，
//			则其前一个结点（设为 small ）必定是选择的另一个结点
//			（若找不到则为最后一个结点）
//		3. 递归完成后，交换 large 和 small 结点的值即可
//
//		由于递归和普通循环方式的空间复杂度都是 O(n)
//		所以需要其他的方式进行中序遍历，
//		普通循环方式需要 O(n) 的空间是因为我们需要先处理左子树，然后退回到当前结点，即保存每一个根结点的信息
//		如果我们知道每一个根结点的前一个遍历的结点（前驱结点），则这个空间就不再需要
//		1. 若当前根结点没有左子结点，则不需要找到前驱结点，因为它已经遍历过了，
//			当前根结点就是需要遍历的下一个结点，直接处理即可
//		2. 若当前根结点有左子结点，则需要找到前驱结点，
//			可由以下方法找到：从左子结点开始，循环找到其右子结点，直到某一右子结点为空，则最后一个结点就是前驱结点
//			我们可将前驱结点的右子结点变为当前根结点，这样就能不需要多余空间
//			但同时我们还需要在处理完当前根结点后复原前驱结点，避免破坏树但结构
//			所以当根结点的前驱结点的右子结点为当前根结点时，需要将前驱结点的右子结点置空
//			即：
//			(1) 若前驱结点的右子结点为空，则将前驱结点右子结点变为当前根结点
//			(2) 若前驱结点的右子结点为当前结点，处理当前根结点，然后将前驱结点右子结点置空
//
//		时间复杂度： O(n) 空间复杂度： O(1)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recoverTree(root *TreeNode)  {
	var large, small, last *TreeNode
	// 1. 用 Morris 找到较换的两个结点
	for cur := root; cur != nil; {
		// 无左子结点，则直接处理当前结点
		if cur.Left == nil {
			// 更新 large 和 small
			large, small = update(cur, large, small, last)
			// 更新最后一个结点
			last = cur
			// 处理右子树
			cur = cur.Right
		} else {
			// 有左子结点，则需找到前驱结点
			pre := cur.Left
			for ; pre.Right != nil && pre.Right != cur; {
				pre = pre.Right
			}
			// 右子结点为空，则当前只需将右子结点指向当前结点即可
			if pre.Right == nil {
				pre.Right = cur
				// 此时需先遍历左子树
				cur = cur.Left
			} else {
				// 右子结点为不为空，则先处理当前结点
				// 更新 large 和 small
				large, small = update(cur, large, small, last)
				// 更新最后一个结点
				last = cur
				// 将右子结点置空
				pre.Right = nil
				// 此时只用遍历右子树
				cur = cur.Right
			}
		}
	}
	// 2. 进行交换
	if small == nil {
		// 如果没找到第一个比 large 值大的结点，则 最后一个结点 是另一个结点
		small = last
	}
	large.Val, small.Val = small.Val, large.Val
}

func update(cur, large, small, last *TreeNode) (*TreeNode, *TreeNode) {
	// 对当前结点 cur 和前面得出的 large, small, last 进行处理，返回新的 large, small
	// large, last 分别表示前面遍历所得的 large 结点和上一个结点
	if large == nil {
		// 若前面还没有找到 large 结点，则需要先找到它
		if last != nil && cur.Val < last.Val {
			large = last
		}
	} else if small == nil {
		// 若前面已找到 large 结点，且未找到 small 结点，则需要找到 small 结点
		if cur.Val > large.Val {
			small = last
		}
	}
	return large, small
}
