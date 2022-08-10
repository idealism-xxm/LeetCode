# 链接：https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/
# 题意：给定一个升序排序的整型数组 nums ，将其转换成一棵高度平衡的二叉搜索树？
#
#      高度平衡的二叉搜索树是指一棵二叉搜索树，其中每个结点的左右子树的深度相差不超过 1。


# 数据限制：
#  1 <= nums.length <= 10 ^ 4
#  -(10 ^ 4) <= nums[i] <= 10 ^ 4
#  nums 严格递增的


# 输入： nums = [-10,-3,0,5,9]
# 输出： [0,-3,9,-10,null,5]
# 解释： [0,-10,5,null,-3,null,9] 也是满足题意的（如下右边图示）。
#     0               0
#    / \             / \
#   -3  9          -10  5 
#  /   /             \    \
# -10 5              -3    9 


# 输入： nums = [1,3]
# 输出： [3,1]
# 解释： [1,3] 也是满足题意的（如下右边图示）。
#     3               1
#    /                 \
#   1                   3


# 思路：递归
#
#      我们使用 dfs(nums, start, end) 对 nums[start..=end] 处理，
#      生成一棵高度平衡的二叉搜索树。
#      
#      那么我们只要按照如下逻辑递归处理即可：
#          1. start > end: 说明用于生产二叉树的数组区间为空，
#              则对应的二叉树为空，直接返回空即可
#          2. start <= end: 
#              那么我们只需要让区间中点 mid = (start + end) / 2 对应的数作为根，
#              这样就能保证根节点是高度平衡的，因为左右子树的结点数相差不超过 1 。
#              然后用 dfs 递归处理 [start, mid - 1] 和 [mid + 1, end] 即可。
#
#
#		时间复杂度： O(n)
#          1. 需要遍历全部 O(n) 个数一次
#		空间复杂度： O(n)
#          1. 栈递归深度就是树高，因为每次都是二分区间，所以树高为 O(logn)
#          2. 需要维护树的全部 O(n) 个结点


# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def sortedArrayToBST(self, nums: List[int]) -> Optional[TreeNode]:
        # 返回 nums 对应的高度平衡的二叉搜索树
        return Solution.dfs(nums, 0, len(nums) - 1)

    @staticmethod
    def dfs(nums: List[int], start: int, end: int) -> Optional[TreeNode]:
        # 如果 start > end ，说明用于生产二叉树的数组区间为空，
        # 则对应的二叉树为空，返回 None
        if start > end:
            return None

        # 根节点就是区间 [start, end] 中间的数，这样就能保证根节点是高度平衡的，
        # 因为左右子树的结点数相差不超过 1
        mid: int = (start + end) // 2
        return TreeNode(
            nums[mid],
            # 递归生成左右子树
            Solution.dfs(nums, start, mid - 1),
            Solution.dfs(nums, mid + 1, end),
        )
