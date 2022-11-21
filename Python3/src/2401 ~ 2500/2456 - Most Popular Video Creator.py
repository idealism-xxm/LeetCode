# 链接：https://leetcode.com/problems/most-popular-video-creator/
# 题意：给定两个字符串数组 creators 和 ids ，以及一个整型数组 views 。
#      第 i 部电影的创作者是 creator[i] ， id 是 ids[i] ，观看数是 views[i] 。
#
#      一个创作者的流行度是其所有创作的电影的观看数之和，
#      找到全部流行度最高的作者，及其观看数最高的电影。
#
#      如果一个创作者观看数最高的电影有多个，则返回 id 字典序最小的那个。


# 数据限制：
#  n == creators.length == ids.length == views.length
#  1 <= n <= 10 ^ 5
#  1 <= creators[i].length, ids[i].length <= 5
#  creators[i] 和 ids[i] 仅由英文小写字母组成
#  0 <= views[i] <= 10 ^ 5


# 输入： creators = ["alice","bob","alice","chris"], ids = ["one","two","three","four"], views = [5,10,5,4]
# 输出： [["alice","one"],["bob","two"]]
# 解释： alice 的流行度为 5 + 5 = 10 ，
#       bob 的流行度为 10 ，
#       chris 的流行度 4 。
#
#       alice 和 bob are 是流行度最高的创作者。
#       bob 观看数最高的电影是 "two" ；
#       alice 观看数最高的电影是 "one" 和 "three" ， "one" 的 id 字典序最小。

# 输入： nums = [1,2,4,7,10]
# 输出： 0
# 解释： 没有能被 3 整除的偶数，平均数为 0


# 思路： Map
#
#      我们用 creator_to_info 维护一个创作者的相关信息，
#      其中 creator_to_info[creator] = (index, view) ：
#          1. creator: 创作者
#          2. index: creator 的电影 ids[index] 是其观看数最高且字典序最小的
#          3. view: 创作者所有电影的观看数之和
#
#      然后遍历第 i 部电影，将观看数 views[i] 计入对应的创作者中，
#      并更新 index ，保证其一直是观看数最高且字典序最小的。
#
#      同时用 max_view 维护创作者所有电影的观看数之和的最大值。
#
#      最后将 creator_to_info 中 view 等于 max_view 的元素收集成结果返回即可。
#
#
#      时间复杂度：O(n)
#          1. 需要遍历 creators, ids, views 中全部 O(n) 个元素
#          2. 需要遍历 creator_to_info 中全部不同的元素，最差情况有 O(n) 个
#      空间复杂度：O(n)
#          1. 需要维护 creator_to_info 中全部不同的元素，最差情况有 O(n) 个


class Solution:
    def mostPopularCreator(self, creators: List[str], ids: List[str], views: List[int]) -> List[List[str]]:
        # creator_to_info[creator] = (index, view)
        #  creator: 创作者
        #  index: creator 的电影 ids[index] 是其观看数最高且字典序最小的
        #  view: 创作者所有电影的观看数之和
        creator_to_info: Dict[str, Tuple[int, int]] = defaultdict(lambda: (i, 0))
        # 创作者所有电影的观看数之和的最大值
        max_view: int = 0
        for i, creator in enumerate(creators):
            # 获取创作者 creator 的相关信息
            index, view = creator_to_info[creator]
            # 将第 i 部电影的观看数计入 view
            view += views[i]
            # 如果第 i 部电影的观看数更多，或者观看次数相等但字典序更小，
            # 则贪心地选择第 i 部电影
            if views[i] > views[index] or (views[i] == views[index] and ids[i] < ids[index]):
                index = i
            creator_to_info[creator] = (index, view)

            # 更新观看数之和的最大值
            max_view = max(max_view, view)
        
        # 收集成列表后返回
        return [
            # 转换成需要的形式
            [creator, ids[index]]
            for creator, (index, view) in creator_to_info.items()
            # 仅过滤出观看数之和最大的
            if view == max_view
        ]
