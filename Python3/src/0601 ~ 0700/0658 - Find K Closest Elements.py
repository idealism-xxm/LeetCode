# 链接：https://leetcode.com/problems/find-k-closest-elements/
# 题意：给定一个升序排序的数组 arr ，返回其中与 x 最接近的 k 个数（结果数组按升序排序）。
#
#      当 a, b, x 满足以下两个条件之一时， a 与 x 比 b 与 x 更接近时。
#          1. |a - x| < |b - x|
#          2. |a - x| == |b - x| && a < b


# 数据限制：
#  1 <= k <= arr.length
#  1 <= arr.length <= 10 ^ 4
#  arr 是升序排序的
#  -(10 ^ 4) <= arr[i], x <= 10 ^ 4


# 输入： arr = [1,2,3,4,5], k = 4, x = 3
# 输出： [1,2,3,4]

# 输入： arr = [1,2,3,4,5], k = 4, x = -1
# 输出： [1,2,3,4]


# 思路： 二分 + 双指针
#
#      按照题意，我们必定是先找 arr 中最接近 x 的那个数，
#      再找次接近 x 的数，直至找满 k 个数。
#
#      但如果直接遍历，时间复杂度为 O(kn) ，在给定的数据范围内无法通过。
#
#      不过可以注意到给定的数组 arr 是有序的，
#      我们可以先用二分找到第一个大于等于 x 的数的下标 i ，
#      那么此时最接近 x 的数是 arr[i - 1] 或 arr[i] 。
#
#      接下来可以通过双指针的方式找到最接近 x 的 k 个数。
#
#      双指针 l, r 表示当前最接近 x 的 r - l 个数是 arr[l:r] ，
#      那么此时下一个最接近 x 的数是 arr[l - 1] 或 arr[r] 。
#
#      初始化 l 和 r 均为 i ，表示结果窗口为空，刚好符合前面的二分结果。
#
#      然后将双指针向两边共移动 k 次，每次将当前最接近 x 的数纳入结果窗口中：
#          1. |arr[l - 1] - x| <= |arr[r] - x|: arr[l - 1] 更接近 x ，
#              则左移左指针，将其纳入结果窗口中，即 l -= 1
#          2. |arr[l - 1] - x| > |arr[r] - x|: arr[r] 更接近 x ，
#              则右移右指针，将其纳入结果窗口中，即 r += 1
#
#
#      时间复杂度：O(logn + k)
#          1. 需要使用二分找到全部 O(n) 个数中最接近 x 的那个
#          2. 需要移动指针 O(k) 次，找到最接近 x 的 O(k) 个数
#      空间复杂度：O(1)
#          1. 只需要维护常数个额外变量即可


class Solution:
    def findClosestElements(self, arr: List[int], k: int, x: int) -> List[int]:
        n: int = len(arr)
        # 用二分找到第一个大于等于 x 的数的下标 i ，
        # 那么此时最接近 x 的数是 arr[i - 1] 或 arr[i]
        i: int = bisect.bisect_left(arr, x)
        # 接下来可以通过双指针的方式找到最接近 x 的 k 个数。
        # 双指针 l, r 表示当前最接近 x 的 r - l 个数是 arr[l:r] ，
        # 那么此时下一个最接近 x 的数是 arr[l - 1] 或 arr[r] 。
        # 初始化 l 和 r 均为 i ，表示结果窗口为空，刚好符合前面的二分结果。
        l, r = i, i
        # 将双指针向两边共移动 k 次，每次将当前最接近 x 的数纳入结果窗口中
        for _ in range(k):
            # 如果左指针 l 已到边界，则 arr[:k] 就是满足题意的结果
            if l == 0:
                return arr[:k]
            # 如果右指针 r 已到边界，则 arr[n - k:] 就是满足题意的结果
            if r == n:
                return arr[n - k:]
            
            if abs(arr[l - 1] - x) <= abs(arr[r] - x):
                # 如果 arr[l - 1] 更接近 x ，则左移左指针，将其纳入结果窗口中
                l -= 1
            else:
                # 此时 arr[r] 更接近 x ，则右移右指针，将其纳入结果窗口中
                r += 1

        # 此时子数组 arr[l:r] 的长度为 k ，且是满足题意的结果
        return arr[l:r]
