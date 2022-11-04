// 链接：https://leetcode.com/problems/word-ladder/
// 题意：给定开始单词、结束单词和一个单词列表（所有单词长度一样），
//      每次可以改变一个字母变成单词列表内的一个单词，
//      求从开始单词变成结束单词最短转换序列的长度？


// 数据限制：
//  1 <= beginWord.length <= 10
//  endWord.length == beginWord.length
//  1 <= wordList.length <= 5000
//  wordList[i].length == beginWord.length
//  beginWord, endWord 和 wordList[i] 均由小写英文字母组成
//  beginWord != endWord
//  wordList 中的所有字符串互不相同


// 输入：beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log","cog"]
// 输出：5
// 解释：一个最短转换序列是 "hit" -> "hot" -> "dot" -> "dog" -> "cog", 返回它的长度 5 。

// 输入：beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log"]
// 输出：0
// 解释：结束单词 "cog" 不在单词列表中，所以无法进行转换。


// 思路：BFS + Map
//
//      本题是 LeetCode 433 加强版，字符集大小从 4 变为 26 ，数据量也变大了。
//      但思路和代码基本一致，只需要修改一下返回值就能直接复用。
//
//
//      本题是单源最短路，而且边长都是 1 ，所以可以直接使用 BFS 搜索即可。
//
//      我们可以维护一个邻接表 adj ，遍历每个单词 wordList[i] 的每个位置 j ，
//      将第 j 个字符替换为 '.' 形成通配字符串 source 。
//
//      然后将 i 放入 adj[source] 中，
//      那么 adj[source] 中的所有下标对应的单词都可以相互转换。
//
//      同时维护一个距离数组 distance ， 
//      distance[i] 表示转换到单词 wordList[i] 时的序列长度，
//      0 表示无法转换至单词 wordList[i] 。
//
//      BFS 每个单词出队时，遍历可替换的字符生成通配字符串 source ，
//      遍历 adj[source] 中所有能转换的单词下标 next ，
//      更新转化序列长度 distance[next] ，并将 next 放入队列中。
//
//      每次出队时，如果当前单词下标 cur 就是结束单词的下标 end_index ，
//      则直接返回 distance[cur] 。
//
//      最后如果 BFS 结束还没有返回，则直接返回 0 ，表示无法转换到结束单词。
//
//      由于本题只用找到最短路径的长度即可，所以也可以使用 双向 BFS 。
//
//
//      设 n 为单词列表长度， L 为单词长度。
//
//      时间复杂度：O(n * L ^ 2)
//          1. 需要计算全部 O(n) 个单词所属的邻接表，
//              每次计算时都需要遍历全部 O(L) 个可替换的字符，
//              每次遍历时都需要生成对应的长度为 O(L) 的通配符字符串。
//              总时间复杂度为 O(n * L ^ 2)
//          2. 全部 O(n) 个单词都会入队列一次
//          3. 全部 O(n) 个单词都会出队列一次，
//              每次出队列都需要遍历全部 O(L) 个可替换的字符，
//              每次遍历时都需要生成对应的长度为 O(L) 的通配符字符串。
//              总时间复杂度为 O(n * L ^ 2)
//      空间复杂度：O(nL) 
//          1. 需要维护邻接表中全部 O(nL) 个单词下标
//          2. 需要维护 distance 全部 O(n) 个状态
//          3. 需要维护队列 q 中全部 O(n) 个单词


func ladderLength(beginWord string, endWord string, wordList []string) int {
    // 先把开始单词放入单词列表中，方便后续使用下标处理
    wordList = append(wordList, beginWord)
    startIndex := len(wordList) - 1

    // 找到结束单词在单词列表中的下标
    endIndex := -1
    for i, word := range wordList {
        if word == endWord {
            endIndex = i
            break
        }
    }
    // 如果结束单词不在单词列表中，则无法转换，直接返回 0
    if endIndex == -1 {
        return 0
    }

    // 构建邻接表
    adj := make(map[string][]int)
    for i, word := range wordList {
        // 枚举 word 替换的字符
        for j := range word {
            // 将第 j 个字符替换为通配符 '.'
            source := word[:j] + "." + word[j+1:]
            // 所有能变为 source 的单词都能相互转换
            adj[source] = append(adj[source], i)
        }
    }

    // 队列 q 存储 BFS 下一次遍历的单词下标
    var q []int
    // 初始只有开始单词的下标在其中
    q = append(q, startIndex)
    // distance[i] 表示从 startIndex 转换到 i 时的序列长度，
    // 初始化为 0 ，表示无法转换
    distance := make([]int, len(wordList))
    // 开始单词本身的无需任何转换就能得到
    distance[startIndex] = 1

    // 不断从 q 中获取下一个单词下标，直至 q 为空
    for len(q) > 0 {
        cur := q[0]
        q = q[1:]
        // 如果当前单词下标就是结束单词的下标，则直接返回
        if cur == endIndex {
            return distance[endIndex]
        }

        // 枚举 cur 替换的字符
        for j := range wordList[cur] {
            // 将第 j 个字符替换为通配符 '.'
            source := wordList[cur][:j] + "." + wordList[cur][j+1:]
            // 遍历邻接表
            for _, next := range adj[source] {
                // 如果 next 还未遍历过，则更新 distance[next] ，
                // 并将 next 放入队列 q 中
                if distance[next] == 0 {
                    distance[next] = distance[cur] + 1
                    q = append(q, next)
                }
            }
        }
    }

    // 最后遍历完还没有找结束单词，则直接返回 0
    return 0
}
