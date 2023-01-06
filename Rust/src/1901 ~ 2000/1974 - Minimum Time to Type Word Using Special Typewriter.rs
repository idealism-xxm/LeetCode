// 链接：https://leetcode.com/problems/minimum-time-to-type-word-using-special-typewriter/
// 题意：给定一个刻着 26 个字母的圆盘，有一个指针初始指向字母 'a' ，
//      每 1 秒可以将指针移向相邻的字母，或者打出指针当前指向的字母。
//
//      求打出给定单词 word 的最小时间？


// 数据限制：
//   1 <= word.length <= 100
//   word 均有英文小写字母组成


// 输入： word = "abc"
// 输出： 5
// 解释： 
//       第 0 秒：指针指向 'a'
//       第 1 秒：打出 'a'
//       第 2 秒：指针顺时针移动至 'b'
//       第 3 秒：打出 'b'
//       第 4 秒：指针顺时针移动至 'c'
//       第 5 秒：打出 'c'

// 输入： word = "bza"
// 输出： 7
// 解释： 
//       第 0 秒：指针指向 'a'
//       第 1 秒：指针顺时针移动至 'b'
//       第 2 秒：打出 'b'
//       第 3 秒：指针逆时针移动至 'a'
//       第 4 秒：指针逆时针移动至 'z'
//       第 5 秒：打出 'z'
//       第 6 秒：指针顺时针移动至 'a'
//       第 7 秒：打出 'a'

// 输入： word = "zjpc"
// 输出： 34
// 解释： 
//       第 0 秒：指针指向 'a'
//       第 1 秒：指针逆时针移动至 'z'
//       第 2 秒：打出 'z'
//       第 3 ~ 12 秒：指针顺时针移动至 'j'
//       第 13 秒：打出 'j'
//       第 14 ~ 19 秒：指针顺时针移动至 'p'
//       第 20 秒：打出 'p'
//       第 21 ~ 33  秒：指针逆时针移动至 'c'
//       第 34 秒：打出 'c'


// 思路： 贪心
//
//       模拟当前指针指向字母的下标 cur ，遍历下一个要打的字母的下标 idx 。
//
//       假设顺时针从 cur 移动至 idx ，则要消耗 cnt = idx - cur 秒，
//       如果 cnt < 0 ，则进入了下一圈，需要加上 26 。
//
//       那么逆时针从 cur 移动至 idx ，则要消耗 26 - cnt 秒，两者取较小值即可。
//
//       即打印出 idx 的字母耗时为 min(cnt, 26 - cnt) + 1 秒。
//
//
//       时间复杂度： O(n)
//          1. 需要遍历 word 中全部 O(n) 个字符
//       空间复杂度： O(1)
//          1. 只需要维护常数个额外变量


impl Solution {
    pub fn min_time_to_type(word: String) -> i32 {
        // ans 维护打完 word 中全部字符所需的时间
        let mut ans = 0;
        // cur 维护指针当前指向的字母下标
        let mut cur = 0;
        // 遍历 word 中每个单词
        for &ch in word.as_bytes() {
            // 计算该字母的下标
            let idx = ch as i32 - b'a' as i32;
            
            // 计算顺时针移动所需的时间
            let mut cnt = idx - cur;
            if cnt < 0 {
                cnt += 26;
            }
            // 移动时间为顺时针和逆时针时间的较小值，
            // 还要加上打印字母的时间
            ans += cnt.min(26 - cnt) + 1;
            // 指针移动至 idx 对应的字母
            cur = idx
        }
        return ans
    }
}