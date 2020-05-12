# 链接：https://leetcode.com/problems/transpose-file/
# 题意：给定一个文件 file.txt ，文件内是 m * n 的单词矩阵，
#       每一个行的单词通过空格分割，
#       现在需要转置后输出为 n * m 的单词矩阵？

# 输入：
#		name age
#   alice 21
#   ryan 30
# 输出：
#		name alice ryan
#   age 21 30

# 思路： head + awk + while + expr
#
#		首先通过 head 获取文件第一行，然后通过计算出列数，
#   再通过 while 循环，每次读取文件，将每行的第 index 列输出，
#   再经由 xargs 将换行转换成空格，然后使用 expr 对 index 执行 +1

# Read from the file file.txt and print its transposed content to stdout.
# 读取第一行，并用统计列数
columnsCount=`head -1 file.txt | awk '{print NF}'`
# 循环 columnsCount 次
index=1
while [ $index -le $columnsCount ]
do
    # 每次将文件每一行的第 index 列输出即可，最后通过 xargs 将换行和空白替换成空格
    cat file.txt | awk -v i=$index '{print $i}' | xargs
    index=`expr $index + 1`
done
