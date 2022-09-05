// 链接：https://leetcode.com/problems/vertical-order-traversal-of-a-binary-tree/
// 题意：给定一棵二叉树 root ，返回其垂直遍历的结果
//      （按列分组，按层升序，同层同列值小的在前）。


// 数据限制：
//  二叉树的结点数范围为 [1, 1000]
//  0 <= Node.val <= 1000


// 输入： root = [3,9,20,null,null,15,7]
// 输出： [[9],[3,15],[20],[7]]
// 解释：     3 (0, 0)
//          / \
// (1, -1) 9  20 (1, 1) 
//           /  \
//   (2, 0) 15   7 (2, 2)
//
//      第 -1 列：只有 9 这个结点
//      第  0 列：有 3 和 15 这两个结点，其中 3 的层更小
//      第  1 列：只有 20 这个结点
//      第  2 列：只有 7 这个结点

// 输入： root = [1,2,3,4,5,6,7]
// 输出： [[4],[2],[1,5,6],[3],[7]]
// 解释：         1 (0, 0)
//            /     \
//   (1, -1) 2       3 (1, 1) 
//          / \     / \
// (2, -2) 4   5   6   7 (2, 2)
//         (2, 0) (2, 0)
//
//      第 -2 列：只有 4 这个结点
//      第 -1 列：只有 2 这个结点
//      第  0 列：有 1, 5, 6 这三个结点，其中 1 的层最小，
//               5 和 6 在同层，但 5 的值更小
//      第  1 列：只有 3 这个结点
//      第  2 列：只有 7 这个结点


// 思路：递归/DFS
//
//      我们可以先用 dfs 将所有结点的信息收集到 result 中，
//      其中 result 的每个元素包含该结点的坐标 (row, col) 及其结点值 val 。
//
//      然后对 result 进行排序，按列升序排序，再按层升序排序，最后按值升序排序。
//
//      这样后续按顺序遍历时，就只用将同一列的按顺序收集在一个列表中，
//      保证相同列的顺序是满足题意的。
//      
//
//      时间复杂度： O(nlogn)
//          1. 需要遍历全部 O(n) 个结点
//          2. 需要对全部 O(n) 个结点的信息排序
//          3. 需要收集全部 O(n) 个结点值
//      空间复杂度： O(n)
//          1. 栈递归深度就是树高 O(h) ，
//              最差情况下，全部 O(n) 个结点在一条链上
//          2. 需要用 result 维护全部 O(n) 个结点的信息
//          3. 需要用 ans 维护结果中 O(n) 个结点值


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func verticalTraversal(root *TreeNode) [][]int {
	// 先用 dfs 收集所有结点的坐标和值
	var result []*NodeInfo
	dfs(root, 0, 0, &result)
	// 然后按列升序排序，再按层升序排序，最后按值升序排序
	sort.Slice(result, func(i, j int) bool {
        if result[i].col != result[j].col {
			return result[i].col < result[j].col
		}
		if result[i].row != result[j].row {
			return result[i].row < result[j].row
		}
		return result[i].val < result[j].val
	})
	// 开始按列分组并将结点值收集到 ans 中
	var ans [][]int
	// cur 表示当前列的结点值列表，初始放入第一个结点值
	cur := []int{ result[0].val }
	// preCol 表示前一列的列坐标，用于判断是否需要开启新分组，
	// 初始为第一个结点的列坐标
	preCol := result[0].col
	// 遍历 result ，由于刚刚已经考虑了第一个结点，所以跳过第一个
	for i := 1; i < len(result); i++ {
        col, val := result[i].col, result[i].val
		if preCol != col {
			// 如果当前结点在新的一列，则将 cur 放入 ans ，
			// 然后重新开始处理新的一列
			ans = append(ans, cur)
			cur = []int{ val }
			preCol = col
		} else {
			// 如果当前结点不在新的一列，则直接放入 cur 中
			cur = append(cur, val)
		}
	}
	// 将最后一列的结点值列表放入 ans 中
	ans = append(ans, cur)

	return ans
}

func dfs(root *TreeNode, row int, col int, result *[]*NodeInfo) {
	// 如果是空结点，则直接返回
	if root == nil {
		return
	}

	// 把当前结点的坐标和值放入 ans 中
	*result = append(*result, &NodeInfo{ col, row, root.Val })
	// 递归处理左右子树
	dfs(root.Left, row + 1, col - 1, result)
	dfs(root.Right, row + 1, col + 1, result)
}

type NodeInfo struct {
	col, row, val int
}
