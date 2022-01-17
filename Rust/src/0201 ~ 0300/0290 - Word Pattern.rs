// 链接：https://leetcode.com/problems/word-pattern/
// 题意：给定一个模式串 pattern 和一个字符串 s ，
//      判断 s 中的非空单词与 pattern 中的字母是否能形成双射？

// 数据限制：
//  1 <= pattern.length <= 300
//  pattern 只含有英文小写字母
//  1 <= s.length <= 3000
//  s 只含有英文小写字母和空格 ' '
//  s 的头尾均不含有空格
//  s 中所有的单词都通过一个空格分隔

// 输入：pattern = "abba", s = "dog cat cat dog"
// 输出：true
// 解释：'a' -> "dog"
//      'b' -> "cat"

// 输入：pattern = "abba", s = "dog cat cat fish"
// 输出：false
// 解释：'a' 无法同时映射成 "dog" 和 "fish"

// 输入：pattern = "aaaa", s = "dog cat cat dog"
// 输出：false
// 解释：'a' 无法同时映射成 "dog" 和 "cat"


// 思路： map
//
//      先将 s 按照空格分割成单词列表 words ，
//      然后判断 words.len() 与 pattern.len() 是否相等，
//          如果不相等，则无法形成双射，直接返回 false
//
//      
//      最后同时遍历 pattern 和 words ，
//      同时维护两个 HashMap ，分别为 ch_to_word 和 word_to_ch ，
//          ch_to_word 用于存储字母 ch 对应的单词 word ，
//          word_to_ch 用于存储字母 word 对应的单词 ch 
//
//      每次对于当前字母 ch 和当前单词 word 进行如下处理，
//          1. 如果 ch 在 ch_to_word 中，则 ch 还未映射过，
//              此时可以进行映射： ch_to_word[ch] = word
//          2. 如果 word 在 word_to_ch 中，则 word 还未映射过，
//              此时可以进行映射： word_to_ch[word] = ch
//          3. 获取并判断双射关系是否成立：
//              (1) word == ch_to_word[ch] && ch == word_to_ch[word]: 
//                  目前仍能形成双射，可以继续遍历处理即可
//              (2) word != ch_to_word[ch] || ch != word_to_ch[word]: 
//                  已无法形成双射，因为 ch 不可能同时映射成两个不同的单词，
//                  此时，直接返回 false
//
//      最后遍历完成后还未返回，则说明双射一直成立，返回 true
//
//      时间复杂度： O(n + m)
//      空间复杂度： O(n + m)

use std::collections::HashMap;

impl Solution {
    pub fn word_pattern(pattern: String, s: String) -> bool {
        // 将 s 按照空格分割成单词列表 words
        let words: Vec<&str> = s.split(' ').collect();
        // 如果单词数不等于 pattern 的字母数，则无法形成双射，直接返回 false
        if words.len() != pattern.len() {
            return false;
        }

        // 遍历 pattern 的字母
        pattern.chars()
            // 同时遍历下标对应的 word
            .zip(words.iter())
            // try_fold 会尝试遍历完迭代器，如果处理函数 f 返回的 accumulator 是 Some ，
            //      则会继续遍历，直至 f 返回 None
            // 1. 如果最后返回值是 Some ，则是 accumulator 的最终值
            // 2. 如果最后返回值是 None ，则说明没有遍历完迭代器， f 返回 None 提前终止了
            .try_fold((HashMap::new(), HashMap::new()), |(mut ch_to_word, mut word_to_ch), (ch, word)| {
                // 获取 ch 对应的 word ，没有则进行插入，即让 ch 映射到当前 word ，
                // 获取 word 对应的 ch ，没有则进行插入，即让 word 映射到当前 ch
                if word == *ch_to_word.entry(ch).or_insert(word) 
                    && ch == *word_to_ch.entry(word).or_insert(ch) {
                    // 如果当前 word 和 ch 实际对应的 word 相同，
                    //  且当前 ch 和 word 实际对应的 ch 相同，
                    //  则目前仍能够形成双射，返回 Some(ch_to_word)
                    Some((ch_to_word, word_to_ch))
                } else {
                    // 当前 word 与 ch 对应的 word 不同，则目前无法形成双射，
                    // 返回 None ，停止继续处理
                    None
                }
            })
            // 最后如果还存在 (ch_to_word, word_to_ch) ，则说明能形成双射
            .is_some()
    }
}
