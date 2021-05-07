# 链接：https://leetcode.com/problems/two-sum/
# 题意：给定一个数组 nums 和一个目标数 target，求和为 target 的两个数的下标
# 思路：用 dict 保存每个数出现的下标，遍历数组求目标数的下标，存在则直接返回，不存在则保存入 map 中

class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        # 记录每个数最后一次出现的下标
        num_to_index: Dict[int, int] = {}
        for i, num in enumerate(nums):
            # 从 dict 中获取需要的数的下标
            j = num_to_index.get(target - num)
            # 该数存在，则直接返回
            if j is not None:
                return [i, j]

            # 需要的数不存在，则更新 num 的下标
            num_to_index[num] = i

        # 题目保证一定有答案，所以不会走到这
        pass
