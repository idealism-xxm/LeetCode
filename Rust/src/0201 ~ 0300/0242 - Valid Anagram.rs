// 链接：https://leetcode.com/problems/valid-anagram/
// 题意：给定两个字符串 s 和 t ，判断它们是不是一对变位词？

// 输入： s = "anagram", t = "nagaram"
// 输出： true

// 输入： s = "rat", t = "car"
// 输出： false

// 思路： map
//
//      我们用一个 map 统计每个字符出现的次数（兼容 unicode 字符），
//      对于 s 中的每个字符，我们给对应的次数 + 1
//      对于 t 中的每个字符，我们给对应的次数 - 1
//      最后判断 map 中每个字符的次数是不是 0 即可
//
//      时间复杂度： O(n)
//      空间复杂度： O(k) （ k 为字符表的大小）

use std::collections::HashMap;

impl Solution {
    pub fn is_anagram(s: String, t: String) -> bool {
        let mut counts: HashMap<char, i32> = HashMap::new();
        // 对于 s 中的每个字符，我们给对应的次数 + 1
        for ch in s.chars() {
            let count = counts.entry(ch).or_insert(0);
            *count += 1;
        }
        // 对于 t 中的每个字符，我们给对应的次数 - 1
        for ch in t.chars() {
            let count = counts.entry(ch).or_insert(0);
            *count -= 1;
        }
        // 判断 map 中每个字符的次数是不是 0
        for count in counts.values() {
            if *count != 0 {
                return false;
            }
        }

        // 全部次数都为 0 ，则是一对异位词
        true
    }
}
