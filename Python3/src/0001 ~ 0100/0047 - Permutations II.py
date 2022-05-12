# 链接：https://leetcode.com/problems/permutations-ii/
# 题意：给定一个可能含有重复数字的数组 nums ，求所有不同的排列。


# 数据限制：
#  1 <= nums.length <= 8
#  -10 <= nums[i] <= 10


# 输入： nums = [1,1,2]
# 输出： [[1,1,2],[1,2,1],[2,1,1]]
# 解释： 只有值不同的排列才认为是不同的排列。

# 输入： nums = [1,2,3]
# 输出： [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]


# 思路： Map + 递归/回溯/DFS
#
#      由于只有值不同的排列才算不同的排列，而 nums 中存在重复的值，
#      所以我们维护每个数可使用次数的 map ，从而避免产生重复的排列。
#
#      我们使用 dfs(num_to_cnt, remain, cur, ans) 回溯收集所有可能的排列，其中：
#          1. num_to_cnt: 当前每个数可使用次数的 map
#          2. remain: 当前还需要选择 remain 个数
#          3. cur: 当前已选择的数字列表
#          4. ans: 当前收集到的所有可能的组合的列表
#
#      在 dfs 中，我们按照如下逻辑处理即可：
#          1. remain == 0: 已选取完所有的数，则当前排列 cur 满足题意，
#              将 cur 放入到 ans 中，然后返回。
#          2. remain != 0: 则还需要继续选取数字，
#              遍历 num_to_cnt 中的每个数字 num ，
#              如果 num 还可以使用，则将 num 加入到 cur 中，
#              并调用 dfs 继续回溯。
#
#
#     时间复杂度：O(n * n!)
#          1. 需要遍历全部 O(n!) 个可能的排列
#          2. 每找到一个排列时，都需要遍历其中的全部 O(n) 个数字
#      空间复杂度：O(n * n!)
#          1. 需要分别维护 O(n) 的数组和 map
#          2. 需要存储全部 O(n!) 个可能的排列，每个排列的都含有 O(n) 个数字，
#              这部分总空间复杂度为 O(n * n!)
#          3. 栈递归深度为 O(n)


class Solution:
    def permuteUnique(self, nums: List[int]) -> List[List[int]]:
        # ans 用于收集所有可能的排列
        ans: List[List[int]] = []
        # cur 用于收集当前的排列
        cur: List[int] = [0] * len(nums)
        # num_to_cnt[num] 表示 num 可使用次数，
        # 初始化每个 num 的可使用次数次数
        num_to_cnt: Counter = Counter(nums)
        # 递归回溯所有可能的排列
        Solution.dfs(num_to_cnt, len(nums), cur, ans)
        return ans

    @staticmethod
    def dfs(num_to_cnt: Counter, remain: int, cur: List[int], ans: List[List[int]]):
        # 如果所有数字都已收集完，则将 cur 放入 ans 中，并返回
        if remain == 0:
            ans.append(cur[:])
            return

        # 计算倒数第 remain 个数字在 cur 中的下标
        i: int = len(cur) - remain
        # 遍历所有可能的数字 num
        for num, cnt in num_to_cnt.items():
            # 如果 num 还有可使用次数，则可以选择 num 加入 cur 中
            if cnt > 0:
                # num 可使用次数 -1
                num_to_cnt[num] -= 1
                # num 放入 cur 中
                cur[i] = num
                # 继续回溯
                Solution.dfs(num_to_cnt, remain - 1, cur, ans)
                # num 可使用次数 +1
                num_to_cnt[num] += 1
