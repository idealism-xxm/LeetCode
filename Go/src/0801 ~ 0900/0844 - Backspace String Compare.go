// 链接：https://leetcode.com/problems/backspace-string-compare/
// 题意：给定两个字符串 s 和 t ，其中 '#' 代表退格。
//      现在将 s 与 t 输入到文本编辑器中，
//      如果最后的文本一致，则返回 true ，否则返回 false 。


// 数据限制：
//  1 <= s.length, t.length <= 200
//  s 和 t only 仅含有英文小写字母和 '#'


// 输入： s = "ab#c", t = "ad#c"
// 输出： true
// 解释： s 和 t 都会变为 "ac"

// 输入： s = "ab##", t = "c#d#"
// 输出： true
// 解释： s 和 t 都会变为 ""

// 输入： s = "a#c", t = "b"
// 输出： false
// 解释： s 会变成 "c", 而 t 会变成 "b"


// 思路： 栈
//
//      由于这个文本编辑器的过程只有两个操作：输入字符和退格，
//      即后进先出，所以可以用栈来模拟这个过程。
//
//      初始化一个空栈 stack ，
//      然后遍历字符串中的每一个字符 ch ：
//          1. ch == '#': 则将 stack 中的栈顶元素弹出，
//              如果栈为空，则不做处理
//          2. ch != '#': 则将 ch 压入 stack 中
//
//      最后栈 stack 中的字符就是字符串对应的文本。
//
//      获取 s 和 t 对应的文本，然后返回两个文本的比较结果即可。
//
//
//      时间复杂度：O(n + m)
//          1. 需要遍历 s 中全部 O(n) 个字符
//          2. 需要遍历 t 中全部 O(m) 个字符
//      空间复杂度：O(n + m)
//          1. 需要维护 s 对应的文本，最差情况下有 O(n) 个字符
//          2. 需要维护 t 对应的文本，最差情况下有 O(m) 个字符         


func backspaceCompare(s string, t string) bool {
    // 获取 s 和 t 对应的文本，然后返回比较结果
    return getText(s) == getText(t)
}

func getText(s string) string {
    // 初始化一个空栈
    var stack []rune
    // 遍历字符串中的每一个字符
    for _, ch := range s {
        if ch == '#' {
            // 如果是 '#' ，则将栈顶元素弹出
            if len(stack) > 0 {
                stack = stack[:len(stack)-1]
            }
        } else {
            // 否则，将字符压入栈中
            stack = append(stack, ch)
        }
    }

    // 最后将所有字符转换为字符串
    return string(stack)
}


// 思路2： 双指针
//
//      从前往后遍历时，无法知道一个字符是否会在最终文本中存在，
//      所以我们可以从后往前遍历。
//
//      从后往前遍历字符串 s ，
//      同时维护还需要跳过的字符数 sSkip ：
//          1. 当遇到一个 '#' 时，对 sSkip 进行加 1
//          2. 当遇到一个英文字母时，根据 sSkip 的值进行处理：
//              (1) sSkip == 0: 当前字符会出现在最终文本中
//              (2) sSkip != 0: 当前字符需要跳过，对 sSkip 进行减 1
//
//      我们可以对 t 进行同样的操作，
//      每次找到 s 和 t 的最终文本中必定存在的字符，
//      如果字符不同 或者 有且仅有一个遍历完成，
//      则说明两个字符串的最终文本不同，直接返回 false 。
//
//      最后 s 和 t 中的所有字符都全部遍历完成，
//      则说明两个字符串的最终文本相同，直接返回 true 。
//
//
//      时间复杂度：O(n + m)
//          1. 需要遍历 s 中全部 O(n) 个字符
//          2. 需要遍历 t 中全部 O(m) 个字符
//      空间复杂度：O(1)
//          1. 只需要维护常数个额外遍历即可       


func backspaceCompare(s string, t string) bool {
    // 初始化 s 对应的指针 sIndex 和当前应该跳过的字符数 sSkip
    sIndex, sSkip := len(s) - 1, 0
    // 初始化 t 对应的指针 tIndex 和当前应该跳过的字符数 tSkip
    tIndex, tSkip := len(t) - 1, 0
    // 当 s 和 t 至少有一个还有字符时，继续处理
    for sIndex >= 0 || tIndex >= 0 {
        // 当 s 还有字符时，需要继续处理
        for sIndex >= 0 {
            if s[sIndex] == '#' {
                // 如果当前字符是 '#' ，则需要跳过的字符数加 1
                sSkip += 1
            } else if sSkip > 0 {
                // 如果当前字符不是 '#' ，且还需要跳过字符，
                // 则跳过当前字符
                sSkip -= 1
            } else {
                // 如果当前字符不是 '#' ，且不需要跳过字符，
                // 则当前字符在最终文本中，跳出循环即可
                break
            }
            // 将 sIndex 向前移动一位
            sIndex -= 1
        }
        // 当 t 还有字符时，需要继续处理
        for tIndex >= 0 {
            if t[tIndex] == '#' {
                // 如果当前字符是 '#' ，则需要跳过的字符数加 1
                tSkip += 1
            } else if tSkip > 0 {
                // 如果当前字符不是 '#' ，且还需要跳过字符，
                // 则跳过当前字符
                tSkip -= 1
            } else {
                // 如果当前字符不是 '#' ，且不需要跳过字符，
                // 则当前字符在最终文本中，跳出循环即可
                break
            }
            // 将 tIndex 向前移动一位
            tIndex -= 1
        }

        // 处理 s 和 t 的当前字符
        if sIndex >= 0 && tIndex >= 0 {
            // 如果两者都有字符，则继续比较当前字符。
            // 两者字符不同，则最终文本不同，直接返回 false
            if s[sIndex] != t[tIndex] {
                return false
            }
        } else if sIndex < 0 && tIndex < 0 {
            // 如果两者都无字符，则最终文本相同，直接返回 true
            return true
        } else {
            // 如果一个有字符，一个无字符，则最终文本不同，直接返回 false
            return false
        }

        // s 和 t 的当前字符处理完毕，
        // 将 sIndex 和 tIndex 向前移动一位
        sIndex -= 1
        tIndex -= 1
    }
    // 此时 s 和 t 都没有字符了，
    // 说明最终文本完全匹配，返回 true
    return true
}
