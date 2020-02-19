// 链接：https://leetcode.com/problems/unique-binary-search-trees-ii/
// 题意：给定一个整数，求所有的二叉搜索树？

// 输入：3
// 输出：
// [
//   [1,null,3,2],
//   [3,2,null,1],
//   [3,1,null,null,2],
//   [2,1,3],
//   [1,null,2,null,3]
// ]

// 思路：DP
//		其实最开始就可以想到递归处理，
//		但是一想递归其实就是枚举该结点的的值，
//		然后将小于的值递归处理成左子树，再将大于的值递归处理成右子树
//		考虑到每次都是 x 个不同的树组成二叉搜索树，可以将 1~x 组成的所有结果存起来，
//		然后按照大小依次克隆，赋值为对应次序的值即可
//		这样以来就有子结构和递推方式了，很容易就可以转换为 DP
//
//		dp[i] 表示 1～i 个数可以构成的所有二叉搜索树
//		初始化：dp[0] = []*TreeNode{nil} 没有任何数字，即为 nil 结点
//		状态转移方程：
//			对于 dp[i] 的根结点 j (1 <= j <= i) ，
//			其左子树可从 dp[j - 1] 中选后可得，
//			其右子树可从 dp[i - j - 1] 中选后并根据大小赋值可得
//			则：dp[i] 则等于每一种情况的乘积之和

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generateTrees(n int) []*TreeNode {
	// 0 应该返回 nil
	if n == 0 {
		return nil
	}

	dp := make([][]*TreeNode, n + 1)
	// 初始化 dp[0] ：没有任何数字，即为 nil 结点
	dp[0] = append(dp[0], nil)
	// values 用于克隆赋值
	values := make([]int, n)
	for i := 1; i <= n; i++ {
		// 第 i 个数是 i
		values[i - 1] = i
		// 枚举根结点的数字
		for j := 1; j <= i; j++ {
			// 左子树可从 dp[j - 1] 中获取
			for _, left := range dp[j - 1] {
				// 右子树可从 dp[i - j - 1] 中获取
				for _, right := range dp[i - j] {
					dp[i] = append(dp[i], &TreeNode{
						Val:   j,
						// 可以复用左子树，但没必要
						Left:  cloneAndSetValue(left, values[:j - 1]),
						Right: cloneAndSetValue(right, values[j:i]),
					})
				}
			}
		}
	}
	return dp[n]
}

// 克隆 cur 子树，克隆后根结点的值为 values[cur.Val - 1]
func cloneAndSetValue(cur *TreeNode, values []int) *TreeNode {
	if cur == nil {
		return nil
	}

	// 克隆根结点，并赋值，且递归克隆左结点和右结点
	return &TreeNode{
		Val:   values[cur.Val - 1],
		Left:  cloneAndSetValue(cur.Left, values),
		Right: cloneAndSetValue(cur.Right, values),
	}
}
