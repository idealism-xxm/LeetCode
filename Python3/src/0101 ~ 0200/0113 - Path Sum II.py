# 链接：https://leetcode.com/problems/path-sum-ii/
# 题意：给定一棵二叉树 root 和一个正整数 targetSum ，
#      求所有和为 targetSum 的根到叶子的路径？


# 数据限制：
#  二叉树的结点数范围为 [0, 5000]
#  -1000 <= Node.val <= 1000
#  -1000 <= targetSum <= 1000


# 输入： root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
# 输出： [[5,4,11,2],[5,8,4,5]]
# 解释： 有两条根到叶子的路径的和为 22 ：
#       5 + 4 + 11 + 2 = 22
#       5 + 8 + 4 + 5 = 22
#
#         (5)       
#         /  \      
#       (4)  (8)     
#       /    / \    
#     (11)  13 (4)
#     /  \     /  \
#    7   (2)  (5)  1

# 输入： root = [1,2,3], targetSum = 5
# 输出： false
# 解释： 只有两条从根到叶子的路径：
#          1. (1 -> 2): 和为 3
#          2. (1 -> 3): 和为 4
#       没有和为 5 的从根到叶子的路径

# 输入： root = [1,2], targetSum = 0
# 输出： false
# 解释： 只有一条从根到叶子的路径： 1 + 2 = 3 ，
#       没有和为 0 的从根到叶子的路径


# 思路：递归/回溯/DFS
#
#      本题是 LeetCode 0112 的加强版，不过需要回溯全部可能的结果，
#      需要记录从根结点到当前结点的所有数，最后在叶子结点处判断收集可能的路径。
#
#      本题需要维护路径上的数，所以需要定义 dfs(root, targetSum, nums, ans) 来辅助处理。
#          1. root: 当前待处理的子树的根结点
#          2. targetSum: root 到其叶子结点的路径上数之和需要为 targetSum
#          3. nums: 从根结点到 root 的所有数
#          4. ans: 用来收集所有可能的路径
#
#      dfs 按照以下逻辑进行处理：
#          1. 如果 root 是空结点，则不存在路径，则直接返回
#          2. 将 root.val 放入 nums 中，并从 targetSum 中减去 root.val
#          3. 如果 targetSum 为 0 且 root 是叶子结点，则当前路径满足题意，将 nums 放入 ans
#          4. 递归处理左子结点，直接调用 dfs(root.left, targetSum, nums, ans)
#          5. 递归处理右子结点，直接调用 dfs(root.right, targetSum, nums, ans)
#          6. 退出递归前需要将本层放入 nums 中的最后一个数移除
#
#
#      时间复杂度： O(n ^ 2)
#          1. 需要遍历全部 O(n) 个结点
#          2. 最差情况下，前一半结点是一条链，后一半结点是完全二叉树，且所有路径都满足题意。
#              那么共有 O(1/4 * n) = O(n) 个叶子结点，
#              每条满足题意的路径有 O(1/2 * n + log(1/2 * n)) = O(n) 个结点，
#              每次到叶子结点时都需要拷贝路径一份，总时间复杂度为 O(n ^ 2)
#      空间复杂度： O(n ^ 2)
#          1. 栈递归深度就是树高，最差情况下，全部 O(n) 个结点在一条链上
#          2. 按照时间复杂度中的分析可知，最差情况下共有 O(n) 个叶子结点，
#              每个叶子结点都会产生一个长度为 O(n) 的路径，
#               总共需要维护结果全部 O(n ^ 2) 个结点的值


# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def pathSum(self, root: Optional[TreeNode], targetSum: int) -> List[List[int]]:
        ans: List[List[int]] = []
        Solution.dfs(root, targetSum, [], ans)
        
        return ans

    @staticmethod
    def dfs(root: Optional[TreeNode], targetSum: int, nums: List[int], ans: List[List[int]]):
        # 空结点不存在路径，直接返回
        if not root:
            return

        # root.val 放入 nums 中，并从 targetSum 中减去 root.val
        nums.append(root.val)
        targetSum -= root.val

        # 如果 targetSum 为 0 且 root 是叶子结点，则当前路径满足题意
        if targetSum == 0 and not root.left and not root.right:
            ans.append(nums[:])

        # 递归处理左右子结点
        Solution.dfs(root.left, targetSum, nums, ans)
        Solution.dfs(root.right, targetSum, nums, ans)

        # 将本层放入 nums 中的最后一个数移除
        nums.pop()
