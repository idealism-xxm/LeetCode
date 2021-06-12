# 链接：https://leetcode.com/problems/minimum-space-wasted-from-packaging/
# 题意：有 n 个包裹需要放到盒子中，每个盒子装一个包裹。
#       现有 m 个供应商分别提供不同尺寸的盒子（每个尺寸无限量），
#       一个包裹只能被放入尺寸大于等于它的盒子。
#
#       packages[i] 表示第 i 个包裹的大小，
#       boxes[j] 表示第 j 个供应商提供的盒子的尺寸列表。
#
#       现在需要选择一个供应商，使得放下所有包裹后浪费的空间最小，
#       返回这个最小值（结果模 1e9 + 7）。
#       浪费的空间 = 盒子尺寸 - 包裹大小

# 数据限制：
#   n == packages.length
#   m == boxes.length
#   1 <= n <= 10 ^ 5
#   1 <= m <= 10 ^ 5
#   1 <= packages[i] <= 10 ^ 5
#   1 <= boxes[j].length <= 10 ^ 5
#   1 <= boxes[j][k] <= 10 ^ 5
#   sum(boxes[j].length) <= 10 ^ 5
#   The elements in boxes[j] are distinct.

# 输入： packages = [2,3,5], boxes = [[4,8],[2,8]]
# 输出： 6
# 解释： 选择第一个供应商，
#       浪费空间 = (4-2) + (4-3) + (8-5) = 6

# 输入： packages = [2,3,5], boxes = [[1,4],[2,3],[3,4]]
# 输出： -1

# 输入： packages = [3,5,8,10,11,12], boxes = [[12],[11,9],[10,5,14]]
# 输出： 9
# 解释： 选择第三个供应商，
#       浪费空间 = (5-3) + (5-5) + (10-8) + (10-10) + (14-11) + (14-12) = 9

# 思路： 二分
#
#       已经知道了需要对包裹和盒子排序，没有想到可以将包裹的大小和盒子的尺寸分开算，
#       所以仍然是纯暴力遍历，没有利用统计的思想
#
#       首先计算好所有包裹大小的和 packages_sum ，
#       然后遍历每个供应商，然后从小到大遍历盒子，
#       找到刚好能使用该盒子的包裹数量，
#       将其算入盒子尺寸和中： boxes_sum += cnt * box_size
#       最后放入所有包裹后，
#       boxes_sum - packages_sum 就是使用当前供应商时浪费的空间和
#
#       找到所有这样的值的最小值即可
#
#
#       时间复杂度： O(nlogn + mlogm + mlogn)
#       空间复杂度： O(1)


class Solution:
    def minWastedSpace(self, packages: List[int], boxes: List[List[int]]) -> int:
        packages.sort()
        packages_sum = sum(packages)
        ans = 10 ** 11
        for cur in boxes:
            cur.sort()
            # 如果最大的盒子都放不下最大的包裹，则直接处理下一个
            if cur[-1] < packages[-1]:
                continue

            # 记录当前使用的盒子尺寸的和
            boxes_sum = 0
            # 记录当前还未放入盒子的第一个包裹的下标
            i = 0
            for box_size in cur:
                # 找到比 box_size 大的第一个包裹的位置
                # packages[i:j] 都会使用 box_size 这个盒子
                j = bisect.bisect(packages, box_size, i)
                boxes_sum += box_size * (j - i)
                i = j

            # 更新最小浪费空间
            ans = min(ans, boxes_sum - packages_sum)

        return ans % (10 ** 9 + 7) if ans != 10 ** 11 else -1
