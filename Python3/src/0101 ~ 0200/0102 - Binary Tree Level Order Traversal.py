# 链接：https://leetcode.com/problems/binary-tree-level-order-traversal/
# 题意：给定一棵二叉树 root ，返回其层序遍历的结果（从左到右，按层分组）。


# 数据限制：
#  树中结点的数量范围是 [0, 2000]
#  -1000 <= Node.val <= 1000


# 输入：root = [3,9,20,null,null,15,7]
# 输出：[[3],[9,20],[15,7]]
# 解释： 3
#      / \
#     9  20
#       /  \
#      15   7

# 输入：root = [1]
# 输出：[[1]]

# 输入：root = []
# 输出：[]


# 思路：BFS
#
#      普通的层序遍历可以直接使用 BFS 处理，但本题需要将同一层的值收集到一个数组中，
#      所以需要特殊处理一下。
#
#      只要我们知道当前层的结点数 num ，从队列 queue 中只取前 num 个结点处理，
#      这样就能将 BFS 转换成满足题意的层序遍历。
#
#      那么如何知道当前层的结点数 num ？
#
#      可以发现，只要我们是一层一层处理的，那么在处理当前层的结点前，
#      queue 中的所有结点都是当前层的结点，即当前层的结点数 num = queue.len() 。
#
#      所以我们在 BFS 循环中再加一层循环进行遍历，外部循环仅控制层序遍历是否结束，
#      内部循环才进行真正的处理逻辑，内部循环中只取 queue 中前 num 个结点，
#      收集这些结点的值，并将其非空左右子结点加入 queue 中。
#
#
#      时间复杂度： O(n)
#          1. 需要遍历全部 O(n) 个结点
#      空间复杂度： O(n)
#          1. 需要维护结果数组 ans ，保存全部 O(n) 个结点的值
#          2. 需要维护队列 queue ，保存全部 O(n) 个结点


# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def levelOrder(self, root: Optional[TreeNode]) -> List[List[int]]:
        ans: List[List[int]] = []
        # 如果根结点为空，则直接返回空的结果
        if not root:
            return ans

        # 把根结点放入队列中，准备进行 BFS
        queue: deque = deque([root])
        # 队列中还有结点，则可以继续进行 BFS
        while len(queue):
            # level 按顺序收集当前层的值
            level: List[int] = []
            # 此时 queue 中的结点数量为当前层的结点数量，
            # 只从队列中这些数量的结点，这样就变成了层序遍历
            for _ in range(len(queue)):
                node = queue.popleft()
                # 收集这些结点的值到 level 中
                level.append(node.val)
                # 将其非空左右子结点放入队列中
                if node.left:
                    queue.append(node.left)
                if node.right:
                    queue.append(node.right)

            ans.append(level)

        return ans
