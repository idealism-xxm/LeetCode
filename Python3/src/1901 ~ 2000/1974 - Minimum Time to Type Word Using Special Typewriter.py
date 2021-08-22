# 链接：https://leetcode.com/problems/minimum-time-to-type-word-using-special-typewriter/
# 题意：给定一个刻着 26 个字母的圆盘，有一个指针初始指向字母 'a' ，
#       每 1 秒可以讲指针移向相邻的字母，或者打出当前指针指向的字母。
#       求打出给定单词 word 的最小时间？

# 数据限制：
#   1 <= word.length <= 100
#   word 均有英文小写字母组成

# 输入： word = "abc"
# 输出： 5
# 解释： 
#       第 0 秒：指针指向 'a'
#       第 1 秒：打出 'a'
#       第 2 秒：指针顺时针移动至 'b'
#       第 3 秒：打出 'b'
#       第 4 秒：指针顺时针移动至 'c'
#       第 5 秒：打出 'c'

# 输入： word = "bza"
# 输出： 7
# 解释： 
#       第 0 秒：指针指向 'a'
#       第 1 秒：指针顺时针移动至 'b'
#       第 2 秒：打出 'b'
#       第 3 秒：指针逆时针移动至 'a'
#       第 4 秒：指针逆时针移动至 'z'
#       第 5 秒：打出 'z'
#       第 6 秒：指针顺时针移动至 'a'
#       第 7 秒：打出 'a'

# 输入： word = "zjpc"
# 输出： 34
# 解释： 
#       第 0 秒：指针指向 'a'
#       第 1 秒：指针逆时针移动至 'z'
#       第 2 秒：打出 'z'
#       第 3 ~ 12 秒：指针顺时针移动至 'j'
#       第 13 秒：打出 'j'
#       第 14 ~ 19 秒：指针顺时针移动至 'p'
#       第 20 秒：打出 'p'
#       第 21 ~ 33  秒：指针逆时针移动至 'c'
#       第 34 秒：打出 'c'


# 思路： 模拟
#
#       模拟当前指针指向字母的下标 cur ，
#       遍历下一个要打的字母的下标 idx ，
#       假设顺时针从 cur 移动至 idx ，则要消耗 cnt = idx - cur 秒，
#       如果 cnt < 0 ，则进入了下一圈，需要加上 26
#       那么逆时针从 cur 移动至 idx ，则要消耗 26 - cnt 秒，
#       两者取较小值即可，
#       那么打印出 idx 的字母耗时为 min(cnt, 26 - cnt) + 1 秒
#       
#       时间复杂度： O(|word|)
#       空间复杂度： O(1)

class Solution:
    def minTimeToType(self, word: str) -> int:
        a_ord = ord('a')
        cur = 0
        ans = 0
        for ch in word:
            # 计算字母的下标
            idx = ord(ch) - a_ord
            
            # 计算顺时针移动所需的时间
            cnt = idx - cur
            if cnt < 0:
                cnt += 26
            # 移动时间为顺时针和逆时针时间的较小值
            # 还要加上打印字母的时间
            ans += min(cnt, 26 - cnt) + 1
            # 指针移动至 idx 对应的字母
            cur = idx
        return ans
