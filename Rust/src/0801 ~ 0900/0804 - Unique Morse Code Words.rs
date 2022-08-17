// 链接：https://leetcode.com/problems/unique-morse-code-words/
// 题意：给定一个单词数组 words ，求这些单词对应的所有莫斯电码中，不同摩斯电码的数量？
//
//      每个英文字母对应的摩斯电码如下：
//      [".-","-...","-.-.","-..",".","..-.","--.","....","..",
//          ".---","-.-",".-..","--","-.","---",".--.","--.-",".-.",
//          "...","-","..-","...-",".--","-..-","-.--","--.."]


// 数据限制：
//  1 <= words.length <= 100
//  1 <= words[i].length <= 12
//  words[i] 仅由英文小写字母组成


// 输入： words = ["gin","zen","gig","msg"]
// 输出： 2
// 解释： 有两种不同的摩斯电码： "--...-." 和 "--...-."
//       "gin" -> "--...-."
//       "zen" -> "--...-."
//       "gig" -> "--...--."
//       "msg" -> "--...--."

// 输入： words = ["a"]
// 输出： 1
// 解释： 有一种不同的摩斯电码： ".-" 
//       "a" -> ".-"


// 思路： Set
//
//      将每个单词都按照转换表转换成莫斯电码，然后放入一个集合中去重，
//      那么集合的大小就是不同摩斯电码的数量。
//
//
//      设单词长度最长为 L 。
//
//      时间复杂度：O(nL)
//          1. 需要遍历全部 O(n) 个单词，每个单词都需要遍历全部 O(L) 个字符。
//      空间复杂度：O(nL)
//          1. 需要维护全部不同的摩斯电码，最差情况下有 O(n) 个不同的摩斯电码；
//              每个莫斯电码的长度都为 O(L)


use std::collections::HashSet;


// 莫斯电码转换表
const MORSE_CODE: [&str; 26] = [
    ".-","-...","-.-.","-..",".","..-.","--.","....","..",
    ".---","-.-",".-..","--","-.","---",".--.","--.-",".-.",
    "...","-","..-","...-",".--","-..-","-.--","--..",
];


impl Solution {
    pub fn unique_morse_representations(words: Vec<String>) -> i32 {
        words
            .iter()
            // 将每个单词转换成莫斯电码
            .map(|word| word.chars().map(|ch| MORSE_CODE[ch as usize - 'a' as usize]).collect::<String>())
            // 转换成集合
            .collect::<HashSet<_>>()
            // 返回集合大小
            .len() as i32
    }
}
