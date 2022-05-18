// 链接：https://leetcode.com/problems/populating-next-right-pointers-in-each-node-ii/
// 题意：给定一棵二叉树，每个结点 cur 有 val, left, right, next 四个字段，
//      其中 cur.next 指向当前层 cur 右侧的结点，
//      如果 cur 右侧没有结点，则设置为 NULL 。
//
//      初始时所有结点的 next 都是 NULL ，我们需要设置正确的 next 。


// 数据限制：
//  树的结点数在 [0, 6000] 内
//  -100 <= Node.val <= 100


// 输入： root = [1,2,3,4,5,null,7]
// 输出： [1,#,2,3,#,4,5,7,#]
// 解释： 1            1    
//      / \          / \
//     2   3   =>   2 → 3
//    / \   \      / \   \
//   4   5   7    4 → 5 → 7
//
//      序列化的结果是按照每层的 next 顺序处理的，
//      每一层的末尾用 '#' 标示。

// 输入： root = []
// 输出： []


// 思路： 迭代
//
//      假设第 i 层的 next 已经全部处理完，
//      且第 i 层对应的最左侧结点为 leftMost 。
//
//      则我们可以通过遍历 leftMost 基于 next 的链表，
//      将第 i + 1 层的所有结点的 next 处理完成。
//
//      第 1 层的 next 不需要处理即符合题意，其 leftMost = root 。
//
//      处理第 i 层的 leftMost 时，
//      我们使用 cur 遍历当前层的结点，初始化 cur = leftMost 。
//
//      同时我们维护第 i + 1 层的链表结点 nextHeadPre 和 nextLast ，
//      其中 nextHeadPre 是第 i + 1 层链表的哨兵结点，方便后续处理，
//      nextLast 是第 i + 1 层链表的最后一个结点，方便使用尾插法插入新结点。
//
//      遍历过程中，判断 cur 的左右子结点：
//         1. cur 的左子结点存在，则将其挂在链上，可以进行更新：
//              last.next = cur.left
//              last = cur.left
//         2. cur 的右子结点存在，则将其挂在链上，可以进行更新：
//              last.next = cur.right
//              last = cur.right
//
//      然后更新 cur = cur.next ，直至 cur 为空。
//
//      此时，我们将 nextHeadPre.next 指向第 i + 1 层的最左侧结点，
//      赋值给 leftMost 后，继续处理第 i + 1 层即可。
// 
//
//      时间复杂度：O(n)
//          1. 需要遍历全部 O(n) 个结点
//      空间复杂度：O(1)
//          1. 只需要维护常数个额外变量即可


/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
    // leftMost 表示每一层最左侧结点，
    // 初始化为第一层最左侧结点 root
    leftMost := root
    // 当前层还有结点时，可以继续处理下一次的 next 指针
    for leftMost != nil {
        // 维护一个下一层基于 next 的链表，
        // 为了方便处理，这里使用一个哨兵节点
        nextHeadPre := &Node{}
        // 下一层基于 next 的链表的最后一个结点，
        // 方便使用尾插法插入新结点
        nextLast := nextHeadPre
        // 当前层的结点
        cur := leftMost
        // 如果当前层还有结点，则继续处理
        for cur != nil {
            // 如果存在左子结点，则将其加入到下一层的链表尾部
            if cur.Left != nil {
                nextLast.Next = cur.Left
                nextLast = nextLast.Next
            }
            // 如果存在左子结点，则将其加入到下一层的链表尾部
            if cur.Right != nil {
                nextLast.Next = cur.Right
                nextLast = nextLast.Next
            }
            // 继续处理当前层的下一个结点
            cur = cur.Next
        }
        // nextHeadPre.Next 就是下一层的最左侧结点
        leftMost = nextHeadPre.Next
    }
    return root
}
