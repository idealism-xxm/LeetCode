// 链接：https://leetcode.com/problems/path-sum-iii/
// 题意：给定一棵二叉树 root 和一个正整数 targetSum ，
//      求所有和为 targetSum 的路径的数量？


// 数据限制：
//  二叉树的结点数范围为 [0, 1000]
//  -(10 ^ 9) <= Node.val <= 10 ^ 9
//  -1000 <= targetSum <= 1000


// 输入： root = [10,5,-3,3,2,null,11,3,-2,null,1], targetSum = 8
// 输出： 3
// 解释： 有三条路径的和为 8 ：
//       5 + 3 = 8
//       5 + 2 + 1 = 8
//       -3 + 11 = 8
//
//          10       
//         /  \      
//       (5)  (-3)
//       /  \   \    
//     (3)  (2)  (11)
//     /  \   \    \
//    3   -2  (1)   11

// 输入： root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
// 输出： 3
// 解释： 有三条路径的和为 22 ：
//       5 + 4 + 11 + 2 = 22
//       5 + 8 + 4 + 5 = 22
//       4 + 11 + 7 = 22
//
//          (5)       
//         /   \      
//       (4)   (8)     
//       /    /  \    
//     (11)  13  (4)
//     /  \      /  \
//    (7) (2)   (5)  1


// 思路：递归/回溯/DFS + 前缀和 + Map
//
//      本题是 LeetCode 0113 的加强版，要统计和为 targetSum 的数量，
//      需要记录从根结点到当前结点的路径的前缀和及其出现次数，并在每个结点处统计满足题意的路径。
//
//      这个 前缀和 + Map 思路就是 LeetCode 560 这题的思路，
//      本题其实就是将给定的数组变为了动态计算的路径。
//
//      本题需要从根结点到当前结点的路径的前缀和及其出现次数，
//      所以需要定义 dfs(root, targetSum, currSum, preSum) 来辅助处理，
//          1. root 当前待处理的子树的根结点
//          2. targetSum 路径上的数之和需要为 targetSum ，直接透传，不做修改
//          3. currSum 从根结点到 root 的路径上的数之和
//          4. preSum preSum[sum] 表示从根结点到 root 的路径中，前缀和为 sum 的路径数
//
//      返回值就是以 root 子树中每个结点为路径截至结点的所有路径中，满足题意的路径数之和。
//
//      dfs 按照以下逻辑进行处理：
//          1. 如果 root 是空结点，则不存在路径，则直接返回 0
//          2. 将 root.Val 计入 currSum 中
//          3. 以 root 为路径截至结点的所有路径中，满足题意的路径必定是后缀和为 targetSum 的路径，
//             这些路径数量必定等于前缀和为 currSum - targetSum 的路径数，
//             即 ans = preSum[currSum - targetSum]
//          4. 将 currSum 计入 preSum 中，【注意】这一步需要在统计完成后，否则可能会错误统计入空路径
//          5. 递归处理左右子结点，将结果计入 ans 中
//          6. 退出递归前需要将本层计入 preSum 中前缀和 currSum 移除
//          7. 此时 ans 就是以 root 子树中每个结点为路径截至结点的所有路径中，满足题意的路径数之和
//
//
//      时间复杂度： O(n)
//          1. 需要遍历全部 O(n) 个结点
//      空间复杂度： O(n)
//          1. 栈递归深度就是树高，最差情况下，全部 O(n) 个结点在一条链上
//          2. 需要维护路径上全部不同的前缀和的出现次数，最差情况下，有 O(n) 个不同的前缀和


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) int {
    // preSum[sum] 表示从根结点到当前结点的路径中，前缀和为 sum 的路径数
    preSum := make(map[int]int)
    // 初始化前缀和为 0 的路径出现 1 次，方便后续计入从根结点到当前结点的路径
    preSum[0] = 1

    return dfs(root, targetSum, 0, preSum)
}

func dfs(root *TreeNode, targetSum int, currSum int, preSum map[int]int) int {
    // 空结点不存在路径，直接返回 0
    if root == nil {
        return 0;
    }

    // 计算根结点到当前结点的路径的数之和 currSum
    currSum += root.Val

    // 此时从根结点到当前结点的路径和为 currSum ，而我们需要求后缀和为 targetSum 的路径数，
    // 那么这个数量必定等于前缀和为 currSum - targetSum 的路径数。
    ans := preSum[currSum - targetSum]

    // 将 currSum 计入 preSum 中
    preSum[currSum] += 1

    // 递归处理左右子结点
    ans += dfs(root.Left, targetSum, currSum, preSum)
    ans += dfs(root.Right, targetSum, currSum, preSum)

    // 将本层计入的 currSum 从 preSum 中移除
    preSum[currSum] -= 1
    
    return ans
}
