# 链接：https://leetcode.com/problems/combination-sum-iii/
# 题意：找到所有和为 num 的 k 个数的组合，同时需要满足以下两个条件：
#      1. 这个组合中的数字范围在 [1, 9] 之间
#      2. 这个组合中的数字不能重复


# 数据限制：
#  2 <= k <= 9
#  1 <= n <= 60


# 输入： k = 3, n = 7
# 输出： [[1,2,4]]
# 解释： 1 + 2 + 4 = 7
#       没有其他合法的组合了。

# 输入： k = 3, n = 9
# 输出： [[1,2,6],[1,3,5],[2,3,4]]
# 解释： 1 + 2 + 6 = 9
#       1 + 3 + 5 = 9
#       2 + 3 + 4 = 9
#       没有其他合法的组合了。

# 输入： k = 4, n = 1
# 输出： []
# 解释： 没有合法的组合，
#       因为最小能得到的数字 1 + 2 + 3 + 4 = 10 > 1


# 思路： 递归
#
#      我们使用 dfs(k, n, digit, cur, ans) 遍历收集所有可能的组合，其中：
#          1. k: 还需要选择的数字个数
#          2. n: 还需要选择的 k 个数字的和
#          3. digit: 要从 [digit, 9] 中选取数字
#          4. cur: 当前已选择的数字列表
#          5. ans: 当前收集到的所有可能的组合的列表
#
#      在 dfs 中，我们按照如下逻辑处理即可：
#          1. k == 0 || n == 0: 已选取完所有的数 或 和已经满足要求，需要返回。
#              如果此时有 k == 0 && n == 0，则说明当前组合 cur 满足题意，
#              需要将 cur 加入到 ans 中。
#          2. k != 0 && n != 0: 则还可以继续选取数字，
#              遍历 [digit, min(9, n)] 中的每个数字 num ，
#              并将 num 加入到 cur 中，并调用 dfs 继续递归。
#
#
#      时间复杂度：O(C(n, k) * k)
#          1. 递归终止时，组合最多选取 k 个数字，有 C(n, k) 种可能
#          2. 最差情况下，每个组合都满足题意，需要拷贝 cur 的 O(k) 个数字
#      空间复杂度：O(k)
#          1. 需要用 cur 保存选取的 O(k) 个数字的组合
#          2. 栈递归深度为 O(k)


class Solution:
    def combinationSum3(self, k: int, n: int) -> List[List[int]]:
        # ans 用于收集所有可能的组合
        ans = []
        # 递归回溯所有可能的组合
        Solution.dfs(k, n, 1, [0] * k, ans)
        return ans

    # 从 [digit, 9] 中选取和为 n 的 k 个数，并放在 cur 的后 k 个位置，
    # 最后结束递归时将合法的 cur 放入到 ans 中
    @staticmethod
    def dfs(k: int, n: int, digit: int, cur: List[int], ans: List[List[int]]):
        # 如果已选取完所有的数 或 和已经满足要求，则需要返回
        if k == 0 or n == 0:
            # 如果已选取完所有的数 且 和已经满足要求，
            # 则当前 cur 是合法的组合，放入 ans 中
            if k == 0 and n == 0:
                ans.append(cur[:])

            return
        
        # 计算倒数第 k 个数在 cur 中的下标
        i: int = len(cur) - k
        # 遍历 [digit, min(9, n)] 中的数字，
        # 这里取 min(9, n) 是使得后续的 n >= 0 ，减少不必要的判断
        for num in range(digit, min(9, n) + 1):
            # 将倒数第 k 个数赋值为 num
            cur[i] = num
            # 递归回溯所有可能的组合
            Solution.dfs(k - 1, n - num, num + 1, cur, ans)
