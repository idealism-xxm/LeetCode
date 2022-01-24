// 链接：https://leetcode.com/problems/detect-capital/
// 题意：给定一个单词 word ，判断是否满足以下三个条件之一：
//          1. 全部字母都是大写，例如 "USA"
//          2. 全部字母都是小写，例如 "leetcode"
//          3. 首字母大写，后续字母都是小写，例如 "Google"

// 数据限制：
//  1 <= word.length <= 100
//  word 仅由英文大小写字母组成

// 输入： word = "USA"
// 输出： true

// 输入： word = "FlaG"
// 输出： false


// 思路：模拟
//
//      直接按照题意模拟判断即可，
//      可以注意到第 2 个和第 3 个条件可以融合成一种情况：除首字母外全是小写字母。
//      （如果首字母是小写，则是第 2 个条件；如果首字母是大写，则是第 3 个条件。）
//
//      时间复杂度： O(n)
//      空间复杂度： O(1)

impl Solution {
    pub fn detect_capital_use(word: String) -> bool {
        // 如果全是大写字母，则满足第 1 个条件
        word.chars().all(|ch| ch.is_uppercase())
        // 如果除了首字母全是小写字母，则符合第 2 个和第 3 个条件
        || word.chars().skip(1).all(|ch| ch.is_lowercase())
    }
}
