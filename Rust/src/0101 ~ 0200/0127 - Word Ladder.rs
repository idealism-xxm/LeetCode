// 链接：https://leetcode.com/problems/word-ladder/
// 题意：给定开始单词、结束单词和一个单词列表（所有单词长度一样），
//		每次可以改变一个字母变成单词列表内的一个单词，
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


// 思路：BFS
//
//      本题是单源最短路，而且边长都是 1 ，所以可以直接使用 BFS 搜索即可。
//
//      我们维护一个距离数组 distance ， 
//      distance[i] 表示转换到单词 word_list[i] 时的转换序列长度，
//      0 表示无法转换至单词 word_list[i] 。
//
//      由于每个单词最多只会入/出队列一次，所以我们无需提前建立无向边。
//      在每个单词出队时，遍历单词列表 word_list ，找到所有可能转换的单词下标 next ，
//      更新转化序列长度 distance[next] ，并将 next 放入到队列中。
//
//      每次出队时，如果当前单词下标 cur 就是结束单词的下标 end_index ，
//      则直接返回 distance[cur] 。
//
//      最后如果 BFS 结束还没有返回，则直接返回 0 ，表示无法转换到结束单词。
//
//      由于本题只用找到最短路径的长度即可，所以也可以使用 双向 BFS 。
//
//      【进阶】找到所有的最短路径，也就是 LeetCode 126 这题。
//
//      设 n 为单词列表长度， L 为单词长度
//
//      时间复杂度：O(L * n ^ 2)
//          每个单词都会入/出队列各一次，时间复杂度为 O(n)
//          每次出队列都会枚举单词列表并判断是否可以转换，时间复杂度为 O(n * L)。
//      空间复杂度：O(n + L) 
//          distance 的空间复杂度为 O(n) 
//          开始单词克隆的空间复杂度为 O(L)

use std::collections::VecDeque;

impl Solution {
    pub fn ladder_length(begin_word: String, end_word: String, mut word_list: Vec<String>) -> i32 {
        // 先把开始单词放入单词列表中，方便后续使用下标处理
        word_list.push(begin_word.clone());
        let start_index = word_list.len() - 1;
        // distance[i] 表示从 start_index 转换到 i 时的转换序列长度，
        // 初始化为 0 ，表示无法转换
        let mut distance = vec![0; word_list.len() + 1];
        // 开始单词本身的转换序列长度是 1
        distance[start_index] = 1;

        // 找到结束单词在单词列表中的下标
        let end_index = word_list
            // 转成迭代器
            .iter()
            // 判断单词是否是结束单词
            .position(|word| word == &end_word);
        // 如果结束单词不在单词列表中，则无法转换，直接返回 0
        if end_index.is_none() {
            return 0;
        }
        // 获取结束单词下标
        let end_index = end_index.unwrap();

        // 队列 q 存储 BFS 下一次遍历的单词下标
        let mut q = VecDeque::new();
        // 初识只有开始单词的下标在其中
        q.push_back(start_index);
        // 不断从 q 中获取，直至 q 为空
        while let Some(cur) = q.pop_front() {
            // 如果当前单词下标就是结束单词的下标，则直接返回
            if cur == end_index {
                return distance[end_index];
            }

            // 枚举下一个单词
            for (next, word) in word_list.iter().enumerate() {
                // 如果下一个单词还未遍历过，且与当前单词可以相互转换，
                // 则更新 distance[next] ，并将 next 放入队列 q 中
                if distance[next] == 0 && Self::is_connected(&word_list[cur], word) {
                    distance[next] = distance[cur] + 1;
                    q.push_back(next);
                }
            }
        }

        // 最后遍历完还没有找结束单词，则直接返回 0
        0
    }

    fn is_connected(start: &String, end: &String) -> bool {
        start
            // 转成字节数组
            .as_bytes()
            // 转成迭代器
            .iter()
            // 同时遍历 end
            .zip(end.as_bytes().iter())
            // 使用 fold 统计 start 和 end 位置相同字符不同的数量
            .fold(0, |cnt, (s, e)| {
                if s != e {
                    cnt + 1
                } else {
                    cnt
                }
            })
            // 如果位置相同字符不同的数量是 1 ，则可以相互转换
            == 1
    }
}
