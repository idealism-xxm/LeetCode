// 链接：https://leetcode.com/problems/words-within-two-edits-of-dictionary/
// 题意：给定两个字符串数组 queries 和 dictionary 。
//      每次修改可以选择 queries 中的一个字符串，将其中一个字母修改成任意其他字母。
//      求哪些字符串经过最多 2 次修改后，等于 dictionary 中的某个字符串？


// 数据限制：
//  1 <= queries.length, dictionary.length <= 100
//  n == queries[i].length == dictionary[j].length
//  1 <= n <= 100
//  queries[i] 和 dictionary[j] 仅由英文小写字母组成


// 输入： queries = ["word","note","ants","wood"], dictionary = ["wood","joke","moat"]
// 输出： ["word","note","wood"]
// 解释： "word": 将 'r' 修改为 'o' ，得到 "wood"
//       "note": 将 'n' 修改为 'j' ，将 't' 修改为 'k' ，得到 "joke"
//       "ants": 最多修改 2 次，无法得到 dictionary 中的字符串
//       "wood": 不做修改，得到 "wood"

// 输入： queries = ["yes"], dictionary = ["not"]
// 输出： []
// 解释： "yes": 最多修改 2 次，无法得到 dictionary 中的字符串


// 思路： 模拟
//
//      queries 和 dictionary 中所有单词的长度都是 n ，
//      所以我们可以通过计算汉明距离（位置相同但字符不同的数量）的方式，
//      得到 s 转换为 target 所需的最少修改次数。
//
//      对于 queries 中的每个字符串 s ，我们遍历 dictionary 中的每个字符串 target ，
//      计算两者的汉明距离。
//
//      如果存在一个 target 与 s 的汉明距离不超过 2 ，则 s 满足题意，放入结果列表中。
//
//
//      设 queries 的长度为 l ， dictionary 的长度为 m 。
//
//      时间复杂度：O(lmn)
//          1. 需要枚举 queries 中全部 O(l) 个字符串，
//              然后每次都需要 dictionary 中全部 O(m) 个字符串，
//              最后每次都需要同时枚举两个字符串全部 O(n) 个字符
//      空间复杂度：O(l)
//          1. 需要维护 queries 中所有满足题意的字符串，最差情况下有 O(l) 个


impl Solution {
    pub fn two_edit_words(queries: Vec<String>, dictionary: Vec<String>) -> Vec<String> {
        queries.into_iter()
            // 如果 s 与 dictionary 中某个字符串的汉明距离不超过 2 ，则满足题意
            .filter(|s| dictionary.iter().any(|target| Self::hamming_distance(s, target) <= 2))
            // 收集成数组返回
            .collect()
    }

    // 计算字符串 a 和 b 的汉明距离
    fn hamming_distance(a: &String, b: &String) -> usize {
        a.chars()
            .zip(b.chars())
            // 如果相同位置的字符不同，则汉明距离 +1
            .filter(|(a_ch, b_ch)| a_ch != b_ch)
            .count()
    }
}
