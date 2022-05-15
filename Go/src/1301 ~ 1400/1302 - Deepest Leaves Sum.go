// 链接：https://leetcode.com/problems/deepest-leaves-sum/
// 题意：给定一棵二叉树，求最深的叶子节点值的和。


// 数据限制：
//  树中的结点数在 [1, 10 ^ 4] 内
//  1 <= Node.val <= 100


// 输入： root = [1,2,3,4,5,null,6,7,null,null,null,null,8]
// 输出： 15
// 解释： 1
//      / \
//     2   3
//    / \   \
//   4   5   6
//  /         \
// 7           8

// 输入： root = [6,7,8,2,7,1,3,9,null,1,4,null,null,null,5]
// 输出： 19
// 解释：   6
//      /     \
//     7       8
//    / \     / \
//   2   7   1   3
//  /   / \       \
// 9   1   4       5


// 思路： BFS
//
//      我们只需要按照 BFS 求出每一层的所有结点值的和 valSum ，
//      那么最后一层的所有结点值的和就是最深的叶子节点值的和。
//
//      定义一个队列 q ，初始化放入第一层的结点 root ；
//      同时维护我们最终的答案 ans ，初始化为 0 ；
//
//      当 q 不为空时，按照以下逻辑循环处理：
//          1. 获取当前 q 中的结点数 length ，
//              则 length 就是当前层所有的结点数。
//          2. 依次取出 q 中前 length 个结点，
//              计算这些结点值的和 valSum ，
//              然后将每个结点的左右子结点放入 q 队尾中。
//          3. 处理完钱 length 个结点后，当前层的所有结点已遍历完成，
//              令 ans = valSum ，继续处理下一层。
//
//      当 q 为空时， ans 就是最后一层的所有结点值的和。
//
//
//      时间复杂度：O(n)
//          1. 需要遍历全部 O(n) 个结点
//      空间复杂度：O(n)
//          1. 需要维护一个队列，最差情况下所有 O(n) 个结点都在队列中


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func deepestLeavesSum(root *TreeNode) int {
    // 维护最后一层的所有结点值的和
    ans := 0
    // 队列 q ，用于 BFS 按层遍历
    var q []*TreeNode
    // 初始化放入第一层的结点 root
    q = append(q, root)
    // 当 q 不为空时，按照以下逻辑循环处理
    for len(q) > 0 {
        // 当前层的所有结点值的和
        valSum := 0
        // 前 q.len() 个结点是当前层的结点
        for i := len(q); i > 0; i-- {
            // 取出队首结点 node
            node := q[0]
            q = q[1:]
            // valSum 加上当前结点的值
            valSum += node.Val
            // 如果 node 有左子结点，则放入队尾
            if node.Left != nil {
                q = append(q, node.Left)
            }
            // 如果 node 有右子结点，则放入队尾
            if node.Right != nil {
                q = append(q, node.Right)
            }
        }
        // ans 更新为当前层的所有结点值的和
        ans = valSum
    }

    // 此时 ans 就是最后一层的所有结点值的和
    return ans
}
