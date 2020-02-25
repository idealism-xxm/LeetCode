// 链接：https://leetcode.com/problems/implement-strstr/
// 题意：给定一个字符串和一个子串，求字符串中第一次出现这个子串的下标，没有则返回 -1

// 输入：nums = [0,1,2,2,3,0,4,2], val = 2
// 输出：5

// 思路1：枚举
//      直接枚举起点下标即可
//      时间复杂度：O(mn)

func strStr(haystack string, needle string) int {
    m, n := len(haystack), len(needle)
    if n == 0 {
        return 0
    }

    for i := 0; i < m; i++ {
        flag := true
        for j := 0; j < n; j++ {
            if  i + j >= m || haystack[i + j] != needle[j] {
                flag = false
                break
            }
        }
        if flag {
            return i
        }
    }
    return -1
}

// 思路2：KMP
//      时间复杂度：O(n + m)

func strStr(haystack string, needle string) int {
    n, m := len(haystack), len(needle)
    if m == 0 {
        return 0
    }

    next := getNext(needle)
    i, j := 0, 0
    for i < n && j < m  {
        if j == -1 || haystack[i] == needle[j] { // 字符相等，则同时移动一个字符
            i++
            j++
        } else {
            j = next[j] // 字符不同，j 移动到 j 之前的前缀相同但当前字符不同的最后一个位置
        }
    }

    if j == m {
        return i - j
    }
    return -1
}

// 求模式串的 next 数组
func getNext(pattern string) []int {
    m := len(pattern)
    next := make([]int, m) // next[j] 表示母串当前匹配字符与 pattern[j] 不匹配时，j 应该变为 next[j]
    next[0] = -1
    j, k := 0, -1

    for j < m - 1{
        if k == -1 || pattern[j] == pattern[k] {
            j++ // 移动 j
            k++ // 移动 k
            // 优化，使 next[j] 指向 j 之前的前缀相同但当前字符不同的最后一个位置
            if pattern[j] == pattern[k] { 
                next[j] = next[k] // 如果后一个字符相同，则需要优化为 next[k]
            } else {
                next[j] = k // 如果后一个字符不相同，则直接赋值 next[j] = k，
            }
        } else {
            k = next[k]
        }
    }
    
    return next
}