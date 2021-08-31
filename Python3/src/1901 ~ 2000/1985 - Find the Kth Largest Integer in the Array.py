# 链接：https://leetcode.com/problems/find-the-kth-largest-integer-in-the-array/
# 题意：给定一个字符串表示的数字数组，返回第 k 大的数？

# 数据限制：
#   1 <= k <= nums.length <= 10 ^ 4
#   1 <= nums[i].length <= 100
#   nums[i] 仅由数字组成
#   nums[i] 不含前导零

# 输入： nums = ["3","6","7","10"], k = 4
# 输出： "3"
# 解释： 
#       升序排序后为 ["3","6","7","10"] ，第 4 大的数为 "3"

# 输入： nums = ["2","21","12","1"], k = 3
# 输出： "2"
# 解释： 
#       升序排序后为 ["1","2","12","21"] ，第 3 大的数为 "2"

# 输入： nums = ["0","0"], k = 2
# 输出： "0"
# 解释： 
#       升序排序后为 ["0","0"] ，第 3 大的数为 "0"

# 思路： 排序
#
#       写一个正数的比较函数，然后按照从大到小排序即可，最后返回第 k 大的数
#       （ Python 支持大数，可以直接转换成 int 类型排序）
#
#       时间复杂度： O(nlogn)
#       空间复杂度： O(1)

def cmp(a: str, b: str) -> int:
    """
    比较两个字符串的大小
    """
    # 如果长度不等，则长度更长的数更大
    if len(a) != len(b):
        return len(a) - len(b)
    # 如果长度相等，则比较字符串的大小
    for ach, bch in zip(a, b):
        # 对应位置字符不等，则数字大的数更大
        if ord(ach) != ord(bch):
            return ord(ach) - ord(bch)
    # 都相等，则两个字符串数字相同
    return 0

# 将比较函数转换成 key
cmp_key = cmp_to_key(cmp)

class Solution:
    def kthLargestNumber(self, nums: List[str], k: int) -> str:
        # 降序排序后返回第 k 个即可
        return sorted(nums, key=cmp_key , reverse=True)[k - 1]
