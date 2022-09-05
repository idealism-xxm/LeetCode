# 链接：https://leetcode.com/problems/n-ary-tree-level-order-traversal/
# 题意：给定一棵多叉树 root ，返回其层序遍历的结果（从左到右，按层分组）。


# 数据限制：
#  树高最大为 1000
#  树中结点的数量范围是 [0, 10 ^ 4]


# 输入：root = [1,null,3,2,4,null,5,6]
# 输出：[[1],[3,2,4],[5,6]]
# 解释：  1
#      / | \
#     3  2  4
#    / \
#   5   6

# 输入：root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
# 输出：[[1],[2,3,4,5],[6,7,8,9,10],[11,12,13],[14]]
# 解释：     1
#    /  /   |   \
#   2  3    4    5
#     / \   |   / \
#    6   7  8  9  10
#        |  |  |
#       11  12 13
#        |
#       14


# 思路：BFS
#
#      本题是 LeetCode 102 的加强版，将二叉树换成了多叉树，
#      可以直接化用对应的思路和代码。
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
#      收集这些结点的值，并将其非空子结点加入 queue 中。
#
#
#      时间复杂度： O(n)
#          1. 需要遍历全部 O(n) 个结点
#      空间复杂度： O(n)
#          1. 需要维护结果数组 ans ，保存全部 O(n) 个结点的值
#          2. 需要维护队列 queue ，保存全部 O(n) 个结点


"""
# Definition for a Node.
class Node:
    def __init__(self, val=None, children=None):
        self.val = val
        self.children = children
"""
class Solution:
    def levelOrder(self, root: 'Node') -> List[List[int]]:
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
                # 将其非空子结点放入队列中
                for child in node.children:
                    queue.append(child)

            ans.append(level)

        return ans
