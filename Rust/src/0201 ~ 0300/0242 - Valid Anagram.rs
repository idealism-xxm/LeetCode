// 链接：https://leetcode.com/problems/valid-anagram/
// 题意：给定两个字符串 s 和 t ，判断它们是不是一对字母异位词？


// 数据限制：
//  1 <= s.length, t.length <= 5 * 10 ^ 4
//  s 和 t 仅含有英文小写字母


// 输入： s = "anagram", t = "nagaram"
// 输出： true

// 输入： s = "rat", t = "car"
// 输出： false


// 思路： Map
//
//      我们用一个 map 统计每个字符出现的次数（兼容 unicode 字符），
//      对于 s 中的每个字符，我们给对应的次数 + 1
//      对于 t 中的每个字符，我们给对应的次数 - 1
//
//      最后判断 map 中所有字符的次数是不是全为 0 即可。
//
//
//      设字符集大小为 C 。      
//
//      时间复杂度： O(n + m + C)
//          1. 需要遍历 s 中全部 O(n) 个字符
//          2. 需要遍历 t 中全部 O(m) 个字符
//          3. 需要遍历全部 O(C) 个不同的字符
//      空间复杂度： O(C)
//          1. 需要维护全部 O(C) 个不同字符的次数


use std::collections::HashMap;
use std::ops::{AddAssign, SubAssign};


impl Solution {
    pub fn is_anagram(s: String, t: String) -> bool {
        // counts[ch] 表示 ch 在 s 中出现的次数 减去 在 t 中出现的次数
        let mut counts: HashMap<char, i32> = HashMap::new();

        // 对于 s 中的每个字符，我们给对应的次数 + 1
        for ch in s.chars() {
            counts.entry(ch).or_insert(0).add_assign(1);
        }
        // 对于 t 中的每个字符，我们给对应的次数 - 1
        for ch in t.chars() {
            counts.entry(ch).or_insert(0).sub_assign(1);
        }

        // 如果 counts 中全部字符的出现次数都为 0 ，则 s 和 t 是一对异位词
        counts.values().all(|&count| count == 0)
    }
}
