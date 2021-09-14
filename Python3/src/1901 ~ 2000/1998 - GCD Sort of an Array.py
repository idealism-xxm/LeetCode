# 链接：https://leetcode.com/problems/gcd-sort-of-an-array/
# 题意：给定一个数组 nums ，你可以进行任意次如下操作，
#       求经过一些操作后， nums 是否能按照升序排列？
#
#       操作：如果两个数的最大公约数大于 1 ，则可以交换这两个数。
#           即： gcd(nums[i], nums[j]) > 1 ，可以交换 nums[i] 和 nums[j]

# 数据限制：
#   1 <= nums.length <= 3 * 10 ^ 4
#   2 <= nums[i] <= 10 ^ 5

# 输入： nums = [7,21,3]
# 输出： true
# 解释：
#   - [7,21,3] -> [21,7,3]: gcd(7,21) = 7 > 1 ，可以交换
#   - [21,7,3] -> [3,7,21]: gcd(21,3) = 3 > 1 ，可以交换

# 输入： nums = [5,2,6,2]
# 输出： false
# 解释：
#   5 不能和任何数交换

# 输入： nums = [10,5,9,3,15]
# 输出： true
# 解释：
#   - [10,5,9,3,15] -> [15,5,9,3,10]: gcd(10,15) = 5 > 1 ，可以交换
#   - [15,5,9,3,10] -> [3,5,9,15,10]: gcd(15,3) = 3 > 1 ，可以交换
#   - [3,5,9,15,10] -> [3,5,9,10,15]: gcd(10,15) = 5 > 1 ，可以交换


# 思路： 并查集
#
#       比赛的时候已经想到了使用并查集判断，但错误地使用了全部数进行预处理，导致结果错误
#
#       如果两个数的最大公约数大于 1 ，则它们可以放入一个集合中，
#       我们可以用并查集维护这个信息，最后在同一个集合中的数字可以任意交换。
#
#       那么我们可以按照升序重新排序同一个集合中的数字，最后检查整个数组是否为升序即可
#
#       时间复杂度： O(nlogn + mlogm)
#       空间复杂度： O(n + m)


class Solution:
    def gcdSort(self, nums: List[int]) -> bool:
        unique_nums = set(nums)
        mx = max(unique_nums)
        # 最开始每个数字和自己是一个集合
        parent = [i for i in range(mx + 1)]

        def find(x: int) -> int:
            while x != parent[x]:
                parent[x] = parent[parent[x]]
                x = parent[x]

            return x

        def union(x: int, y: int):
            x, y = find(x), find(y)
            if x != y:
                parent[y] = x

        # is_prime[i] 表示 i 是否为素数
        is_prime = [False, False] + [True] * (mx - 1)
        for i in range(2, mx + 1):
            # 如果 i 是素数，则可以将他的倍数都设置为非素数
            if is_prime[i]:
                # lst 存储在 nums 中且为 i 的倍数的数字
                lst = [i] if i in unique_nums else []
                for j in range(i + i, mx + 1, i):
                    is_prime[j] = False
                    if j in unique_nums:
                        lst.append(j)
                
                # 这些数字都属于同一个集合，可以进行合并
                for j in range(1, len(lst)):
                    union(lst[0], lst[j])

        # 如果所有相同位置的数与排序后期望的数在同一个集合中
        # 那么他们必定能通过一些操作后互换
        return all(
            find(num) == find(expected_num)
            for num, expected_num in zip(nums, sorted(nums))
        )
