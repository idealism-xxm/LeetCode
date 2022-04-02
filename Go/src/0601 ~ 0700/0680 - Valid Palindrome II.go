// 链接：https://leetcode.com/problems/valid-palindrome-ii/
// 题意：给定一个字符串 s ，判断最多删除 1 个字符后，
//      是否能够转换成回文串？


// 数据限制：
//  1 <= s.length <= 10 ^ 5
//  s[i] 是英文小写字母


// 输入： s = "aba"
// 输出： true
// 解释： s 本身就是回文串

// 输入： s = "abca"
// 输出： true
// 解释： 删除 'c' 后，得到回文串 "aba"

// 输入： s = "abc"
// 输出： false
// 解释： 至少需要删除 2 个字符后，才能得到回文串


// 思路： 双指针
//
//      定义左指针 l 和右指针 r，初始化为 0 和 s.length - 1 。
//
//      当 l < r 时，说明还需要比较：
//          1. s[l] != s[r] 时，则需要删除一个字符，有三种情况：
//              (1) 删除 s[l] ，那么只要 s[l + 1..=r] 是回文串，
//                  就满足题意，直接返回 true
//              (2) 删除 s[r] ，那么只要 s[l..=r - 1] 是回文串，
//                  就满足题意，直接返回 true
//              (3) 无论删除哪个字符，都不满足题意，返回 false
//          2. s[l] == s[r] 时，则需要继续比较
//
//      最后如果还未返回，说明 s 本身就是回文串，直接返回 true
//
//
//      时间复杂度：O(n)
//          1. 需要遍历 s 中全部 O(n) 个字符
//      空间复杂度：O(1)
//          1. 只需要维护常数个额外变量


func validPalindrome(s string) bool {
    // 定义左指针 l ，初始化为 0
    l := 0
    // 定义右指针 r ，初始化为 s.length - 1
    r := len(s) - 1
    // 当还有字符需要比较时，继续处理
    for l < r {
        // 如果 s[l] 和 s[r] 不相等，则需要删除字符
        if s[l] != s[r] {
            // 如果删除 左指针 或 右指针 指向的字符后，
            // 能形成回文，则直接返回 true ；
            // 否则，返回 false
            return isPalindrome(s, l + 1, r) || isPalindrome(s, l, r - 1)
        }

        // 此时 s[l] 和 s[r] 相等，可以继续处理。
        // 将 l 向右移动一位
        l += 1
        // 将 r 向左移动一位
        r -= 1
    }

    // 此时说明 s 本身就是回文，直接返回 true
    return true
}

func isPalindrome(s string, l int, r int) bool {
    // 当还有字符需要比较时，继续处理
    for l < r {
        // 如果 s[l] 和 s[r] 不相等，则不是回文，
        // 直接返回 false
        if s[l] != s[r] {
            return false
        }

        // 此时 s[l] 和 s[r] 相等，可以继续处理。
        // 将 l 向右移动一位
        l += 1
        // 将 r 向左移动一位
        r -= 1
    }

    // 此时说明 s 删除 1 个字符后是回文，
    // 直接返回 true
    return true
}
