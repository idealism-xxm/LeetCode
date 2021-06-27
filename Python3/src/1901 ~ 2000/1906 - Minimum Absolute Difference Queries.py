# 链接：https://leetcode.com/problems/minimum-absolute-difference-queries/
# 题意：给定一个整型数组 arr 和一组查询 queries ，
#       每个查询 query = [l, r] 查询 arr[l:r + 1] 内
#       所有数对 (arr[i], arr[j]) 差的最小差，
#       （如果 arr[i] == arr[j] ，则不计入，即最小的差为 1 ，
#           如果所有的数都相等，则用 -1 表示）

# 数据限制：
#   2 <= nums.length <= 10 ^ 5
#   1 <= nums[i] <= 100
#   1 <= queries.length <= 2 * 10 ^ 4
#   0 <= l_i < r_i < nums.length

# 输入： nums = [1,3,4,8], queries = [[0,1],[1,2],[2,3],[0,3]]
# 输出： [2,1,4,1]
# 解释：
#   queries[0] = [0,1]: 子数组 [1,3] 的最小差为 |1 - 3| = 2
#   queries[1] = [1,2]: 子数组 [3,4] 的最小差为 |3 - 4| = 1.
#   queries[2] = [2,3]: 子数组 [4,8] 的最小差为 |4 - 8| = 4.
#   queries[3] = [0,3]: 子数组 [1,3,4,8] 的最小差为 |3 - 4| = 1.

# 输入： nums = [4,5,2,2,7,10], queries = [[2,3],[0,2],[0,5],[3,5]]
# 输出： [-1,1,1,3]
# 解释：
#   queries[0] = [2,3]: 子数组 [2,2] 数字都相同，所以结果是 -1
#   queries[1] = [0,2]: 子数组 [4,5,2] 的最小差为 |4 - 5| = 1.
#   queries[2] = [0,5]: 子数组 [4,5,2,2,7,10] 的最小差为 |4 - 5| = 1.
#   queries[3] = [3,5]: 子数组 [2,7,10] 的最小差为 |7 - 10| = 3.

# 思路： 前缀和
#
#       比赛时读题时理解错了，也没注意到数组值的范围，就直接写了个线段树，
#       写完发现错了，重新读题才搞清楚，也刚好看到了数组值的范围。
#       发现改造一下线段树即可，
#       复杂度略微高一点： O(nlogn * k + qlogn * k), 其中 k = 100
#       （不过最后当前数据下，发现时间差不多，内存高一倍）
#
#       赛后发现可以直接反转一下思路，由于数组值的范围很小，
#       所以我们每次枚举 1 ~ 100 是否在区间内即可，然后和前一个在区间的数求差，
#       再更新答案即可
#
#       可以使用前缀和维护 [1, i] 内某数字出现的次数，
#       然后就可以在 O(1) 内查询到一个数是否在区间内
#
#       时间复杂度： O(n * k + q * k), 其中 k = 100
#       空间复杂度： O(n * k)


class Solution:
    def minDifference(self, nums: List[int], queries: List[List[int]]) -> List[int]:
        sm: List[List[int]] = [None] * 101
        for i in range(1, 101):
            sm[i] = [0] * (len(nums) + 1)
        for j in range(1, len(nums) + 1):
            # 所有的出现的次数，都等于上次出现的次数
            for i in range(1, 101):
                sm[i][j] = sm[i][j - 1]
            # 只有本次出现的那个数字需要 +1
            sm[nums[j - 1]][j] += 1

        # 记录每次查询的结果
        ans = []
        for l, r in queries:
            # 初始化
            mn = 101
            pre = -101
            # 遍历每个数
            for i in range(1, 101):
                # 如果该数在 [l, r] 内
                if sm[i][r + 1] - sm[i][l] > 0:
                    # 更新最小差
                    mn = min(mn, i - pre)
                    # 更新上一个出现的值
                    pre = i
            # 将当前找打的最小值放入
            if mn == 101:
                ans.append(-1)
            else:
                ans.append(mn)
        return ans
