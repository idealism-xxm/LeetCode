# 链接：https://leetcode.com/problems/maximum-units-on-a-truck/
# 题意：给定一个二维数组 boxTypes ，
#      其中 boxTypes[i] = [numberOfBoxes_i, numberOfUnitsPerBox_i] ，
#      numberOfBoxes_i 表示第 i 种箱子的数量， 
#      numberOfUnitsPerBox_i 表示第 i 种箱子每个箱子的单元数。
#
#      从中选取 truckSize 个箱子，求这些箱子单元数之和的最大值？


# 数据限制：
#  1 <= boxTypes.length <= 1000
#  1 <= numberOfBoxes_i, numberOfUnitsPerBox_i <= 1000
#  1 <= truckSize <= 10 ^ 6


# 输入： boxTypes = [[1,3],[2,2],[3,1]], truckSize = 4
# 输出： 8
# 解释： 第一种箱子选取 1 个，单元数之和为 1 * 3 = 3
#       第二种箱子选取 2 个，单元数之和为 2 * 2 = 4
#       第三种箱子选取 1 个，单元数之和为 1 * 1 = 1
#
#       全部所选箱子的单元数之和为 3 + 4 + 1 = 8

# 输入： boxTypes = [[5,10],[2,5],[4,7],[3,9]], truckSize = 10
# 输出： 91
# 解释： 第一种箱子选取 5 个，单元数之和为 5 * 10 = 50
#       第二种箱子选取 0 个，单元数之和为 0 * 5 = 0
#       第三种箱子选取 2 个，单元数之和为 2 * 7 = 14
#       第四种箱子选取 3 个，单元数之和为 3 * 9 = 27
#
#       全部所选箱子的单元数之和为 50 + 0 + 14 + 27 = 91


# 思路： 贪心 + 排序
#
#      要让最终所选箱子的单元数之和最大，我们可以贪心地优先选取单元数最大的箱子。
#
#      只需要对 boxTypes 按照 boxTypes[i][1] 倒序排序，
#      然后按顺序选取前 truckSize 个箱子即可。
#
#
#      时间复杂度：O(nlogn)
#          1. 需要对 boxTypes 进行排序，时间复杂度为 O(nlogn)
#          2. 需要遍历所有能选取的箱子，最差情况下需要选取全部 O(n) 种箱子
#      空间复杂度：O(1)
#          1. 只需要使用常数个额外变量即可


class Solution:
    def maximumUnits(self, boxTypes: List[List[int]], truckSize: int) -> int:
        # 按每个箱子的单元数倒序排序，贪心地优先选取单元数最大的箱子
        boxTypes.sort(key=lambda box_type: box_type[1], reverse=True)

        ans: int = 0
        for box_type in boxTypes:
            if truckSize <= box_type[0]:
                # 如果当前类型的箱子能用完选取次数，则选取 truckSize 个箱子后，跳出循环
                ans += truckSize * box_type[1]
                break
            else:
                # 如果当前类型的箱子不能用完选取次数，则选取全部箱子
                ans += box_type[0] * box_type[1]
                truckSize -= box_type[0]

        return ans
