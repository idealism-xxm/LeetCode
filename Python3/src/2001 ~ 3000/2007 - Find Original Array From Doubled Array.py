# 链接：https://leetcode.com/problems/find-original-array-from-doubled-array/
# 题意：有一个整型数组 original ，现在将其每个元素乘以 2 后再放到原数组的末尾形成 changed 数组，
#       最后给出 changed 数组随机改变顺序的数组，求一个可能的 original ？

# 数据限制：
#   1 <= changed.length <= 10 ^ 5
#   0 <= changed[i] <= 10 ^ 5

# 输入： changed = [1,3,4,2,6,8]
# 输出： [1,3,4]
# 解释：
#   [1,3,4] 每个元素乘以 2 后得 [2,6,8]
#   [1,3,4,2,6,8] 就是题目给出的 changed

# 输入： changed = [6,3,0,1]
# 输出： []
# 解释： changed 不是一个变换后的数组

# 输入： changed = [1]
# 输出： []
# 解释： changed 不是一个变换后的数组


# 思路： 贪心
#
#       changed 排序后，然后统计每个数字的个数，
#       再从小枚举每个还有的数字 num ，如果 2 * num 还有则可以减少 2 * num 的一次出现次数，
#       如果有一次发现 2 * num 没有，则直接返回 [] ，否则最后枚举的所有 num 则构成 original
#       
#
#       时间复杂度： O(nlogn)
#       空间复杂度： O(n)


class Solution:
    def findOriginalArray(self, changed: List[int]) -> List[int]:
        # 按从小到大排序
        changed.sort()
        # 统计每个数字出现的次数
        cnt = defaultdict(int)
        for num in changed:
            cnt[num] += 1
        
        ans = []
        for num in changed:
            # 如果 num 还有，则其可以作为 original 的一个数字
            while cnt[num]:
                # num 可用次数减 1
                cnt[num] -= 1
                ans.append(num)
                # 如果 num * 2 还有，则当前 num 是合法的
                if cnt[num << 1]:
                    # num * 2 可用次数减 1
                    cnt[num << 1] -= 1
                else:
                    # 如果 num * 2 没有，则当前 num 不合法，直接返回 []
                    return []
        return ans
