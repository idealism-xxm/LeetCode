// 链接：https://leetcode.com/problems/remove-duplicate-letters/
// 题意：给定一个字符串 s ，删除掉所有重复的字符，使得每个字符只出现一次。
//      返回所有可能的结果中字典序最小的一个。


// 数据限制：
//  1 <= s.length <= 10 ^ 4
//  s 仅由英文小写字母组成


// 输入： s = "bcabc"
// 输出： "abc"
// 解释： 所有可能的删除结果是 "bca"、"bac"、"abc" 和 "cab" ，
//       其中 "abc" 的字典序最小

// 输入： s = "cbacdcbc"
// 输出： "acdb"
// 解释： "acdb" 是所有结果中字典序最小的一个


// 思路： 贪心 + 单调栈
//
//      我们可以用 lastIndex 维护每个字符最后一次出现的位置，
//      然后定义 stack 记录栈中的字符，
//      再用 isInStack 维护某个字符是否在栈中。
//
//      重新遍历 s 中的第 i 个字符 ch ，
//      如果 ch 不在栈中，则需要把 ch 入栈。
//
//      此时，如果栈顶字符 top 大于 ch 且之后还有字符 top ，
//      那么我们可以贪心地弹出栈顶字符 top ，并标记 top 不在栈中，
//      这样后续再放入 top 时，字典序必定更小。
//
//      不断重复这个过程，直至栈为空 或者 top 小于 ch 或者 之后无字符 top 。
//
//      然后我们把 ch 入栈，并标记 ch 在栈中。
//
//
//      设 K 为字符集大小
//      
//      时间复杂度：O(n)
//          1. 需要遍历 s 中全部 O(n) 个字符
//      空间复杂度：O(K)
//          1. 需要维护 lastIndex ，存储全部 O(K) 个不同的字符
//          2. 需要维护 isInStack ，存储全部 O(K) 个不同的字符
//          3. 需要维护一个栈 stack ，存储全部 O(K) 个不同的字符
//          4. 最后的结果有 O(K) 个不同的字符


func removeDuplicateLetters(s string) string {
    // lastIndex[ch] 表示 ch 在 s 中的最后一个出现的位置
    lastIndex := make(map[rune]int)
    // 带下标遍历 s 中的字符
    for i, ch := range s {
        // 更新每个字符最后一次出现的位置
        lastIndex[ch] = i
    }

    // isInStack[ch] 表示 ch 是否在栈中
    isInStack := make(map[rune]bool)
    // 用栈 stack 收集当前的结果
    stack := make([]rune, 0)
    // 带下标遍历 s 中的字符
    for i, ch := range s {
        // 如果当前字符 ch 不在栈中，则需要入栈
        if !isInStack[ch] {
            // 如果栈顶字符 top 比 ch 大，且 top 不是 s 中最后一个字符，
            // 那么那可以先出栈 top ，这样后续再放入时，字典序必定更小
            for ; len(stack) != 0 && stack[len(stack) - 1] > ch && lastIndex[stack[len(stack) - 1]] > i; {
                // 栈顶字符 top 出栈，并标记为不在栈中
                isInStack[stack[len(stack) - 1]] = false
                stack = stack[:len(stack) - 1]
            }

            // 将当前字符 ch 入栈
            stack = append(stack, ch)
            // 标记当前字符 ch 为在栈中
            isInStack[ch] = true
        }
    }
    
    // 最后栈中的字符组成的字符串就是字典序最小的
    return string(stack)
}
