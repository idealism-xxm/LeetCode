// 链接：https://leetcode.com/problems/maximum-difference-between-node-and-ancestor/
// 题意：给定一棵二叉树 root ，找出满足以下条件的 diff 的最大值：
//          1. diff = |a.val - b.val|
//          2. a 是 b 的祖先结点


// 数据限制：
//  树的结点数范围为 [2, 5000]
//  0 <= Node.val <= 10 ^ 5


// 输入： root = [8,3,10,1,6,null,14,null,null,4,7,13]
// 输出： 8
// 解释：     8
//         /   \
//        3    10
//       / \     \
//      1   6    14
//         / \   /
//        4  7  13
//
//      diff 的最大值为 |8 - 1| = 7

// 输入： root = [1,null,2,null,0,3]
// 输出： 3
// 解释： 1
//        \
//         2
//          \
//           0
//          /
//         3
//
//      diff 的最大值为 |0 - 3| = 3


// 思路： 递归/DFS + 贪心
//
//      遍历到结点 b 时，我们已经遍历过其全部祖先结点，
//      所以最朴素的想法就是维护祖先结点列表，
//      枚举全部祖先结点 a 计算结点 b 相关的 diff 的最大值。
//
//      但如果全部 O(n) 个结点在一条链上，时间复杂度就会变为 O(n ^ 2) ，
//      在当前数据范围下无法通过。
//
//      假设祖先结点中，结点值的最小值为 low ，最大值为 high ，
//      对于祖先结点 a 有 low <= a.val <= high ，则有以下三种情况：
//          1. b.val <= low: 则 diff 最大值为 high - b.val
//          2. low < b.val < high: 则 diff 最大值为 max(b.val - low, high - b.val)
//          3. high <= b.val: 则 diff 最大值为 b.val - high
//
//      综上， diff 的最大值一定是 max(abs(low - b.val), abs(high, b.val)) 。
//
//      所以我们可以在递归时维护祖先结点值的最小值和最大值，
//      这样就能在 O(1) 内求出结点 b 相关的 diff 的最大值，
//      整体时间复杂度为 O(n) 。
//
//
//      时间复杂度：O(n)
//          1. 需要遍历树中全部 O(n) 个结点
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
 func maxAncestorDiff(root *TreeNode) int {
    // root 没有祖先结点，为了方便处理，初始化为根结点值
    return dfs(root, root.Val, root.Val)
}

func dfs(root *TreeNode, low int, high int) int {
    // 如果 root 为空，则直接返回 0
    if root == nil {
        return 0
    }

    // 如果 root 不为空，则继续计算结点 root 相关的 diff 的最大值
    ans := max(abs(low - root.Val), abs(high - root.Val))
    // 更新祖先结点值的最小值和最大值
    low = min(low, root.Val)
    high = max(high, root.Val)
    // 递归处理左右子树，并更新 ans 的最大值
    ans = max(ans, dfs(root.Left, low, high))
    ans = max(ans, dfs(root.Right, low, high))

    return ans
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
