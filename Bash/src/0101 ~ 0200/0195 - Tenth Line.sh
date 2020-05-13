# 链接：https://leetcode.com/problems/tenth-line/
# 题意：给定一个文件 file.txt ，输出第 10 行第内容？

# 输入：
#		Line 1
#   Line 2
#   Line 3
#   Line 4
#   Line 5
#   Line 6
#   Line 7
#   Line 8
#   Line 9
#   Line 10
# 输出：
#		Line 10

# 思路1： sed
#
#		sed -n NUMp 可以仅输出第 NUM 行第内容

# Read from the file file.txt and print its transposed content to stdout.
sed -n 10p file.txt

# 思路2： wc + head + tail
#
#		wc -l 可以统计行数
#   当 file.txt 行树大于 10 时，先获取头 10 行，在获取其中当最后一行即可

[ `wc -l file.txt` -ge 10 ] && head -10 file.txt | tail -1

# 思路3： awk
#
#		awk 'NR==NUM' 可以输出第 NUM 行的内容
#   awk 'NR<=NUM' 可以输出前 NUM 行的内容

awk 'NR==10' file.txt
