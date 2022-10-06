# 链接：https://leetcode.com/problems/path-sum-iii/
# 题意：给定一棵二叉树 root 和一个正整数 targetSum ，
#      求所有和为 targetSum 的路径的数量？


# 数据限制：
#  二叉树的结点数范围为 [0, 1000]
#  -(10 ^ 9) <= Node.val <= 10 ^ 9
#  -1000 <= targetSum <= 1000


# 输入： root = [10,5,-3,3,2,null,11,3,-2,null,1], targetSum = 8
# 输出： 3
# 解释： 有三条路径的和为 8 ：
#       5 + 3 = 8
#       5 + 2 + 1 = 8
#       -3 + 11 = 8
#
#          10       
#         /  \      
#       (5)  (-3)
#       /  \   \    
#     (3)  (2)  (11)
#     /  \   \    \
#    3   -2  (1)   11

# 输入： root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
# 输出： 3
# 解释： 有三条路径的和为 22 ：
#       5 + 4 + 11 + 2 = 22
#       5 + 8 + 4 + 5 = 22
#       4 + 11 + 7 = 22
#
#          (5)       
#         /   \      
#       (4)   (8)     
#       /    /  \    
#     (11)  13  (4)
#     /  \      /  \
#    (7) (2)   (5)  1


# 思路：递归/回溯/DFS + 前缀和 + Map
#
#      本题是 LeetCode 0113 的加强版，要统计和为 target_sum 的数量，
#      需要记录从根结点到当前结点的路径的前缀和及其出现次数，并在每个结点处统计满足题意的路径。
#
#      这个 前缀和 + Map 思路就是 LeetCode 560 这题的思路，
#      本题其实就是将给定的数组变为了动态计算的路径。
#
#      本题需要从根结点到当前结点的路径的前缀和及其出现次数，
#      所以需要定义 dfs(root, target_sum, curr_sum, pre_sum) 来辅助处理，
#          1. root: 当前待处理的子树的根结点
#          2. target_sum: 路径上的数之和需要为 target_sum ，直接透传，不做修改
#          3. curr_sum: 从根结点到 root 的路径上的数之和
#          4. pre_sum: pre_sum[sum] 表示从根结点到 root 的路径中，前缀和为 sum 的路径数
#
#      返回值就是以 root 子树中每个结点为路径截至结点的所有路径中，满足题意的路径数之和。
#
#      dfs 按照以下逻辑进行处理：
#          1. 如果 root 是空结点，则不存在路径，则直接返回 0
#          2. 将 root.val 计入 curr_sum 中
#          3. 以 root 为路径截至结点的所有路径中，满足题意的路径必定是后缀和为 target_sum 的路径，
#             这些路径数量必定等于前缀和为 curr_sum - target_sum 的路径数，
#             即 ans = pre_sum[curr_sum - target_sum]
#          4. 将 curr_sum 计入 pre_sum 中，【注意】这一步需要在统计完成后，否则可能会错误统计入空路径
#          5. 递归处理左右子结点，将结果计入 ans 中
#          6. 退出递归前需要将本层计入 pre_sum 中前缀和 curr_sum 移除
#          7. 此时 ans 就是以 root 子树中每个结点为路径截至结点的所有路径中，满足题意的路径数之和
#
#
#      时间复杂度： O(n)
#          1. 需要遍历全部 O(n) 个结点
#      空间复杂度： O(n)
#          1. 栈递归深度就是树高，最差情况下，全部 O(n) 个结点在一条链上
#          2. 需要维护路径上全部不同的前缀和的出现次数，最差情况下，有 O(n) 个不同的前缀和


# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def pathSum(self, root: Optional[TreeNode], targetSum: int) -> int:
        # pre_sum[sum] 表示从根结点到当前结点的路径中，前缀和为 sum 的路径数
        pre_sum: Dict[int, int] = defaultdict(int)
        # 初始化前缀和为 0 的路径出现 1 次，方便后续计入从根结点到当前结点的路径
        pre_sum[0] = 1

        return Solution.dfs(root, targetSum, 0, pre_sum)

    @staticmethod
    def dfs(root: Optional[TreeNode], target_sum: int, curr_sum: int, pre_sum: Dict[int, int]) -> int:
        # 空结点不存在路径，直接返回 0
        if not root:
            return 0

        # 计算根结点到当前结点的路径的数之和 curr_sum
        curr_sum += root.val

        # 此时从根结点到当前结点的路径和为 curr_sum ，而我们需要求后缀和为 target_sum 的路径数，
        # 那么这个数量必定等于前缀和为 curr_sum - target_sum 的路径数。
        ans: int = pre_sum.get(curr_sum - target_sum, 0)

        # 将 curr_sum 计入 pre_sum 中
        pre_sum[curr_sum] += 1

        # 递归处理左右子结点，将结果计入 ans 中
        ans += Solution.dfs(root.left, target_sum, curr_sum, pre_sum)
        ans += Solution.dfs(root.right, target_sum, curr_sum, pre_sum)

        # 将本层计入的 curr_sum 从 pre_sum 中移除
        pre_sum[curr_sum] -= 1
        
        return ans
