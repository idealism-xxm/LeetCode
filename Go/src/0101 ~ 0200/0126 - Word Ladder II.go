// 链接：https://leetcode.com/problems/word-ladder-ii/
// 题意：给定开始单词、结束单词和一个单词列表（所有单词长度一样），
//		每次个以改变一个字母变成单词列表内的一个单词，求所有从开始单词变成结束单词最短变换序列？

// 输入：
// beginWord = "hit",
// endWord = "cog",
// wordList = ["hot","dot","dog","lot","log","cog"]
// 输出：
// [
//   ["hit","hot","dot","dog","cog"],
//   ["hit","hot","lot","log","cog"]
// ]

// 输入：
// beginWord = "hit"
// endWord = "cog"
// wordList = ["hot","dot","dog","lot","log"]
// 输出： []
// 解释：结束单词 "cog" 不在单词列表中，因此不存在可能的转换序列

// 思路： BFS + DFS
//		刚开始第一反应就是在 O(n ^ 2) 将可相互转换的单词标记，视为无向边，变长均为 1 ，
//		【0127 的题解用 map 和字符串变换，可以将相互可转换的单词在 O(n) 内找出来，
//		可以将时间复杂度优化为 O(m * n) , m 是单词长度】
//		这样就转换成了边长都相同的最短路，直接 BFS 标记每个点可由哪些点抵达，最后 DFS 收集所有转换序列即可
//		后来发现其实不需要预处理无向边，直接在 BFS 的时候判断即可，
//		因为每个单词只会入队一次，入队前和一个单词判断一次，出队后和列表中的单词判断一次
//
//		时间复杂度： O(m * n ^ 2)
//		空间复杂度： O(m * n)

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	// 1. 预处理
	// 将开始单词放入，方便后续操作
	wordList = append(wordList, beginWord)
	// 将单词映射成下标，方便后续操作
	wordToIndex := make(map[string]int)
	for i, word := range wordList {
		wordToIndex[word] = i
	}
	// 如果结束单词不在单词列表中，则直接返回 nil
	if _, exists := wordToIndex[endWord]; !exists {
		return nil
	}
	// 获取开始单词和结束单词的下标
	beginIndex, endIndex := wordToIndex[beginWord], wordToIndex[endWord]
	// 记录每个下标对应的单词是第几个（0 表示还未遍历过）
	length := len(wordList)
	count := make([]int, length)
	// 记录每个下标对应的单词可由哪些下标对应的单词转换而来
	pre := make([][]int, length)

	// 2. BFS 找出可能路径
	// BFS 所用队列
	queue := []int{beginIndex}
	// 开始单词是第一个
	count[beginIndex] = 1
	for ; len(queue) != 0; {
		// 队首下标出队
		curIndex := queue[0]
		queue = queue[1:]
		// 找下一个入队单词
		for i := 0; i < length; i++ {
			// 如果这两个单词有一条无向边，则可以继续处理
			if isConnected(wordList[curIndex], wordList[i]) {
				// 如果 i 对应的单词还未遍历过，则入队
				if count[i] == 0 {
					queue = append(queue, i)
					count[i] = count[curIndex] + 1
				}
				// 如果当前转换在最短路径上，则记录前驱
				if count[i] == count[curIndex] + 1 {
					pre[i] = append(pre[i], curIndex)
				}
			}
		}
	}

	// 3. 如果结束单词无法抵达，则直接返回
	if count[endIndex] == 0 {
		return nil
	}

	// 4. DFS 从结束单词开始收集最短路径的答案
	// 提前分配好足够的空间
	list := make([]string, count[endIndex])
	// 第一个单词必定是开始单词
	list[0] = beginWord
	return dfs(wordList, pre, endIndex, list, count[endIndex])
}

func dfs(wordList []string, pre [][]int, curIndex int, list []string, remainWordCount int) [][]string {
	// 如果只剩最后一个单词了，则必定是开始单词，直接放入答案列表
	if remainWordCount == 1 {
		return [][]string {append(list[:0:0], list...)}
	}

	// 倒序将路径上的单词放入
	list[remainWordCount - 1] = wordList[curIndex]
	// 找到所有可能前驱，并递归收集对应的答案
	var result [][]string
	for _, preIndex := range pre[curIndex] {
		result = append(result, dfs(wordList, pre, preIndex, list, remainWordCount - 1)...)
	}

	// 返回收集到的答案
	return result
}

func isConnected(start string, end string) bool {
	count := 0
	for i, ch := range start {
		if ch != int32(end[i]) {
			count++
		}
	}
	return count == 1
}
