// 链接：https://leetcode.com/problems/word-ladder/
// 题意：给定开始单词、结束单词和一个单词列表（所有单词长度一样），
//		每次个以改变一个字母变成单词列表内的一个单词，求从开始单词变成结束单词最短变换序列的长度？

// 输入：
// beginWord = "hit",
// endWord = "cog",
// wordList = ["hot","dot","dog","lot","log","cog"]
// 输出： 5
// 解释： "hit" -> "hot" -> "dot" -> "dog" -> "cog"

// 输入：
// beginWord = "hit"
// endWord = "cog"
// wordList = ["hot","dot","dog","lot","log"]
// 输出： 0
// 解释：结束单词 "cog" 不在单词列表中，因此不存在可能的转换序列

// 思路： BFS
//		0126 的简化版，直接使用 BFS 即可
//		刚开始第一反应就是在 O(n ^ 2) 将可相互转换的单词标记，视为无向边，变长均为 1 ，
//		【题解用 map 和字符串变换，可以将相互可转换的单词在 O(n) 内找出来，感觉非常巧妙
//		可以将时间复杂度优化为 O(m * n) , m 是单词长度】
//		这样就转换成了边长都相同的最短路，直接 BFS 并标记转换到每个单词最短路径长度
//		后来发现其实不需要预处理无向边，直接在 BFS 的时候判断即可，
//		因为每个单词只会入队一次，入队前和一个单词判断一次，出队后和列表中的单词判断一次
//
//		由于本题只用找到最短路径的长度即可，所以也可以使用 双向 BFS
//
//		时间复杂度： O(m * n)
//		空间复杂度： O(m * n)

func ladderLength(beginWord string, endWord string, wordList []string) int {
	// 1. 预处理
	// 将开始单词放入，方便后续操作
	wordList = append(wordList, beginWord)
	// 获取开始单词和结束单词的下标
	length := len(wordList)
	beginIndex, endIndex := length - 1, -1
	// 将单词 word 变换成 word[:j] + "*" + word[j + 1:] ，然后映射到的均可以互相转换
	m := len(beginWord)
	connectedIndex := make(map[string][]int)
	for i, word := range wordList {
		// 找到结束单词的下标
		if endWord == word {
			endIndex = i
		}
		// 枚举可转换字母的位置
		for j := 0; j < m; j++ {
			commonWord := word[:j] + "*" + word[j + 1:]
			connectedIndex[commonWord] = append(connectedIndex[commonWord], i)
		}
	}
	// 如果结束单词不在单词列表中，则直接返回 0
	if endIndex == -1 {
		return 0
	}
	// 记录每个下标对应的单词是第几个（0 表示还未遍历过）
	count := make([]int, length)

	// 2. BFS 找出可能路径
	// BFS 所用队列
	queue := []int{beginIndex}
	// 开始单词是第一个
	count[beginIndex] = 1
	for ; len(queue) != 0; {
		// 队首下标出队
		curIndex := queue[0]
		queue = queue[1:]
		curWord := wordList[curIndex]
		// 找下一个入队单词，枚举可转换字母的位置
		for j := 0; j < m; j++ {
			commonWord := curWord[:j] + "*" + curWord[j + 1:]
			for _, index := range connectedIndex[commonWord] {
				// 如果 index 对应的单词还未遍历过，则入队
				if count[index] == 0 {
					queue = append(queue, index)
					count[index] = count[curIndex] + 1
				}
			}
		}
	}

	// 3. 返回转换成结束单词的所需长度
	return count[endIndex]
}
