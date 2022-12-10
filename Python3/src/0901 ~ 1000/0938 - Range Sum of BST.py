# 链接：https://leetcode.com/problems/range-sum-of-bst/
# 题意：给定一颗二叉搜索树 root 和两个整数 low, high ，
#      求树中所有结点值在 [low, high] 内结点值之和？


# 数据限制：
#  树的结点数范围为 [1, 2 * 10 ^ 4]
#  1 <= Node.val <= 10 ^ 5
#  1 <= low <= high <= 10 ^ 5
#  所有 Node.val 都各不相同


# 输入： root = [10,5,15,3,7,null,18], low = 7, high = 15
# 输出： 32
# 解释： 结点值 7, 10, 15 在 [7, 15] 内， 7 + 10 + 15 = 32

# 输入： root = [10,5,15,3,7,13,18,1,null,6], low = 6, high = 10
# 输出： 23
# 解释： 结点值 6, 7, 10 在 [6, 10] 内， 6 + 7 + 10 = 23


# 思路： 递归/DFS
#
#      最朴素的想法就是使用先序遍历全部结点，
#      统计所有结点值在 [low, high] 内结点值之和。
#
#      本题给定的是一棵二叉搜索树，保证了以下两点：
#          1. 左子树中所有结点值都小于 root.val
#          2. 右子树中所有结点值都大于 root.val
#
#      所以我们可以利用这个特性进行剪枝优化，降低时间复杂度的常数：
#          1. 只有 low < root.val 时，才递归处理左子树
#          2. 只有 root.val < high 时，才递归处理右子树
#
#
#      时间复杂度：O(n)
#          1. 需要遍历树中全部结点值在 [low, high] 内的结点，
#              最差情况下全部 O(n) 个结点的值都在 [low, high] 内
#      空间复杂度：O(n)
#          1. 栈递归深度就是树高，最差情况下，全部 O(n) 个结点在一条链上


# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def rangeSumBST(self, root: Optional[TreeNode], low: int, high: int) -> int:
        # 如果 root 子树为空，则直接返回 0
        if not root:
            return 0

        # ans 维护 root 子树中，所有结点值在 [low, high] 内结点值之和
        ans: int = 0
        # 如果 root 结点值在 [low, high] 内，则计入 ans
        if low <= root.val and root.val <= high:
            ans += root.val

        # 如果 root 结点值大于 low ，则其左子树可能存在需要统计的结点，递归处理左子树
        if low < root.val:
            ans += self.rangeSumBST(root.left, low, high)
        # 如果 root 结点值小于 high ，则其右子树可能存在需要统计的结点，递归处理右子树
        if root.val < high:
            ans += self.rangeSumBST(root.right, low, high)

        return ans
