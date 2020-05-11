# 链接：https://leetcode.com/problems/word-frequency/
# 题意：给定一个文件 words.txt ，文件内包含一些单词，
#       每行的单词用一个或多个空格分割，现在需要统计每个单词都出现次数，
#       并按出现次数降序输出每个单词及其出现次数，每个一行？

# 输入：
#		the day is sunny the the
#   the sunny is is
# 输出：
#		the 4
#   is 3
#   sunny 2
#   day 1

# 思路： tr + sort + uniq + awk
#
#		第一反应还是用高级编程语言那套，直接模拟题意即可，
#		但是发现写起来太痛了，而且很难一行搞定
#
#		看了答案发现还是知道的太少了
#   先用 cat words.txt 获取文件中的内容
#   然后用 tr -s ' ' '\n' 可以将所有的空格替换成换行符，
#     -s 表示换行符如果多次出现，则会压缩成一个
#     这样处理后每行有且仅有一个单词
#   再用 sort 对每行的单词进行排序，因为 uniq 只能对连续出现的字符串进行统计
#   再用 uniq -c 统计单词出现次数， -c 表示还需要输出每个单词的出现次数
#   再用 sort -r 按照单词出现次数进行倒序排序
#   最后在使用 awk '{print $2, $1}' 按照题目要求进行输出即可

# Read from the file file.txt and output all valid phone numbers to stdout.
cat words.txt | tr -s ' ' '\n' | sort | uniq -c | sort -r | awk '{print $2, $1}'
