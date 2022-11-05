# 链接：https://leetcode.com/problems/word-ladder-ii/
# 题意：给定开始单词、结束单词和一个单词列表（所有单词长度一样），
#      每次可以改变一个字母变成单词列表内的一个单词，
#      求从开始单词变成结束单词所有最短转换序列？


# 数据限制：
#  1 <= beginWord.length <= 5
#  endWord.length == beginWord.length
#  1 <= wordList.length <= 500
#  wordList[i].length == beginWord.length
#  beginWord, endWord 和 wordList[i] 均由小写英文字母组成
#  beginWord != endWord
#  wordList 中所有的单词都各不相同
#  所有最短转换序列的个数不超过 10 ^ 5


# 输入：beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log","cog"]
# 输出：[["hit","hot","dot","dog","cog"],["hit","hot","lot","log","cog"]]
# 解释：有两个最短转换序列：
#          "hit" -> "hot" -> "dot" -> "dog" -> "cog"
#          "hit" -> "hot" -> "lot" -> "log" -> "cog"

# 输入：beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log"]
# 输出：[]
# 解释：结束单词 "cog" 不在单词列表中，所以无法进行转换。


# 思路：BFS + Map + 递归/回溯/DFS
#
#      本题是 LeetCode 127 加强版，需要找到所有最短转换序列。
#      但前半部分思路基本一致，需要先用 BFS 找到最短序列的长度，
#      然后再用 DFS 收集即可。
#
#
#      本题是单源最短路，而且边长都是 1 ，所以可以直接使用 BFS 搜索即可。
#
#      我们可以维护一个邻接表 adj ，遍历每个单词 wordList[i] 的每个位置 j ，
#      将第 j 个字符替换为 '.' 形成通配字符串 source 。
#
#      然后将 i 放入 adj[source] 中，
#      那么 adj[source] 中的所有下标对应的单词都可以相互转换。
#
#      同时维护一个距离数组 distance ， 
#      distance[i] 表示转换到单词 wordList[i] 时的转换序列长度，
#      0 表示无法转换至单词 wordList[i] 。
#
#      再维护一个前驱二维数组 pres ，
#      pres[i] 表示最短转换序列中， wordList[i] 所有可能的前一个单词下标的列表。
#
#      BFS 每个单词出队时，遍历可替换的字符生成通配字符串 source ，
#      遍历 adj[source] 中所有能转换的单词下标 next 。
#          1. distance[next] == 0: 则表明到第一次遍历到 next ，
#              更新 distance[next] = distance[cur] + 1 ，
#              并将 cur 放入 pres[next] 中，再将 next 放入队列即可
#          2. distance[next] == distance[cur] + 1: 则表明非第一次遍历到 next ，
#              但 cur 也是 next 最短序列的前驱，仅将 cur 放入 pres[next] 中
#
#      最后如果 BFS 结束后， distance[end_index] 为 0 ，
#      表示无法转换到结束单词，直接返回空列表。
#
#      否则使用 DFS 根据 pres 收集所有可能的最短转换序列即可。
#
#
#      设 n 为单词列表长度， L 为单词长度， C 为最短转换序列个数， M 为最短转换序列长度。
#
#      时间复杂度：O(n * L ^ 2 + CM)
#          1. 需要计算全部 O(n) 个单词所属的邻接表，
#              每次计算时都需要遍历全部 O(L) 个可替换的字符，
#              每次遍历时都需要生成对应的长度为 O(L) 的通配符字符串。
#              总时间复杂度为 O(n * L ^ 2)
#          2. 全部 O(n) 个单词都会入队列一次
#          3. 全部 O(n) 个单词都会出队列一次，
#              每次出队列都需要遍历全部 O(L) 个可替换的字符，
#              每次遍历时都需要生成对应的长度为 O(L) 的通配符字符串。
#              总时间复杂度为 O(n * L ^ 2)
#          4. 需要收集全部 O(C) 个最短转换序列，
#              每次都需要克隆每个最短转换序列全部 O(M) 个字符串。
#              总时间复杂度为 O(CM)
#      空间复杂度：O(nL + n ^ 2 + CM)
#          1. 需要维护邻接表中全部 O(nL) 个单词下标
#          2. 需要维护 distance 全部 O(n) 个状态
#          3. 需要维护队列 q 中全部 O(n) 个单词
#          4. 需要维护 pres 中全部 O(n ^ 2) 个前驱
#          5. 需要维护 seq 中全部 O(M) 个字符串
#          6. 需要维护 ans 中全部 O(C) 个最短转换序列，
#              每个最短转换序列有 O(M) 个字符串。
#              总空间复杂度为 O(CM)


class Solution:
    def findLadders(self, beginWord: str, endWord: str, wordList: List[str]) -> List[List[str]]:
        # 先把开始单词放入单词列表中，方便后续使用下标处理
        wordList.append(beginWord)
        start_index: int = len(wordList) - 1

        # 找到结束单词在单词列表中的下标
        end_index: int = -1
        for i, word in enumerate(wordList):
            if word == endWord:
                end_index = i
                break
        # 如果结束单词不在单词列表中，则无法转换，直接返回空列表
        if end_index == -1:
            return []

        # 构建邻接表
        adj: Dict[str, List[int]] = defaultdict(list)
        for i, word in enumerate(wordList):
            # 枚举 word 替换的字符
            for j in range(len(word)):
                # 将第 j 个字符替换为通配符 '.'
                source: str = word[:j] + '.' + word[j+1:]
                # 所有能变为 source 的单词都能相互转换
                adj[source].append(i)

        # 队列 q 存储 BFS 下一次遍历的单词下标
        q: Deque = deque()
        # 初始只有开始单词的下标在其中
        q.append(start_index)
        # distance[i] 表示从 start_index 转换到 i 时的序列长度，
        # 初始化为 0 ，表示无法转换
        distance: List[int] = [0] * len(wordList)
        # 开始单词本身的无需任何转换就能得到
        distance[start_index] = 1
        # pres 表示最短转换序列中， word_list[i] 所有可能的前一个单词下标的列表
        pres: List[List[int]] = [[] for _ in range(len(wordList))]

        # 不断从 q 中获取下一个单词下标，直至 q 为空
        while q:
            cur: int = q.popleft()

            # 枚举 cur 替换的字符
            for j in range(len(wordList[cur])):
                # 将第 j 个字符替换为通配符 '.'
                source: str = wordList[cur][:j] + '.' + wordList[cur][j+1:]
                # 遍历邻接表
                for next in adj[source]:
                    if distance[next] == 0:
                        # 如果 next 还未遍历过，则更新 distance[next] ，
                        # 将 cur 放入 pres[next] 中，并将 next 放入队列 q 中
                        distance[next] = distance[cur] + 1
                        pres[next].append(cur)
                        q.append(next)
                    elif distance[next] == distance[cur] + 1:
                        # 如果 next 已遍历过，但 cur 也是最短序列中 next 的前驱，
                        # 则将 cur 放入 pres[next] 中
                        pres[next].append(cur)


        # distance[end_index] 为 0 ，表示无法转换到结束单词，直接返回空列表
        if distance[end_index] == 0:
            return []

        # seq 用于收集每一个最短转换序列，长度为 distance[end_index]
        seq: List[str] = [''] * distance[end_index]
        # 最短转换序列的第一个单词必定是开始单词
        seq[0] = wordList[start_index]
        # ans 用于收集全部可能的最短转换序列
        ans: List[List[str]] = []
        # 使用 dfs 根据 pres 进行收集
        Solution.dfs(wordList, pres, end_index, distance[end_index], seq, ans)
        return ans


    # 用于收集所有可能的最短序列列表
    #  word_list: 单词列表，用于定位具体单词
    #  pres: 每个单词的前驱列表
    #  cur: 最短序列上当前单词的下标
    #  remain: 最短序列中剩余需要遍历的单词数
    #  seq: 当前最短序列
    #  ans: 所有最短序列的列表
    @staticmethod
    def dfs(word_list: List[str], pres: List[List[int]], cur: int, remain: int, seq: List[str], ans: List[List[str]]):
        # 如果剩余的单词只有 1 个，那么必定是开始单词，直接将 seq 收集进 ans 即可
        if remain == 1:
            ans.append(seq[:])
            return

        # 当前最短序列中第 remain - 1 个单词为 word_list[cur]
        seq[remain - 1] = word_list[cur]
        # 遍历 cur 前驱，递归收集
        for pre in pres[cur]:
            Solution.dfs(word_list, pres, pre, remain - 1, seq, ans)
