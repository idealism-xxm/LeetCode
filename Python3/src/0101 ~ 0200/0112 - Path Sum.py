# 链接：https://leetcode.com/problems/path-sum/
# 题意：给定一棵二叉树 root 和一个正整数 targetSum ，
#      判断是否存在一条根到叶子路径上所有数字之和为 targetSum ？


# 数据限制：
#  二叉树的结点数范围为 [0, 5000]
#  -1000 <= Node.val <= 1000
#  -1000 <= targetSum <= 1000


# 输入： root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22
# 输出： true
# 解释：
#         (5)       
#         /  \      
#       (4)   8     
#       /  \   \    
#     (11)  13  4
#     /  \        \
#    7   (2)       1

# 输入： root = [1,2,3], targetSum = 5
# 输出： false
# 解释： 只有两条从根到叶子的路径：
#          1. (1 -> 2): 和为 3
#          2. (1 -> 3): 和为 4
#       没有和为 5 的从根到叶子的路径

# 输入： root = [], targetSum = 0
# 输出： false
# 解释： 树为空，没有从根到叶子的路径


# 思路：递归/DFS
#
#      本题可以直接使用递归/DFS 进行处理，可以直接复用题目给定的函数定义即可，
#      按照以下逻辑进行处理：
#
#      1. 如果当前是空结点，则必定不存在路径，直接返回 false
#      2. 如果当前是叶子结点，则返回 root.val == targetSum
#      3. 如果左子结点存在，递归处理左子结点，并将 sum 减去当前结点的值
#          (1) 递归处理结果为 true ，直接返回 true
#          (2) 递归处理结果为 false ，继续执行以下逻辑
#      4. 若果右子结点存在，递归处理右子结点，并将 sum 减去当前结点的值
#          (1) 递归处理结果为 true ，直接返回 true
#          (2) 递归处理结果为 false ，继续执行以下逻辑
#      5. 左右子结点都不满足题意，返回 false
#
#      时间复杂度： O(n)
#          1. 需要遍历找到一条合法的路径，最差情况下，无合法的路径，
#              需要遍历全部 O(n) 个结点
#      空间复杂度： O(n)
#          1. 栈递归深度就是树高，最差情况下，全部 O(n) 个结点在一条链上


# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def hasPathSum(self, root: Optional[TreeNode], targetSum: int) -> bool:
        # 空结点不存在路径，直接返回 false
        if not root:
            return False

        # 如果是叶子结点，则仅当值等于 targetSum 时，才满足题意
        if not root.left and not root.right:
            return root.val == targetSum

        remainSum: int = targetSum - root.val
        # 当左子结点存在，且存在一条左子结点到叶子路径上所有值到和为 remainSum ，则满足题意
        if root.left and self.hasPathSum(root.left, remainSum):
            return True
        # 当右子结点存在，且存在一条右子结点到叶子路径上所有值到和为 remainSum ，则满足题意
        if root.right and self.hasPathSum(root.right, remainSum):
            return True

        # 左右子结点都不满足题意，返回 false
        return False
