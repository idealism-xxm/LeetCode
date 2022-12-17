# 链接：https://leetcode.com/problems/house-robber-iii/
# 题意：给定一棵二叉树，不能选择相邻的两个结点的数，求选择某些数的和的最大值？


# 数据限制：
#  树的结点数范围为 [1, 10 ^ 4]
#  0 <= Node.val <= 10 ^ 4


# 输入： root = [3,2,3,null,3,null,1]
# 输出： 7
# 解释： 选择 3, 3, 1 ，最大和为 3 + 3 + 1 = 7 
#      (3)
#      / \
#     2   3
#      \   \
#      (3) (1)

# 输入： root = [3,4,5,1,3,null,1]
# 输出： 9
# 解释： 选择 4, 5 ，最大和为 4 + 5 = 9
#       3
#      / \
#    (4) (5)
#    / \   \
#   1   3   1


# 思路： 树形 DP
#
#      本题是 LeetCode 213 的加强版，将环变成了树。
#
#      原来在环中只需要考虑相邻的数，而在树中需要考虑相邻层的数。
#
#      设 dp[root] = (pick[root], skip[root]) ：
#          1. pick[root] 表示在子树 root 中，选择 root 时，选择的数的最大和
#          2. skip[root] 表示在子树 root 中，不选 root 时，选择的数的最大和
#
#      对树的空结点初始化为 (0, 0) ，方便状态转移。
#            
#      状态转移：
#          1. 选择 root 时，则必定不能选其子结点：
#              pick[root] = root.val + skip[root.left] + skip[root.right]
#              
#          2. 不选 root 时，则子结点可选可不选，即子结点取两者最大值：
#              skip[root] = max(pick[root.left], skip[root.left]) +
#                           max(pick[root.right], skip[root.right])           
#
#      则最终结果就是 max(pick[root], skip[root])
#
#
#      时间复杂度： O(n)
#          1. 需要遍历计算树中全部 O(n) 个状态
#      空间复杂度： O(n)
#          2. 栈递归深度就是树高，最差情况下，全部 O(n) 个结点在一条链上


# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def rob(self, root: Optional[TreeNode]) -> int:
        # pick 表示在子树 root 中，选择 root 时，选择的数的最大和
        # skip 表示在子树 root 中，不选 root 时，选择的数的最大和
        pick, skip = Solution.dfs(root)
        # 两者取最大值返回即可
        return max(pick, skip)
    
    @staticmethod
    def dfs(root: Optional[TreeNode]) -> Tuple[int, int]:
        # 如果是空结点，则直接返回 (0, 0)
        if not root:
            return 0, 0

        # 递归计算左右子结点的状态
        left_pick, left_skip = Solution.dfs(root.left)
        right_pick, right_skip = Solution.dfs(root.right)
        
        # 选择 root 时，则必定不能选其子结点
        pick = root.val + left_skip + right_skip
        # 不选 root 时，则子结点可选可不选，即子结点取两者最大值
        skip = max(left_pick, left_skip) + max(right_pick, right_skip)
        
        return pick, skip
