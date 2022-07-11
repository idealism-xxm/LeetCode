# 链接：https:#leetcode.com/problems/binary-tree-right-side-view/
# 题意：给定一棵二叉树，想象你站在它右侧，
#      按照从顶部到底部的顺序，返回你能看见的结点值的数组。


# 数据限制：
#  二叉树的结点数的范围为 [0, 100]
#  -100 <= Node.val <= 100


# 输入： root = [1,2,3,null,5,null,4]
# 输出： [1,3,4]
# 解释： 1
#      / \
#     2   3
#      \   \
#       5   4

# 输入： root = [1,null,3]
# 输出： [1,3]
# 解释： 1
#        \
#         3

# 输入： root = []
# 输出： []


# 思路： 递归/DFS
#
#      用 ans 维护每一层最右侧结点的值，
#      然后使用 dfs(root, depth, ans) 先序遍历收集每一层最右侧结点的值即可。
#
#      dfs(root, depth, ans) 的参数含义如下：
#          1. root: 当前处理的 root 子树，初始为二叉树的根
#          2. depth: root 结点的深度，初始二叉树的根的深度为 1
#          3. ans: 已收集的每一层最右侧结点的值，初始为空
#
#      在 dfs 中，如果 root 不为空，则 root 有可能是当前层最右侧的结点，
#      可以通过 depth 和 len(ans) 的关系来判断：
#          1. depth <= len(ans): 已收集了前 len(ans) 最右侧结点的值，
#              则此时 root 不是当前层最右侧的结点，不做处理
#          2. depth > len(ans): 则必有 depth = len(ans) + 1 ，
#              此时 root 就是当前层最右侧的结点，将其值收集到 ans 中
#
#      然后递归遍历子树，注意先递归遍历右子树，再递归遍历左子树，
#      这样后续遇到每一层的第一个结点，必定是该层最右侧结点。
#
#
#		时间复杂度： O(n)
#          1. 需要遍历全部 O(n) 个结点一次
#		空间复杂度： O(n)
#          1. 栈递归深度就是树高，最差情况下，全部 O(n) 个结点在一条链上
#          2. 需要记录每一层最右侧结点的值，也就是树高，
#              最差情况下，全部 O(n) 个结点在一条链上


# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def rightSideView(self, root: Optional[TreeNode]) -> List[int]:
        # ans 用于收集每一层最右侧结点的值
        ans: List[int] = []
        # dfs 先序遍历收集每一层最右侧结点的值
        Solution.dfs(root, 1, ans)

        return ans

    @staticmethod
    def dfs(root: Optional[TreeNode], depth: int, ans: List[int]) -> None:
        # 如果 root 为空，则直接返回
        if not root:
            return

        # 如果深度大于 ans 的长度，则 root.val 是当前层最右侧结点的值，需要收集
        if depth > len(ans):
            ans.append(root.val)

        # 先递归遍历右子树，在递归遍历左子树，
        # 这样后续遇到每一层的第一个结点，必定是该层最右侧结点
        Solution.dfs(root.right, depth + 1, ans)
        Solution.dfs(root.left, depth + 1, ans)
