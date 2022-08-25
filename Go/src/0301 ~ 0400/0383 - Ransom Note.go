// 链接：https://leetcode.com/problems/ransom-note/
// 题意：给定两个字符串 ransomNote 和 magazine ，
//      判断 ransomNote 是否能由 magazine 中的字母组成？
//
//      magazine 中的每个字母只能在 ransomNote 中使用一次。


// 数据限制：
//  1 <= ransomNote.length, magazine.length <= 10 ^ 5
//  ransomNote 和 magazine 仅由英文小写字母组成


// 输入： ransomNote = "a", magazine = "b"
// 输出： false

// 输入： ransomNote = "aa", magazine = "ab"
// 输出： false

// 输入： ransomNote = "aa", magazine = "aab"
// 输出： true


// 思路： Map
//
//      本题和 LeetCode 242 基本一致，所以同样可以采用统计字符出现次数的方式处理。
//
//      我们用一个 map 统计每个字符出现的次数（兼容 unicode 字符），
//      对于 magazine 中的每个字符，我们给对应的次数 + 1
//      对于 ransomNote 中的每个字符，我们给对应的次数 - 1
//
//      最后判断 map 中所有字符的次数是不是全都大于等于 0 即可。
//
//
//      设字符集大小为 C 。      
//
//      时间复杂度： O(n + m + C)
//          1. 需要遍历 ransomNote 中全部 O(n) 个字符
//          2. 需要遍历 magazine 中全部 O(m) 个字符
//          3. 需要遍历全部 O(C) 个不同的字符
//      空间复杂度： O(C)
//          1. 需要维护全部 O(C) 个不同字符的次数


func canConstruct(ransomNote string, magazine string) bool {
    // counts[ch] 表示 ch 在 magazine 中出现的次数 减去 在 ransomNote 中出现的次数
    counts := make(map[rune]int)

    // 对于 magazine 中的每个字符，我们给对应的次数 + 1
    for _, ch := range magazine {
        counts[ch]++
    }
    // 对于 ransomNote 中的每个字符，我们给对应的次数 - 1
    for _, ch := range ransomNote {
        counts[ch]--
    }

    // 如果 counts 中全部字符的出现次数都大于等于 0 ，
    // 则 ransomNote 能由 magazine 中的字母组成
    for _, count := range counts {
        if count < 0 {
            return false
        }
    }
    return true
}
