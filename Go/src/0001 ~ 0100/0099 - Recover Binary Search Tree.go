// 链接：https://leetcode.com/problems/recover-binary-search-tree/
// 题意：给定一棵二叉搜索树 root ，其中恰好有两个结点被互换了，
//      将这颗二叉搜索树复原。


// 数据限制：
//  树的结点数为 n
//  1 <= k <= n <= 10 ^ 4
//  0 <= Node.Val <= 10 ^ 4


// 输入： root = [1,3,null,null,2]
// 输出： [3,1,null,null,2]
// 解释： 3 不能作为 1 的左子结点，因为 3 > 1 
//       1            3
//      /            /
//     3      →     1
//      \            \
//       2            2
//

// 输入： root = [3,1,4,null,null,2]
// 输出： [2,1,4,null,null,3]
// 解释： 2 不能在 3 的右子树中，因为 2 < 3
//       3           2
//      / \         / \
//     1   4   →   1   4
//        /           /
//       2           3


// 思路1： 递归
//
//      一个二叉搜索树的中序遍历值必定严格递增，
//      所以我们只要进行中序遍历，然后比较相邻两个结点的值。
//
//      我们可以使用 dfs 闭包递归中序遍历处理，
//      该闭包能引用三个外部变量：
//          1. previous: 表示中序遍历的前一个结点
//          2. first:    表示互换结点的前者，该结点必定比后一个结点大
//          3. second:   表示互换结点的后者，前一个结点必定比该结点大
//
//      可以发现两个互换的结点，必定出现在中序遍历时大小不对的位置处，
//      所以在 dfs 中，如果前一个结点 previous 的值大于当前结点 root 的值，
//      则找到了一个互换的结点。
//
//      1. 如果这样的位置有 1 处，那么 first 必定是 previous ，
//          second 必定是 root 
//      2. 如果这样的位置有 2 处，那么 first 一定是第一处的 previous ，
//          second 一定是第二处的 root
//
//      综上： first 必定是第一处的 previous ，
//           second 必定是最后一处的 root
//
//
//      时间复杂度：O(n)
//          1. 需要遍历找到两个互换的结点，最差情况下，
//              最后一个结点被换了，需要遍历全部 O(n) 个结点
//      空间复杂度：O(n)
//          1. 栈递归深度就是树高，最差情况下，全部 O(n) 个结点在一条链上


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recoverTree(root *TreeNode)  {
	// 定义 first 和 second ，用于维护交换的两个结点
	// previous: 表示中序遍历的前一个结点
	// first:    表示互换结点的前者，该结点必定比后一个结点大
	var first, second *TreeNode
	// previous: 表示中序遍历的前一个结点
	var previous *TreeNode
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		// 对 root 子树进行递归中序遍历，找到两个互换的结点。

		// 如果当前结点为空，则直接返回
		if root == nil {
			return
		}

		// 先递归处理左子树
		dfs(root.Left)

		// 如果前一个结点的值大于当前结点，则找到了一个互换的结点
		if previous != nil && previous.Val > root.Val {
			if first == nil {
				// 如果第 1 个结点未找到，则设置第 1 个结点为前一个结点
				first = previous
			}
			if first != nil {
				// 如果第 1 个结点已找到，则当前找到的是第 2 个结点，
				// 设置第 2 个结点为当前结点
				second = root
			}
		}

		// 设置前一个结点为当前结点
		previous = root
		// 继续递归处理右子树
		dfs(root.Right)
	}

	// 递归中序遍历找到互换的结点
	dfs(root)
	// 交换 first 和 second 两个结点的值
	first.Val, second.Val = second.Val, first.Val
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
